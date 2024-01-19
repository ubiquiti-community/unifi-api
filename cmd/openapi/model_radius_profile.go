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

type RADIUSProfile struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	AccountingEnabled     bool                       `json:"accounting_enabled"`
	AcctServers           []RADIUSProfileAcctServers `json:"acct_servers,omitempty"`
	AuthServers           []RADIUSProfileAuthServers `json:"auth_servers,omitempty"`
	InterimUpdateEnabled  bool                       `json:"interim_update_enabled"`
	InterimUpdateInterval int                        `json:"interim_update_interval,omitempty"` // ^([6-9][0-9]|[1-9][0-9]{2,3}|[1-7][0-9]{4}|8[0-5][0-9]{3}|86[0-3][0-9][0-9]|86400)$
	Name                  string                     `json:"name,omitempty"`                    // .{1,128}
	UseUsgAcctServer      bool                       `json:"use_usg_acct_server"`
	UseUsgAuthServer      bool                       `json:"use_usg_auth_server"`
	VLANEnabled           bool                       `json:"vlan_enabled"`
	VLANWLANMode          string                     `json:"vlan_wlan_mode,omitempty"` // disabled|optional|required
}

func (dst *RADIUSProfile) UnmarshalJSON(b []byte) error {
	type Alias RADIUSProfile
	aux := &struct {
		InterimUpdateInterval emptyStringInt `json:"interim_update_interval"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.InterimUpdateInterval = int(aux.InterimUpdateInterval)

	return nil
}

type RADIUSProfileAcctServers struct {
	IP      string `json:"ip,omitempty"`   // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$
	Port    int    `json:"port,omitempty"` // ^([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])$|^$
	XSecret string `json:"x_secret,omitempty"`
}

func (dst *RADIUSProfileAcctServers) UnmarshalJSON(b []byte) error {
	type Alias RADIUSProfileAcctServers
	aux := &struct {
		Port emptyStringInt `json:"port"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Port = int(aux.Port)

	return nil
}

type RADIUSProfileAuthServers struct {
	IP      string `json:"ip,omitempty"`   // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$
	Port    int    `json:"port,omitempty"` // ^([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])$|^$
	XSecret string `json:"x_secret,omitempty"`
}

func (dst *RADIUSProfileAuthServers) UnmarshalJSON(b []byte) error {
	type Alias RADIUSProfileAuthServers
	aux := &struct {
		Port emptyStringInt `json:"port"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Port = int(aux.Port)

	return nil
}

type RADIUSProfileGetRequest struct {
	ID string `path:"id"`
}

type RADIUSProfileDeleteRequest struct {
	ID string `path:"id"`
}

type RADIUSProfileUpdateRequest struct {
	*RADIUSProfile
	ID string `path:"id",json:"_id,omitempty"`
}

type RADIUSProfileResponse struct {
	Meta meta            `json:"meta"`
	Data []RADIUSProfile `json:"data"`
}

func addRADIUSProfile() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/radiusprofile/{id}")
	getOp.AddReqStructure(new(RADIUSProfileGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetRADIUSProfile")
	getOp.SetTags("RADIUSProfile")
	getOp.AddRespStructure(new(RADIUSProfileResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/rest/radiusprofile/{id}")
	updateOp.AddReqStructure(new(RADIUSProfileUpdateRequest))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateRADIUSProfile")
	updateOp.SetTags("RADIUSProfile")
	updateOp.AddRespStructure(new(RADIUSProfileResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/radiusprofile")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListRADIUSProfile")
	listOp.SetTags("RADIUSProfile")
	listOp.AddReqStructure(nil)
	listOp.AddRespStructure(new(RADIUSProfileResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/rest/radiusprofile")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateRADIUSProfile")
	createOp.SetTags("RADIUSProfile")
	createOp.AddReqStructure(new(RADIUSProfile))
	createOp.AddRespStructure(new(RADIUSProfileResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/rest/radiusprofile/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteRADIUSProfile")
	deleteOp.SetTags("RADIUSProfile")
	deleteOp.AddReqStructure(new(RADIUSProfileDeleteRequest))
	deleteOp.AddRespStructure(new(RADIUSProfileResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
