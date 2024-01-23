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

type FirewallGroup struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	GroupMembers []string `json:"group_members,omitempty"`
	GroupType    string   `json:"group_type,omitempty"` // address-group|port-group|ipv6-address-group
	Name         string   `json:"name,omitempty"`       // .{1,64}
}

func (dst *FirewallGroup) UnmarshalJSON(b []byte) error {
	type Alias FirewallGroup
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

type FirewallGroupGetRequest struct {
	ID string `path:"id"`
}

type FirewallGroupDeleteRequest struct {
	ID string `path:"id"`
}

type FirewallGroupUpdateRequest struct {
	*FirewallGroup
	ID string `path:"id",json:"_id,omitempty"`
}

type FirewallGroupResponse struct {
	Meta meta            `json:"meta"`
	Data []FirewallGroup `json:"data"`
}

func addFirewallGroup() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/firewallgroup/{id}")
	getOp.AddReqStructure(new(FirewallGroupGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetFirewallGroup")
	getOp.SetTags("FirewallGroup")
	getOp.AddRespStructure(new(FirewallGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/rest/firewallgroup/{id}")
	updateOp.AddReqStructure(new(FirewallGroupUpdateRequest))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateFirewallGroup")
	updateOp.SetTags("FirewallGroup")

	updateOp.AddRespStructure(new(FirewallGroupResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/firewallgroup")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListFirewallGroup")
	listOp.SetTags("FirewallGroup")
	listOp.AddReqStructure(nil)

	listOp.AddRespStructure(new(FirewallGroupResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/rest/firewallgroup")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateFirewallGroup")
	createOp.SetTags("FirewallGroup")
	createOp.AddReqStructure(new(FirewallGroup))

	getOp.AddRespStructure(new(FirewallGroupResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/rest/firewallgroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteFirewallGroup")
	deleteOp.SetTags("FirewallGroup")
	deleteOp.AddReqStructure(new(FirewallGroupDeleteRequest))

	deleteOp.AddRespStructure(new(FirewallGroupResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
