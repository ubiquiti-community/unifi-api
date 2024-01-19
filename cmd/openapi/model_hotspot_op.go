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
	Site string `path:"site"`
	ID   string `path:"id"`
}

type HotspotOpDeleteRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type HotspotOpUpdateRequest struct {
	*HotspotOp
	Site string `path:"site"`
	ID   string `path:"id",json:"_id,omitempty"`
}

type HotspotOpListRequest struct {
	Site string `path:"site"`
}

type HotspotOpCreateRequest struct {
	*HotspotOp
	Site string `path:"site"`
}

type HotspotOpResponse struct {
	Meta meta        `json:"meta"`
	Data []HotspotOp `json:"data"`
}

func addHotspotOp() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/hotspotop/{id}")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetHotspotOp")
	getOp.SetTags("HotspotOp")
	getOp.AddReqStructure(new(HotspotOpGetRequest))
	getOp.AddRespStructure(new(HotspotOpResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/rest/hotspotop/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateHotspotOp")
	updateOp.SetTags("HotspotOp")
	updateOp.AddReqStructure(new(HotspotOpUpdateRequest))
	updateOp.AddRespStructure(new(HotspotOpResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/hotspotop")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListHotspotOp")
	listOp.SetTags("HotspotOp")
	listOp.AddReqStructure(new(HotspotOpListRequest))
	listOp.AddRespStructure(new(HotspotOpResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{site}/rest/hotspotop")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateHotspotOp")
	createOp.SetTags("HotspotOp")
	createOp.AddReqStructure(new(HotspotOpCreateRequest))
	createOp.AddRespStructure(new(HotspotOpResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{site}/get/setting/hotspotop/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteHotspotOp")
	deleteOp.SetTags("HotspotOp")
	deleteOp.AddReqStructure(new(HotspotOpDeleteRequest))
	deleteOp.AddRespStructure(new(HotspotOpResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
