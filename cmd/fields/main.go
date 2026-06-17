package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/hashicorp/go-version"
	"github.com/iancoleman/strcase"
)

type replacement struct {
	Old string
	New string
}

var fieldReps = []replacement{
	{"Dhcpdv6", "DHCPDV6"},

	{"Dhcpd", "DHCPD"},
	{"Idx", "IDX"},
	{"Ipsec", "IPSec"},
	{"Ipv6", "IPV6"},
	{"Openvpn", "OpenVPN"},
	{"Tftp", "TFTP"},
	{"Wlangroup", "WLANGroup"},

	{"FrrBgpdConfig", "Config"},
	{"BgpConfig", "BGPConfig"},

	{"Bc", "Broadcast"},
	{"Dhcp", "DHCP"},
	{"Dns", "DNS"},
	{"Dpi", "DPI"},
	{"Dtim", "DTIM"},
	{"Firewallgroup", "FirewallGroup"},
	{"Fixedip", "FixedIP"},
	{"Icmp", "ICMP"},
	{"Id", "ID"},
	{"Igmp", "IGMP"},
	{"Ip", "IP"},
	{"Leasetime", "LeaseTime"},
	{"Mac", "MAC"},
	{"Mcastenhance", "MulticastEnhance"},
	{"Minrssi", "MinRSSI"},
	{"Monthdays", "MonthDays"},
	{"Nat", "NAT"},
	{"Networkconf", "Network"},
	{"Networkgroup", "NetworkGroup"},
	{"Pd", "PD"},
	{"Pmf", "PMF"},
	{"pnp", "PnP"},
	{"Portconf", "PortProfile"},
	{"Qos", "QOS"},
	{"Radiusprofile", "RADIUSProfile"},
	{"Radius", "RADIUS"},
	{"Ssid", "SSID"},
	{"Smartq", "SmartQ"},
	{"Startdate", "StartDate"},
	{"Starttime", "StartTime"},
	{"Stopdate", "StopDate"},
	{"Stoptime", "StopTime"},
	{"Supression", "Suppression"}, //nolint:misspell
	{"Tcp", "TCP"},
	{"Udp", "UDP"},
	{"Usergroup", "UserGroup"},
	{"Utc", "UTC"},
	{"Vlan", "VLAN"},
	{"Vpn", "VPN"},
	{"Wan", "WAN"},
	{"Wep", "WEP"},
	{"Wlan", "WLAN"},
	{"Wpa", "WPA"},
	{"XWireguardPrivateKey", "WireguardPrivateKey"},
	{"XSsh", "SSH"},
	{"XMgmt", "Mgmt"},
	{"UnifiIDp", "UniFiIdentityProvider"},
	{"PortStop", "Stop"},
	{"PortStart", "Start"},
	{"IPStart", "Start"},
	{"IPStop", "Stop"},
	{"IPVersion", "Version"},
	{"IPOrSubnet", "Address"},
}

var fileReps = []replacement{
	{"WlanConf", "WLAN"},
	{"Dhcp", "DHCP"},
	{"Wlan", "WLAN"},
	{"NetworkConf", "Network"},
	{"PortConf", "PortProfile"},
	{"RadiusProfile", "RADIUSProfile"},
	{"ApGroups", "APGroup"},
	{"DnsRecord", "DNSRecord"},
	{"BgpConfig", "BGPConfig"},
	{"User", "Client"},
	{"UserGroup", "ClientGroup"},
}

type ResourceInfo struct {
	StructName       string
	ResourcePath     string
	ItemResourcePath string
	Types            map[string]*FieldInfo
	FieldProcessor   func(name string, f *FieldInfo) error
}

type FieldInfo struct {
	FieldName           string
	JSONName            string
	FieldType           string
	IsPointer           bool
	FieldValidation     string
	OmitEmpty           bool
	IsArray             bool
	Fields              map[string]*FieldInfo
	CustomUnmarshalType string
	CustomUnmarshalFunc string
}

