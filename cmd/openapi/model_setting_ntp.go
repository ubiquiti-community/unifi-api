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

type SettingNtp struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	NtpServer1        string `json:"ntp_server_1,omitempty"`
	NtpServer2        string `json:"ntp_server_2,omitempty"`
	NtpServer3        string `json:"ntp_server_3,omitempty"`
	NtpServer4        string `json:"ntp_server_4,omitempty"`
	SettingPreference string `json:"setting_preference,omitempty"` // auto|manual
}

func (dst *SettingNtp) UnmarshalJSON(b []byte) error {
	type Alias SettingNtp
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

type SettingNtpGetRequest struct {
	Site string `path:"site"`
}

type SettingNtpUpdateRequest struct {
	*SettingNtp
	Site string `path:"site"`
}

type SettingNtpResponse struct {
	Meta meta         `json:"meta"`
	Data []SettingNtp `json:"data"`
}

func addSettingNtp() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/get/setting/ntp")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingNtp")
	getOp.SetTags("SettingNtp")
	getOp.AddReqStructure(new(SettingNtpGetRequest))
	getOp.AddRespStructure(new(SettingNtpResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/set/setting/ntp")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingNtp")
	updateOp.SetTags("SettingNtp")
	updateOp.AddReqStructure(new(SettingNtpUpdateRequest))
	updateOp.AddRespStructure(new(SettingNtpResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
