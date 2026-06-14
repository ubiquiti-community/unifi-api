package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBundleModelsAndWriteAliases(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	writeTestFile(t, filepath.Join(dir, "model_device.go"), `package models

import "encoding/json"

type Device struct{}

var _ = json.Valid
`)
	writeTestFile(t, filepath.Join(dir, "model_device_state.go"), `package models

import "fmt"

type DeviceState int

var _ = fmt.Sprint
`)
	writeTestFile(t, filepath.Join(dir, "model_client_group.go"), `package models

type ClientGroup struct{}
`)

	if err := bundleModels(dir, map[string]string{"DeviceState": "Device"}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, "model_device_state.go")); !os.IsNotExist(err) {
		t.Fatalf("child model file still exists: %v", err)
	}

	device := readTestFile(t, filepath.Join(dir, "model_device.go"))
	for _, want := range []string{`"encoding/json"`, `"fmt"`, "type Device struct", "type DeviceState int"} {
		if !strings.Contains(device, want) {
			t.Errorf("bundled model missing %q", want)
		}
	}

	aliases := filepath.Join(dir, "model_aliases.go")
	if err := writeAliases(dir, aliases, "unifi", "example.com/client/models"); err != nil {
		t.Fatal(err)
	}
	gotAliases := readTestFile(t, aliases)
	for _, want := range []string{
		`import models "example.com/client/models"`,
		"type ClientGroup = models.ClientGroup",
		"type Device = models.Device",
		"type DeviceState = models.DeviceState",
	} {
		if !strings.Contains(gotAliases, want) {
			t.Errorf("aliases missing %q", want)
		}
	}
}

func writeTestFile(t *testing.T, path, contents string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatal(err)
	}
}

func readTestFile(t *testing.T, path string) string {
	t.Helper()
	contents, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(contents)
}
