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

// InterfacePoeTypeValue * `type1-ieee802.3af` - 802.3af (Type 1) * `type2-ieee802.3at` - 802.3at (Type 2) * `type3-ieee802.3bt` - 802.3bt (Type 3) * `type4-ieee802.3bt` - 802.3bt (Type 4) * `passive-24v-2pair` - Passive 24V (2-pair) * `passive-24v-4pair` - Passive 24V (4-pair) * `passive-48v-2pair` - Passive 48V (2-pair) * `passive-48v-4pair` - Passive 48V (4-pair)
type InterfacePoeTypeValue string

// List of Interface_poe_type_value
const (
	INTERFACEPOETYPEVALUE_TYPE1_IEEE802_3AF InterfacePoeTypeValue = "type1-ieee802.3af"
	INTERFACEPOETYPEVALUE_TYPE2_IEEE802_3AT InterfacePoeTypeValue = "type2-ieee802.3at"
	INTERFACEPOETYPEVALUE_TYPE3_IEEE802_3BT InterfacePoeTypeValue = "type3-ieee802.3bt"
	INTERFACEPOETYPEVALUE_TYPE4_IEEE802_3BT InterfacePoeTypeValue = "type4-ieee802.3bt"
	INTERFACEPOETYPEVALUE_PASSIVE_24V_2PAIR InterfacePoeTypeValue = "passive-24v-2pair"
	INTERFACEPOETYPEVALUE_PASSIVE_24V_4PAIR InterfacePoeTypeValue = "passive-24v-4pair"
	INTERFACEPOETYPEVALUE_PASSIVE_48V_2PAIR InterfacePoeTypeValue = "passive-48v-2pair"
	INTERFACEPOETYPEVALUE_PASSIVE_48V_4PAIR InterfacePoeTypeValue = "passive-48v-4pair"
	INTERFACEPOETYPEVALUE_EMPTY             InterfacePoeTypeValue = ""
)

// All allowed values of InterfacePoeTypeValue enum
var AllowedInterfacePoeTypeValueEnumValues = []InterfacePoeTypeValue{
	"type1-ieee802.3af",
	"type2-ieee802.3at",
	"type3-ieee802.3bt",
	"type4-ieee802.3bt",
	"passive-24v-2pair",
	"passive-24v-4pair",
	"passive-48v-2pair",
	"passive-48v-4pair",
	"",
}

func (v *InterfacePoeTypeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfacePoeTypeValue(value)
	for _, existing := range AllowedInterfacePoeTypeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfacePoeTypeValue", value)
}

// NewInterfacePoeTypeValueFromValue returns a pointer to a valid InterfacePoeTypeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfacePoeTypeValueFromValue(v string) (*InterfacePoeTypeValue, error) {
	ev := InterfacePoeTypeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfacePoeTypeValue: valid values are %v", v, AllowedInterfacePoeTypeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfacePoeTypeValue) IsValid() bool {
	for _, existing := range AllowedInterfacePoeTypeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interface_poe_type_value value
func (v InterfacePoeTypeValue) Ptr() *InterfacePoeTypeValue {
	return &v
}

type NullableInterfacePoeTypeValue struct {
	value *InterfacePoeTypeValue
	isSet bool
}

func (v NullableInterfacePoeTypeValue) Get() *InterfacePoeTypeValue {
	return v.value
}

func (v *NullableInterfacePoeTypeValue) Set(val *InterfacePoeTypeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfacePoeTypeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfacePoeTypeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfacePoeTypeValue(val *InterfacePoeTypeValue) *NullableInterfacePoeTypeValue {
	return &NullableInterfacePoeTypeValue{value: val, isSet: true}
}

func (v NullableInterfacePoeTypeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfacePoeTypeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
