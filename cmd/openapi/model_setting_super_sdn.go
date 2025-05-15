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

type SettingSuperSdn struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	AuthToken       string `json:"auth_token,omitempty"`
	DeviceID        string `json:"device_id"`
	Enabled         bool   `json:"enabled"`
	Migrated        bool   `json:"migrated"`
	SsoLoginEnabled string `json:"sso_login_enabled,omitempty"`
	UbicUuid        string `json:"ubic_uuid,omitempty"`
}

func (dst *SettingSuperSdn) UnmarshalJSON(b []byte) error {
	type Alias SettingSuperSdn
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

type SettingSuperSdnUpdateRequest struct {
	*SettingSuperSdn
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type SettingSuperSdnResponse struct {
	Meta meta              `json:"meta"`
	Data []SettingSuperSdn `json:"data"`
}

func addSettingSuperSdn() {
	resourceName := strcase.SnakeCase("SettingSuperSdn")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/get/setting/super_sdn")
	getOp.AddReqStructure(new(SiteRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/get/setting/super_sdn",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingSuperSdn")
	getOp.SetTags("SettingSuperSdn")
	getOp.AddRespStructure(new(SettingSuperSdnResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/set/setting/super_sdn")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingSuperSdn")
	updateOp.SetTags("SettingSuperSdn")
	updateOp.AddReqStructure(new(SettingSuperSdnUpdateRequest))

	updateOp.AddRespStructure(new(SettingSuperSdnResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
