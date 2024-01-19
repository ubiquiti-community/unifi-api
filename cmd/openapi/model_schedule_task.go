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
	Site string `path:"site"`
	ID   string `path:"id"`
}

type ScheduleTaskDeleteRequest struct {
	Site string `path:"site"`
	ID   string `path:"id"`
}

type ScheduleTaskUpdateRequest struct {
	*ScheduleTask
	Site string `path:"site"`
	ID   string `path:"id",json:"_id,omitempty"`
}

type ScheduleTaskListRequest struct {
	Site string `path:"site"`
}

type ScheduleTaskCreateRequest struct {
	*ScheduleTask
	Site string `path:"site"`
}

type ScheduleTaskResponse struct {
	Meta meta           `json:"meta"`
	Data []ScheduleTask `json:"data"`
}

func addScheduleTask() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/scheduletask/{id}")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetScheduleTask")
	getOp.SetTags("ScheduleTask")
	getOp.AddReqStructure(new(ScheduleTaskGetRequest))
	getOp.AddRespStructure(new(ScheduleTaskResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/rest/scheduletask/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateScheduleTask")
	updateOp.SetTags("ScheduleTask")
	updateOp.AddReqStructure(new(ScheduleTaskUpdateRequest))
	updateOp.AddRespStructure(new(ScheduleTaskResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/rest/scheduletask")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListScheduleTask")
	listOp.SetTags("ScheduleTask")
	listOp.AddReqStructure(new(ScheduleTaskListRequest))
	listOp.AddRespStructure(new(ScheduleTaskResponse), openapi.WithHTTPStatus(http.StatusOK))
	listOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{site}/rest/scheduletask")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateScheduleTask")
	createOp.SetTags("ScheduleTask")
	createOp.AddReqStructure(new(ScheduleTaskCreateRequest))
	createOp.AddRespStructure(new(ScheduleTaskResponse), openapi.WithHTTPStatus(http.StatusOK))
	createOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{site}/get/setting/scheduletask/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteScheduleTask")
	deleteOp.SetTags("ScheduleTask")
	deleteOp.AddReqStructure(new(ScheduleTaskDeleteRequest))
	deleteOp.AddRespStructure(new(ScheduleTaskResponse), openapi.WithHTTPStatus(http.StatusOK))
	deleteOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
