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

type DpiGroup struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	DPIappIDs []string `json:"dpiapp_ids,omitempty"` // [\d\w]+
	Enabled   bool     `json:"enabled"`
	Name      string   `json:"name,omitempty"` // .{1,128}
}

func (dst *DpiGroup) UnmarshalJSON(b []byte) error {
	type Alias DpiGroup
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

type DpiGroupGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type DpiGroupListRequest struct {
	SiteID string `path:"siteId"`
}

type DpiGroupCreateRequest struct {
	*DpiGroup
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type DpiGroupDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type DpiGroupUpdateRequest struct {
	*DpiGroup
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type DpiGroupResponse struct {
	Meta meta       `json:"meta"`
	Data []DpiGroup `json:"data"`
}

func addDpiGroup() {
	resourceName := strcase.SnakeCase("DpiGroup")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/dpigroup/{id}")
	getOp.AddReqStructure(new(DpiGroupGetRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/rest/dpigroup/{id}",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetDpiGroup")
	getOp.SetTags("DpiGroup")
	getOp.AddRespStructure(new(DpiGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/dpigroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateDpiGroup")
	updateOp.SetTags("DpiGroup")
	updateOp.AddReqStructure(new(DpiGroupUpdateRequest))

	updateOp.AddRespStructure(new(DpiGroupResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/dpigroup")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListDpiGroup")
	listOp.SetTags("DpiGroup")
	listOp.AddReqStructure(new(DpiGroupListRequest))

	listOp.AddRespStructure(new(DpiGroupResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/dpigroup")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateDpiGroup")
	createOp.SetTags("DpiGroup")
	createOp.AddReqStructure(new(DpiGroupCreateRequest))

	getOp.AddRespStructure(new(DpiGroupResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/dpigroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteDpiGroup")
	deleteOp.SetTags("DpiGroup")
	deleteOp.AddReqStructure(new(DpiGroupDeleteRequest))

	deleteOp.AddRespStructure(new(DpiGroupResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
