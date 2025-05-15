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

type SettingTrafficFlow struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	EnabledAllowedTraffic        bool `json:"enabled_allowed_traffic"`
	GatewayDNSEnabled            bool `json:"gateway_dns_enabled"`
	UnifiDeviceManagementEnabled bool `json:"unifi_device_management_enabled"`
	UnifiServicesEnabled         bool `json:"unifi_services_enabled"`
}

func (dst *SettingTrafficFlow) UnmarshalJSON(b []byte) error {
	type Alias SettingTrafficFlow
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

type SettingTrafficFlowUpdateRequest struct {
	*SettingTrafficFlow
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type SettingTrafficFlowResponse struct {
	Meta meta                 `json:"meta"`
	Data []SettingTrafficFlow `json:"data"`
}

func addSettingTrafficFlow() {
	resourceName := strcase.SnakeCase("SettingTrafficFlow")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/get/setting/traffic_flow")
	getOp.AddReqStructure(new(SiteRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/get/setting/traffic_flow",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingTrafficFlow")
	getOp.SetTags("SettingTrafficFlow")
	getOp.AddRespStructure(new(SettingTrafficFlowResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/set/setting/traffic_flow")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingTrafficFlow")
	updateOp.SetTags("SettingTrafficFlow")
	updateOp.AddReqStructure(new(SettingTrafficFlowUpdateRequest))

	updateOp.AddRespStructure(new(SettingTrafficFlowResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
