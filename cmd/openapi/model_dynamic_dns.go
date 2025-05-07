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

type DynamicDNS struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	CustomService string   `json:"custom_service,omitempty"` // ^[^"' ]+$
	HostName      string   `json:"host_name,omitempty"`      // ^[^"' ]+$
	Interface     string   `json:"interface,omitempty"`      // wan[2-8]?
	Login         string   `json:"login,omitempty"`          // ^[^"' ]+$
	Options       []string `json:"options,omitempty"`        // ^[^"' ]+$
	Server        string   `json:"server"`                   // ^[^"' ]+$|^$
	Service       string   `json:"service,omitempty"`        // afraid|changeip|cloudflare|cloudxns|ddnss|dhis|dnsexit|dnsomatic|dnspark|dnspod|dslreports|dtdns|duckdns|duiadns|dyn|dyndns|dynv6|easydns|freemyip|googledomains|loopia|namecheap|noip|nsupdate|ovh|sitelutions|spdyn|strato|tunnelbroker|zoneedit|cloudflare|custom
	XPassword     string   `json:"x_password,omitempty"`     // ^[^"' ]+$
}

func (dst *DynamicDNS) UnmarshalJSON(b []byte) error {
	type Alias DynamicDNS
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

type DynamicDNSGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type DynamicDNSListRequest struct {
	SiteID string `path:"siteId"`
}

type DynamicDNSCreateRequest struct {
	*DynamicDNS
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type DynamicDNSDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type DynamicDNSUpdateRequest struct {
	*DynamicDNS
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type DynamicDNSResponse struct {
	Meta meta         `json:"meta"`
	Data []DynamicDNS `json:"data"`
}

func addDynamicDNS() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/dynamicdns/{id}")
	getOp.AddReqStructure(new(DynamicDNSGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetDynamicDNS")
	getOp.SetTags("DynamicDNS")
	getOp.AddRespStructure(new(DynamicDNSResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/dynamicdns/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateDynamicDNS")
	updateOp.SetTags("DynamicDNS")
	updateOp.AddReqStructure(new(DynamicDNSUpdateRequest))

	updateOp.AddRespStructure(new(DynamicDNSResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/dynamicdns")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListDynamicDNS")
	listOp.SetTags("DynamicDNS")
	listOp.AddReqStructure(new(DynamicDNSListRequest))

	listOp.AddRespStructure(new(DynamicDNSResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/dynamicdns")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateDynamicDNS")
	createOp.SetTags("DynamicDNS")
	createOp.AddReqStructure(new(DynamicDNSCreateRequest))

	getOp.AddRespStructure(new(DynamicDNSResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/dynamicdns/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteDynamicDNS")
	deleteOp.SetTags("DynamicDNS")
	deleteOp.AddReqStructure(new(DynamicDNSDeleteRequest))

	deleteOp.AddRespStructure(new(DynamicDNSResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
