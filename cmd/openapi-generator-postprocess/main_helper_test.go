package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestHelperProcess is not a real test. It is re-executed as a subprocess by
// runMain so that main() (which calls flag.Parse() against os.Args and may
// call os.Exit) can be exercised without tearing down the real test binary.
// See https://pkg.go.dev/os/exec#Cmd for the pattern this follows.
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
	os.Args = append([]string{"openapi-generator-postprocess"}, args...)

	main()
}

// runMain re-executes the current test binary, invoking main() with args in
// a fresh process so os.Exit calls inside main/fatal don't kill the real
// test run.
func runMain(t *testing.T, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()

	cs := append([]string{"-test.run=TestHelperProcess", "--"}, args...)
	cmd := exec.Command(os.Args[0], cs...) //nolint:gosec
	cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")

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

func TestMainMissingFlags(t *testing.T) {
	t.Parallel()

	_, _, code := runMain(t)
	if code != 2 {
		t.Fatalf("exit code = %d, want 2", code)
	}
}

func TestMainSuccess(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	modelsDir := filepath.Join(dir, "models")
	if err := os.MkdirAll(modelsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	writeTestFile(t, filepath.Join(modelsDir, "model_device.go"), "package models\n\ntype Device struct{}\n")

	mappingsPath := filepath.Join(dir, "mappings.cfg")
	writeTestFile(t, mappingsPath, "")

	aliasesPath := filepath.Join(dir, "model_aliases.go")

	_, stderr, code := runMain(t,
		"-models", modelsDir,
		"-mappings", mappingsPath,
		"-aliases", aliasesPath,
		"-alias-package", "unifi",
		"-alias-import", "example.com/client/models",
	)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0 (stderr: %s)", code, stderr)
	}

	got := readTestFile(t, aliasesPath)
	if !strings.Contains(got, "type Device = models.Device") {
		t.Errorf("aliases file missing Device alias: %s", got)
	}
}

func TestMainFatalOnReadMappingsError(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	modelsDir := filepath.Join(dir, "models")
	if err := os.MkdirAll(modelsDir, 0o755); err != nil {
		t.Fatal(err)
	}

	_, stderr, code := runMain(t,
		"-models", modelsDir,
		"-mappings", filepath.Join(dir, "does-not-exist.cfg"),
		"-aliases", filepath.Join(dir, "model_aliases.go"),
		"-alias-package", "unifi",
		"-alias-import", "example.com/client/models",
	)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if stderr == "" {
		t.Error("expected fatal() to write the error to stderr")
	}
}

func TestMainFatalOnBundleModelsError(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	modelsDir := filepath.Join(dir, "models")
	if err := os.MkdirAll(modelsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	writeTestFile(t, filepath.Join(modelsDir, "model_bad.go"), "package models\n\nfunc {{{ invalid\n")

	mappingsPath := filepath.Join(dir, "mappings.cfg")
	writeTestFile(t, mappingsPath, "")

	_, stderr, code := runMain(t,
		"-models", modelsDir,
		"-mappings", mappingsPath,
		"-aliases", filepath.Join(dir, "model_aliases.go"),
		"-alias-package", "unifi",
		"-alias-import", "example.com/client/models",
	)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if stderr == "" {
		t.Error("expected fatal() to write the error to stderr")
	}
}
