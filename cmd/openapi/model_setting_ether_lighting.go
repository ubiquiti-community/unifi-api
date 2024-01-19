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

type SettingEtherLighting struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	NetworkOverrides []SettingEtherLightingNetworkOverrides `json:"network_overrides,omitempty"`
	SpeedOverrides   []SettingEtherLightingSpeedOverrides   `json:"speed_overrides,omitempty"`
}

func (dst *SettingEtherLighting) UnmarshalJSON(b []byte) error {
	type Alias SettingEtherLighting
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

type SettingEtherLightingNetworkOverrides struct {
	Key         string `json:"key,omitempty"`
	RawColorHex string `json:"raw_color_hex,omitempty"` // [0-9A-Fa-f]{6}
}

func (dst *SettingEtherLightingNetworkOverrides) UnmarshalJSON(b []byte) error {
	type Alias SettingEtherLightingNetworkOverrides
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

type SettingEtherLightingSpeedOverrides struct {
	Key         string `json:"key,omitempty"`           // FE|GbE|2.5GbE|5GbE|10GbE|25GbE|40GbE|100GbE
	RawColorHex string `json:"raw_color_hex,omitempty"` // [0-9A-Fa-f]{6}
}

func (dst *SettingEtherLightingSpeedOverrides) UnmarshalJSON(b []byte) error {
	type Alias SettingEtherLightingSpeedOverrides
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

type SettingEtherLightingGetRequest struct {
	Site string `path:"site"`
}

type SettingEtherLightingUpdateRequest struct {
	*SettingEtherLighting
	Site string `path:"site"`
}

type SettingEtherLightingResponse struct {
	Meta meta                   `json:"meta"`
	Data []SettingEtherLighting `json:"data"`
}

func addSettingEtherLighting() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/get/setting/ether_lighting")
	if err != nil {
		log.Fatal(err)
	}
	getOp.AddReqStructure(new(SettingEtherLightingGetRequest))
	getOp.AddRespStructure(new(SettingEtherLightingResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/set/setting/ether_lighting")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.AddReqStructure(new(SettingEtherLightingUpdateRequest))
	updateOp.AddRespStructure(new(SettingEtherLightingResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
