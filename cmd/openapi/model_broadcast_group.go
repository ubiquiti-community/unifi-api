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
	ID string `path:"id"`
}

type BroadcastGroupDeleteRequest struct {
	ID string `path:"id"`
}

type BroadcastGroupUpdateRequest struct {
	*BroadcastGroup
	ID string `path:"id",json:"_id,omitempty"`
}

type BroadcastGroupResponse struct {
	Meta meta             `json:"meta"`
	Data []BroadcastGroup `json:"data"`
}

func addBroadcastGroup() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/broadcastgroup/{id}")
	getOp.AddReqStructure(new(BroadcastGroupGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetBroadcastGroup")
	getOp.SetTags("BroadcastGroup")
	getOp.AddRespStructure(new(BroadcastGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/rest/broadcastgroup/{id}")
	updateOp.AddReqStructure(new(BroadcastGroupUpdateRequest))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateBroadcastGroup")
	updateOp.SetTags("BroadcastGroup")

	updateOp.AddRespStructure(new(BroadcastGroupResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/broadcastgroup")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListBroadcastGroup")
	listOp.SetTags("BroadcastGroup")
	listOp.AddReqStructure(nil)

	listOp.AddRespStructure(new(BroadcastGroupResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/rest/broadcastgroup")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateBroadcastGroup")
	createOp.SetTags("BroadcastGroup")
	createOp.AddReqStructure(new(BroadcastGroup))

	getOp.AddRespStructure(new(BroadcastGroupResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/rest/broadcastgroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteBroadcastGroup")
	deleteOp.SetTags("BroadcastGroup")
	deleteOp.AddReqStructure(new(BroadcastGroupDeleteRequest))

	deleteOp.AddRespStructure(new(BroadcastGroupResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
