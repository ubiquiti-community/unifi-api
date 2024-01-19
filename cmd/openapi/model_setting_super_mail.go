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

type SettingSuperMail struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Provider string `json:"provider,omitempty"` // smtp|cloud|disabled
}

func (dst *SettingSuperMail) UnmarshalJSON(b []byte) error {
	type Alias SettingSuperMail
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

type SettingSuperMailGetRequest struct {
	Site string `path:"site"`
}

type SettingSuperMailUpdateRequest struct {
	*SettingSuperMail
	Site string `path:"site"`
}

type SettingSuperMailResponse struct {
	Meta meta               `json:"meta"`
	Data []SettingSuperMail `json:"data"`
}

func addSettingSuperMail() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{site}/get/setting/super_mail")
	if err != nil {
		log.Fatal(err)
	}
	getOp.AddReqStructure(new(SettingSuperMailGetRequest))
	getOp.AddRespStructure(new(SettingSuperMailResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{site}/set/setting/super_mail")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.AddReqStructure(new(SettingSuperMailUpdateRequest))
	updateOp.AddRespStructure(new(SettingSuperMailResponse), openapi.WithHTTPStatus(http.StatusCreated))
	updateOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
