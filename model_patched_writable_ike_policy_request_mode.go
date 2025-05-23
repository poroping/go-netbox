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

// PatchedWritableIKEPolicyRequestMode * `aggressive` - Aggressive * `main` - Main
type PatchedWritableIKEPolicyRequestMode string

// List of PatchedWritableIKEPolicyRequest_mode
const (
	PATCHEDWRITABLEIKEPOLICYREQUESTMODE_AGGRESSIVE PatchedWritableIKEPolicyRequestMode = "aggressive"
	PATCHEDWRITABLEIKEPOLICYREQUESTMODE_MAIN       PatchedWritableIKEPolicyRequestMode = "main"
	PATCHEDWRITABLEIKEPOLICYREQUESTMODE_EMPTY      PatchedWritableIKEPolicyRequestMode = ""
)

// All allowed values of PatchedWritableIKEPolicyRequestMode enum
var AllowedPatchedWritableIKEPolicyRequestModeEnumValues = []PatchedWritableIKEPolicyRequestMode{
	"aggressive",
	"main",
	"",
}

func (v *PatchedWritableIKEPolicyRequestMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableIKEPolicyRequestMode(value)
	for _, existing := range AllowedPatchedWritableIKEPolicyRequestModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableIKEPolicyRequestMode", value)
}

// NewPatchedWritableIKEPolicyRequestModeFromValue returns a pointer to a valid PatchedWritableIKEPolicyRequestMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableIKEPolicyRequestModeFromValue(v string) (*PatchedWritableIKEPolicyRequestMode, error) {
	ev := PatchedWritableIKEPolicyRequestMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableIKEPolicyRequestMode: valid values are %v", v, AllowedPatchedWritableIKEPolicyRequestModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableIKEPolicyRequestMode) IsValid() bool {
	for _, existing := range AllowedPatchedWritableIKEPolicyRequestModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableIKEPolicyRequest_mode value
func (v PatchedWritableIKEPolicyRequestMode) Ptr() *PatchedWritableIKEPolicyRequestMode {
	return &v
}

type NullablePatchedWritableIKEPolicyRequestMode struct {
	value *PatchedWritableIKEPolicyRequestMode
	isSet bool
}

func (v NullablePatchedWritableIKEPolicyRequestMode) Get() *PatchedWritableIKEPolicyRequestMode {
	return v.value
}

func (v *NullablePatchedWritableIKEPolicyRequestMode) Set(val *PatchedWritableIKEPolicyRequestMode) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableIKEPolicyRequestMode) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableIKEPolicyRequestMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableIKEPolicyRequestMode(val *PatchedWritableIKEPolicyRequestMode) *NullablePatchedWritableIKEPolicyRequestMode {
	return &NullablePatchedWritableIKEPolicyRequestMode{value: val, isSet: true}
}

func (v NullablePatchedWritableIKEPolicyRequestMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableIKEPolicyRequestMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
