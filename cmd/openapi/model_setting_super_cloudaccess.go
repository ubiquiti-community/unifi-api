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

type SettingSuperCloudaccess struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	DeviceAuth      string `json:"device_auth,omitempty"`
	DeviceID        string `json:"device_id"`
	Enabled         bool   `json:"enabled"`
	UbicUuid        string `json:"ubic_uuid,omitempty"`
	XCertificateArn string `json:"x_certificate_arn,omitempty"`
	XCertificatePem string `json:"x_certificate_pem,omitempty"`
	XPrivateKey     string `json:"x_private_key,omitempty"`
}

func (dst *SettingSuperCloudaccess) UnmarshalJSON(b []byte) error {
	type Alias SettingSuperCloudaccess
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

type SettingSuperCloudaccessUpdateRequest struct {
	*SettingSuperCloudaccess
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type SettingSuperCloudaccessResponse struct {
	Meta meta                      `json:"meta"`
	Data []SettingSuperCloudaccess `json:"data"`
}

func addSettingSuperCloudaccess() {
	resourceName := strcase.SnakeCase("SettingSuperCloudaccess")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/get/setting/super_cloudaccess")
	getOp.AddReqStructure(new(SiteRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/get/setting/super_cloudaccess",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingSuperCloudaccess")
	getOp.SetTags("SettingSuperCloudaccess")
	getOp.AddRespStructure(new(SettingSuperCloudaccessResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/set/setting/super_cloudaccess")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingSuperCloudaccess")
	updateOp.SetTags("SettingSuperCloudaccess")
	updateOp.AddReqStructure(new(SettingSuperCloudaccessUpdateRequest))

	updateOp.AddRespStructure(new(SettingSuperCloudaccessResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
