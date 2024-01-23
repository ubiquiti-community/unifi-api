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

type UserGroup struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Name           string `json:"name,omitempty"`              // .{1,128}
	QOSRateMaxDown int    `json:"qos_rate_max_down,omitempty"` // -1|[2-9]|[1-9][0-9]{1,4}|100000
	QOSRateMaxUp   int    `json:"qos_rate_max_up,omitempty"`   // -1|[2-9]|[1-9][0-9]{1,4}|100000
}

func (dst *UserGroup) UnmarshalJSON(b []byte) error {
	type Alias UserGroup
	aux := &struct {
		QOSRateMaxDown emptyStringInt `json:"qos_rate_max_down"`
		QOSRateMaxUp   emptyStringInt `json:"qos_rate_max_up"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.QOSRateMaxDown = int(aux.QOSRateMaxDown)
	dst.QOSRateMaxUp = int(aux.QOSRateMaxUp)

	return nil
}

type UserGroupGetRequest struct {
	ID string `path:"id"`
}

type UserGroupDeleteRequest struct {
	ID string `path:"id"`
}

type UserGroupUpdateRequest struct {
	*UserGroup
	ID string `path:"id",json:"_id,omitempty"`
}

type UserGroupResponse struct {
	Meta meta        `json:"meta"`
	Data []UserGroup `json:"data"`
}

func addUserGroup() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/usergroup/{id}")
	getOp.AddReqStructure(new(UserGroupGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetUserGroup")
	getOp.SetTags("UserGroup")
	getOp.AddRespStructure(new(UserGroupResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/rest/usergroup/{id}")
	updateOp.AddReqStructure(new(UserGroupUpdateRequest))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateUserGroup")
	updateOp.SetTags("UserGroup")

	updateOp.AddRespStructure(new(UserGroupResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/usergroup")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListUserGroup")
	listOp.SetTags("UserGroup")
	listOp.AddReqStructure(nil)

	listOp.AddRespStructure(new(UserGroupResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/rest/usergroup")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateUserGroup")
	createOp.SetTags("UserGroup")
	createOp.AddReqStructure(new(UserGroup))

	getOp.AddRespStructure(new(UserGroupResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/rest/usergroup/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteUserGroup")
	deleteOp.SetTags("UserGroup")
	deleteOp.AddReqStructure(new(UserGroupDeleteRequest))

	deleteOp.AddRespStructure(new(UserGroupResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
