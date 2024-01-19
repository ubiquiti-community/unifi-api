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
	Site string `path:"site"`
	ID   string `path:"id"`
}

type MediaFileDeleteRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type MediaFileUpdateRequest struct {
	*MediaFile
	Site string `path:"site"`
	ID   string `path:"id",json:"_id,omitempty"`
}

type MediaFileListRequest struct {
	Site string `path:"site"`
}

type MediaFileCreateRequest struct {
	*MediaFile
	Site string `path:"site"`
}

type MediaFileResponse struct {
	Meta meta        `json:"meta"`
	Data []MediaFile `json:"data"`
}

func addMediaFile() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/mediafile/{id}")
	if err != nil {
		log.Fatal(err)
	}
	getOp.AddReqStructure(new(MediaFileGetRequest))
	getOp.AddRespStructure(new(MediaFileResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/rest/mediafile/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.AddReqStructure(new(MediaFileUpdateRequest))
	updateOp.AddRespStructure(new(MediaFileResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/mediafile")
	if err != nil {
		log.Fatal(err)
	}
	listOp.AddReqStructure(new(MediaFileListRequest))
	listOp.AddRespStructure(new(MediaFileResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{site}/rest/mediafile")
	if err != nil {
		log.Fatal(err)
	}
	createOp.AddReqStructure(new(MediaFileCreateRequest))
	createOp.AddRespStructure(new(MediaFileResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{site}/get/setting/mediafile/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.AddReqStructure(new(MediaFileDeleteRequest))
	deleteOp.AddRespStructure(new(MediaFileResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
