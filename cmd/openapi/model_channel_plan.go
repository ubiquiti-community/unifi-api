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

type ChannelPlan struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	ApBlacklistedChannels   []ChannelPlanApBlacklistedChannels   `json:"ap_blacklisted_channels,omitempty"`
	ConfSource              string                               `json:"conf_source,omitempty"` // manual|radio-ai
	Coupling                []ChannelPlanCoupling                `json:"coupling,omitempty"`
	Date                    string                               `json:"date"` // ^$|^(20[0-9]{2}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])Z?$
	Fitness                 float64                              `json:"fitness,omitempty"`
	Note                    string                               `json:"note,omitempty"`  // .{0,1024}
	Radio                   string                               `json:"radio,omitempty"` // na|ng|ng\+na
	RadioTable              []ChannelPlanRadioTable              `json:"radio_table,omitempty"`
	Satisfaction            float64                              `json:"satisfaction,omitempty"`
	SatisfactionTable       []ChannelPlanSatisfactionTable       `json:"satisfaction_table,omitempty"`
	SiteBlacklistedChannels []ChannelPlanSiteBlacklistedChannels `json:"site_blacklisted_channels,omitempty"`
}

func (dst *ChannelPlan) UnmarshalJSON(b []byte) error {
	type Alias ChannelPlan
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

type ChannelPlanApBlacklistedChannels struct {
	Channel   int    `json:"channel,omitempty"`   // 36|38|40|42|44|46|48|52|56|60|64|100|104|108|112|116|120|124|128|132|136|140|144|149|153|157|161|165|183|184|185|187|188|189|192|196
	MAC       string `json:"mac,omitempty"`       // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	Timestamp int    `json:"timestamp,omitempty"` // [1-9][0-9]{12}
}

func (dst *ChannelPlanApBlacklistedChannels) UnmarshalJSON(b []byte) error {
	type Alias ChannelPlanApBlacklistedChannels
	aux := &struct {
		Channel   emptyStringInt `json:"channel"`
		Timestamp emptyStringInt `json:"timestamp"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Channel = int(aux.Channel)
	dst.Timestamp = int(aux.Timestamp)

	return nil
}

type ChannelPlanCoupling struct {
	Rssi   int    `json:"rssi,omitempty"`
	Source string `json:"source,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2}).*$
	Target string `json:"target,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2}).*$
}

func (dst *ChannelPlanCoupling) UnmarshalJSON(b []byte) error {
	type Alias ChannelPlanCoupling
	aux := &struct {
		Rssi emptyStringInt `json:"rssi"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Rssi = int(aux.Rssi)

	return nil
}

type ChannelPlanRadioTable struct {
	BackupChannel string `json:"backup_channel,omitempty"` // [0-9]|[1][0-4]|16|34|36|38|40|42|44|46|48|52|56|60|64|100|104|108|112|116|120|124|128|132|136|140|144|149|153|157|161|165|183|184|185|187|188|189|192|196|auto
	Channel       string `json:"channel,omitempty"`        // [0-9]|[1][0-4]|16|34|36|38|40|42|44|46|48|52|56|60|64|100|104|108|112|116|120|124|128|132|136|140|144|149|153|157|161|165|183|184|185|187|188|189|192|196|auto
	DeviceMAC     string `json:"device_mac,omitempty"`     // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	Name          string `json:"name,omitempty"`           // [a-z]*[0-9]*
	TxPower       string `json:"tx_power,omitempty"`       // [\d]+|auto
	TxPowerMode   string `json:"tx_power_mode,omitempty"`  // auto|medium|high|low|custom
	Width         int    `json:"width,omitempty"`          // 20|40|80|160
}

func (dst *ChannelPlanRadioTable) UnmarshalJSON(b []byte) error {
	type Alias ChannelPlanRadioTable
	aux := &struct {
		BackupChannel numberOrString `json:"backup_channel"`
		Channel       numberOrString `json:"channel"`
		TxPower       numberOrString `json:"tx_power"`
		Width         emptyStringInt `json:"width"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.BackupChannel = string(aux.BackupChannel)
	dst.Channel = string(aux.Channel)
	dst.TxPower = string(aux.TxPower)
	dst.Width = int(aux.Width)

	return nil
}

type ChannelPlanSatisfactionTable struct {
	DeviceMAC    string  `json:"device_mac,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	Satisfaction float64 `json:"satisfaction,omitempty"`
}

func (dst *ChannelPlanSatisfactionTable) UnmarshalJSON(b []byte) error {
	type Alias ChannelPlanSatisfactionTable
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

type ChannelPlanSiteBlacklistedChannels struct {
	Channel   int `json:"channel,omitempty"`   // 36|38|40|42|44|46|48|52|56|60|64|100|104|108|112|116|120|124|128|132|136|140|144|149|153|157|161|165|183|184|185|187|188|189|192|196
	Timestamp int `json:"timestamp,omitempty"` // [1-9][0-9]{12}
}

func (dst *ChannelPlanSiteBlacklistedChannels) UnmarshalJSON(b []byte) error {
	type Alias ChannelPlanSiteBlacklistedChannels
	aux := &struct {
		Channel   emptyStringInt `json:"channel"`
		Timestamp emptyStringInt `json:"timestamp"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Channel = int(aux.Channel)
	dst.Timestamp = int(aux.Timestamp)

	return nil
}

type ChannelPlanGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type ChannelPlanListRequest struct {
	SiteID string `path:"siteId"`
}

type ChannelPlanCreateRequest struct {
	*ChannelPlan
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type ChannelPlanDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type ChannelPlanUpdateRequest struct {
	*ChannelPlan
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type ChannelPlanResponse struct {
	Meta meta          `json:"meta"`
	Data []ChannelPlan `json:"data"`
}

func addChannelPlan() {
	resourceName := strcase.SnakeCase("ChannelPlan")

	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/channelplan/{id}")
	getOp.AddReqStructure(new(ChannelPlanGetRequest))
	generatorConfig.DataSources[resourceName] = map[string]any{
		"read": map[string]any{
			"path":   "/s/{siteId}/rest/channelplan/{id}",
			"method": "GET",
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetChannelPlan")
	getOp.SetTags("ChannelPlan")
	getOp.AddRespStructure(new(ChannelPlanResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/channelplan/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateChannelPlan")
	updateOp.SetTags("ChannelPlan")
	updateOp.AddReqStructure(new(ChannelPlanUpdateRequest))

	updateOp.AddRespStructure(new(ChannelPlanResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/channelplan")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListChannelPlan")
	listOp.SetTags("ChannelPlan")
	listOp.AddReqStructure(new(ChannelPlanListRequest))

	listOp.AddRespStructure(new(ChannelPlanResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/channelplan")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateChannelPlan")
	createOp.SetTags("ChannelPlan")
	createOp.AddReqStructure(new(ChannelPlanCreateRequest))

	getOp.AddRespStructure(new(ChannelPlanResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/channelplan/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteChannelPlan")
	deleteOp.SetTags("ChannelPlan")
	deleteOp.AddReqStructure(new(ChannelPlanDeleteRequest))

	deleteOp.AddRespStructure(new(ChannelPlanResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
