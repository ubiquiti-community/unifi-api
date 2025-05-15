// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/stoewer/go-strcase"
	"github.com/swaggest/openapi-go"
)

// just to fix compile issues with the import
var (
	_ context.Context
	_ fmt.Formatter
	_ json.Marshaler
)

type Routing struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Enabled              bool   `json:"enabled"`
	GatewayDevice        string `json:"gateway_device,omitempty"`        // ^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$
	GatewayType          string `json:"gateway_type,omitempty"`          // default|switch
	Name                 string `json:"name,omitempty"`                  // .{1,128}
	StaticRouteDistance  int    `json:"static-route_distance,omitempty"` // ^[1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]$|^$
	StaticRouteInterface string `json:"static-route_interface"`          // WAN[1-8]?|[\d\w]+|^$
	StaticRouteNetwork   string `json:"static-route_network,omitempty"`  // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\/([1-9]|[1-2][0-9]|3[0-2])$|^([a-fA-F0-9:]+\/(([1-9]|[1-8][0-9]|9[0-9]|1[01][0-9]|12[0-8])))$
	StaticRouteNexthop   string `json:"static-route_nexthop"`            // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^([a-fA-F0-9:]+)$|^$
	StaticRouteType      string `json:"static-route_type,omitempty"`     // nexthop-route|interface-route|blackhole
	Type                 string `json:"type,omitempty"`                  // static-route
}

func (dst *Routing) UnmarshalJSON(b []byte) error {
	type Alias Routing
	aux := &struct {
		StaticRouteDistance emptyStringInt `json:"static-route_distance"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.StaticRouteDistance = int(aux.StaticRouteDistance)

	return nil
}

type RoutingGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type RoutingListRequest struct {
	SiteID string `path:"siteId"`
}

type RoutingCreateRequest struct {
	*Routing
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type RoutingDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type RoutingUpdateRequest struct {
	*Routing
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type RoutingResponse struct {
	Meta meta      `json:"meta"`
	Data []Routing `json:"data"`
}

func addRouting() {
	resourceName := strcase.SnakeCase("Routing")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/routing/{id}")
	getOp.AddReqStructure(new(RoutingGetRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/rest/routing/{id}",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetRouting")
	getOp.SetTags("Routing")
	getOp.AddRespStructure(new(RoutingResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/routing/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateRouting")
	updateOp.SetTags("Routing")
	updateOp.AddReqStructure(new(RoutingUpdateRequest))

	updateOp.AddRespStructure(new(RoutingResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/routing")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListRouting")
	listOp.SetTags("Routing")
	listOp.AddReqStructure(new(RoutingListRequest))

	listOp.AddRespStructure(new(RoutingResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/routing")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateRouting")
	createOp.SetTags("Routing")
	createOp.AddReqStructure(new(RoutingCreateRequest))

	getOp.AddRespStructure(new(RoutingResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/routing/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteRouting")
	deleteOp.SetTags("Routing")
	deleteOp.AddReqStructure(new(RoutingDeleteRequest))

	deleteOp.AddRespStructure(new(RoutingResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
