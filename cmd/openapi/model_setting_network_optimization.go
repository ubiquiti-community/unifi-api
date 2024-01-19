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

type SettingNetworkOptimization struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Enabled bool `json:"enabled"`
}

func (dst *SettingNetworkOptimization) UnmarshalJSON(b []byte) error {
	type Alias SettingNetworkOptimization
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

type SettingNetworkOptimizationGetRequest struct {
	Site string `path:"site"`
}

type SettingNetworkOptimizationUpdateRequest struct {
	*SettingNetworkOptimization
	Site string `path:"site"`
}

type SettingNetworkOptimizationResponse struct {
	Meta meta                         `json:"meta"`
	Data []SettingNetworkOptimization `json:"data"`
}

func addSettingNetworkOptimization() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/get/setting/network_optimization")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingNetworkOptimization")
	getOp.SetTags("SettingNetworkOptimization")
	getOp.AddReqStructure(new(SettingNetworkOptimizationGetRequest))
	getOp.AddRespStructure(new(SettingNetworkOptimizationResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/set/setting/network_optimization")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingNetworkOptimization")
	updateOp.SetTags("SettingNetworkOptimization")
	updateOp.AddReqStructure(new(SettingNetworkOptimizationUpdateRequest))
	updateOp.AddRespStructure(new(SettingNetworkOptimizationResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
