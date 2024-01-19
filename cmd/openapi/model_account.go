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

type Account struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	IP               string `json:"ip,omitempty"`   // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$
	Name             string `json:"name,omitempty"` // ^[^"' ]+$
	NetworkID        string `json:"networkconf_id,omitempty"`
	TunnelConfigType string `json:"tunnel_config_type,omitempty"` // vpn|802.1x|custom
	TunnelMediumType int    `json:"tunnel_medium_type,omitempty"` // [1-9]|1[0-5]|^$
	TunnelType       int    `json:"tunnel_type,omitempty"`        // [1-9]|1[0-3]|^$
	VLAN             int    `json:"vlan,omitempty"`               // [2-9]|[1-9][0-9]{1,2}|[1-3][0-9]{3}|400[0-9]|^$
	XPassword        string `json:"x_password,omitempty"`
}

func (dst *Account) UnmarshalJSON(b []byte) error {
	type Alias Account
	aux := &struct {
		TunnelMediumType emptyStringInt `json:"tunnel_medium_type"`
		TunnelType       emptyStringInt `json:"tunnel_type"`
		VLAN             emptyStringInt `json:"vlan"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.TunnelMediumType = int(aux.TunnelMediumType)
	dst.TunnelType = int(aux.TunnelType)
	dst.VLAN = int(aux.VLAN)

	return nil
}

type AccountGetRequest struct {
	ID string `path:"id"`
}

type AccountDeleteRequest struct {
	ID string `path:"id"`
}

type AccountUpdateRequest struct {
	*Account
	ID string `path:"id",json:"_id,omitempty"`
}

type AccountResponse struct {
	Meta meta      `json:"meta"`
	Data []Account `json:"data"`
}

func addAccount() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/account/{id}")
	getOp.AddReqStructure(new(AccountGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetAccount")
	getOp.SetTags("Account")
	getOp.AddRespStructure(new(AccountResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/rest/account/{id}")
	updateOp.AddReqStructure(new(AccountUpdateRequest))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateAccount")
	updateOp.SetTags("Account")
	updateOp.AddRespStructure(new(AccountResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/account")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListAccount")
	listOp.SetTags("Account")
	listOp.AddReqStructure(nil)
	listOp.AddRespStructure(new(AccountResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/rest/account")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateAccount")
	createOp.SetTags("Account")
	createOp.AddReqStructure(new(Account))
	createOp.AddRespStructure(new(AccountResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/rest/account/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteAccount")
	deleteOp.SetTags("Account")
	deleteOp.AddReqStructure(new(AccountDeleteRequest))
	deleteOp.AddRespStructure(new(AccountResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
