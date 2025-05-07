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

type SettingLcm struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Brightness  int  `json:"brightness,omitempty"` // [1-9]|[1-9][0-9]|100
	Enabled     bool `json:"enabled"`
	IDleTimeout int  `json:"idle_timeout,omitempty"` // [1-9][0-9]|[1-9][0-9][0-9]|[1-2][0-9][0-9][0-9]|3[0-5][0-9][0-9]|3600
	Sync        bool `json:"sync"`
	TouchEvent  bool `json:"touch_event"`
}

func (dst *SettingLcm) UnmarshalJSON(b []byte) error {
	type Alias SettingLcm
	aux := &struct {
		Brightness  emptyStringInt `json:"brightness"`
		IDleTimeout emptyStringInt `json:"idle_timeout"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Brightness = int(aux.Brightness)
	dst.IDleTimeout = int(aux.IDleTimeout)

	return nil
}

type SettingLcmUpdateRequest struct {
	*SettingLcm
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type SettingLcmResponse struct {
	Meta meta         `json:"meta"`
	Data []SettingLcm `json:"data"`
}

func addSettingLcm() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/get/setting/lcm")
	getOp.AddReqStructure(new(SiteRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingLcm")
	getOp.SetTags("SettingLcm")
	getOp.AddRespStructure(new(SettingLcmResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/set/setting/lcm")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingLcm")
	updateOp.SetTags("SettingLcm")
	updateOp.AddReqStructure(new(SettingLcmUpdateRequest))

	updateOp.AddRespStructure(new(SettingLcmResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
