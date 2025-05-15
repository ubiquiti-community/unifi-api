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

type VirtualDevice struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	HeightInMeters float64 `json:"heightInMeters,omitempty"`
	Locked         bool    `json:"locked"`
	MapID          string  `json:"map_id"`
	Type           string  `json:"type,omitempty"` // uap|usg|usw
	X              string  `json:"x,omitempty"`
	Y              string  `json:"y,omitempty"`
}

func (dst *VirtualDevice) UnmarshalJSON(b []byte) error {
	type Alias VirtualDevice
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

type VirtualDeviceGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type VirtualDeviceListRequest struct {
	SiteID string `path:"siteId"`
}

type VirtualDeviceCreateRequest struct {
	*VirtualDevice
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type VirtualDeviceDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type VirtualDeviceUpdateRequest struct {
	*VirtualDevice
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type VirtualDeviceResponse struct {
	Meta meta            `json:"meta"`
	Data []VirtualDevice `json:"data"`
}

func addVirtualDevice() {
	resourceName := strcase.SnakeCase("VirtualDevice")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/virtualdevice/{id}")
	getOp.AddReqStructure(new(VirtualDeviceGetRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/rest/virtualdevice/{id}",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetVirtualDevice")
	getOp.SetTags("VirtualDevice")
	getOp.AddRespStructure(new(VirtualDeviceResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/virtualdevice/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateVirtualDevice")
	updateOp.SetTags("VirtualDevice")
	updateOp.AddReqStructure(new(VirtualDeviceUpdateRequest))

	updateOp.AddRespStructure(new(VirtualDeviceResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/virtualdevice")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListVirtualDevice")
	listOp.SetTags("VirtualDevice")
	listOp.AddReqStructure(new(VirtualDeviceListRequest))

	listOp.AddRespStructure(new(VirtualDeviceResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/virtualdevice")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateVirtualDevice")
	createOp.SetTags("VirtualDevice")
	createOp.AddReqStructure(new(VirtualDeviceCreateRequest))

	getOp.AddRespStructure(new(VirtualDeviceResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/virtualdevice/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteVirtualDevice")
	deleteOp.SetTags("VirtualDevice")
	deleteOp.AddReqStructure(new(VirtualDeviceDeleteRequest))

	deleteOp.AddRespStructure(new(VirtualDeviceResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
