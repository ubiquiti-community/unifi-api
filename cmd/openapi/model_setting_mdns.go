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

type SettingMdns struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	CustomServices     []SettingMdnsCustomServices     `json:"custom_services,omitempty"`
	PredefinedServices []SettingMdnsPredefinedServices `json:"predefined_services,omitempty"`
}

func (dst *SettingMdns) UnmarshalJSON(b []byte) error {
	type Alias SettingMdns
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

type SettingMdnsCustomServices struct {
	Address string `json:"address,omitempty"` // ^_[a-zA-Z0-9._-]+\._(tcp|udp)(\.local)?$
	Name    string `json:"name,omitempty"`
}

func (dst *SettingMdnsCustomServices) UnmarshalJSON(b []byte) error {
	type Alias SettingMdnsCustomServices
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

type SettingMdnsPredefinedServices struct {
	Code string `json:"code,omitempty"` // amazon_devices|apple_airDrop|apple_airPlay|apple_file_sharing|apple_iChat|apple_iTunes|dns_service_discovery|ftp_servers|google_chromecast|homeKit|matter_network|printers|scanners|spotify_connect|ssh_servers|web_servers|windows_file_sharing_samba
}

func (dst *SettingMdnsPredefinedServices) UnmarshalJSON(b []byte) error {
	type Alias SettingMdnsPredefinedServices
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

type SettingMdnsResponse struct {
	Meta meta          `json:"meta"`
	Data []SettingMdns `json:"data"`
}

func addSettingMdns() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/get/setting/mdns")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingMdns")
	getOp.SetTags("SettingMdns")
	getOp.AddRespStructure(new(SettingMdnsResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/set/setting/mdns")
	updateOp.AddReqStructure(new(SettingMdns))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingMdns")
	updateOp.SetTags("SettingMdns")

	updateOp.AddRespStructure(new(SettingMdnsResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
