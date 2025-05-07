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

type SettingElementAdopt struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Enabled       bool   `json:"enabled"`
	XElementEssid string `json:"x_element_essid,omitempty"`
	XElementPsk   string `json:"x_element_psk,omitempty"`
}

func (dst *SettingElementAdopt) UnmarshalJSON(b []byte) error {
	type Alias SettingElementAdopt
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

type SettingElementAdoptUpdateRequest struct {
	*SettingElementAdopt
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type SettingElementAdoptResponse struct {
	Meta meta                  `json:"meta"`
	Data []SettingElementAdopt `json:"data"`
}

func addSettingElementAdopt() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/get/setting/element_adopt")
	getOp.AddReqStructure(new(SiteRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingElementAdopt")
	getOp.SetTags("SettingElementAdopt")
	getOp.AddRespStructure(new(SettingElementAdoptResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/set/setting/element_adopt")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingElementAdopt")
	updateOp.SetTags("SettingElementAdopt")
	updateOp.AddReqStructure(new(SettingElementAdoptUpdateRequest))

	updateOp.AddRespStructure(new(SettingElementAdoptResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
