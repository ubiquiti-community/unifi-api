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

type SettingTeleport struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Enabled    bool   `json:"enabled"`
	SubnetCidr string `json:"subnet_cidr"` // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\/([8-9]|[1-2][0-9]|3[0-2])$|^$
}

func (dst *SettingTeleport) UnmarshalJSON(b []byte) error {
	type Alias SettingTeleport
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

type SettingTeleportResponse struct {
	Meta meta              `json:"meta"`
	Data []SettingTeleport `json:"data"`
}

func addSettingTeleport() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/get/setting/teleport")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingTeleport")
	getOp.SetTags("SettingTeleport")
	getOp.AddRespStructure(new(SettingTeleportResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/set/setting/teleport")
	updateOp.AddReqStructure(new(SettingTeleport))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingTeleport")
	updateOp.SetTags("SettingTeleport")
	updateOp.AddRespStructure(new(SettingTeleportResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
