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

type User struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	DevIdOverride int    `json:"dev_id_override,omitempty"` // non-generated field
	IP            string `json:"ip,omitempty"`              // non-generated field

	Blocked                       bool   `json:"blocked,omitempty"`
	FixedApEnabled                bool   `json:"fixed_ap_enabled"`
	FixedApMAC                    string `json:"fixed_ap_mac,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	FixedIP                       string `json:"fixed_ip,omitempty"`
	Hostname                      string `json:"hostname,omitempty"`
	LastSeen                      int    `json:"last_seen,omitempty"`
	LocalDNSRecord                string `json:"local_dns_record,omitempty"`
	LocalDNSRecordEnabled         bool   `json:"local_dns_record_enabled"`
	MAC                           string `json:"mac,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	Name                          string `json:"name,omitempty"`
	NetworkID                     string `json:"network_id"`
	Note                          string `json:"note,omitempty"`
	UseFixedIP                    bool   `json:"use_fixedip"`
	UserGroupID                   string `json:"usergroup_id"`
	VirtualNetworkOverrideEnabled bool   `json:"virtual_network_override_enabled"`
	VirtualNetworkOverrideID      string `json:"virtual_network_override_id"`
}

func (dst *User) UnmarshalJSON(b []byte) error {
	type Alias User
	aux := &struct {
		LastSeen emptyStringInt `json:"last_seen"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.LastSeen = int(aux.LastSeen)

	return nil
}

type UserGetRequest struct {
	ID string `path:"id"`
}

type UserDeleteRequest struct {
	ID string `path:"id"`
}

type UserUpdateRequest struct {
	*User
	ID string `path:"id",json:"_id,omitempty"`
}

type UserResponse struct {
	Meta meta   `json:"meta"`
	Data []User `json:"data"`
}

func addUser() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/user/{id}")
	getOp.AddReqStructure(new(UserGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetUser")
	getOp.SetTags("User")
	getOp.AddRespStructure(new(UserResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/rest/user/{id}")
	updateOp.AddReqStructure(new(UserUpdateRequest))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateUser")
	updateOp.SetTags("User")

	updateOp.AddRespStructure(new(UserResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/user")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListUser")
	listOp.SetTags("User")
	listOp.AddReqStructure(nil)

	listOp.AddRespStructure(new(UserResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/rest/user")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateUser")
	createOp.SetTags("User")
	createOp.AddReqStructure(new(User))

	getOp.AddRespStructure(new(UserResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/rest/user/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteUser")
	deleteOp.SetTags("User")
	deleteOp.AddReqStructure(new(UserDeleteRequest))

	deleteOp.AddRespStructure(new(UserResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
