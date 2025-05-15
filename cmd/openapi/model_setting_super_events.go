// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/stoewer/go-strcase"
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

type SettingSuperEventsUpdateRequest struct {
	*SettingSuperEvents
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type SettingSuperEventsResponse struct {
	Meta meta                 `json:"meta"`
	Data []SettingSuperEvents `json:"data"`
}

func addSettingSuperEvents() {
	resourceName := strcase.SnakeCase("SettingSuperEvents")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/get/setting/super_events")
	getOp.AddReqStructure(new(SiteRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/get/setting/super_events",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingSuperEvents")
	getOp.SetTags("SettingSuperEvents")
	getOp.AddRespStructure(new(SettingSuperEventsResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/set/setting/super_events")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingSuperEvents")
	updateOp.SetTags("SettingSuperEvents")
	updateOp.AddReqStructure(new(SettingSuperEventsUpdateRequest))

	updateOp.AddRespStructure(new(SettingSuperEventsResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
