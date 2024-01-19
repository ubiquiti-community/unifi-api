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

type SettingBroadcast struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	SoundAfterEnabled   bool   `json:"sound_after_enabled"`
	SoundAfterResource  string `json:"sound_after_resource,omitempty"`
	SoundAfterType      string `json:"sound_after_type,omitempty"` // sample|media
	SoundBeforeEnabled  bool   `json:"sound_before_enabled"`
	SoundBeforeResource string `json:"sound_before_resource,omitempty"`
	SoundBeforeType     string `json:"sound_before_type,omitempty"` // sample|media
}

func (dst *SettingBroadcast) UnmarshalJSON(b []byte) error {
	type Alias SettingBroadcast
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

type SettingBroadcastResponse struct {
	Meta meta               `json:"meta"`
	Data []SettingBroadcast `json:"data"`
}

func addSettingBroadcast() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/get/setting/broadcast")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingBroadcast")
	getOp.SetTags("SettingBroadcast")
	getOp.AddRespStructure(new(SettingBroadcastResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/set/setting/broadcast")
	updateOp.AddReqStructure(new(SettingBroadcast))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingBroadcast")
	updateOp.SetTags("SettingBroadcast")
	updateOp.AddRespStructure(new(SettingBroadcastResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
