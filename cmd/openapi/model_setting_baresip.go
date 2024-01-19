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

type SettingBaresip struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Enabled       bool   `json:"enabled"`
	OutboundProxy string `json:"outbound_proxy,omitempty"`
	PackageUrl    string `json:"package_url,omitempty"`
	Server        string `json:"server,omitempty"`
}

func (dst *SettingBaresip) UnmarshalJSON(b []byte) error {
	type Alias SettingBaresip
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

type SettingBaresipResponse struct {
	Meta meta             `json:"meta"`
	Data []SettingBaresip `json:"data"`
}

func addSettingBaresip() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/get/setting/baresip")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingBaresip")
	getOp.SetTags("SettingBaresip")
	getOp.AddRespStructure(new(SettingBaresipResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/set/setting/baresip")
	updateOp.AddReqStructure(new(SettingBaresip))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingBaresip")
	updateOp.SetTags("SettingBaresip")
	updateOp.AddRespStructure(new(SettingBaresipResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
