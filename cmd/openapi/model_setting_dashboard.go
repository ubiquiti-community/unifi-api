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

type SettingDashboard struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	LayoutPreference string                    `json:"layout_preference,omitempty"` // auto|custom
	Widgets          []SettingDashboardWidgets `json:"widgets,omitempty"`
}

func (dst *SettingDashboard) UnmarshalJSON(b []byte) error {
	type Alias SettingDashboard
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

type SettingDashboardWidgets struct {
	Name string `json:"name,omitempty"` // traffic_identification|connection_types|wifi_technology|most_active_clients|most_active_aps|meshing|network_activity|wireless_experience|internet|wifi_activity|wifi_channels|wifi_client_experience|wifi_tx_retries|admin_activity|device_client_count|server_ip
}

func (dst *SettingDashboardWidgets) UnmarshalJSON(b []byte) error {
	type Alias SettingDashboardWidgets
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

type SettingDashboardResponse struct {
	Meta meta               `json:"meta"`
	Data []SettingDashboard `json:"data"`
}

func addSettingDashboard() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/get/setting/dashboard")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingDashboard")
	getOp.SetTags("SettingDashboard")
	getOp.AddRespStructure(new(SettingDashboardResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/set/setting/dashboard")
	updateOp.AddReqStructure(new(SettingDashboard))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingDashboard")
	updateOp.SetTags("SettingDashboard")

	updateOp.AddRespStructure(new(SettingDashboardResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
