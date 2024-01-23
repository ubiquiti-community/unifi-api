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

type ScheduleTask struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Action          string                       `json:"action,omitempty"` // upgrade
	CronExpr        string                       `json:"cron_expr,omitempty"`
	ExecuteOnlyOnce bool                         `json:"execute_only_once"`
	Name            string                       `json:"name,omitempty"`
	UpgradeTargets  []ScheduleTaskUpgradeTargets `json:"upgrade_targets,omitempty"`
}

func (dst *ScheduleTask) UnmarshalJSON(b []byte) error {
	type Alias ScheduleTask
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

type ScheduleTaskUpgradeTargets struct {
	MAC string `json:"mac,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
}

func (dst *ScheduleTaskUpgradeTargets) UnmarshalJSON(b []byte) error {
	type Alias ScheduleTaskUpgradeTargets
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

type ScheduleTaskGetRequest struct {
	ID string `path:"id"`
}

type ScheduleTaskDeleteRequest struct {
	ID string `path:"id"`
}

type ScheduleTaskUpdateRequest struct {
	*ScheduleTask
	ID string `path:"id",json:"_id,omitempty"`
}

type ScheduleTaskResponse struct {
	Meta meta           `json:"meta"`
	Data []ScheduleTask `json:"data"`
}

func addScheduleTask() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/scheduletask/{id}")
	getOp.AddReqStructure(new(ScheduleTaskGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetScheduleTask")
	getOp.SetTags("ScheduleTask")
	getOp.AddRespStructure(new(ScheduleTaskResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/rest/scheduletask/{id}")
	updateOp.AddReqStructure(new(ScheduleTaskUpdateRequest))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateScheduleTask")
	updateOp.SetTags("ScheduleTask")

	updateOp.AddRespStructure(new(ScheduleTaskResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/rest/scheduletask")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListScheduleTask")
	listOp.SetTags("ScheduleTask")
	listOp.AddReqStructure(nil)

	listOp.AddRespStructure(new(ScheduleTaskResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/rest/scheduletask")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateScheduleTask")
	createOp.SetTags("ScheduleTask")
	createOp.AddReqStructure(new(ScheduleTask))

	getOp.AddRespStructure(new(ScheduleTaskResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/rest/scheduletask/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteScheduleTask")
	deleteOp.SetTags("ScheduleTask")
	deleteOp.AddReqStructure(new(ScheduleTaskDeleteRequest))

	deleteOp.AddRespStructure(new(ScheduleTaskResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
