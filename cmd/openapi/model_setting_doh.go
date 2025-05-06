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

	CustomServers []SettingDohCustomServers `json:"custom_servers,omitempty"`
	ServerNames   []string                  `json:"server_names,omitempty"`
	State         string                    `json:"state,omitempty"` // off|auto|manual|custom
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

type SettingDohCustomServers struct {
	Enabled    bool   `json:"enabled"`
	SdnsStamp  string `json:"sdns_stamp,omitempty"`
	ServerName string `json:"server_name,omitempty"`
}

func (dst *SettingDohCustomServers) UnmarshalJSON(b []byte) error {
	type Alias SettingDohCustomServers
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

type SettingDohResponse struct {
	Meta meta         `json:"meta"`
	Data []SettingDoh `json:"data"`
}

func addSettingDoh() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/get/setting/doh")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingDoh")
	getOp.SetTags("SettingDoh")
	getOp.AddRespStructure(new(SettingDohResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/set/setting/doh")
	updateOp.AddReqStructure(new(SettingDoh))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingDoh")
	updateOp.SetTags("SettingDoh")

	updateOp.AddRespStructure(new(SettingDohResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
