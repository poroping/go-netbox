/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.3.0-Docker-3.3.0 (4.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// PatchedWritableRackRequestStatus * `reserved` - Reserved * `available` - Available * `planned` - Planned * `active` - Active * `deprecated` - Deprecated
type PatchedWritableRackRequestStatus string

// List of PatchedWritableRackRequest_status
const (
	PATCHEDWRITABLERACKREQUESTSTATUS_RESERVED   PatchedWritableRackRequestStatus = "reserved"
	PATCHEDWRITABLERACKREQUESTSTATUS_AVAILABLE  PatchedWritableRackRequestStatus = "available"
	PATCHEDWRITABLERACKREQUESTSTATUS_PLANNED    PatchedWritableRackRequestStatus = "planned"
	PATCHEDWRITABLERACKREQUESTSTATUS_ACTIVE     PatchedWritableRackRequestStatus = "active"
	PATCHEDWRITABLERACKREQUESTSTATUS_DEPRECATED PatchedWritableRackRequestStatus = "deprecated"
)

// All allowed values of PatchedWritableRackRequestStatus enum
var AllowedPatchedWritableRackRequestStatusEnumValues = []PatchedWritableRackRequestStatus{
	"reserved",
	"available",
	"planned",
	"active",
	"deprecated",
}

func (v *PatchedWritableRackRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableRackRequestStatus(value)
	for _, existing := range AllowedPatchedWritableRackRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableRackRequestStatus", value)
}

// NewPatchedWritableRackRequestStatusFromValue returns a pointer to a valid PatchedWritableRackRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableRackRequestStatusFromValue(v string) (*PatchedWritableRackRequestStatus, error) {
	ev := PatchedWritableRackRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableRackRequestStatus: valid values are %v", v, AllowedPatchedWritableRackRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableRackRequestStatus) IsValid() bool {
	for _, existing := range AllowedPatchedWritableRackRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableRackRequest_status value
func (v PatchedWritableRackRequestStatus) Ptr() *PatchedWritableRackRequestStatus {
	return &v
}

type NullablePatchedWritableRackRequestStatus struct {
	value *PatchedWritableRackRequestStatus
	isSet bool
}

func (v NullablePatchedWritableRackRequestStatus) Get() *PatchedWritableRackRequestStatus {
	return v.value
}

func (v *NullablePatchedWritableRackRequestStatus) Set(val *PatchedWritableRackRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableRackRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableRackRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableRackRequestStatus(val *PatchedWritableRackRequestStatus) *NullablePatchedWritableRackRequestStatus {
	return &NullablePatchedWritableRackRequestStatus{value: val, isSet: true}
}

func (v NullablePatchedWritableRackRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableRackRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
