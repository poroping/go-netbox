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

// PowerOutletRequestFeedLeg * `A` - A * `B` - B * `C` - C
type PowerOutletRequestFeedLeg string

// List of PowerOutletRequest_feed_leg
const (
	POWEROUTLETREQUESTFEEDLEG_A     PowerOutletRequestFeedLeg = "A"
	POWEROUTLETREQUESTFEEDLEG_B     PowerOutletRequestFeedLeg = "B"
	POWEROUTLETREQUESTFEEDLEG_C     PowerOutletRequestFeedLeg = "C"
	POWEROUTLETREQUESTFEEDLEG_EMPTY PowerOutletRequestFeedLeg = ""
)

// All allowed values of PowerOutletRequestFeedLeg enum
var AllowedPowerOutletRequestFeedLegEnumValues = []PowerOutletRequestFeedLeg{
	"A",
	"B",
	"C",
	"",
}

func (v *PowerOutletRequestFeedLeg) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerOutletRequestFeedLeg(value)
	for _, existing := range AllowedPowerOutletRequestFeedLegEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerOutletRequestFeedLeg", value)
}

// NewPowerOutletRequestFeedLegFromValue returns a pointer to a valid PowerOutletRequestFeedLeg
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerOutletRequestFeedLegFromValue(v string) (*PowerOutletRequestFeedLeg, error) {
	ev := PowerOutletRequestFeedLeg(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerOutletRequestFeedLeg: valid values are %v", v, AllowedPowerOutletRequestFeedLegEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerOutletRequestFeedLeg) IsValid() bool {
	for _, existing := range AllowedPowerOutletRequestFeedLegEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerOutletRequest_feed_leg value
func (v PowerOutletRequestFeedLeg) Ptr() *PowerOutletRequestFeedLeg {
	return &v
}

type NullablePowerOutletRequestFeedLeg struct {
	value *PowerOutletRequestFeedLeg
	isSet bool
}

func (v NullablePowerOutletRequestFeedLeg) Get() *PowerOutletRequestFeedLeg {
	return v.value
}

func (v *NullablePowerOutletRequestFeedLeg) Set(val *PowerOutletRequestFeedLeg) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerOutletRequestFeedLeg) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerOutletRequestFeedLeg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerOutletRequestFeedLeg(val *PowerOutletRequestFeedLeg) *NullablePowerOutletRequestFeedLeg {
	return &NullablePowerOutletRequestFeedLeg{value: val, isSet: true}
}

func (v NullablePowerOutletRequestFeedLeg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerOutletRequestFeedLeg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
