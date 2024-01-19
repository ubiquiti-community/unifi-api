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

type PortProfile struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Autoneg                      bool     `json:"autoneg"`
	Dot1XCtrl                    string   `json:"dot1x_ctrl,omitempty"`             // auto|force_authorized|force_unauthorized|mac_based|multi_host
	Dot1XIDleTimeout             int      `json:"dot1x_idle_timeout,omitempty"`     // [0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5]
	EgressRateLimitKbps          int      `json:"egress_rate_limit_kbps,omitempty"` // 6[4-9]|[7-9][0-9]|[1-9][0-9]{2,6}
	EgressRateLimitKbpsEnabled   bool     `json:"egress_rate_limit_kbps_enabled"`
	ExcludedNetworkIDs           []string `json:"excluded_networkconf_ids,omitempty"`
	Forward                      string   `json:"forward,omitempty"` // all|native|customize|disabled
	FullDuplex                   bool     `json:"full_duplex"`
	Isolation                    bool     `json:"isolation"`
	LldpmedEnabled               bool     `json:"lldpmed_enabled"`
	LldpmedNotifyEnabled         bool     `json:"lldpmed_notify_enabled"`
	NATiveNetworkID              string   `json:"native_networkconf_id"`
	Name                         string   `json:"name,omitempty"`
	OpMode                       string   `json:"op_mode,omitempty"`  // switch
	PoeMode                      string   `json:"poe_mode,omitempty"` // auto|off
	PortKeepaliveEnabled         bool     `json:"port_keepalive_enabled"`
	PortSecurityEnabled          bool     `json:"port_security_enabled"`
	PortSecurityMACAddress       []string `json:"port_security_mac_address,omitempty"` // ^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$
	PriorityQueue1Level          int      `json:"priority_queue1_level,omitempty"`     // [0-9]|[1-9][0-9]|100
	PriorityQueue2Level          int      `json:"priority_queue2_level,omitempty"`     // [0-9]|[1-9][0-9]|100
	PriorityQueue3Level          int      `json:"priority_queue3_level,omitempty"`     // [0-9]|[1-9][0-9]|100
	PriorityQueue4Level          int      `json:"priority_queue4_level,omitempty"`     // [0-9]|[1-9][0-9]|100
	SettingPreference            string   `json:"setting_preference,omitempty"`        // auto|manual
	Speed                        int      `json:"speed,omitempty"`                     // 10|100|1000|2500|5000|10000|20000|25000|40000|50000|100000
	StormctrlBroadcastastEnabled bool     `json:"stormctrl_bcast_enabled"`
	StormctrlBroadcastastLevel   int      `json:"stormctrl_bcast_level,omitempty"` // [0-9]|[1-9][0-9]|100
	StormctrlBroadcastastRate    int      `json:"stormctrl_bcast_rate,omitempty"`  // [0-9]|[1-9][0-9]{1,6}|1[0-3][0-9]{6}|14[0-7][0-9]{5}|148[0-7][0-9]{4}|14880000
	StormctrlMcastEnabled        bool     `json:"stormctrl_mcast_enabled"`
	StormctrlMcastLevel          int      `json:"stormctrl_mcast_level,omitempty"` // [0-9]|[1-9][0-9]|100
	StormctrlMcastRate           int      `json:"stormctrl_mcast_rate,omitempty"`  // [0-9]|[1-9][0-9]{1,6}|1[0-3][0-9]{6}|14[0-7][0-9]{5}|148[0-7][0-9]{4}|14880000
	StormctrlType                string   `json:"stormctrl_type,omitempty"`        // level|rate
	StormctrlUcastEnabled        bool     `json:"stormctrl_ucast_enabled"`
	StormctrlUcastLevel          int      `json:"stormctrl_ucast_level,omitempty"` // [0-9]|[1-9][0-9]|100
	StormctrlUcastRate           int      `json:"stormctrl_ucast_rate,omitempty"`  // [0-9]|[1-9][0-9]{1,6}|1[0-3][0-9]{6}|14[0-7][0-9]{5}|148[0-7][0-9]{4}|14880000
	StpPortMode                  bool     `json:"stp_port_mode"`
	TaggedVLANMgmt               string   `json:"tagged_vlan_mgmt,omitempty"` // auto|block_all|custom
	VoiceNetworkID               string   `json:"voice_networkconf_id"`
}

