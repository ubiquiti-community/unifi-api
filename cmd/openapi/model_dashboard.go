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

type Dashboard struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	ControllerVersion string             `json:"controller_version,omitempty"`
	Desc              string             `json:"desc,omitempty"`
	IsPublic          bool               `json:"is_public"`
	Modules           []DashboardModules `json:"modules,omitempty"`
	Name              string             `json:"name,omitempty"`
}

func (dst *Dashboard) UnmarshalJSON(b []byte) error {
	type Alias Dashboard
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

type DashboardModules struct {
	Config       string `json:"config,omitempty"`
	ID           string `json:"id"`
	ModuleID     string `json:"module_id"`
	Restrictions string `json:"restrictions,omitempty"`
}

func (dst *DashboardModules) UnmarshalJSON(b []byte) error {
	type Alias DashboardModules
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

type DashboardGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type DashboardListRequest struct {
	SiteID string `path:"siteId"`
}

type DashboardCreateRequest struct {
	*Dashboard
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type DashboardDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type DashboardUpdateRequest struct {
	*Dashboard
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type DashboardResponse struct {
	Meta meta        `json:"meta"`
	Data []Dashboard `json:"data"`
}

func addDashboard() {
	resourceName := strcase.SnakeCase("Dashboard")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/dashboard/{id}")
	getOp.AddReqStructure(new(DashboardGetRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/rest/dashboard/{id}",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetDashboard")
	getOp.SetTags("Dashboard")
	getOp.AddRespStructure(new(DashboardResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/dashboard/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateDashboard")
	updateOp.SetTags("Dashboard")
	updateOp.AddReqStructure(new(DashboardUpdateRequest))

	updateOp.AddRespStructure(new(DashboardResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/dashboard")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListDashboard")
	listOp.SetTags("Dashboard")
	listOp.AddReqStructure(new(DashboardListRequest))

	listOp.AddRespStructure(new(DashboardResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/dashboard")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateDashboard")
	createOp.SetTags("Dashboard")
	createOp.AddReqStructure(new(DashboardCreateRequest))

	getOp.AddRespStructure(new(DashboardResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/dashboard/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteDashboard")
	deleteOp.SetTags("Dashboard")
	deleteOp.AddReqStructure(new(DashboardDeleteRequest))

	deleteOp.AddRespStructure(new(DashboardResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
