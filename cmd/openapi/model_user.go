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
	Site string `path:"site"`
	ID   string `path:"id"`
}

type UserDeleteRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type UserUpdateRequest struct {
	*User
	Site string `path:"site"`
	ID   string `path:"id",json:"_id,omitempty"`
}

type UserListRequest struct {
	Site string `path:"site"`
}

type UserCreateRequest struct {
	*User
	Site string `path:"site"`
}

type UserResponse struct {
	Meta meta   `json:"meta"`
	Data []User `json:"data"`
}

func addUser() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/user/{id}")
	if err != nil {
		log.Fatal(err)
	}
	getOp.AddReqStructure(new(UserGetRequest))
	getOp.AddRespStructure(new(UserResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/rest/user/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.AddReqStructure(new(UserUpdateRequest))
	updateOp.AddRespStructure(new(UserResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/user")
	if err != nil {
		log.Fatal(err)
	}
	listOp.AddReqStructure(new(UserListRequest))
	listOp.AddRespStructure(new(UserResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{site}/rest/user")
	if err != nil {
		log.Fatal(err)
	}
	createOp.AddReqStructure(new(UserCreateRequest))
	createOp.AddRespStructure(new(UserResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{site}/get/setting/user/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.AddReqStructure(new(UserDeleteRequest))
	deleteOp.AddRespStructure(new(UserResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
