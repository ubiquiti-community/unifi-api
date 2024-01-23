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

type SettingMgmt struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	AdvancedFeatureEnabled  bool                  `json:"advanced_feature_enabled"`
	AlertEnabled            bool                  `json:"alert_enabled"`
	AutoUpgrade             bool                  `json:"auto_upgrade"`
	AutoUpgradeHour         int                   `json:"auto_upgrade_hour,omitempty"` // [0-9]|1[0-9]|2[0-3]|^$
	BootSound               bool                  `json:"boot_sound"`
	DebugToolsEnabled       bool                  `json:"debug_tools_enabled"`
	DirectConnectEnabled    bool                  `json:"direct_connect_enabled"`
	LedEnabled              bool                  `json:"led_enabled"`
	OutdoorModeEnabled      bool                  `json:"outdoor_mode_enabled"`
	UnifiIDpEnabled         bool                  `json:"unifi_idp_enabled"`
	WifimanEnabled          bool                  `json:"wifiman_enabled"`
	XMgmtKey                string                `json:"x_mgmt_key,omitempty"` // [0-9a-f]{32}
	XSshAuthPasswordEnabled bool                  `json:"x_ssh_auth_password_enabled"`
	XSshBindWildcard        bool                  `json:"x_ssh_bind_wildcard"`
	XSshEnabled             bool                  `json:"x_ssh_enabled"`
	XSshKeys                []SettingMgmtXSshKeys `json:"x_ssh_keys,omitempty"`
	XSshMd5Passwd           string                `json:"x_ssh_md5passwd,omitempty"`
	XSshPassword            string                `json:"x_ssh_password,omitempty"` // .{1,128}
	XSshSha512Passwd        string                `json:"x_ssh_sha512passwd,omitempty"`
	XSshUsername            string                `json:"x_ssh_username,omitempty"` // ^[_A-Za-z0-9][-_.A-Za-z0-9]{0,29}$
}

func (dst *SettingMgmt) UnmarshalJSON(b []byte) error {
	type Alias SettingMgmt
	aux := &struct {
		AutoUpgradeHour emptyStringInt `json:"auto_upgrade_hour"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.AutoUpgradeHour = int(aux.AutoUpgradeHour)

	return nil
}

type SettingMgmtXSshKeys struct {
	Comment     string `json:"comment"`
	Date        string `json:"date"`
	Fingerprint string `json:"fingerprint"`
	Key         string `json:"key"`
	KeyType     string `json:"type"`
	Name        string `json:"name"`
}

func (dst *SettingMgmtXSshKeys) UnmarshalJSON(b []byte) error {
	type Alias SettingMgmtXSshKeys
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

type SettingMgmtResponse struct {
	Meta meta          `json:"meta"`
	Data []SettingMgmt `json:"data"`
}

func addSettingMgmt() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/get/setting/mgmt")
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetSettingMgmt")
	getOp.SetTags("SettingMgmt")
	getOp.AddRespStructure(new(SettingMgmtResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/set/setting/mgmt")
	updateOp.AddReqStructure(new(SettingMgmt))
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateSettingMgmt")
	updateOp.SetTags("SettingMgmt")

	updateOp.AddRespStructure(new(SettingMgmtResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

}
