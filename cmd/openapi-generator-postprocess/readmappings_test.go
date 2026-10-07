package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestReadMappings(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "mappings.cfg")
	writeTestFile(t, path, "# a comment\n\nDeviceState=Device\nClientGroup=Client\n")

	got, err := readMappings(path)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"DeviceState": "Device", "ClientGroup": "Client"}
	if len(got) != len(want) {
		t.Fatalf("readMappings = %v, want %v", got, want)
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("mapping[%q] = %q, want %q", k, got[k], v)
		}
	}
}

func TestReadMappingsMissingFile(t *testing.T) {
	t.Parallel()

	if _, err := readMappings(filepath.Join(t.TempDir(), "missing.cfg")); err == nil {
		t.Fatal("expected error for missing mappings file")
	}
}

func TestReadMappingsInvalidLine(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "mappings.cfg")
	writeTestFile(t, path, "not-a-valid-mapping-line\n")

	if _, err := readMappings(path); err == nil {
		t.Fatal("expected error for invalid mapping line")
	} else if !strings.Contains(err.Error(), "invalid mapping") {
		t.Errorf("error = %v, want it to mention 'invalid mapping'", err)
	}
}

func TestBundleModelsSkipsUnknownMappings(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	writeTestFile(t, filepath.Join(dir, "model_device.go"), "package models\n\ntype Device struct{}\n")

	// Neither side of either mapping resolves to a known+distinct pair, so
	// bundleModels should leave the model file untouched.
	err := bundleModels(dir, map[string]string{
		"Ghost":  "Device",
		"Device": "AlsoGhost",
	})
	if err != nil {
		t.Fatal(err)
	}

	content := readTestFile(t, filepath.Join(dir, "model_device.go"))
	if !strings.Contains(content, "type Device struct") {
		t.Fatalf("model file unexpectedly changed: %s", content)
	}
}

func TestBundleModelsSkipsSameFileMapping(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	writeTestFile(t, filepath.Join(dir, "model_device.go"), "package models\n\ntype Device struct{}\ntype DeviceExtra struct{}\n")

	// DeviceExtra and Device are declared in the same file, so the mapping is
	// a no-op (childPath == ownerPath).
	if err := bundleModels(dir, map[string]string{"DeviceExtra": "Device"}); err != nil {
		t.Fatal(err)
	}

	content := readTestFile(t, filepath.Join(dir, "model_device.go"))
	for _, want := range []string{"type Device struct", "type DeviceExtra struct"} {
		if !strings.Contains(content, want) {
			t.Errorf("model file missing %q: %s", want, content)
		}
	}
}

func TestBundleModelsMergesNamedImports(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	writeTestFile(t, filepath.Join(dir, "model_device.go"), "package models\n\ntype Device struct{}\n")
	writeTestFile(t, filepath.Join(dir, "model_device_state.go"), `package models

import j "encoding/json"

type DeviceState int

var _ = j.Valid
`)

	if err := bundleModels(dir, map[string]string{"DeviceState": "Device"}); err != nil {
		t.Fatal(err)
	}

	content := readTestFile(t, filepath.Join(dir, "model_device.go"))
	if !strings.Contains(content, `j "encoding/json"`) {
		t.Errorf("merged file missing named import: %s", content)
	}
}

func TestParseModelFilesInvalidSyntax(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	writeTestFile(t, filepath.Join(dir, "model_bad.go"), "package models\n\nfunc {{{ invalid\n")

	if _, _, err := parseModelFiles(dir); err == nil {
		t.Fatal("expected a parse error for invalid Go syntax")
	}
}
