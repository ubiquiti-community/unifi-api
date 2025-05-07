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

type SettingSnmp struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Community string `json:"community,omitempty"` // .{1,256}
	Enabled   bool   `json:"enabled"`
	EnabledV3 bool   `json:"enabledV3"`
	Username  string `json:"username,omitempty"`   // [a-zA-Z0-9_-]{1,30}
	XPassword string `json:"x_password,omitempty"` // [^'"]{8,32}
}

func (dst *SettingSnmp) UnmarshalJSON(b []byte) error {
	type Alias SettingSnmp
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

type SettingSnmpUpdateRequest struct {
	*SettingSnmp
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type SettingSnmpResponse struct {
	Meta meta          `json:"meta"`
	Data []SettingSnmp `json:"data"`
}

func addSettingSnmp() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/get/setting/snmp")
	getOp.AddReqStructure(new(SiteRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingSnmp")
	getOp.SetTags("SettingSnmp")
	getOp.AddRespStructure(new(SettingSnmpResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/set/setting/snmp")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingSnmp")
	updateOp.SetTags("SettingSnmp")
	updateOp.AddReqStructure(new(SettingSnmpUpdateRequest))

	updateOp.AddRespStructure(new(SettingSnmpResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
