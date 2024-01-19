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

type HeatMapPoint struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	DownloadSpeed float64 `json:"download_speed,omitempty"`
	HeatmapID     string  `json:"heatmap_id"`
	UploadSpeed   float64 `json:"upload_speed,omitempty"`
	X             float64 `json:"x,omitempty"`
	Y             float64 `json:"y,omitempty"`
}

func (dst *HeatMapPoint) UnmarshalJSON(b []byte) error {
	type Alias HeatMapPoint
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

type HeatMapPointGetRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type HeatMapPointDeleteRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type HeatMapPointUpdateRequest struct {
	*HeatMapPoint
	Site string `path:"site"`
	ID   string `path:"id",json:"_id,omitempty"`
}

type HeatMapPointListRequest struct {
	Site string `path:"site"`
}

type HeatMapPointCreateRequest struct {
	*HeatMapPoint
	Site string `path:"site"`
}

type HeatMapPointResponse struct {
	Meta meta           `json:"meta"`
	Data []HeatMapPoint `json:"data"`
}

func addHeatMapPoint() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/heatmappoint/{id}")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetHeatMapPoint")
	getOp.SetTags("HeatMapPoint")
	getOp.AddReqStructure(new(HeatMapPointGetRequest))
	getOp.AddRespStructure(new(HeatMapPointResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/rest/heatmappoint/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateHeatMapPoint")
	updateOp.SetTags("HeatMapPoint")
	updateOp.AddReqStructure(new(HeatMapPointUpdateRequest))
	updateOp.AddRespStructure(new(HeatMapPointResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/heatmappoint")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListHeatMapPoint")
	listOp.SetTags("HeatMapPoint")
	listOp.AddReqStructure(new(HeatMapPointListRequest))
	listOp.AddRespStructure(new(HeatMapPointResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{site}/rest/heatmappoint")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateHeatMapPoint")
	createOp.SetTags("HeatMapPoint")
	createOp.AddReqStructure(new(HeatMapPointCreateRequest))
	createOp.AddRespStructure(new(HeatMapPointResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{site}/get/setting/heatmappoint/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteHeatMapPoint")
	deleteOp.SetTags("HeatMapPoint")
	deleteOp.AddReqStructure(new(HeatMapPointDeleteRequest))
	deleteOp.AddRespStructure(new(HeatMapPointResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
