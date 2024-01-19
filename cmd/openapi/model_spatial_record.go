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
	Site string `path:"site"`
	ID   string `path:"id"`
}

type SpatialRecordDeleteRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type SpatialRecordUpdateRequest struct {
	*SpatialRecord
	Site string `path:"site"`
	ID   string `path:"id",json:"_id,omitempty"`
}

type SpatialRecordListRequest struct {
	Site string `path:"site"`
}

type SpatialRecordCreateRequest struct {
	*SpatialRecord
	Site string `path:"site"`
}

type SpatialRecordResponse struct {
	Meta meta            `json:"meta"`
	Data []SpatialRecord `json:"data"`
}

func addSpatialRecord() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/spatialrecord/{id}")
	if err != nil {
		log.Fatal(err)
	}
	getOp.AddReqStructure(new(SpatialRecordGetRequest))
	getOp.AddRespStructure(new(SpatialRecordResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/rest/spatialrecord/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.AddReqStructure(new(SpatialRecordUpdateRequest))
	updateOp.AddRespStructure(new(SpatialRecordResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/spatialrecord")
	if err != nil {
		log.Fatal(err)
	}
	listOp.AddReqStructure(new(SpatialRecordListRequest))
	listOp.AddRespStructure(new(SpatialRecordResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{site}/rest/spatialrecord")
	if err != nil {
		log.Fatal(err)
	}
	createOp.AddReqStructure(new(SpatialRecordCreateRequest))
	createOp.AddRespStructure(new(SpatialRecordResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{site}/get/setting/spatialrecord/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.AddReqStructure(new(SpatialRecordDeleteRequest))
	deleteOp.AddRespStructure(new(SpatialRecordResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
