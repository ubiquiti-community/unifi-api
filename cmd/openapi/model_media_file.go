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

type MediaFile struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Name string `json:"name,omitempty"`
}

func (dst *MediaFile) UnmarshalJSON(b []byte) error {
	type Alias MediaFile
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

type MediaFileGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type MediaFileListRequest struct {
	SiteID string `path:"siteId"`
}

type MediaFileCreateRequest struct {
	*MediaFile
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type MediaFileDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type MediaFileUpdateRequest struct {
	*MediaFile
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type MediaFileResponse struct {
	Meta meta        `json:"meta"`
	Data []MediaFile `json:"data"`
}

func addMediaFile() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/mediafile/{id}")
	getOp.AddReqStructure(new(MediaFileGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetMediaFile")
	getOp.SetTags("MediaFile")
	getOp.AddRespStructure(new(MediaFileResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/mediafile/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateMediaFile")
	updateOp.SetTags("MediaFile")
	updateOp.AddReqStructure(new(MediaFileUpdateRequest))

	updateOp.AddRespStructure(new(MediaFileResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/mediafile")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListMediaFile")
	listOp.SetTags("MediaFile")
	listOp.AddReqStructure(new(MediaFileListRequest))

	listOp.AddRespStructure(new(MediaFileResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/mediafile")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateMediaFile")
	createOp.SetTags("MediaFile")
	createOp.AddReqStructure(new(MediaFileCreateRequest))

	getOp.AddRespStructure(new(MediaFileResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/mediafile/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteMediaFile")
	deleteOp.SetTags("MediaFile")
	deleteOp.AddReqStructure(new(MediaFileDeleteRequest))

	deleteOp.AddRespStructure(new(MediaFileResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
