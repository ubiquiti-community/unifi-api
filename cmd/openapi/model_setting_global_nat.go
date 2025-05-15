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

type SettingGlobalNat struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	ExcludedNetworkIDs []string `json:"excluded_network_ids,omitempty"`
	Mode               string   `json:"mode,omitempty"` // auto|custom|off
}

func (dst *SettingGlobalNat) UnmarshalJSON(b []byte) error {
	type Alias SettingGlobalNat
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

type SettingGlobalNatUpdateRequest struct {
	*SettingGlobalNat
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type SettingGlobalNatResponse struct {
	Meta meta               `json:"meta"`
	Data []SettingGlobalNat `json:"data"`
}

func addSettingGlobalNat() {
	resourceName := strcase.SnakeCase("SettingGlobalNat")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/get/setting/global_nat")
	getOp.AddReqStructure(new(SiteRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/get/setting/global_nat",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingGlobalNat")
	getOp.SetTags("SettingGlobalNat")
	getOp.AddRespStructure(new(SettingGlobalNatResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/set/setting/global_nat")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingGlobalNat")
	updateOp.SetTags("SettingGlobalNat")
	updateOp.AddReqStructure(new(SettingGlobalNatUpdateRequest))

	updateOp.AddRespStructure(new(SettingGlobalNatResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
