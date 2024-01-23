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

type SettingConnectivity struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	EnableIsolatedWLAN bool   `json:"enable_isolated_wlan"`
	Enabled            bool   `json:"enabled"`
	UplinkHost         string `json:"uplink_host,omitempty"`
	UplinkType         string `json:"uplink_type,omitempty"`
	XMeshEssid         string `json:"x_mesh_essid,omitempty"`
	XMeshPsk           string `json:"x_mesh_psk,omitempty"`
}

func (dst *SettingConnectivity) UnmarshalJSON(b []byte) error {
	type Alias SettingConnectivity
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

type SettingConnectivityResponse struct {
	Meta meta                  `json:"meta"`
	Data []SettingConnectivity `json:"data"`
}

func addSettingConnectivity() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/get/setting/connectivity")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingConnectivity")
	getOp.SetTags("SettingConnectivity")
	getOp.AddRespStructure(new(SettingConnectivityResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/set/setting/connectivity")
	updateOp.AddReqStructure(new(SettingConnectivity))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingConnectivity")
	updateOp.SetTags("SettingConnectivity")

	updateOp.AddRespStructure(new(SettingConnectivityResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
