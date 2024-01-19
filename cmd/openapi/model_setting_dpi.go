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

type SettingDpi struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Enabled               bool `json:"enabled"`
	FingerprintingEnabled bool `json:"fingerprintingEnabled"`
}

func (dst *SettingDpi) UnmarshalJSON(b []byte) error {
	type Alias SettingDpi
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

type SettingDpiGetRequest struct {
	Site string `path:"site"`
}

type SettingDpiUpdateRequest struct {
	*SettingDpi
	Site string `path:"site"`
}

type SettingDpiResponse struct {
	Meta meta         `json:"meta"`
	Data []SettingDpi `json:"data"`
}

func addSettingDpi() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/get/setting/dpi")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingDpi")
	getOp.SetTags("SettingDpi")
	getOp.AddReqStructure(new(SettingDpiGetRequest))
	getOp.AddRespStructure(new(SettingDpiResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/set/setting/dpi")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingDpi")
	updateOp.SetTags("SettingDpi")
	updateOp.AddReqStructure(new(SettingDpiUpdateRequest))
	updateOp.AddRespStructure(new(SettingDpiResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
