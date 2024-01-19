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

type WLANGroup struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Name string `json:"name,omitempty"` // .{1,128}
}

func (dst *WLANGroup) UnmarshalJSON(b []byte) error {
	type Alias WLANGroup
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

type WLANGroupGetRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type WLANGroupDeleteRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type WLANGroupUpdateRequest struct {
	*WLANGroup
	Site string `path:"site"`
	ID   string `path:"id",json:"_id,omitempty"`
}

type WLANGroupListRequest struct {
	Site string `path:"site"`
}

type WLANGroupCreateRequest struct {
	*WLANGroup
	Site string `path:"site"`
}

type WLANGroupResponse struct {
	Meta meta        `json:"meta"`
	Data []WLANGroup `json:"data"`
}

func addWLANGroup() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/wlangroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	getOp.AddReqStructure(new(WLANGroupGetRequest))
	getOp.AddRespStructure(new(WLANGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/rest/wlangroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.AddReqStructure(new(WLANGroupUpdateRequest))
	updateOp.AddRespStructure(new(WLANGroupResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/wlangroup")
	if err != nil {
		log.Fatal(err)
	}
	listOp.AddReqStructure(new(WLANGroupListRequest))
	listOp.AddRespStructure(new(WLANGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{site}/rest/wlangroup")
	if err != nil {
		log.Fatal(err)
	}
	createOp.AddReqStructure(new(WLANGroupCreateRequest))
	createOp.AddRespStructure(new(WLANGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{site}/get/setting/wlangroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.AddReqStructure(new(WLANGroupDeleteRequest))
	deleteOp.AddRespStructure(new(WLANGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
