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

type DpiApp struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Apps           []int  `json:"apps,omitempty"`
	Blocked        bool   `json:"blocked"`
	Cats           []int  `json:"cats,omitempty"`
	Enabled        bool   `json:"enabled"`
	Log            bool   `json:"log"`
	Name           string `json:"name,omitempty"`              // .{1,128}
	QOSRateMaxDown int    `json:"qos_rate_max_down,omitempty"` // -1|[2-9]|[1-9][0-9]{1,4}|100000|10[0-1][0-9]{3}|102[0-3][0-9]{2}|102400
	QOSRateMaxUp   int    `json:"qos_rate_max_up,omitempty"`   // -1|[2-9]|[1-9][0-9]{1,4}|100000|10[0-1][0-9]{3}|102[0-3][0-9]{2}|102400
}

func (dst *DpiApp) UnmarshalJSON(b []byte) error {
	type Alias DpiApp
	aux := &struct {
		Apps           []emptyStringInt `json:"apps"`
		Cats           []emptyStringInt `json:"cats"`
		QOSRateMaxDown emptyStringInt   `json:"qos_rate_max_down"`
		QOSRateMaxUp   emptyStringInt   `json:"qos_rate_max_up"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Apps = make([]int, len(aux.Apps))
	for i, v := range aux.Apps {
		dst.Apps[i] = int(v)
	}
	dst.Cats = make([]int, len(aux.Cats))
	for i, v := range aux.Cats {
		dst.Cats[i] = int(v)
	}
	dst.QOSRateMaxDown = int(aux.QOSRateMaxDown)
	dst.QOSRateMaxUp = int(aux.QOSRateMaxUp)

	return nil
}

type DpiAppGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type DpiAppListRequest struct {
	SiteID string `path:"siteId"`
}

type DpiAppCreateRequest struct {
	*DpiApp
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type DpiAppDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type DpiAppUpdateRequest struct {
	*DpiApp
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type DpiAppResponse struct {
	Meta meta     `json:"meta"`
	Data []DpiApp `json:"data"`
}

func addDpiApp() {
	resourceName := strcase.SnakeCase("DpiApp")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/dpiapp/{id}")
	getOp.AddReqStructure(new(DpiAppGetRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/rest/dpiapp/{id}",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetDpiApp")
	getOp.SetTags("DpiApp")
	getOp.AddRespStructure(new(DpiAppResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/dpiapp/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateDpiApp")
	updateOp.SetTags("DpiApp")
	updateOp.AddReqStructure(new(DpiAppUpdateRequest))

	updateOp.AddRespStructure(new(DpiAppResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/dpiapp")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListDpiApp")
	listOp.SetTags("DpiApp")
	listOp.AddReqStructure(new(DpiAppListRequest))

	listOp.AddRespStructure(new(DpiAppResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/dpiapp")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateDpiApp")
	createOp.SetTags("DpiApp")
	createOp.AddReqStructure(new(DpiAppCreateRequest))

	getOp.AddRespStructure(new(DpiAppResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/dpiapp/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteDpiApp")
	deleteOp.SetTags("DpiApp")
	deleteOp.AddReqStructure(new(DpiAppDeleteRequest))

	deleteOp.AddRespStructure(new(DpiAppResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
