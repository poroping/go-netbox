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

// PatchedWritableTunnelRequestGroup - struct for PatchedWritableTunnelRequestGroup
type PatchedWritableTunnelRequestGroup struct {
	BriefTunnelGroupRequest *BriefTunnelGroupRequest
	Int32                   *int32
}

// BriefTunnelGroupRequestAsPatchedWritableTunnelRequestGroup is a convenience function that returns BriefTunnelGroupRequest wrapped in PatchedWritableTunnelRequestGroup
func BriefTunnelGroupRequestAsPatchedWritableTunnelRequestGroup(v *BriefTunnelGroupRequest) PatchedWritableTunnelRequestGroup {
	return PatchedWritableTunnelRequestGroup{
		BriefTunnelGroupRequest: v,
	}
}

// int32AsPatchedWritableTunnelRequestGroup is a convenience function that returns int32 wrapped in PatchedWritableTunnelRequestGroup
func Int32AsPatchedWritableTunnelRequestGroup(v *int32) PatchedWritableTunnelRequestGroup {
	return PatchedWritableTunnelRequestGroup{
		Int32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchedWritableTunnelRequestGroup) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into BriefTunnelGroupRequest
	err = newStrictDecoder(data).Decode(&dst.BriefTunnelGroupRequest)
	if err == nil {
		jsonBriefTunnelGroupRequest, _ := json.Marshal(dst.BriefTunnelGroupRequest)
		if string(jsonBriefTunnelGroupRequest) == "{}" { // empty struct
			dst.BriefTunnelGroupRequest = nil
		} else {
			if err = validator.Validate(dst.BriefTunnelGroupRequest); err != nil {
				dst.BriefTunnelGroupRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.BriefTunnelGroupRequest = nil
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
		dst.BriefTunnelGroupRequest = nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PatchedWritableTunnelRequestGroup)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PatchedWritableTunnelRequestGroup)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchedWritableTunnelRequestGroup) MarshalJSON() ([]byte, error) {
	if src.BriefTunnelGroupRequest != nil {
		return json.Marshal(&src.BriefTunnelGroupRequest)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchedWritableTunnelRequestGroup) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BriefTunnelGroupRequest != nil {
		return obj.BriefTunnelGroupRequest
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullablePatchedWritableTunnelRequestGroup struct {
	value *PatchedWritableTunnelRequestGroup
	isSet bool
}

func (v NullablePatchedWritableTunnelRequestGroup) Get() *PatchedWritableTunnelRequestGroup {
	return v.value
}

func (v *NullablePatchedWritableTunnelRequestGroup) Set(val *PatchedWritableTunnelRequestGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableTunnelRequestGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableTunnelRequestGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableTunnelRequestGroup(val *PatchedWritableTunnelRequestGroup) *NullablePatchedWritableTunnelRequestGroup {
	return &NullablePatchedWritableTunnelRequestGroup{value: val, isSet: true}
}

func (v NullablePatchedWritableTunnelRequestGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableTunnelRequestGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
