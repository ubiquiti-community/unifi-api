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

type SettingMagicSiteToSiteVpn struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Enabled bool `json:"enabled"`
}

func (dst *SettingMagicSiteToSiteVpn) UnmarshalJSON(b []byte) error {
	type Alias SettingMagicSiteToSiteVpn
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

type SettingMagicSiteToSiteVpnResponse struct {
	Meta meta                        `json:"meta"`
	Data []SettingMagicSiteToSiteVpn `json:"data"`
}

func addSettingMagicSiteToSiteVpn() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/get/setting/magic_site_to_site_vpn")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingMagicSiteToSiteVpn")
	getOp.SetTags("SettingMagicSiteToSiteVpn")
	getOp.AddRespStructure(new(SettingMagicSiteToSiteVpnResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/set/setting/magic_site_to_site_vpn")
	updateOp.AddReqStructure(new(SettingMagicSiteToSiteVpn))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingMagicSiteToSiteVpn")
	updateOp.SetTags("SettingMagicSiteToSiteVpn")
	updateOp.AddRespStructure(new(SettingMagicSiteToSiteVpnResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