func (dst *PortProfile) UnmarshalJSON(b []byte) error {
	type Alias PortProfile
	aux := &struct {
		Dot1XIDleTimeout           emptyStringInt `json:"dot1x_idle_timeout"`
		EgressRateLimitKbps        emptyStringInt `json:"egress_rate_limit_kbps"`
		PriorityQueue1Level        emptyStringInt `json:"priority_queue1_level"`
		PriorityQueue2Level        emptyStringInt `json:"priority_queue2_level"`
		PriorityQueue3Level        emptyStringInt `json:"priority_queue3_level"`
		PriorityQueue4Level        emptyStringInt `json:"priority_queue4_level"`
		Speed                      emptyStringInt `json:"speed"`
		StormctrlBroadcastastLevel emptyStringInt `json:"stormctrl_bcast_level"`
		StormctrlBroadcastastRate  emptyStringInt `json:"stormctrl_bcast_rate"`
		StormctrlMcastLevel        emptyStringInt `json:"stormctrl_mcast_level"`
		StormctrlMcastRate         emptyStringInt `json:"stormctrl_mcast_rate"`
		StormctrlUcastLevel        emptyStringInt `json:"stormctrl_ucast_level"`
		StormctrlUcastRate         emptyStringInt `json:"stormctrl_ucast_rate"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Dot1XIDleTimeout = int(aux.Dot1XIDleTimeout)
	dst.EgressRateLimitKbps = int(aux.EgressRateLimitKbps)
	dst.PriorityQueue1Level = int(aux.PriorityQueue1Level)
	dst.PriorityQueue2Level = int(aux.PriorityQueue2Level)
	dst.PriorityQueue3Level = int(aux.PriorityQueue3Level)
	dst.PriorityQueue4Level = int(aux.PriorityQueue4Level)
	dst.Speed = int(aux.Speed)
	dst.StormctrlBroadcastastLevel = int(aux.StormctrlBroadcastastLevel)
	dst.StormctrlBroadcastastRate = int(aux.StormctrlBroadcastastRate)
	dst.StormctrlMcastLevel = int(aux.StormctrlMcastLevel)
	dst.StormctrlMcastRate = int(aux.StormctrlMcastRate)
	dst.StormctrlUcastLevel = int(aux.StormctrlUcastLevel)
	dst.StormctrlUcastRate = int(aux.StormctrlUcastRate)

	return nil
}

type PortProfileGetRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type PortProfileDeleteRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type PortProfileUpdateRequest struct {
	*PortProfile
	Site string `path:"site"`
	ID   string `path:"id",json:"_id,omitempty"`
}

type PortProfileListRequest struct {
	Site string `path:"site"`
}

type PortProfileCreateRequest struct {
	*PortProfile
	Site string `path:"site"`
}

type PortProfileResponse struct {
	Meta meta          `json:"meta"`
	Data []PortProfile `json:"data"`
}

func addPortProfile() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/portconf/{id}")
	if err != nil {
		log.Fatal(err)
	}
	getOp.AddReqStructure(new(PortProfileGetRequest))
	getOp.AddRespStructure(new(PortProfileResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/rest/portconf/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.AddReqStructure(new(PortProfileUpdateRequest))
	updateOp.AddRespStructure(new(PortProfileResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/portconf")
	if err != nil {
		log.Fatal(err)
	}
	listOp.AddReqStructure(new(PortProfileListRequest))
	listOp.AddRespStructure(new(PortProfileResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{site}/rest/portconf")
	if err != nil {
		log.Fatal(err)
	}
	createOp.AddReqStructure(new(PortProfileCreateRequest))
	createOp.AddRespStructure(new(PortProfileResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{site}/get/setting/portconf/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.AddReqStructure(new(PortProfileDeleteRequest))
	deleteOp.AddRespStructure(new(PortProfileResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
