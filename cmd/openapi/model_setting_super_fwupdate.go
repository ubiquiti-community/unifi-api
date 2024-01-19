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

type SettingSuperFwupdate struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	ControllerChannel string `json:"controller_channel,omitempty"` // internal|alpha|beta|release-candidate|release
	FirmwareChannel   string `json:"firmware_channel,omitempty"`   // internal|alpha|beta|release-candidate|release
	SsoEnabled        bool   `json:"sso_enabled"`
}

func (dst *SettingSuperFwupdate) UnmarshalJSON(b []byte) error {
	type Alias SettingSuperFwupdate
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

type SettingSuperFwupdateGetRequest struct {
	Site string `path:"site"`
}

type SettingSuperFwupdateUpdateRequest struct {
	*SettingSuperFwupdate
	Site string `path:"site"`
}

type SettingSuperFwupdateResponse struct {
	Meta meta                   `json:"meta"`
	Data []SettingSuperFwupdate `json:"data"`
}

func addSettingSuperFwupdate() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/get/setting/super_fwupdate")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingSuperFwupdate")
	getOp.SetTags("SettingSuperFwupdate")
	getOp.AddReqStructure(new(SettingSuperFwupdateGetRequest))
	getOp.AddRespStructure(new(SettingSuperFwupdateResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/set/setting/super_fwupdate")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingSuperFwupdate")
	updateOp.SetTags("SettingSuperFwupdate")
	updateOp.AddReqStructure(new(SettingSuperFwupdateUpdateRequest))
	updateOp.AddRespStructure(new(SettingSuperFwupdateResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
