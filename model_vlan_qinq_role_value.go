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

// VLANQinqRoleValue * `svlan` - Service * `cvlan` - Customer
type VLANQinqRoleValue string

// List of VLAN_qinq_role_value
const (
	VLANQINQROLEVALUE_SVLAN VLANQinqRoleValue = "svlan"
	VLANQINQROLEVALUE_CVLAN VLANQinqRoleValue = "cvlan"
)

// All allowed values of VLANQinqRoleValue enum
var AllowedVLANQinqRoleValueEnumValues = []VLANQinqRoleValue{
	"svlan",
	"cvlan",
}

func (v *VLANQinqRoleValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VLANQinqRoleValue(value)
	for _, existing := range AllowedVLANQinqRoleValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VLANQinqRoleValue", value)
}

// NewVLANQinqRoleValueFromValue returns a pointer to a valid VLANQinqRoleValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVLANQinqRoleValueFromValue(v string) (*VLANQinqRoleValue, error) {
	ev := VLANQinqRoleValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VLANQinqRoleValue: valid values are %v", v, AllowedVLANQinqRoleValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VLANQinqRoleValue) IsValid() bool {
	for _, existing := range AllowedVLANQinqRoleValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VLAN_qinq_role_value value
func (v VLANQinqRoleValue) Ptr() *VLANQinqRoleValue {
	return &v
}

type NullableVLANQinqRoleValue struct {
	value *VLANQinqRoleValue
	isSet bool
}

func (v NullableVLANQinqRoleValue) Get() *VLANQinqRoleValue {
	return v.value
}

func (v *NullableVLANQinqRoleValue) Set(val *VLANQinqRoleValue) {
	v.value = val
	v.isSet = true
}

func (v NullableVLANQinqRoleValue) IsSet() bool {
	return v.isSet
}

func (v *NullableVLANQinqRoleValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVLANQinqRoleValue(val *VLANQinqRoleValue) *NullableVLANQinqRoleValue {
	return &NullableVLANQinqRoleValue{value: val, isSet: true}
}

func (v NullableVLANQinqRoleValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVLANQinqRoleValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
