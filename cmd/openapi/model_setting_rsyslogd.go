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

type SettingRsyslogd struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Contents                    []string `json:"contents,omitempty"` // device|client|firewall_default_policy|triggers|updates|admin_activity|critical|security_detections|vpn
	Debug                       bool     `json:"debug"`
	Enabled                     bool     `json:"enabled"`
	IP                          string   `json:"ip,omitempty"`
	LogAllContents              bool     `json:"log_all_contents"`
	NetconsoleEnabled           bool     `json:"netconsole_enabled"`
	NetconsoleHost              string   `json:"netconsole_host,omitempty"`
	NetconsolePort              int      `json:"netconsole_port,omitempty"` // [1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]
	Port                        int      `json:"port,omitempty"`            // [1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]
	ThisController              bool     `json:"this_controller"`
	ThisControllerEncryptedOnly bool     `json:"this_controller_encrypted_only"`
}

func (dst *SettingRsyslogd) UnmarshalJSON(b []byte) error {
	type Alias SettingRsyslogd
	aux := &struct {
		NetconsolePort emptyStringInt `json:"netconsole_port"`
		Port           emptyStringInt `json:"port"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.NetconsolePort = int(aux.NetconsolePort)
	dst.Port = int(aux.Port)

	return nil
}

type SettingRsyslogdUpdateRequest struct {
	*SettingRsyslogd
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type SettingRsyslogdResponse struct {
	Meta meta              `json:"meta"`
	Data []SettingRsyslogd `json:"data"`
}

func addSettingRsyslogd() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/get/setting/rsyslogd")
	getOp.AddReqStructure(new(SiteRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingRsyslogd")
	getOp.SetTags("SettingRsyslogd")
	getOp.AddRespStructure(new(SettingRsyslogdResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/set/setting/rsyslogd")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingRsyslogd")
	updateOp.SetTags("SettingRsyslogd")
	updateOp.AddReqStructure(new(SettingRsyslogdUpdateRequest))

	updateOp.AddRespStructure(new(SettingRsyslogdResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
