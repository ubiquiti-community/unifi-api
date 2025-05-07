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

type FirewallRule struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Action                string   `json:"action,omitempty"` // drop|reject|accept
	DstAddress            string   `json:"dst_address,omitempty"`
	DstAddressIPV6        string   `json:"dst_address_ipv6,omitempty"`
	DstFirewallGroupIDs   []string `json:"dst_firewallgroup_ids,omitempty"` // [\d\w]+
	DstNetworkID          string   `json:"dst_networkconf_id"`              // [\d\w]+|^$
	DstNetworkType        string   `json:"dst_networkconf_type,omitempty"`  // ADDRv4|NETv4
	DstPort               string   `json:"dst_port,omitempty"`
	Enabled               bool     `json:"enabled"`
	ICMPTypename          string   `json:"icmp_typename"`   // ^$|address-mask-reply|address-mask-request|any|communication-prohibited|destination-unreachable|echo-reply|echo-request|fragmentation-needed|host-precedence-violation|host-prohibited|host-redirect|host-unknown|host-unreachable|ip-header-bad|network-prohibited|network-redirect|network-unknown|network-unreachable|parameter-problem|port-unreachable|precedence-cutoff|protocol-unreachable|redirect|required-option-missing|router-advertisement|router-solicitation|source-quench|source-route-failed|time-exceeded|timestamp-reply|timestamp-request|TOS-host-redirect|TOS-host-unreachable|TOS-network-redirect|TOS-network-unreachable|ttl-zero-during-reassembly|ttl-zero-during-transit
	ICMPv6Typename        string   `json:"icmpv6_typename"` // ^$|address-unreachable|bad-header|beyond-scope|communication-prohibited|destination-unreachable|echo-reply|echo-request|failed-policy|neighbor-advertisement|neighbor-solicitation|no-route|packet-too-big|parameter-problem|port-unreachable|redirect|reject-route|router-advertisement|router-solicitation|time-exceeded|ttl-zero-during-reassembly|ttl-zero-during-transit|unknown-header-type|unknown-option
	IPSec                 string   `json:"ipsec"`           // match-ipsec|match-none|^$
	Logging               bool     `json:"logging"`
	Name                  string   `json:"name,omitempty"` // .{1,128}
	Protocol              string   `json:"protocol"`       // ^$|all|([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])|tcp_udp|ah|ax.25|dccp|ddp|egp|eigrp|encap|esp|etherip|fc|ggp|gre|hip|hmp|icmp|idpr-cmtp|idrp|igmp|igp|ip|ipcomp|ipencap|ipip|ipv6|ipv6-frag|ipv6-icmp|ipv6-nonxt|ipv6-opts|ipv6-route|isis|iso-tp4|l2tp|manet|mobility-header|mpls-in-ip|ospf|pim|pup|rdp|rohc|rspf|rsvp|sctp|shim6|skip|st|tcp|udp|udplite|vmtp|vrrp|wesp|xns-idp|xtp
	ProtocolMatchExcepted bool     `json:"protocol_match_excepted"`
	ProtocolV6            string   `json:"protocol_v6"`                  // ^$|([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])|ah|all|dccp|eigrp|esp|gre|icmpv6|ipcomp|ipv6|ipv6-frag|ipv6-icmp|ipv6-nonxt|ipv6-opts|ipv6-route|isis|l2tp|manet|mobility-header|mpls-in-ip|ospf|pim|rsvp|sctp|shim6|tcp|tcp_udp|udp|vrrp
	RuleIndex             int      `json:"rule_index,omitempty"`         // 2[0-9]{3,4}|4[0-9]{3,4}
	Ruleset               string   `json:"ruleset,omitempty"`            // WAN_IN|WAN_OUT|WAN_LOCAL|LAN_IN|LAN_OUT|LAN_LOCAL|GUEST_IN|GUEST_OUT|GUEST_LOCAL|WANv6_IN|WANv6_OUT|WANv6_LOCAL|LANv6_IN|LANv6_OUT|LANv6_LOCAL|GUESTv6_IN|GUESTv6_OUT|GUESTv6_LOCAL
	SettingPreference     string   `json:"setting_preference,omitempty"` // auto|manual
	SrcAddress            string   `json:"src_address,omitempty"`
	SrcAddressIPV6        string   `json:"src_address_ipv6,omitempty"`
	SrcFirewallGroupIDs   []string `json:"src_firewallgroup_ids,omitempty"` // [\d\w]+
	SrcMACAddress         string   `json:"src_mac_address"`                 // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$|^$
	SrcNetworkID          string   `json:"src_networkconf_id"`              // [\d\w]+|^$
	SrcNetworkType        string   `json:"src_networkconf_type,omitempty"`  // ADDRv4|NETv4
	SrcPort               string   `json:"src_port,omitempty"`
	StateEstablished      bool     `json:"state_established"`
	StateInvalid          bool     `json:"state_invalid"`
	StateNew              bool     `json:"state_new"`
	StateRelated          bool     `json:"state_related"`
}

