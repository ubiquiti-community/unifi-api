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

type HotspotPackage struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Amount                         float64 `json:"amount,omitempty"`
	ChargedAs                      string  `json:"charged_as,omitempty"`
	Currency                       string  `json:"currency,omitempty"` // [A-Z]{3}
	CustomPaymentFieldsEnabled     bool    `json:"custom_payment_fields_enabled"`
	Hours                          int     `json:"hours,omitempty"`
	Index                          int     `json:"index,omitempty"`
	LimitDown                      int     `json:"limit_down,omitempty"`
	LimitOverwrite                 bool    `json:"limit_overwrite"`
	LimitQuota                     int     `json:"limit_quota,omitempty"`
	LimitUp                        int     `json:"limit_up,omitempty"`
	Name                           string  `json:"name,omitempty"`
	PaymentFieldsAddressEnabled    bool    `json:"payment_fields_address_enabled"`
	PaymentFieldsAddressRequired   bool    `json:"payment_fields_address_required"`
	PaymentFieldsCityEnabled       bool    `json:"payment_fields_city_enabled"`
	PaymentFieldsCityRequired      bool    `json:"payment_fields_city_required"`
	PaymentFieldsCountryEnabled    bool    `json:"payment_fields_country_enabled"`
	PaymentFieldsCountryRequired   bool    `json:"payment_fields_country_required"`
	PaymentFieldsEmailEnabled      bool    `json:"payment_fields_email_enabled"`
	PaymentFieldsEmailRequired     bool    `json:"payment_fields_email_required"`
	PaymentFieldsFirstNameEnabled  bool    `json:"payment_fields_first_name_enabled"`
	PaymentFieldsFirstNameRequired bool    `json:"payment_fields_first_name_required"`
	PaymentFieldsLastNameEnabled   bool    `json:"payment_fields_last_name_enabled"`
	PaymentFieldsLastNameRequired  bool    `json:"payment_fields_last_name_required"`
	PaymentFieldsStateEnabled      bool    `json:"payment_fields_state_enabled"`
	PaymentFieldsStateRequired     bool    `json:"payment_fields_state_required"`
	PaymentFieldsZipEnabled        bool    `json:"payment_fields_zip_enabled"`
	PaymentFieldsZipRequired       bool    `json:"payment_fields_zip_required"`
	TrialDurationMinutes           int     `json:"trial_duration_minutes,omitempty"`
	TrialReset                     float64 `json:"trial_reset,omitempty"`
}

func (dst *HotspotPackage) UnmarshalJSON(b []byte) error {
	type Alias HotspotPackage
	aux := &struct {
		Hours                emptyStringInt `json:"hours"`
		Index                emptyStringInt `json:"index"`
		LimitDown            emptyStringInt `json:"limit_down"`
		LimitQuota           emptyStringInt `json:"limit_quota"`
		LimitUp              emptyStringInt `json:"limit_up"`
		TrialDurationMinutes emptyStringInt `json:"trial_duration_minutes"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Hours = int(aux.Hours)
	dst.Index = int(aux.Index)
	dst.LimitDown = int(aux.LimitDown)
	dst.LimitQuota = int(aux.LimitQuota)
	dst.LimitUp = int(aux.LimitUp)
	dst.TrialDurationMinutes = int(aux.TrialDurationMinutes)

	return nil
}

type HotspotPackageGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type HotspotPackageListRequest struct {
	SiteID string `path:"siteId"`
}

type HotspotPackageCreateRequest struct {
	*HotspotPackage
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type HotspotPackageDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type HotspotPackageUpdateRequest struct {
	*HotspotPackage
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type HotspotPackageResponse struct {
	Meta meta             `json:"meta"`
	Data []HotspotPackage `json:"data"`
}

func addHotspotPackage() {
	resourceName := strcase.SnakeCase("HotspotPackage")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/hotspotpackage/{id}")
	getOp.AddReqStructure(new(HotspotPackageGetRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/rest/hotspotpackage/{id}",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetHotspotPackage")
	getOp.SetTags("HotspotPackage")
	getOp.AddRespStructure(new(HotspotPackageResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/hotspotpackage/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateHotspotPackage")
	updateOp.SetTags("HotspotPackage")
	updateOp.AddReqStructure(new(HotspotPackageUpdateRequest))

	updateOp.AddRespStructure(new(HotspotPackageResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/hotspotpackage")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListHotspotPackage")
	listOp.SetTags("HotspotPackage")
	listOp.AddReqStructure(new(HotspotPackageListRequest))

	listOp.AddRespStructure(new(HotspotPackageResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/hotspotpackage")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateHotspotPackage")
	createOp.SetTags("HotspotPackage")
	createOp.AddReqStructure(new(HotspotPackageCreateRequest))

	getOp.AddRespStructure(new(HotspotPackageResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/hotspotpackage/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteHotspotPackage")
	deleteOp.SetTags("HotspotPackage")
	deleteOp.AddReqStructure(new(HotspotPackageDeleteRequest))

	deleteOp.AddRespStructure(new(HotspotPackageResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
