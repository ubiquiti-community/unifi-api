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

type Map struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Lat        string  `json:"lat,omitempty"` // ^([-]?[\d]+[.]?[\d]*([eE][-+]?[\d]+)?)$
	Lng        string  `json:"lng,omitempty"` // ^([-]?[\d]+[.]?[\d]*([eE][-+]?[\d]+)?)$
	MapTypeID  string  `json:"mapTypeId"`     // satellite|roadmap|hybrid|terrain
	Name       string  `json:"name,omitempty"`
	OffsetLeft float64 `json:"offset_left,omitempty"`
	OffsetTop  float64 `json:"offset_top,omitempty"`
	Opacity    float64 `json:"opacity,omitempty"` // ^(0(\.[\d]{1,2})?|1)$|^$
	Selected   bool    `json:"selected"`
	Tilt       int     `json:"tilt,omitempty"`
	Type       string  `json:"type,omitempty"` // designerMap|imageMap|googleMap
	Unit       string  `json:"unit,omitempty"` // m|f
	Upp        float64 `json:"upp,omitempty"`
	Zoom       int     `json:"zoom,omitempty"`
}

func (dst *Map) UnmarshalJSON(b []byte) error {
	type Alias Map
	aux := &struct {
		Tilt emptyStringInt `json:"tilt"`
		Zoom emptyStringInt `json:"zoom"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Tilt = int(aux.Tilt)
	dst.Zoom = int(aux.Zoom)

	return nil
}

type MapGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type MapListRequest struct {
	SiteID string `path:"siteId"`
}

type MapCreateRequest struct {
	*Map
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type MapDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type MapUpdateRequest struct {
	*Map
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type MapResponse struct {
	Meta meta  `json:"meta"`
	Data []Map `json:"data"`
}

func addMap() {
	resourceName := strcase.SnakeCase("Map")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/map/{id}")
	getOp.AddReqStructure(new(MapGetRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/rest/map/{id}",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetMap")
	getOp.SetTags("Map")
	getOp.AddRespStructure(new(MapResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/map/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateMap")
	updateOp.SetTags("Map")
	updateOp.AddReqStructure(new(MapUpdateRequest))

	updateOp.AddRespStructure(new(MapResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/map")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListMap")
	listOp.SetTags("Map")
	listOp.AddReqStructure(new(MapListRequest))

	listOp.AddRespStructure(new(MapResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/map")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateMap")
	createOp.SetTags("Map")
	createOp.AddReqStructure(new(MapCreateRequest))

	getOp.AddRespStructure(new(MapResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/map/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteMap")
	deleteOp.SetTags("Map")
	deleteOp.AddReqStructure(new(MapDeleteRequest))

	deleteOp.AddRespStructure(new(MapResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
