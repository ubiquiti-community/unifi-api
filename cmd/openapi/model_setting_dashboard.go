// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/stoewer/go-strcase"
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

	LayoutPreference string                    `json:"layout_preference,omitempty"` // auto|manual
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
	Enabled bool   `json:"enabled"`
	Name    string `json:"name,omitempty"` // critical_traffic_prioritization|cybersecure|traffic_identification|wifi_technology|wifi_channels|wifi_client_experience|wifi_tx_retries|most_active_apps_aps_clients|most_active_apps_clients|most_active_aps_clients|most_active_apps_aps|most_active_apps|v2_most_active_aps|v2_most_active_clients|wifi_connectivity|ap_radio_density|wifi_channel_preset_configuration
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

type SettingDashboardUpdateRequest struct {
	*SettingDashboard
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type SettingDashboardResponse struct {
	Meta meta               `json:"meta"`
	Data []SettingDashboard `json:"data"`
}

func addSettingDashboard() {
	resourceName := strcase.SnakeCase("SettingDashboard")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/get/setting/dashboard")
	getOp.AddReqStructure(new(SiteRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/get/setting/dashboard",
			"method": "GET",
		},
	}
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

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/set/setting/dashboard")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingDashboard")
	updateOp.SetTags("SettingDashboard")
	updateOp.AddReqStructure(new(SettingDashboardUpdateRequest))

	updateOp.AddRespStructure(new(SettingDashboardResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
