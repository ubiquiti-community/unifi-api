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

type RADIUSProfile struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	AccountingEnabled         bool                       `json:"accounting_enabled"`
	AcctServers               []RADIUSProfileAcctServers `json:"acct_servers,omitempty"`
	AuthServers               []RADIUSProfileAuthServers `json:"auth_servers,omitempty"`
	InterimUpdateEnabled      bool                       `json:"interim_update_enabled"`
	InterimUpdateInterval     int                        `json:"interim_update_interval,omitempty"` // ^([6-9][0-9]|[1-9][0-9]{2,3}|[1-7][0-9]{4}|8[0-5][0-9]{3}|86[0-3][0-9][0-9]|86400)$
	Name                      string                     `json:"name,omitempty"`                    // .{1,128}
	TlsEnabled                bool                       `json:"tls_enabled"`
	UseUsgAcctServer          bool                       `json:"use_usg_acct_server"`
	UseUsgAuthServer          bool                       `json:"use_usg_auth_server"`
	VLANEnabled               bool                       `json:"vlan_enabled"`
	VLANWLANMode              string                     `json:"vlan_wlan_mode,omitempty"` // disabled|optional|required
	XCaCrts                   []RADIUSProfileXCaCrts     `json:"x_ca_crts,omitempty"`
	XClientCrt                string                     `json:"x_client_crt,omitempty"`
	XClientCrtFilename        string                     `json:"x_client_crt_filename,omitempty"`
	XClientPrivateKey         string                     `json:"x_client_private_key,omitempty"`
	XClientPrivateKeyFilename string                     `json:"x_client_private_key_filename,omitempty"`
	XClientPrivateKeyPassword string                     `json:"x_client_private_key_password,omitempty"`
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

type RADIUSProfileXCaCrts struct {
	Filename string `json:"filename,omitempty"`
	XCaCrt   string `json:"x_ca_crt,omitempty"`
}

func (dst *RADIUSProfileXCaCrts) UnmarshalJSON(b []byte) error {
	type Alias RADIUSProfileXCaCrts
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

type RADIUSProfileGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type RADIUSProfileListRequest struct {
	SiteID string `path:"siteId"`
}

type RADIUSProfileCreateRequest struct {
	*RADIUSProfile
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type RADIUSProfileDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type RADIUSProfileUpdateRequest struct {
	*RADIUSProfile
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type RADIUSProfileResponse struct {
	Meta meta            `json:"meta"`
	Data []RADIUSProfile `json:"data"`
}

func addRADIUSProfile() {
	resourceName := strcase.SnakeCase("RADIUSProfile")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/radiusprofile/{id}")
	getOp.AddReqStructure(new(RADIUSProfileGetRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/rest/radiusprofile/{id}",
			"method": "GET",
		},
	}
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

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/radiusprofile/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateRADIUSProfile")
	updateOp.SetTags("RADIUSProfile")
	updateOp.AddReqStructure(new(RADIUSProfileUpdateRequest))

	updateOp.AddRespStructure(new(RADIUSProfileResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/radiusprofile")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListRADIUSProfile")
	listOp.SetTags("RADIUSProfile")
	listOp.AddReqStructure(new(RADIUSProfileListRequest))

	listOp.AddRespStructure(new(RADIUSProfileResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/radiusprofile")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateRADIUSProfile")
	createOp.SetTags("RADIUSProfile")
	createOp.AddReqStructure(new(RADIUSProfileCreateRequest))

	getOp.AddRespStructure(new(RADIUSProfileResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/radiusprofile/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteRADIUSProfile")
	deleteOp.SetTags("RADIUSProfile")
	deleteOp.AddReqStructure(new(RADIUSProfileDeleteRequest))

	deleteOp.AddRespStructure(new(RADIUSProfileResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
