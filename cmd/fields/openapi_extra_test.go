package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagKind(t *testing.T) {
	tests := []struct {
		structName string
		want       string
	}{
		{"SettingMgmt", "settings"},
		{"PortForward", "network"},
		{"RADIUSProfile", "security"},
		{"FirewallPolicy", "firewall"},
		{"Nat", "firewall"},
		{"Client", "clients"},
		{"Account", "clients"},
		{"WLAN", "wireless"},
		{"APGroup", "wireless"},
		{"ChannelPlan", "wireless"},
		{"Device", "devices"},
		{"PortProfile", "devices"},
		{"TrafficRoute", "routing"},
		{"BGPConfig", "routing"},
		{"OSPFRouter", "routing"},
		{"DynamicDNS", "routing"},
		{"Network", "network"},
		{"DNSRecord", "network"},
		{"HeatMap", "resource"},
	}
	for _, tt := range tests {
		t.Run(tt.structName, func(t *testing.T) {
			r := NewResource(tt.structName, "path")
			assert.Equal(t, tt.want, tagKind(r))
		})
	}
}

func TestTagParent(t *testing.T) {
	names := map[string]bool{"Client": true, "ClientGroup": true, "Foo": true, "FooProfile": true, "Standalone": true}

	cg := &ResourceInfo{StructName: "ClientGroup"}
	assert.Equal(t, "Client", tagParent(cg, names))

	fp := &ResourceInfo{StructName: "FooProfile"}
	assert.Equal(t, "Foo", tagParent(fp, names))

	// Explicit parent present in the map, but not among the emitted tag
	// names -> no reference is emitted (avoids a dangling tag parent).
	orphan := map[string]bool{"ClientGroup": true}
	assert.Equal(t, "", tagParent(cg, orphan))

	// No explicit or suffix-derived parent.
	assert.Equal(t, "", tagParent(&ResourceInfo{StructName: "Standalone"}, names))
}

func TestOpenapiType(t *testing.T) {
	assert.Equal(t, "string", openapiType(String))
	assert.Equal(t, "integer", openapiType(Int))
	assert.Equal(t, "boolean", openapiType(Bool))
	assert.Equal(t, "number", openapiType(Number))
	assert.Equal(t, "number", openapiType("float64"))
	assert.Equal(t, "integer", openapiType("int"))
	assert.Equal(t, "object", openapiType("SomeStruct"))
	assert.Equal(t, "string", openapiType("[]"+String))
}

func TestOpenapiFormat(t *testing.T) {
	assert.Equal(t, "int64", openapiFormat(Int))
	assert.Equal(t, "int64", openapiFormat("int"))
	assert.Equal(t, "double", openapiFormat("float64"))
	assert.Equal(t, "", openapiFormat(String))
	assert.Equal(t, "", openapiFormat("SomeStruct"))
}

func TestItemAndCollectionPaths(t *testing.T) {
	// Device resources always use /stat/ paths.
	dev := NewResource("Device", "device")
	assert.Equal(t, "/api/s/{site}/stat/device", collectionPath(dev))
	assert.Equal(t, "/api/s/{site}/stat/device/{id}", itemPath(dev))

	// v2 resource without an explicit ItemResourcePath falls back to
	// collectionPath + "/{id}".
	fp := NewResource("FirewallPolicy", "firewall-policies")
	assert.Equal(t, "/v2/api/site/{site}/firewall-policies", collectionPath(fp))
	assert.Equal(t, "/v2/api/site/{site}/firewall-policies/{id}", itemPath(fp))
}

func TestItemPathAllBranches(t *testing.T) {
	// v2 + explicit ItemResourcePath.
	v2 := &ResourceInfo{StructName: "FirewallPolicy", ResourcePath: "firewall-policies", ItemResourcePath: "firewall-policy"}
	assert.Equal(t, "/v2/api/site/{site}/firewall-policy/{id}", itemPath(v2))

	// Device + explicit ItemResourcePath.
	dev := &ResourceInfo{StructName: "Device", ResourcePath: "device", ItemResourcePath: "device-item"}
	assert.Equal(t, "/api/s/{site}/stat/device-item/{id}", itemPath(dev))

	// v1 + explicit ItemResourcePath.
	v1 := &ResourceInfo{StructName: "Widget", ResourcePath: "widgets", ItemResourcePath: "widget"}
	assert.Equal(t, "/api/s/{site}/rest/widget/{id}", itemPath(v1))
}

func TestHumanize(t *testing.T) {
	assert.Equal(t, "WLAN Group", humanize("WLANGroup"))
	assert.Equal(t, "Client Group", humanize("ClientGroup"))
	assert.Equal(t, "DNS Record", humanize("DNSRecord"))
}

func TestCleanStructName(t *testing.T) {
	assert.Equal(t, "Mgmt", NewResource("SettingMgmt", "mgmt").CleanStructName())
	assert.Equal(t, "Network", NewResource("Network", "networkconf").CleanStructName())
}

func TestWriteOpenAPIInvalidPath(t *testing.T) {
	err := WriteOpenAPI(sampleResources(), "1.0.0", filepath.Join(t.TempDir(), "missing-dir", "openapi.yaml"))
	assert.Error(t, err)
}
