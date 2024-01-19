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

type SettingDoh struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	ServerNames []string `json:"server_names,omitempty"`
	State       string   `json:"state,omitempty"` // off|auto|manual
}

func (dst *SettingDoh) UnmarshalJSON(b []byte) error {
	type Alias SettingDoh
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

type SettingDohGetRequest struct {
	Site string `path:"site"`
}

type SettingDohUpdateRequest struct {
	*SettingDoh
	Site string `path:"site"`
}

type SettingDohResponse struct {
	Meta meta         `json:"meta"`
	Data []SettingDoh `json:"data"`
}

func addSettingDoh() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/get/setting/doh")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingDoh")
	getOp.SetTags("SettingDoh")
	getOp.AddReqStructure(new(SettingDohGetRequest))
	getOp.AddRespStructure(new(SettingDohResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/set/setting/doh")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingDoh")
	updateOp.SetTags("SettingDoh")
	updateOp.AddReqStructure(new(SettingDohUpdateRequest))
	updateOp.AddRespStructure(new(SettingDohResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
