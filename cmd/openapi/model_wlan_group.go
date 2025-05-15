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

type WLANGroup struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Name string `json:"name,omitempty"` // .{1,128}
}

func (dst *WLANGroup) UnmarshalJSON(b []byte) error {
	type Alias WLANGroup
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

type WLANGroupGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type WLANGroupListRequest struct {
	SiteID string `path:"siteId"`
}

type WLANGroupCreateRequest struct {
	*WLANGroup
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type WLANGroupDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type WLANGroupUpdateRequest struct {
	*WLANGroup
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type WLANGroupResponse struct {
	Meta meta        `json:"meta"`
	Data []WLANGroup `json:"data"`
}

func addWLANGroup() {
	resourceName := strcase.SnakeCase("WLANGroup")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/wlangroup/{id}")
	getOp.AddReqStructure(new(WLANGroupGetRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/rest/wlangroup/{id}",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetWLANGroup")
	getOp.SetTags("WLANGroup")
	getOp.AddRespStructure(new(WLANGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/wlangroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateWLANGroup")
	updateOp.SetTags("WLANGroup")
	updateOp.AddReqStructure(new(WLANGroupUpdateRequest))

	updateOp.AddRespStructure(new(WLANGroupResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/wlangroup")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListWLANGroup")
	listOp.SetTags("WLANGroup")
	listOp.AddReqStructure(new(WLANGroupListRequest))

	listOp.AddRespStructure(new(WLANGroupResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/wlangroup")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateWLANGroup")
	createOp.SetTags("WLANGroup")
	createOp.AddReqStructure(new(WLANGroupCreateRequest))

	getOp.AddRespStructure(new(WLANGroupResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/wlangroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteWLANGroup")
	deleteOp.SetTags("WLANGroup")
	deleteOp.AddReqStructure(new(WLANGroupDeleteRequest))

	deleteOp.AddRespStructure(new(WLANGroupResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
