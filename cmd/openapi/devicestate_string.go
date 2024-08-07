// Code generated by "stringer -trimprefix DeviceState -type DeviceState"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DeviceStateUnknown-0]
	_ = x[DeviceStateConnected-1]
	_ = x[DeviceStatePending-2]
	_ = x[DeviceStateFirmwareMismatch-3]
	_ = x[DeviceStateUpgrading-4]
	_ = x[DeviceStateProvisioning-5]
	_ = x[DeviceStateHeartbeatMissed-6]
	_ = x[DeviceStateAdopting-7]
	_ = x[DeviceStateDeleting-8]
	_ = x[DeviceStateInformError-9]
	_ = x[DeviceStateAdoptFailed-10]
	_ = x[DeviceStateIsolated-11]
}

const _DeviceState_name = "UnknownConnectedPendingFirmwareMismatchUpgradingProvisioningHeartbeatMissedAdoptingDeletingInformErrorAdoptFailedIsolated"

var _DeviceState_index = [...]uint8{0, 7, 16, 23, 39, 48, 60, 75, 83, 91, 102, 113, 121}

func (i DeviceState) String() string {
	if i < 0 || i >= DeviceState(len(_DeviceState_index)-1) {
		return "DeviceState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DeviceState_name[_DeviceState_index[i]:_DeviceState_index[i+1]]
}
