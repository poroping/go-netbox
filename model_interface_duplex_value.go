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

// InterfaceDuplexValue * `half` - Half * `full` - Full * `auto` - Auto
type InterfaceDuplexValue string

// List of Interface_duplex_value
const (
	INTERFACEDUPLEXVALUE_HALF  InterfaceDuplexValue = "half"
	INTERFACEDUPLEXVALUE_FULL  InterfaceDuplexValue = "full"
	INTERFACEDUPLEXVALUE_AUTO  InterfaceDuplexValue = "auto"
	INTERFACEDUPLEXVALUE_EMPTY InterfaceDuplexValue = ""
)

// All allowed values of InterfaceDuplexValue enum
var AllowedInterfaceDuplexValueEnumValues = []InterfaceDuplexValue{
	"half",
	"full",
	"auto",
	"",
}

func (v *InterfaceDuplexValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceDuplexValue(value)
	for _, existing := range AllowedInterfaceDuplexValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceDuplexValue", value)
}

// NewInterfaceDuplexValueFromValue returns a pointer to a valid InterfaceDuplexValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceDuplexValueFromValue(v string) (*InterfaceDuplexValue, error) {
	ev := InterfaceDuplexValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceDuplexValue: valid values are %v", v, AllowedInterfaceDuplexValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceDuplexValue) IsValid() bool {
	for _, existing := range AllowedInterfaceDuplexValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interface_duplex_value value
func (v InterfaceDuplexValue) Ptr() *InterfaceDuplexValue {
	return &v
}

type NullableInterfaceDuplexValue struct {
	value *InterfaceDuplexValue
	isSet bool
}

func (v NullableInterfaceDuplexValue) Get() *InterfaceDuplexValue {
	return v.value
}

func (v *NullableInterfaceDuplexValue) Set(val *InterfaceDuplexValue) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceDuplexValue) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceDuplexValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceDuplexValue(val *InterfaceDuplexValue) *NullableInterfaceDuplexValue {
	return &NullableInterfaceDuplexValue{value: val, isSet: true}
}

func (v NullableInterfaceDuplexValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceDuplexValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
