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

type DpiApp struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Apps           []int  `json:"apps,omitempty"`
	Blocked        bool   `json:"blocked"`
	Cats           []int  `json:"cats,omitempty"`
	Enabled        bool   `json:"enabled"`
	Log            bool   `json:"log"`
	Name           string `json:"name,omitempty"`              // .{1,128}
	QOSRateMaxDown int    `json:"qos_rate_max_down,omitempty"` // -1|[2-9]|[1-9][0-9]{1,4}|100000|10[0-1][0-9]{3}|102[0-3][0-9]{2}|102400
	QOSRateMaxUp   int    `json:"qos_rate_max_up,omitempty"`   // -1|[2-9]|[1-9][0-9]{1,4}|100000|10[0-1][0-9]{3}|102[0-3][0-9]{2}|102400
}

func (dst *DpiApp) UnmarshalJSON(b []byte) error {
	type Alias DpiApp
	aux := &struct {
		Apps           []emptyStringInt `json:"apps"`
		Cats           []emptyStringInt `json:"cats"`
		QOSRateMaxDown emptyStringInt   `json:"qos_rate_max_down"`
		QOSRateMaxUp   emptyStringInt   `json:"qos_rate_max_up"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Apps = make([]int, len(aux.Apps))
	for i, v := range aux.Apps {
		dst.Apps[i] = int(v)
	}
	dst.Cats = make([]int, len(aux.Cats))
	for i, v := range aux.Cats {
		dst.Cats[i] = int(v)
	}
	dst.QOSRateMaxDown = int(aux.QOSRateMaxDown)
	dst.QOSRateMaxUp = int(aux.QOSRateMaxUp)

	return nil
}

type DpiAppGetRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type DpiAppDeleteRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type DpiAppUpdateRequest struct {
	*DpiApp
	Site string `path:"site"`
	ID   string `path:"id",json:"_id,omitempty"`
}

type DpiAppListRequest struct {
	Site string `path:"site"`
}

type DpiAppCreateRequest struct {
	*DpiApp
	Site string `path:"site"`
}

type DpiAppResponse struct {
	Meta meta     `json:"meta"`
	Data []DpiApp `json:"data"`
}

func addDpiApp() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/dpiapp/{id}")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetDpiApp")
	getOp.SetTags("DpiApp")
	getOp.AddReqStructure(new(DpiAppGetRequest))
	getOp.AddRespStructure(new(DpiAppResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/rest/dpiapp/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateDpiApp")
	updateOp.SetTags("DpiApp")
	updateOp.AddReqStructure(new(DpiAppUpdateRequest))
	updateOp.AddRespStructure(new(DpiAppResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/dpiapp")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListDpiApp")
	listOp.SetTags("DpiApp")
	listOp.AddReqStructure(new(DpiAppListRequest))
	listOp.AddRespStructure(new(DpiAppResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{site}/rest/dpiapp")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateDpiApp")
	createOp.SetTags("DpiApp")
	createOp.AddReqStructure(new(DpiAppCreateRequest))
	createOp.AddRespStructure(new(DpiAppResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{site}/get/setting/dpiapp/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteDpiApp")
	deleteOp.SetTags("DpiApp")
	deleteOp.AddReqStructure(new(DpiAppDeleteRequest))
	deleteOp.AddRespStructure(new(DpiAppResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
