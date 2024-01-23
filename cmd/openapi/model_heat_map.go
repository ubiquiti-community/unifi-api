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

type HeatMap struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Description string `json:"description,omitempty"`
	MapID       string `json:"map_id"`
	Name        string `json:"name,omitempty"` // .*[^\s]+.*
	Type        string `json:"type,omitempty"` // download|upload
}

func (dst *HeatMap) UnmarshalJSON(b []byte) error {
	type Alias HeatMap
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

type HeatMapGetRequest struct {
	ID string `path:"id"`
}

type HeatMapDeleteRequest struct {
	ID string `path:"id"`
}

type HeatMapUpdateRequest struct {
	*HeatMap
	ID string `path:"id",json:"_id,omitempty"`
}

type HeatMapResponse struct {
	Meta meta      `json:"meta"`
	Data []HeatMap `json:"data"`
}

func addHeatMap() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/heatmap/{id}")
	getOp.AddReqStructure(new(HeatMapGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetHeatMap")
	getOp.SetTags("HeatMap")
	getOp.AddRespStructure(new(HeatMapResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/rest/heatmap/{id}")
	updateOp.AddReqStructure(new(HeatMapUpdateRequest))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateHeatMap")
	updateOp.SetTags("HeatMap")

	updateOp.AddRespStructure(new(HeatMapResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/heatmap")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListHeatMap")
	listOp.SetTags("HeatMap")
	listOp.AddReqStructure(nil)

	listOp.AddRespStructure(new(HeatMapResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/rest/heatmap")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateHeatMap")
	createOp.SetTags("HeatMap")
	createOp.AddReqStructure(new(HeatMap))

	getOp.AddRespStructure(new(HeatMapResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/rest/heatmap/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteHeatMap")
	deleteOp.SetTags("HeatMap")
	deleteOp.AddReqStructure(new(HeatMapDeleteRequest))

	deleteOp.AddRespStructure(new(HeatMapResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