func NewResource(structName string, resourcePath string) *ResourceInfo {
	baseType := NewFieldInfo(structName, resourcePath, "struct", "", false, false, false, "")
	resource := &ResourceInfo{
		StructName:   structName,
		ResourcePath: resourcePath,
		Types: map[string]*FieldInfo{
			structName: baseType,
		},
		FieldProcessor: func(name string, f *FieldInfo) error { return nil },
	}

	// Since template files iterate through map keys in sorted order, these initial fields
	// are named such that they stay at the top for consistency. The spacer items create a
	// blank line in the resulting generated file.
	//
	// This hack is here for stability of the generatd code, but can be removed if desired.
	baseType.Fields = map[string]*FieldInfo{
		"   ID":      NewFieldInfo("ID", "_id", String, "", true, false, false, ""),
		"   SiteID":  NewFieldInfo("SiteID", "site_id", String, "", true, false, false, ""),
		"   _Spacer": nil,

		"  Hidden":   NewFieldInfo("Hidden", "attr_hidden", Bool, "", true, false, false, ""),
		"  HiddenID": NewFieldInfo("HiddenID", "attr_hidden_id", String, "", true, false, false, ""),
		"  NoDelete": NewFieldInfo("NoDelete", "attr_no_delete", Bool, "", true, false, false, ""),
		"  NoEdit":   NewFieldInfo("NoEdit", "attr_no_edit", Bool, "", true, false, false, ""),
		"  _Spacer":  nil,

		" _Spacer": nil,
	}

	switch {
	case resource.IsSetting():
		resource.ResourcePath = strcase.ToSnake(strings.TrimPrefix(structName, "Setting"))
		baseType.Fields[" Key"] = NewFieldInfo("Key", "key", String, "", false, false, false, "")
		if resource.StructName == "SettingUsg" {
			// Removed in v7, retaining for backwards compatibility
			baseType.Fields["MdnsEnabled"] = NewFieldInfo("MdnsEnabled", "mdns_enabled", Bool, "", false, false, false, "")
		}
	case resource.StructName == "DNSRecord":
		resource.ResourcePath = "static-dns"
	case resource.StructName == "FirewallZone":
		resource.ResourcePath = "firewall/zone"
	case resource.StructName == "OSPFRouter":
		resource.ResourcePath = "ospf/router"
	case resource.StructName == "FirewallPolicy":
		resource.ResourcePath = "firewall-policies"
	case resource.StructName == "TrafficRoute":
		resource.ResourcePath = "trafficroutes"
	case resource.StructName == "APGroup":
		resource.ResourcePath = "apgroups"
	case resource.StructName == "NetworkMembersGroup":
		resource.ResourcePath = "network-members-groups"
		resource.ItemResourcePath = "network-members-group"
	case resource.StructName == "PowerSupervisor":
		resource.ResourcePath = "power-supervisors"
	case resource.StructName == "Network":
		baseType.Fields["WANEgressQOSEnabled"] = NewFieldInfo("WANEgressQOSEnabled", "wan_egress_qos_enabled", Bool, "", true, false, true, "")
		baseType.Fields["UPnPEnabled"] = NewFieldInfo("UPnPEnabled", "upnp_enabled", Bool, "", true, false, true, "")
		baseType.Fields["UPnPWANInterface"] = NewFieldInfo("UPnPWANInterface", "upnp_wan_interface", String, "", true, false, true, "")
		baseType.Fields["UPnPNatPMPEnabled"] = NewFieldInfo("UPnPNatPMPEnabled", "upnp_nat_pmp_enabled", Bool, "", true, false, true, "")
		baseType.Fields["UPnPSecureMode"] = NewFieldInfo("UPnPSecureMode", "upnp_secure_mode", Bool, "", true, false, true, "")
		baseType.Fields["IPAliases"] = NewFieldInfo("IPAliases", "ip_aliases", String, "", true, true, false, "")
		baseType.Fields["DHCPRelayServers"] = NewFieldInfo("DHCPRelayServers", "dhcp_relay_servers", String, "", true, true, false, "")
		baseType.Fields["WireguardInterfaceBindingModeIPVersion"] = NewFieldInfo(
			"WireguardInterfaceBindingModeIPVersion",
			"wireguard_interface_binding_mode_ip_version",
			String,
			"v4|v6",
			true,
			false,
			true,
			"",
		)
	case resource.StructName == "Device":
		baseType.Fields["PortTable"] = NewFieldInfo("PortTable", "port_table", "[]DevicePortTable", "", true, false, false, "")
		baseType.Fields[" MAC"] = NewFieldInfo("MAC", "mac", String, "", true, false, false, "")
		baseType.Fields["Adopted"] = NewFieldInfo("Adopted", "adopted", Bool, "", false, false, false, "")
		baseType.Fields["Model"] = NewFieldInfo("Model", "model", String, "", true, false, false, "")
		baseType.Fields["State"] = NewFieldInfo("State", "state", "DeviceState", "", false, false, false, "")
		baseType.Fields["Type"] = NewFieldInfo("Type", "type", String, "", true, false, false, "")
		baseType.Fields["InformIP"] = NewFieldInfo("InformIP", "inform_ip", String, "", true, false, false, "")
		baseType.Fields["IP"] = NewFieldInfo("IP", "ip", String, "", true, false, false, "")
	case resource.StructName == "Client":
		baseType.Fields[" DisplayName"] = NewFieldInfo("DisplayName", "display_name", String, "non-generated field", true, false, false, "")
	case resource.StructName == "WLAN":
		// this field removed in v6, retaining for backwards compatibility
		baseType.Fields["WLANGroupID"] = NewFieldInfo("WLANGroupID", "wlangroup_id", String, "", true, false, false, "")
	case resource.StructName == "BGPConfig":
		resource.ResourcePath = "bgp/config"
	}

	return resource
}

