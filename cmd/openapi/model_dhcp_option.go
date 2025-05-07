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

type DHCPOption struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Code   string `json:"code,omitempty"` // ^(?!(?:15|42|43|44|51|66|67|252)$)([7-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-4])$
	Name   string `json:"name,omitempty"` // ^[A-Za-z0-9-_]{1,25}$
	Signed bool   `json:"signed"`
	Type   string `json:"type,omitempty"`  // ^(boolean|hexarray|integer|ipaddress|macaddress|text)$
	Width  int    `json:"width,omitempty"` // ^(8|16|32)$
}

func (dst *DHCPOption) UnmarshalJSON(b []byte) error {
	type Alias DHCPOption
	aux := &struct {
		Width emptyStringInt `json:"width"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Width = int(aux.Width)

	return nil
}

type DHCPOptionGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type DHCPOptionListRequest struct {
	SiteID string `path:"siteId"`
}

type DHCPOptionCreateRequest struct {
	*DHCPOption
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type DHCPOptionDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type DHCPOptionUpdateRequest struct {
	*DHCPOption
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type DHCPOptionResponse struct {
	Meta meta         `json:"meta"`
	Data []DHCPOption `json:"data"`
}

func addDHCPOption() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/dhcpoption/{id}")
	getOp.AddReqStructure(new(DHCPOptionGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetDHCPOption")
	getOp.SetTags("DHCPOption")
	getOp.AddRespStructure(new(DHCPOptionResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/dhcpoption/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateDHCPOption")
	updateOp.SetTags("DHCPOption")
	updateOp.AddReqStructure(new(DHCPOptionUpdateRequest))

	updateOp.AddRespStructure(new(DHCPOptionResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/dhcpoption")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListDHCPOption")
	listOp.SetTags("DHCPOption")
	listOp.AddReqStructure(new(DHCPOptionListRequest))

	listOp.AddRespStructure(new(DHCPOptionResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/dhcpoption")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateDHCPOption")
	createOp.SetTags("DHCPOption")
	createOp.AddReqStructure(new(DHCPOptionCreateRequest))

	getOp.AddRespStructure(new(DHCPOptionResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/dhcpoption/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteDHCPOption")
	deleteOp.SetTags("DHCPOption")
	deleteOp.AddReqStructure(new(DHCPOptionDeleteRequest))

	deleteOp.AddRespStructure(new(DHCPOptionResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
