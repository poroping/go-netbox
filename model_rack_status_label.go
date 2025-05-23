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

// RackStatusLabel the model 'RackStatusLabel'
type RackStatusLabel string

// List of Rack_status_label
const (
	RACKSTATUSLABEL_RESERVED   RackStatusLabel = "Reserved"
	RACKSTATUSLABEL_AVAILABLE  RackStatusLabel = "Available"
	RACKSTATUSLABEL_PLANNED    RackStatusLabel = "Planned"
	RACKSTATUSLABEL_ACTIVE     RackStatusLabel = "Active"
	RACKSTATUSLABEL_DEPRECATED RackStatusLabel = "Deprecated"
)

// All allowed values of RackStatusLabel enum
var AllowedRackStatusLabelEnumValues = []RackStatusLabel{
	"Reserved",
	"Available",
	"Planned",
	"Active",
	"Deprecated",
}

func (v *RackStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackStatusLabel(value)
	for _, existing := range AllowedRackStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackStatusLabel", value)
}

// NewRackStatusLabelFromValue returns a pointer to a valid RackStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackStatusLabelFromValue(v string) (*RackStatusLabel, error) {
	ev := RackStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackStatusLabel: valid values are %v", v, AllowedRackStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackStatusLabel) IsValid() bool {
	for _, existing := range AllowedRackStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Rack_status_label value
func (v RackStatusLabel) Ptr() *RackStatusLabel {
	return &v
}

type NullableRackStatusLabel struct {
	value *RackStatusLabel
	isSet bool
}

func (v NullableRackStatusLabel) Get() *RackStatusLabel {
	return v.value
}

func (v *NullableRackStatusLabel) Set(val *RackStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableRackStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableRackStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackStatusLabel(val *RackStatusLabel) *NullableRackStatusLabel {
	return &NullableRackStatusLabel{value: val, isSet: true}
}

func (v NullableRackStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
