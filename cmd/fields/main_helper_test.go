package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestHelperProcess is not a real test. It is re-executed as a subprocess by
// runFieldsMain so that main() (which calls flag.Parse() against os.Args and
// may call os.Exit) can be exercised without tearing down the real test
// binary. See https://pkg.go.dev/os/exec#Cmd for the pattern this follows.
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	args := os.Args
	for i, a := range args {
		if a == "--" {
			args = args[i+1:]
			break
		}
	}
	os.Args = append([]string{"fields"}, args...)

	main()
}

// runFieldsMain re-executes the current test binary with dir as its working
// directory, invoking main() with args in a fresh process so os.Exit calls
// inside main() don't kill the real test run.
func runFieldsMain(t *testing.T, dir string, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()

	cs := append([]string{"-test.run=TestHelperProcess", "--"}, args...)
	cmd := exec.Command(os.Args[0], cs...) //nolint:gosec
	cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
	cmd.Dir = dir

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if err == nil {
		return outBuf.String(), errBuf.String(), 0
	}
	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) {
		t.Fatalf("failed to run helper process: %v", err)
	}
	return outBuf.String(), errBuf.String(), exitErr.ExitCode()
}

// seedFieldsCache creates a minimal, already-extracted fields cache for
// version under the package directory, which is where main() looks for cached
// definitions (next to its own source file). Setting.json is the cache marker
// main() checks before deciding to download, so seeding it keeps these tests
// off the network. The directory is removed when the test ends.
func seedFieldsCache(t *testing.T, version string) string {
	t.Helper()

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("unable to locate the package directory")
	}
	dir := filepath.Join(filepath.Dir(file), "v"+version)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })

	for name, content := range map[string]string{
		"Setting.json":     testSettingJSON,
		"SettingMgmt.json": `{"x_ssh_enabled": "true|false"}`,
		"User.json":        testUserJSON,
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return version
}

func TestUsage(t *testing.T) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	usage()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "Usage:") {
		t.Errorf("usage() output = %q, want it to contain %q", buf.String(), "Usage:")
	}
}

func TestMainNoVersionOrLatest(t *testing.T) {
	t.Parallel()

	stdout, _, code := runFieldsMain(t, t.TempDir())
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if !strings.Contains(stdout, "must specify version or latest") {
		t.Errorf("stdout = %q, want it to mention the missing version", stdout)
	}
}

func TestMainVersionAndLatestConflict(t *testing.T) {
	t.Parallel()

	stdout, _, code := runFieldsMain(t, t.TempDir(), "-latest", "9.5.21")
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if !strings.Contains(stdout, "cannot specify version with latest") {
		t.Errorf("stdout = %q, want it to mention the conflict", stdout)
	}
}

func TestMainInvalidVersion(t *testing.T) {
	t.Parallel()

	_, _, code := runFieldsMain(t, t.TempDir(), "not-a-version")
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
}

func TestMainDownloadOnlyUsesCachedFields(t *testing.T) {
	t.Parallel()

	ver := seedFieldsCache(t, "0.0.0-test-download-only")
	stdout, stderr, code := runFieldsMain(t, t.TempDir(), "-download-only", ver)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0 (stderr: %s)", code, stderr)
	}
	if !strings.Contains(stdout, "Fields JSON ready!") {
		t.Errorf("stdout = %q, want it to report readiness", stdout)
	}
}

func TestMainFullGenerationUsesCachedFields(t *testing.T) {
	t.Parallel()

	ver := seedFieldsCache(t, "0.0.0-test-generate")
	dir := t.TempDir()
	stdout, stderr, code := runFieldsMain(t, dir, "-assets-dir", "out", ver)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0 (stderr: %s)", code, stderr)
	}
	if !strings.Contains(stdout, "Generated OpenAPI spec:") {
		t.Errorf("stdout = %q, want it to report the generated spec", stdout)
	}

	for _, want := range []string{
		filepath.Join("out", "openapi.yaml"),
		filepath.Join("out", "openapi-generator", "go-name-mappings.cfg"),
		filepath.Join("out", "oapi-codegen-exp", "unifi.yaml"),
	} {
		if _, err := os.Stat(filepath.Join(dir, want)); err != nil {
			t.Errorf("expected generated file %s: %v", want, err)
		}
	}
}
