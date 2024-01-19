package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/swaggest/openapi-go/openapi3"
)

type meta struct {
	RC      string `json:"rc"`
	Message string `json:"msg"`
}

var reflector = openapi3.Reflector{
	Spec: &openapi3.Spec{
		Openapi: "3.0.3",
		Info: openapi3.Info{
			Title:   "Unifi API",
			Version: UnifiVersion,
		},
	},
}

func main() {
	reflector.Spec.SetDescription("Unifi Controller API")

	addOperations()

	schema, err := reflector.Spec.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(filepath.Join("bin", "openapi", "openapi.yaml"), schema, 0777)

	fmt.Println(string(schema))
}

var PathMappings = map[string]map[string]string{
	"Id": {
		"site": "Site",
		"id":   "_id",
	},
	"Site": {},
}

var ErrorResponse = struct {
	Meta meta `json:"meta"`
	Data []struct {
		Meta meta `json:"meta"`
	} `json:"data"`
}{}

func addOperations() {

	addAccount()
	addBroadcastGroup()
	addChannelPlan()
	addDashboard()
	addDevice()
	addDHCPOption()
	addDpiApp()
	addDpiGroup()
	addDynamicDNS()
	addFirewallGroup()
	addFirewallRule()
	addHeatMap()
	addHeatMapPoint()
	addHotspot2Conf()
	addHotspotOp()
	addHotspotPackage()
	addMap()
	addMediaFile()
	addNetwork()
	addPortProfile()
	addPortForward()
	addRADIUSProfile()
	addRouting()
	addScheduleTask()
	addSettingAutoSpeedtest()
	addSettingBaresip()
	addSettingBroadcast()
	addSettingConnectivity()
	addSettingCountry()
	addSettingDoh()
	addSettingDpi()
	addSettingElementAdopt()
	addSettingEtherLighting()
	addSettingGlobalAp()
	addSettingGlobalSwitch()
	addSettingGuestAccess()
	addSettingIps()
	addSettingLcm()
	addSettingLocale()
	addSettingMagicSiteToSiteVpn()
	addSettingMgmt()
	addSettingNetworkOptimization()
	addSettingNtp()
	addSettingPorta()
	addSettingRadioAi()
	addSettingRadius()
	addSettingRsyslogd()
	addSettingSnmp()
	addSettingSuperCloudaccess()
	addSettingSuperEvents()
	addSettingSuperFwupdate()
	addSettingSuperIdentity()
	addSettingSuperMail()
	addSettingSuperMgmt()
	addSettingSuperSdn()
	addSettingSuperSmtp()
	addSettingTeleport()
	addSettingUsg()
	addSettingUsw()
	addSpatialRecord()
	addTag()
	addUser()
	addUserGroup()
	addVirtualDevice()
	addWLAN()
	addWLANGroup()
}

func emptyBoolToTrue(b *bool) bool {
	if b == nil {
		return true
	}
	return *b
}

// numberOrString handles strings that can also accept JSON numbers.
// For example a field may contain a number or the string "auto".
type numberOrString string

func (e *numberOrString) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	s := string(b)
	if s == `""` {
		*e = ""
		return nil
	}
	var err error
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		s, err = strconv.Unquote(s)
		if err != nil {
			return err
		}
		*e = numberOrString(s)
		return nil
	}
	*e = numberOrString(string(b))
	return nil
}

// emptyStringInt was created due to the behavior change in
// Go 1.14 with json.Number's handling of empty string.
type emptyStringInt int

func (e *emptyStringInt) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	s := string(b)
	if s == `""` {
		*e = 0
		return nil
	}
	var err error
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		s, err = strconv.Unquote(s)
		if err != nil {
			return err
		}
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*e = emptyStringInt(i)
	return nil
}

func (e *emptyStringInt) MarshalJSON() ([]byte, error) {
	if e == nil || *e == 0 {
		return []byte(`""`), nil
	}

	return []byte(strconv.Itoa(int(*e))), nil
}

type booleanishString bool

func (e *booleanishString) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == `"enabled"` {
		*e = booleanishString(true)
		return nil
	} else if s == `"disabled"` {
		*e = booleanishString(false)
		return nil
	}
	return errors.New("Could not unmarshal JSON value.")
}

//go:generate go run golang.org/x/tools/cmd/stringer -trimprefix DeviceState -type DeviceState
type DeviceState int

const (
	DeviceStateUnknown          DeviceState = 0
	DeviceStateConnected        DeviceState = 1
	DeviceStatePending          DeviceState = 2
	DeviceStateFirmwareMismatch DeviceState = 3
	DeviceStateUpgrading        DeviceState = 4
	DeviceStateProvisioning     DeviceState = 5
	DeviceStateHeartbeatMissed  DeviceState = 6
	DeviceStateAdopting         DeviceState = 7
	DeviceStateDeleting         DeviceState = 8
	DeviceStateInformError      DeviceState = 9
	DeviceStateAdoptFailed      DeviceState = 10
	DeviceStateIsolated         DeviceState = 11
)
