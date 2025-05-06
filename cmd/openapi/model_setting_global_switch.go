// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/swaggest/openapi-go"
)

// just to fix compile issues with the import
var (
	_ context.Context
	_ fmt.Formatter
	_ json.Marshaler
)

type SettingGlobalSwitch struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	AclDeviceIsolation             []string                            `json:"acl_device_isolation,omitempty"`
	AclL3Isolation                 []SettingGlobalSwitchAclL3Isolation `json:"acl_l3_isolation,omitempty"`
	DHCPSnoop                      bool                                `json:"dhcp_snoop"`
	Dot1XFallbackNetworkID         string                              `json:"dot1x_fallback_networkconf_id"` // [\d\w]+|
	Dot1XPortctrlEnabled           bool                                `json:"dot1x_portctrl_enabled"`
	FloodKnownProtocols            bool                                `json:"flood_known_protocols"`
	FlowctrlEnabled                bool                                `json:"flowctrl_enabled"`
	ForwardUnknownMcastRouterPorts bool                                `json:"forward_unknown_mcast_router_ports"`
	JumboframeEnabled              bool                                `json:"jumboframe_enabled"`
	RADIUSProfileID                string                              `json:"radiusprofile_id"`
	StpVersion                     string                              `json:"stp_version,omitempty"`       // stp|rstp|disabled
	SwitchExclusions               []string                            `json:"switch_exclusions,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
}

func (dst *SettingGlobalSwitch) UnmarshalJSON(b []byte) error {
	type Alias SettingGlobalSwitch
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}

	return nil
}

type SettingGlobalSwitchAclL3Isolation struct {
	DestinationNetworks []string `json:"destination_networks,omitempty"`
	SourceNetwork       string   `json:"source_network,omitempty"`
}

func (dst *SettingGlobalSwitchAclL3Isolation) UnmarshalJSON(b []byte) error {
	type Alias SettingGlobalSwitchAclL3Isolation
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}

	return nil
}

type SettingGlobalSwitchResponse struct {
	Meta meta                  `json:"meta"`
	Data []SettingGlobalSwitch `json:"data"`
}

func addSettingGlobalSwitch() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/get/setting/global_switch")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingGlobalSwitch")
	getOp.SetTags("SettingGlobalSwitch")
	getOp.AddRespStructure(new(SettingGlobalSwitchResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/set/setting/global_switch")
	updateOp.AddReqStructure(new(SettingGlobalSwitch))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingGlobalSwitch")
	updateOp.SetTags("SettingGlobalSwitch")

	updateOp.AddRespStructure(new(SettingGlobalSwitchResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
