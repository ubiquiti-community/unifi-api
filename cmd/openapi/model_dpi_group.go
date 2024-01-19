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
	Site string `path:"site"`
	ID   string `path:"id"`
}

type DpiGroupDeleteRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type DpiGroupUpdateRequest struct {
	*DpiGroup
	Site string `path:"site"`
	ID   string `path:"id",json:"_id,omitempty"`
}

type DpiGroupListRequest struct {
	Site string `path:"site"`
}

type DpiGroupCreateRequest struct {
	*DpiGroup
	Site string `path:"site"`
}

type DpiGroupResponse struct {
	Meta meta       `json:"meta"`
	Data []DpiGroup `json:"data"`
}

func addDpiGroup() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/dpigroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetDpiGroup")
	getOp.SetTags("DpiGroup")
	getOp.AddReqStructure(new(DpiGroupGetRequest))
	getOp.AddRespStructure(new(DpiGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/rest/dpigroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateDpiGroup")
	updateOp.SetTags("DpiGroup")
	updateOp.AddReqStructure(new(DpiGroupUpdateRequest))
	updateOp.AddRespStructure(new(DpiGroupResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/dpigroup")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListDpiGroup")
	listOp.SetTags("DpiGroup")
	listOp.AddReqStructure(new(DpiGroupListRequest))
	listOp.AddRespStructure(new(DpiGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{site}/rest/dpigroup")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateDpiGroup")
	createOp.SetTags("DpiGroup")
	createOp.AddReqStructure(new(DpiGroupCreateRequest))
	createOp.AddRespStructure(new(DpiGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{site}/get/setting/dpigroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteDpiGroup")
	deleteOp.SetTags("DpiGroup")
	deleteOp.AddReqStructure(new(DpiGroupDeleteRequest))
	deleteOp.AddRespStructure(new(DpiGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
