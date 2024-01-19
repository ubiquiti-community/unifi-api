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

type BroadcastGroup struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	MemberTable []string `json:"member_table,omitempty"`
	Name        string   `json:"name,omitempty"`
}

func (dst *BroadcastGroup) UnmarshalJSON(b []byte) error {
	type Alias BroadcastGroup
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

type BroadcastGroupGetRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type BroadcastGroupDeleteRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type BroadcastGroupUpdateRequest struct {
	*BroadcastGroup
	Site string `path:"site"`
	ID   string `path:"id",json:"_id,omitempty"`
}

type BroadcastGroupListRequest struct {
	Site string `path:"site"`
}

type BroadcastGroupCreateRequest struct {
	*BroadcastGroup
	Site string `path:"site"`
}

type BroadcastGroupResponse struct {
	Meta meta             `json:"meta"`
	Data []BroadcastGroup `json:"data"`
}

func addBroadcastGroup() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/broadcastgroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	getOp.AddReqStructure(new(BroadcastGroupGetRequest))
	getOp.AddRespStructure(new(BroadcastGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/rest/broadcastgroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.AddReqStructure(new(BroadcastGroupUpdateRequest))
	updateOp.AddRespStructure(new(BroadcastGroupResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/broadcastgroup")
	if err != nil {
		log.Fatal(err)
	}
	listOp.AddReqStructure(new(BroadcastGroupListRequest))
	listOp.AddRespStructure(new(BroadcastGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{site}/rest/broadcastgroup")
	if err != nil {
		log.Fatal(err)
	}
	createOp.AddReqStructure(new(BroadcastGroupCreateRequest))
	createOp.AddRespStructure(new(BroadcastGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{site}/get/setting/broadcastgroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.AddReqStructure(new(BroadcastGroupDeleteRequest))
	deleteOp.AddRespStructure(new(BroadcastGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
