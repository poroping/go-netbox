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

// PowerFeedStatusLabel the model 'PowerFeedStatusLabel'
type PowerFeedStatusLabel string

// List of PowerFeed_status_label
const (
	POWERFEEDSTATUSLABEL_OFFLINE PowerFeedStatusLabel = "Offline"
	POWERFEEDSTATUSLABEL_ACTIVE  PowerFeedStatusLabel = "Active"
	POWERFEEDSTATUSLABEL_PLANNED PowerFeedStatusLabel = "Planned"
	POWERFEEDSTATUSLABEL_FAILED  PowerFeedStatusLabel = "Failed"
)

// All allowed values of PowerFeedStatusLabel enum
var AllowedPowerFeedStatusLabelEnumValues = []PowerFeedStatusLabel{
	"Offline",
	"Active",
	"Planned",
	"Failed",
}

func (v *PowerFeedStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerFeedStatusLabel(value)
	for _, existing := range AllowedPowerFeedStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerFeedStatusLabel", value)
}

// NewPowerFeedStatusLabelFromValue returns a pointer to a valid PowerFeedStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerFeedStatusLabelFromValue(v string) (*PowerFeedStatusLabel, error) {
	ev := PowerFeedStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerFeedStatusLabel: valid values are %v", v, AllowedPowerFeedStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerFeedStatusLabel) IsValid() bool {
	for _, existing := range AllowedPowerFeedStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerFeed_status_label value
func (v PowerFeedStatusLabel) Ptr() *PowerFeedStatusLabel {
	return &v
}

type NullablePowerFeedStatusLabel struct {
	value *PowerFeedStatusLabel
	isSet bool
}

func (v NullablePowerFeedStatusLabel) Get() *PowerFeedStatusLabel {
	return v.value
}

func (v *NullablePowerFeedStatusLabel) Set(val *PowerFeedStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedStatusLabel(val *PowerFeedStatusLabel) *NullablePowerFeedStatusLabel {
	return &NullablePowerFeedStatusLabel{value: val, isSet: true}
}

func (v NullablePowerFeedStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