func NewFieldInfo(
	fieldName string,
	jsonName string,
	fieldType string,
	fieldValidation string,
	omitempty bool,
	isArray bool,
	isPointer bool,
	customUnmarshalType string,
) *FieldInfo {
	return &FieldInfo{
		FieldName:           fieldName,
		JSONName:            jsonName,
		FieldType:           fieldType,
		FieldValidation:     fieldValidation,
		OmitEmpty:           omitempty,
		IsArray:             isArray,
		IsPointer:           isPointer,
		CustomUnmarshalType: customUnmarshalType,
	}
}

func cleanName(name string, reps []replacement) string {
	for _, rep := range reps {
		name = strings.ReplaceAll(name, rep.Old, rep.New)
	}

	if strings.HasPrefix(name, "X") && len(name) > 1 && unicode.IsUpper(rune(name[1])) {
		name = name[1:]
	}

	return name
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] version\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	assetsDirFlag := flag.String(
		"assets-dir",
		"assets",
		"Output directory (relative to the working dir) for the OpenAPI spec and generator config",
	)
	downloadOnly := flag.Bool(
		"download-only",
		false,
		"Only download and build the fields JSON directory, do not generate",
	)
	useLatestVersion := flag.Bool("latest", false, "Use the latest available version")

	flag.Parse()

	specifiedVersion := flag.Arg(0)
	if specifiedVersion != "" && *useLatestVersion {
		fmt.Print("error: cannot specify version with latest\n\n")
		usage()
		os.Exit(1)
	} else if specifiedVersion == "" && !*useLatestVersion {
		fmt.Print("error: must specify version or latest\n\n")
		usage()
		os.Exit(1)
	}

	var unifiVersion *version.Version
	var unifiDownloadUrl *url.URL
	var err error

	if *useLatestVersion {
		unifiVersion, unifiDownloadUrl, err = latestUnifiVersion()
		if err != nil {
			panic(err)
		}
	} else {
		unifiVersion, err = version.NewVersion(specifiedVersion)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		unifiDownloadUrl, err = url.Parse(fmt.Sprintf("https://dl.ui.com/unifi/%s/unifi_sysvinit_all.deb", unifiVersion))
		if err != nil {
			panic(err)
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Unable to get the current filename")
	}

	versionBaseDir := filepath.Dir(filename)

	fieldsDir := filepath.Join(versionBaseDir, fmt.Sprintf("v%s", unifiVersion))

	assetsDir := filepath.Join(wd, *assetsDirFlag)

	fieldsInfo, err := os.Stat(fieldsDir)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}

		err = os.MkdirAll(fieldsDir, 0o755)
		if err != nil {
			panic(err)
		}

		// download fields, create
		jarFile, err := downloadJar(unifiDownloadUrl, fieldsDir)
		if err != nil {
			panic(err)
		}

		err = extractJSON(jarFile, fieldsDir)
		if err != nil {
			panic(err)
		}

		// defer func() {
		// 	err = os.RemoveAll(fieldsDir)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// }()

		err = copyCustom(fieldsDir)
		if err != nil {
			panic(err)
		}

		fieldsInfo, err = os.Stat(fieldsDir)
		if err != nil {
			panic(err)
		}
	}
	if !fieldsInfo.IsDir() {
		panic("version info isn't a directory")
	}

	if *downloadOnly {
		fmt.Println("Fields JSON ready!")
		os.Exit(0)
	}

	fieldsFiles, err := os.ReadDir(fieldsDir)
	if err != nil {
		panic(err)
	}

	// Every resource parsed from the field definitions; fed to both outputs.
	resources := []*ResourceInfo{}

	for _, fieldsFile := range fieldsFiles {
		name := fieldsFile.Name()
		ext := filepath.Ext(name)

		switch name {
		case "AuthenticationRequest.json", "Setting.json", "Wall.json":
			continue
		}

		if filepath.Ext(name) != ".json" {
			continue
		}

		name = name[:len(name)-len(ext)]

		urlPath := strings.ToLower(name)
		structName := cleanName(name, fileReps)

		fieldsFilePath := filepath.Join(fieldsDir, fieldsFile.Name())
		b, err := os.ReadFile(fieldsFilePath)
		if err != nil {
			fmt.Printf("skipping file %s: %s", fieldsFile.Name(), err)
			continue
		}

		resource := NewResource(structName, urlPath)

		switch resource.StructName {
		case "Account":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				switch name {
				case "IP", "NetworkID":
					f.OmitEmpty = true
				}
				return nil
			}
		case "ChannelPlan":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				switch name {
				case "Channel", "BackupChannel", "TxPower":
					if f.FieldType == String {
						f.CustomUnmarshalType = Number
					}
				}
				return nil
			}
		case "Device":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				switch name {
				case "X", "Y":
					f.FieldType = "float64"
				case "StpPriority":
					f.FieldType = Int
					f.CustomUnmarshalType = Number
				case "ConfigNetwork", "EtherLighting", "MbbOverrides", "NutServer", "RpsOverride", "QOSProfile":
					f.IsPointer = true
				case "Ht":
					// Field within DeviceRadioTable nested type
					f.CustomUnmarshalType = "types.Number"
					f.CustomUnmarshalFunc = "types.ToInt64Pointer"
				case "Channel", "TxPower":
					// String fields in DeviceRadioTable that newer controllers
					// (UniFi 10.x) return as numbers; accept either form.
					if f.FieldType == String {
						f.CustomUnmarshalType = Number
					}
				}

				f.OmitEmpty = true
				switch name {
				case "PortOverrides":
					f.OmitEmpty = false
				}

				return nil
			}
		case "Network":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				switch name {
				case "InternetAccessEnabled", "IntraNetworkAccessEnabled":
					if f.FieldType == Bool {
						f.CustomUnmarshalType = "*bool"
						f.CustomUnmarshalFunc = "emptyBoolToTrue"
					}
				case "DHCPDEnabled":
					// Some controllers (UniFi Network 10.x) return "true"/"false"
					// as JSON strings for this flag, which breaks a plain bool.
					// Decode through the tolerant types.Bool. See
					// terraform-provider-unifi #65.
					if f.FieldType == Bool {
						f.CustomUnmarshalType = "*types.Bool"
						f.CustomUnmarshalFunc = "boolValue"
					}
				case "IPSecEspLifetime", "IPSecIkeLifetime":
					f.FieldType = Int
					f.IsPointer = true
				case "WANDNS1", "WANDNS2", "WANIPV6DNS1", "WANIPV6DNS2", "DHCPDStart", "DHCPDStop", "DHCPDUnifiController",
					"DHCPDTFTPServer", "DHCPDWins1", "DHCPDWins2", "DHCPDWPAdUrl", "DomainName", "DHCPDGateway", "DHCPDNtp1", "DHCPDNtp2":
					f.OmitEmpty = true
					f.IsPointer = true
				case "Purpose":
					f.OmitEmpty = false
					f.IsPointer = false
				}
				if f.OmitEmpty && !f.IsArray {
					switch f.FieldType {
					case Bool, String:
						f.IsPointer = true
					}
				}
				return nil
			}
		case "SettingGlobalAp":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				if strings.HasPrefix(name, "6E") {
					f.FieldName = strings.Replace(f.FieldName, "6E", "SixE", 1)
				}

				return nil
			}
		case "SettingMgmt":
			sshKeyField := NewFieldInfo(resource.StructName+"SSHKeys", "x_ssh_keys", "struct", "", false, false, false, "")
			sshKeyField.Fields = map[string]*FieldInfo{
				"name":        NewFieldInfo("Name", "name", String, "", false, false, false, ""),
				"keyType":     NewFieldInfo("KeyType", "type", String, "", false, false, false, ""),
				"key":         NewFieldInfo("Key", "key", String, "", false, false, false, ""),
				"comment":     NewFieldInfo("Comment", "comment", String, "", false, false, false, ""),
				"date":        NewFieldInfo("Date", "date", String, "", false, false, false, ""),
				"fingerprint": NewFieldInfo("Fingerprint", "fingerprint", String, "", false, false, false, ""),
			}
			resource.Types[sshKeyField.FieldName] = sshKeyField

			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				if name == "SSHKeys" {
					f.FieldType = sshKeyField.FieldName
				}
				return nil
			}
		case "SettingUsg":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				if strings.HasSuffix(name, "Timeout") && name != "ArpCacheTimeout" {
					f.FieldType = Int
					f.CustomUnmarshalType = Number
				}
				return nil
			}
		case "Nat":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				switch name {
				case "SourceFilter":
					f.IsPointer = true
				case "DestinationFilter":
					f.IsPointer = true
				}
				return nil
			}
		case "Client":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				switch name {
				case "Blocked":
					f.FieldType = Bool
					f.IsPointer = true
					// Some controllers return "true"/"false" as JSON strings
					// for this flag, which breaks a plain *bool. Decode through
					// the tolerant types.Bool. See terraform-provider-unifi #132.
					f.CustomUnmarshalType = "*types.Bool"
					f.CustomUnmarshalFunc = "boolPtrValue"
				case "VirtualNetworkOverrideEnabled":
					f.FieldType = Bool
					f.IsPointer = true
					f.OmitEmpty = true
				case "LastSeen":
					f.FieldType = Int
					f.IsPointer = true
				}
				return nil
			}
		case "WLAN":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				switch name {
				case "ScheduleWithDuration":
					// always send schedule, so we can empty it if we want to
					f.OmitEmpty = false
				}
				return nil
			}
		case "DNSRecord":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				switch name {
				case "Hidden", "NoDelete", "NoEdit", "Enabled":
					f.FieldType = Bool
				case "Priority", "Ttl", "Weight":
					f.FieldType = Int
					f.CustomUnmarshalType = Number
				}
				return nil
			}
		}

		err = resource.processJSON(b)
		if err != nil {
			fmt.Printf("skipping file %s: %s", fieldsFile.Name(), err)
			continue
		}

		// Add fields not present in the JAR schema to nested types.
		if resource.StructName == "Device" {
			if portOverrides, ok := resource.Types["DevicePortOverrides"]; ok {
				portOverrides.Fields["TaggedNetworkIDs"] = NewFieldInfo("TaggedNetworkIDs", "tagged_networkconf_ids", String, "", true, true, false, "")
			}
		}

		resources = append(resources, resource)
	}

	if err := os.MkdirAll(assetsDir, 0o755); err != nil {
		panic(err)
	}

	// Output #1: the detailed OpenAPI specification.
	openAPIFile := filepath.Join(assetsDir, "openapi.yaml")
	if err := WriteOpenAPI(resources, unifiVersion.String(), openAPIFile); err != nil {
		panic(err)
	}
	fmt.Printf("Generated OpenAPI spec: %s\n", openAPIFile)

	// Output #2: OpenAPI Generator config that reproduces go-unifi's naming
	// cleanup in generated clients.
	generatorDir := filepath.Join(assetsDir, "openapi-generator")
	if err := WriteGeneratorConfig(resources, generatorDir); err != nil {
		panic(err)
	}
	fmt.Printf("Generated OpenAPI Generator config: %s\n", generatorDir)

	// Output #3: two consolidated oapi-codegen-exp (V3) configs.
	// assets/oapi-codegen-exp/unifi.yaml    → clients/go/unifi.gen.go
	// assets/oapi-codegen-exp/settings.yaml → clients/go/settings/settings.gen.go
	// Consumed by the static //go:generate directives in the repo-root generate.go.
	oapiConfigDir := filepath.Join(assetsDir, "oapi-codegen-exp")
	clientsDir := filepath.Join(wd, "clients", "go")
	if err := WriteOAPICodegenConfigs(resources, oapiConfigDir, openAPIFile, clientsDir); err != nil {
		panic(err)
	}
	fmt.Printf("Generated oapi-codegen-exp configs: %s/{unifi,settings}.yaml\n", oapiConfigDir)
}

