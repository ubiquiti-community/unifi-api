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

type Tag struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	MemberTable []string `json:"member_table,omitempty"`
	Name        string   `json:"name,omitempty"`
}

func (dst *Tag) UnmarshalJSON(b []byte) error {
	type Alias Tag
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

type TagGetRequest struct {
	ID string `path:"id"`
}

type TagDeleteRequest struct {
	ID string `path:"id"`
}

type TagUpdateRequest struct {
	*Tag
	ID string `path:"id",json:"_id,omitempty"`
}

type TagResponse struct {
	Meta meta  `json:"meta"`
	Data []Tag `json:"data"`
}

func addTag() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/tag/{id}")
	getOp.AddReqStructure(new(TagGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetTag")
	getOp.SetTags("Tag")
	getOp.AddRespStructure(new(TagResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/rest/tag/{id}")
	updateOp.AddReqStructure(new(TagUpdateRequest))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateTag")
	updateOp.SetTags("Tag")
	updateOp.AddRespStructure(new(TagResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/tag")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListTag")
	listOp.SetTags("Tag")
	listOp.AddReqStructure(nil)
	listOp.AddRespStructure(new(TagResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/rest/tag")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateTag")
	createOp.SetTags("Tag")
	createOp.AddReqStructure(new(Tag))
	createOp.AddRespStructure(new(TagResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/rest/tag/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteTag")
	deleteOp.SetTags("Tag")
	deleteOp.AddReqStructure(new(TagDeleteRequest))
	deleteOp.AddRespStructure(new(TagResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
