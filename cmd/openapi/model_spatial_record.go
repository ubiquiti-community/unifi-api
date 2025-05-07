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

type SpatialRecord struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Devices []SpatialRecordDevices `json:"devices,omitempty"`
	Name    string                 `json:"name,omitempty"` // .{1,128}
}

func (dst *SpatialRecord) UnmarshalJSON(b []byte) error {
	type Alias SpatialRecord
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

type SpatialRecordDevices struct {
	MAC      string                `json:"mac,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	Position SpatialRecordPosition `json:"position,omitempty"`
}

func (dst *SpatialRecordDevices) UnmarshalJSON(b []byte) error {
	type Alias SpatialRecordDevices
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

type SpatialRecordPosition struct {
	X float64 `json:"x,omitempty"` // (^([-]?[\d]+)$)|(^([-]?[\d]+[.]?[\d]+)$)
	Y float64 `json:"y,omitempty"` // (^([-]?[\d]+)$)|(^([-]?[\d]+[.]?[\d]+)$)
	Z float64 `json:"z,omitempty"` // (^([-]?[\d]+)$)|(^([-]?[\d]+[.]?[\d]+)$)
}

func (dst *SpatialRecordPosition) UnmarshalJSON(b []byte) error {
	type Alias SpatialRecordPosition
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

type SpatialRecordGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type SpatialRecordListRequest struct {
	SiteID string `path:"siteId"`
}

type SpatialRecordCreateRequest struct {
	*SpatialRecord
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type SpatialRecordDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type SpatialRecordUpdateRequest struct {
	*SpatialRecord
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type SpatialRecordResponse struct {
	Meta meta            `json:"meta"`
	Data []SpatialRecord `json:"data"`
}

func addSpatialRecord() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/spatialrecord/{id}")
	getOp.AddReqStructure(new(SpatialRecordGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSpatialRecord")
	getOp.SetTags("SpatialRecord")
	getOp.AddRespStructure(new(SpatialRecordResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/spatialrecord/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSpatialRecord")
	updateOp.SetTags("SpatialRecord")
	updateOp.AddReqStructure(new(SpatialRecordUpdateRequest))

	updateOp.AddRespStructure(new(SpatialRecordResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/spatialrecord")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListSpatialRecord")
	listOp.SetTags("SpatialRecord")
	listOp.AddReqStructure(new(SpatialRecordListRequest))

	listOp.AddRespStructure(new(SpatialRecordResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/spatialrecord")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateSpatialRecord")
	createOp.SetTags("SpatialRecord")
	createOp.AddReqStructure(new(SpatialRecordCreateRequest))

	getOp.AddRespStructure(new(SpatialRecordResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/spatialrecord/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteSpatialRecord")
	deleteOp.SetTags("SpatialRecord")
	deleteOp.AddReqStructure(new(SpatialRecordDeleteRequest))

	deleteOp.AddRespStructure(new(SpatialRecordResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
