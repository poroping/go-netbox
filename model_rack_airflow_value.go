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

// RackAirflowValue * `front-to-rear` - Front to rear * `rear-to-front` - Rear to front
type RackAirflowValue string

// List of Rack_airflow_value
const (
	RACKAIRFLOWVALUE_FRONT_TO_REAR RackAirflowValue = "front-to-rear"
	RACKAIRFLOWVALUE_REAR_TO_FRONT RackAirflowValue = "rear-to-front"
	RACKAIRFLOWVALUE_EMPTY         RackAirflowValue = ""
)

// All allowed values of RackAirflowValue enum
var AllowedRackAirflowValueEnumValues = []RackAirflowValue{
	"front-to-rear",
	"rear-to-front",
	"",
}

func (v *RackAirflowValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackAirflowValue(value)
	for _, existing := range AllowedRackAirflowValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackAirflowValue", value)
}

// NewRackAirflowValueFromValue returns a pointer to a valid RackAirflowValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackAirflowValueFromValue(v string) (*RackAirflowValue, error) {
	ev := RackAirflowValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackAirflowValue: valid values are %v", v, AllowedRackAirflowValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackAirflowValue) IsValid() bool {
	for _, existing := range AllowedRackAirflowValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Rack_airflow_value value
func (v RackAirflowValue) Ptr() *RackAirflowValue {
	return &v
}

type NullableRackAirflowValue struct {
	value *RackAirflowValue
	isSet bool
}

func (v NullableRackAirflowValue) Get() *RackAirflowValue {
	return v.value
}

func (v *NullableRackAirflowValue) Set(val *RackAirflowValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRackAirflowValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRackAirflowValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackAirflowValue(val *RackAirflowValue) *NullableRackAirflowValue {
	return &NullableRackAirflowValue{value: val, isSet: true}
}

func (v NullableRackAirflowValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackAirflowValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