func (dst *FirewallRule) UnmarshalJSON(b []byte) error {
	type Alias FirewallRule
	aux := &struct {
		RuleIndex emptyStringInt `json:"rule_index"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.RuleIndex = int(aux.RuleIndex)

	return nil
}

type FirewallRuleGetRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type FirewallRuleListRequest struct {
	SiteID string `path:"siteId"`
}

type FirewallRuleCreateRequest struct {
	*FirewallRule
	SiteID string `path:"siteId" json:"site_id,omitempty"`
}

type FirewallRuleDeleteRequest struct {
	SiteID string `path:"siteId"`
	ID     string `path:"id"`
}

type FirewallRuleUpdateRequest struct {
	*FirewallRule
	SiteID string `path:"siteId" json:"site_id,omitempty"`
	ID     string `path:"id" json:"_id,omitempty"`
}

type FirewallRuleResponse struct {
	Meta meta           `json:"meta"`
	Data []FirewallRule `json:"data"`
}

func addFirewallRule() {
	// Get

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/firewallrule/{id}")
	getOp.AddReqStructure(new(FirewallRuleGetRequest))
	if err != nil {
		log.Fatal(err)
	}
	getOp.SetID("GetFirewallRule")
	getOp.SetTags("FirewallRule")
	getOp.AddRespStructure(new(FirewallRuleResponse), openapi.WithHTTPStatus(http.StatusOK))
	getOp.AddRespStructure(ErrorResponse, func(cu *openapi.ContentUnit) {
		cu.IsDefault = true
	})
	err = reflector.AddOperation(getOp)
	if err != nil {
		log.Fatal(err)
	}

	// Update

	updateOp, err := reflector.NewOperationContext(http.MethodPut, "/s/{siteId}/rest/firewallrule/{id}")
	if err != nil {
		log.Fatal(err)
	}
	updateOp.SetID("UpdateFirewallRule")
	updateOp.SetTags("FirewallRule")
	updateOp.AddReqStructure(new(FirewallRuleUpdateRequest))

	updateOp.AddRespStructure(new(FirewallRuleResponse), openapi.WithHTTPStatus(http.StatusCreated), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		updateOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(updateOp)
	if err != nil {
		log.Fatal(err)
	}

	// List
	listOp, err := reflector.NewOperationContext(http.MethodGet, "/s/{siteId}/rest/firewallrule")
	if err != nil {
		log.Fatal(err)
	}
	listOp.SetID("ListFirewallRule")
	listOp.SetTags("FirewallRule")
	listOp.AddReqStructure(new(FirewallRuleListRequest))

	listOp.AddRespStructure(new(FirewallRuleResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		listOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(listOp)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	createOp, err := reflector.NewOperationContext(http.MethodPost, "/s/{siteId}/rest/firewallrule")
	if err != nil {
		log.Fatal(err)
	}
	createOp.SetID("CreateFirewallRule")
	createOp.SetTags("FirewallRule")
	createOp.AddReqStructure(new(FirewallRuleCreateRequest))

	getOp.AddRespStructure(new(FirewallRuleResponse), openapi.WithContentType("application/json"), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		getOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(createOp)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, "/s/{siteId}/rest/firewallrule/{id}")
	if err != nil {
		log.Fatal(err)
	}
	deleteOp.SetID("DeleteFirewallRule")
	deleteOp.SetTags("FirewallRule")
	deleteOp.AddReqStructure(new(FirewallRuleDeleteRequest))

	deleteOp.AddRespStructure(new(FirewallRuleResponse), openapi.WithHTTPStatus(http.StatusOK), func(cu *openapi.ContentUnit) { cu.IsDefault = true })

	for _, status := range []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests, http.StatusInternalServerError} {
		deleteOp.AddRespStructure(ErrorResponse, openapi.WithContentType("application/json"), openapi.WithHTTPStatus(status))
	}

	err = reflector.AddOperation(deleteOp)
	if err != nil {
		log.Fatal(err)
	}
}
