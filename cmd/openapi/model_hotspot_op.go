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

type HotspotOp struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Name      string `json:"name,omitempty"` // .{1,256}
	Note      string `json:"note,omitempty"`
	XPassword string `json:"x_password,omitempty"` // .{1,256}
}

func (dst *HotspotOp) UnmarshalJSON(b []byte) error {
	type Alias HotspotOp
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

type HotspotOpGetRequest struct {
	ID string `path:"id"`
}

type HotspotOpDeleteRequest struct {
	ID string `path:"id"`
}

type HotspotOpUpdateRequest struct {
	*HotspotOp
	ID string `path:"id",json:"_id,omitempty"`
}

type HotspotOpResponse struct {
	Meta meta        `json:"meta"`
	Data []HotspotOp `json:"data"`
}

func addHotspotOp() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/hotspotop/{id}")
	getOp.AddReqStructure(new(HotspotOpGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetHotspotOp")
	getOp.SetTags("HotspotOp")
	getOp.AddRespStructure(new(HotspotOpResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/rest/hotspotop/{id}")
	updateOp.AddReqStructure(new(HotspotOpUpdateRequest))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateHotspotOp")
	updateOp.SetTags("HotspotOp")

	updateOp.AddRespStructure(new(HotspotOpResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/hotspotop")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListHotspotOp")
	listOp.SetTags("HotspotOp")
	listOp.AddReqStructure(nil)

	listOp.AddRespStructure(new(HotspotOpResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/rest/hotspotop")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateHotspotOp")
	createOp.SetTags("HotspotOp")
	createOp.AddReqStructure(new(HotspotOp))

	getOp.AddRespStructure(new(HotspotOpResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/rest/hotspotop/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteHotspotOp")
	deleteOp.SetTags("HotspotOp")
	deleteOp.AddReqStructure(new(HotspotOpDeleteRequest))

	deleteOp.AddRespStructure(new(HotspotOpResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
