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

	"gopkg.in/validator.v2"
)

// PatchedWritableTunnelRequestIpsecProfile - struct for PatchedWritableTunnelRequestIpsecProfile
type PatchedWritableTunnelRequestIpsecProfile struct {
	BriefIPSecProfileRequest *BriefIPSecProfileRequest
	Int32                    *int32
}

// BriefIPSecProfileRequestAsPatchedWritableTunnelRequestIpsecProfile is a convenience function that returns BriefIPSecProfileRequest wrapped in PatchedWritableTunnelRequestIpsecProfile
func BriefIPSecProfileRequestAsPatchedWritableTunnelRequestIpsecProfile(v *BriefIPSecProfileRequest) PatchedWritableTunnelRequestIpsecProfile {
	return PatchedWritableTunnelRequestIpsecProfile{
		BriefIPSecProfileRequest: v,
	}
}

// int32AsPatchedWritableTunnelRequestIpsecProfile is a convenience function that returns int32 wrapped in PatchedWritableTunnelRequestIpsecProfile
func Int32AsPatchedWritableTunnelRequestIpsecProfile(v *int32) PatchedWritableTunnelRequestIpsecProfile {
	return PatchedWritableTunnelRequestIpsecProfile{
		Int32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchedWritableTunnelRequestIpsecProfile) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into BriefIPSecProfileRequest
	err = newStrictDecoder(data).Decode(&dst.BriefIPSecProfileRequest)
	if err == nil {
		jsonBriefIPSecProfileRequest, _ := json.Marshal(dst.BriefIPSecProfileRequest)
		if string(jsonBriefIPSecProfileRequest) == "{}" { // empty struct
			dst.BriefIPSecProfileRequest = nil
		} else {
			if err = validator.Validate(dst.BriefIPSecProfileRequest); err != nil {
				dst.BriefIPSecProfileRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.BriefIPSecProfileRequest = nil
	}

	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			if err = validator.Validate(dst.Int32); err != nil {
				dst.Int32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int32 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BriefIPSecProfileRequest = nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PatchedWritableTunnelRequestIpsecProfile)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PatchedWritableTunnelRequestIpsecProfile)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchedWritableTunnelRequestIpsecProfile) MarshalJSON() ([]byte, error) {
	if src.BriefIPSecProfileRequest != nil {
		return json.Marshal(&src.BriefIPSecProfileRequest)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchedWritableTunnelRequestIpsecProfile) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BriefIPSecProfileRequest != nil {
		return obj.BriefIPSecProfileRequest
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullablePatchedWritableTunnelRequestIpsecProfile struct {
	value *PatchedWritableTunnelRequestIpsecProfile
	isSet bool
}

func (v NullablePatchedWritableTunnelRequestIpsecProfile) Get() *PatchedWritableTunnelRequestIpsecProfile {
	return v.value
}

func (v *NullablePatchedWritableTunnelRequestIpsecProfile) Set(val *PatchedWritableTunnelRequestIpsecProfile) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableTunnelRequestIpsecProfile) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableTunnelRequestIpsecProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableTunnelRequestIpsecProfile(val *PatchedWritableTunnelRequestIpsecProfile) *NullablePatchedWritableTunnelRequestIpsecProfile {
	return &NullablePatchedWritableTunnelRequestIpsecProfile{value: val, isSet: true}
}

func (v NullablePatchedWritableTunnelRequestIpsecProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableTunnelRequestIpsecProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
