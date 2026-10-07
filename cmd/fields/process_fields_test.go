package main

import (
	"encoding/json"
	"testing"
)

// A controller that ships a property the generator also adds by hand (for
// older controllers) must not emit it twice: the hand-added field keeps its
// go-unifi name and the derived duplicate is skipped.
func TestProcessFieldsKeepsHandAddedFieldOnNativeDuplicate(t *testing.T) {
	r := NewResource("Network", "networkconf")

	var fields map[string]any
	err := json.Unmarshal([]byte(`{
		"wireguard_interface_binding_mode_ip_version": "^(v4|v6)$",
		"upnp_nat_pmp_enabled": "true|false",
		"name": ".{1,128}"
	}`), &fields)
	if err != nil {
		t.Fatal(err)
	}
	r.processFields(fields)

	base := r.Types["Network"]
	seen := map[string]string{}
	for key, f := range base.Fields {
		if f == nil {
			continue
		}
		if prev, ok := seen[f.JSONName]; ok {
			t.Errorf("property %q emitted twice: as %s and %s", f.JSONName, prev, key)
		}
		seen[f.JSONName] = key
	}
	for _, want := range []string{"WireguardInterfaceBindingModeIPVersion", "UPnPNatPMPEnabled", "Name"} {
		if _, ok := base.Fields[want]; !ok {
			t.Errorf("field %s missing", want)
		}
	}
}
