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

type SettingUsw struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	DHCPSnoop bool `json:"dhcp_snoop"`
}

func (dst *SettingUsw) UnmarshalJSON(b []byte) error {
	type Alias SettingUsw
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

type SettingUswResponse struct {
	Meta meta         `json:"meta"`
	Data []SettingUsw `json:"data"`
}

func addSettingUsw() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/get/setting/usw")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingUsw")
	getOp.SetTags("SettingUsw")
	getOp.AddRespStructure(new(SettingUswResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/set/setting/usw")
	updateOp.AddReqStructure(new(SettingUsw))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingUsw")
	updateOp.SetTags("SettingUsw")

	updateOp.AddRespStructure(new(SettingUswResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
