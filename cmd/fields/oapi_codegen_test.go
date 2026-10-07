package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestWriteOAPICodegenConfigs(t *testing.T) {
	resources := append(sampleResources(), NewResource("SettingMgmt", "mgmt"))

	dir := t.TempDir()
	configDir := filepath.Join(dir, "oapi-codegen-exp")
	clientsDir := filepath.Join(dir, "clients", "go")

	require.NoError(t, WriteOAPICodegenConfigs(resources, configDir, filepath.Join(dir, "openapi.yaml"), clientsDir))

	for _, dirPath := range []string{clientsDir, filepath.Join(clientsDir, "settings")} {
		info, err := os.Stat(dirPath)
		require.NoError(t, err)
		assert.True(t, info.IsDir())
	}

	var unifiCfg oapiConfig
	data, err := os.ReadFile(filepath.Join(configDir, "unifi.yaml"))
	require.NoError(t, err)
	require.NoError(t, yaml.Unmarshal(data, &unifiCfg))
	assert.Equal(t, "unifi", unifiCfg.Package)
	assert.Equal(t, filepath.Join(clientsDir, "unifi.gen.go"), unifiCfg.Output)
	assert.Contains(t, unifiCfg.OutputOptions.IncludeTags, "Network")
	assert.Contains(t, unifiCfg.OutputOptions.IncludeTags, "FirewallPolicy")
	assert.NotContains(t, unifiCfg.OutputOptions.IncludeTags, "SettingMgmt")
	assert.Equal(t, "Name", unifiCfg.NameSubstitutions.PropertyNames["name"])
	assert.Equal(t, "Network", unifiCfg.NameSubstitutions.TypeNames["Network"])

	var settingsCfg oapiConfig
	data, err = os.ReadFile(filepath.Join(configDir, "settings.yaml"))
	require.NoError(t, err)
	require.NoError(t, yaml.Unmarshal(data, &settingsCfg))
	assert.Equal(t, "settings", settingsCfg.Package)
	assert.Contains(t, settingsCfg.OutputOptions.IncludeTags, "SettingMgmt")
	assert.NotContains(t, settingsCfg.OutputOptions.IncludeTags, "Network")
}

func TestWriteOAPICodegenConfigsIDCollision(t *testing.T) {
	r := NewResource("Widget", "widgets")
	r.Types[r.StructName].Fields["ExtraID"] = NewFieldInfo("ID", "id", String, "", true, false, false, "")

	dir := t.TempDir()
	require.NoError(t, WriteOAPICodegenConfigs(
		[]*ResourceInfo{r},
		filepath.Join(dir, "cfg"),
		filepath.Join(dir, "openapi.yaml"),
		filepath.Join(dir, "clients"),
	))

	data, err := os.ReadFile(filepath.Join(dir, "cfg", "unifi.yaml"))
	require.NoError(t, err)
	var cfg oapiConfig
	require.NoError(t, yaml.Unmarshal(data, &cfg))
	assert.Equal(t, "Id", cfg.NameSubstitutions.PropertyNames["id"])
	assert.Equal(t, "ID", cfg.NameSubstitutions.PropertyNames["_id"])
}

func TestWriteOAPICodegenConfigsInvalidConfigDir(t *testing.T) {
	dir := t.TempDir()
	blocker := filepath.Join(dir, "blocker")
	require.NoError(t, os.WriteFile(blocker, []byte("x"), 0o644))

	err := WriteOAPICodegenConfigs(sampleResources(), filepath.Join(blocker, "cfg"), "spec.yaml", filepath.Join(dir, "clients"))
	assert.ErrorContains(t, err, "mkdir")
}

func TestWriteOAPICodegenConfigsInvalidClientsDir(t *testing.T) {
	dir := t.TempDir()
	blocker := filepath.Join(dir, "blocker")
	require.NoError(t, os.WriteFile(blocker, []byte("x"), 0o644))

	err := WriteOAPICodegenConfigs(sampleResources(), filepath.Join(dir, "cfg"), "spec.yaml", filepath.Join(blocker, "clients"))
	assert.ErrorContains(t, err, "mkdir")
}
