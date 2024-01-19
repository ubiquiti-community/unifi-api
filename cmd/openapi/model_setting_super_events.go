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

type SettingSuperEvents struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Ignored string `json:"_ignored,omitempty"`
}

func (dst *SettingSuperEvents) UnmarshalJSON(b []byte) error {
	type Alias SettingSuperEvents
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

type SettingSuperEventsGetRequest struct {
	Site string `path:"site"`
}

type SettingSuperEventsUpdateRequest struct {
	*SettingSuperEvents
	Site string `path:"site"`
}

type SettingSuperEventsResponse struct {
	Meta meta                 `json:"meta"`
	Data []SettingSuperEvents `json:"data"`
}

func addSettingSuperEvents() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/get/setting/super_events")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingSuperEvents")
	getOp.SetTags("SettingSuperEvents")
	getOp.AddReqStructure(new(SettingSuperEventsGetRequest))
	getOp.AddRespStructure(new(SettingSuperEventsResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/set/setting/super_events")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingSuperEvents")
	updateOp.SetTags("SettingSuperEvents")
	updateOp.AddReqStructure(new(SettingSuperEventsUpdateRequest))
	updateOp.AddRespStructure(new(SettingSuperEventsResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
