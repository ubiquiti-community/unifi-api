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

type Dashboard struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	ControllerVersion string             `json:"controller_version,omitempty"`
	Desc              string             `json:"desc,omitempty"`
	IsPublic          bool               `json:"is_public"`
	Modules           []DashboardModules `json:"modules,omitempty"`
	Name              string             `json:"name,omitempty"`
}

func (dst *Dashboard) UnmarshalJSON(b []byte) error {
	type Alias Dashboard
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

type DashboardModules struct {
	Config       string `json:"config,omitempty"`
	ID           string `json:"id"`
	ModuleID     string `json:"module_id"`
	Restrictions string `json:"restrictions,omitempty"`
}

func (dst *DashboardModules) UnmarshalJSON(b []byte) error {
	type Alias DashboardModules
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

type DashboardGetRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type DashboardDeleteRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type DashboardUpdateRequest struct {
	*Dashboard
	Site string `path:"site"`
	ID   string `path:"id",json:"_id,omitempty"`
}

type DashboardListRequest struct {
	Site string `path:"site"`
}

type DashboardCreateRequest struct {
	*Dashboard
	Site string `path:"site"`
}

type DashboardResponse struct {
	Meta meta        `json:"meta"`
	Data []Dashboard `json:"data"`
}

func addDashboard() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/dashboard/{id}")
	if err != nil {
		log.Fatal(err)
	}
	getOp.AddReqStructure(new(DashboardGetRequest))
	getOp.AddRespStructure(new(DashboardResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/rest/dashboard/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.AddReqStructure(new(DashboardUpdateRequest))
	updateOp.AddRespStructure(new(DashboardResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/dashboard")
	if err != nil {
		log.Fatal(err)
	}
	listOp.AddReqStructure(new(DashboardListRequest))
	listOp.AddRespStructure(new(DashboardResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{site}/rest/dashboard")
	if err != nil {
		log.Fatal(err)
	}
	createOp.AddReqStructure(new(DashboardCreateRequest))
	createOp.AddRespStructure(new(DashboardResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{site}/get/setting/dashboard/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.AddReqStructure(new(DashboardDeleteRequest))
	deleteOp.AddRespStructure(new(DashboardResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