func (r *ResourceInfo) IsSetting() bool {
	return strings.HasPrefix(r.StructName, "Setting")
}

func (r *ResourceInfo) IsDevice() bool {
	return r.StructName == "Device"
}

func (r *ResourceInfo) IsV2() bool {
	return slices.Contains([]string{
		"APGroup",
		"ApGroup",
		"BGPConfig",
		"DNSRecord",
		"FirewallPolicy",
		"FirewallZone",
		"NetworkMembersGroup",
		"Nat",
		"OSPFRouter",
		"PowerSupervisor",
		"TrafficRoute",
	}, r.StructName)
}

func (r *ResourceInfo) CleanStructName() string {
	if r.IsSetting() {
		return strings.TrimPrefix(r.StructName, "Setting")
	}
	return r.StructName
}

func (r *ResourceInfo) processFields(fields map[string]any) {
	t := r.Types[r.StructName]
	for name, validation := range fields {
		fieldInfo, err := r.fieldInfoFromValidation(name, validation)
		if err != nil {
			continue
		}

		t.Fields[fieldInfo.FieldName] = fieldInfo
	}
}

func (r *ResourceInfo) fieldInfoFromValidation(name string, validation any) (*FieldInfo, error) {
	fieldName := strcase.ToCamel(name)
	fieldName = cleanName(fieldName, fieldReps)

	empty := &FieldInfo{}
	var fieldInfo *FieldInfo

	switch validation := validation.(type) {
	case []any:
		if len(validation) == 0 {
			fieldInfo = NewFieldInfo(fieldName, name, String, "", false, true, false, "")
			err := r.FieldProcessor(fieldName, fieldInfo)
			return fieldInfo, err
		}
		if len(validation) > 1 {
			return empty, fmt.Errorf("unknown validation %#v", validation)
		}

		fieldInfo, err := r.fieldInfoFromValidation(name, validation[0])
		if err != nil {
			return empty, err
		}

		fieldInfo.OmitEmpty = true
		fieldInfo.IsArray = true
		fieldInfo.IsPointer = false

		err = r.FieldProcessor(fieldName, fieldInfo)
		return fieldInfo, err

	case map[string]any:
		typeName := r.StructName + fieldName

		result := NewFieldInfo(fieldName, name, typeName, "", true, false, true, "")
		result.Fields = make(map[string]*FieldInfo)

		for name, fv := range validation {
			child, err := r.fieldInfoFromValidation(name, fv)
			if err != nil {
				return empty, err
			}

			result.Fields[child.FieldName] = child
		}

		err := r.FieldProcessor(fieldName, result)
		r.Types[typeName] = result
		return result, err

	case string:
		fieldValidation := validation
		normalized := normalizeValidation(validation)

		omitEmpty := false

		switch normalized {
		case "falsetrue", "truefalse":
			fieldInfo = NewFieldInfo(fieldName, name, Bool, "", omitEmpty, false, false, "")
			return fieldInfo, r.FieldProcessor(fieldName, fieldInfo)
		default:
			if _, err := strconv.ParseFloat(normalized, 64); err == nil {
				if normalized == "09" || normalized == "09.09" {
					fieldValidation = ""
				}

				if strings.Contains(normalized, ".") {
					if strings.Contains(validation, "\\.){3}") {
						break
					}

					omitEmpty = true
					fieldInfo = NewFieldInfo(fieldName, name, "float64", fieldValidation, omitEmpty, false, false, "")
					return fieldInfo, r.FieldProcessor(fieldName, fieldInfo)
				}

				omitEmpty = true
				fieldInfo = NewFieldInfo(fieldName, name, Int, fieldValidation, omitEmpty, false, true, Number)
				return fieldInfo, r.FieldProcessor(fieldName, fieldInfo)
			}
		}
		if validation != "" && normalized != "" {
			fmt.Printf("normalize %q to %q\n", validation, normalized)
		}

		omitEmpty = omitEmpty || (!strings.Contains(validation, "^$") && !strings.HasSuffix(fieldName, "Id"))
		fieldInfo = NewFieldInfo(fieldName, name, String, fieldValidation, omitEmpty, false, false, "")
		return fieldInfo, r.FieldProcessor(fieldName, fieldInfo)
	}

	return empty, fmt.Errorf("unable to determine type from validation %q", validation)
}

func (r *ResourceInfo) processJSON(b []byte) error {
	var fields map[string]any
	err := json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}

	r.processFields(fields)

	return nil
}

func normalizeValidation(re string) string {
	re = strings.ReplaceAll(re, "\\d", "[0-9]")
	re = strings.ReplaceAll(re, "[-+]?", "")
	re = strings.ReplaceAll(re, "[+-]?", "")
	re = strings.ReplaceAll(re, "[-]?", "")
	re = strings.ReplaceAll(re, "\\.", ".")
	re = strings.ReplaceAll(re, "[.]?", ".")

	quants := regexp.MustCompile(`\{\d*,?\d*\}|\*|\+|\?`)
	re = quants.ReplaceAllString(re, "")

	control := regexp.MustCompile(`[\(\[\]\)\|\-\$\^]`)
	re = control.ReplaceAllString(re, "")

	re = strings.TrimPrefix(re, "^")
	re = strings.TrimSuffix(re, "$")

	return re
}
