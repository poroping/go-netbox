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

// FHRPGroupAssignmentRequestGroup - struct for FHRPGroupAssignmentRequestGroup
type FHRPGroupAssignmentRequestGroup struct {
	BriefFHRPGroupRequest *BriefFHRPGroupRequest
	Int32                 *int32
}

// BriefFHRPGroupRequestAsFHRPGroupAssignmentRequestGroup is a convenience function that returns BriefFHRPGroupRequest wrapped in FHRPGroupAssignmentRequestGroup
func BriefFHRPGroupRequestAsFHRPGroupAssignmentRequestGroup(v *BriefFHRPGroupRequest) FHRPGroupAssignmentRequestGroup {
	return FHRPGroupAssignmentRequestGroup{
		BriefFHRPGroupRequest: v,
	}
}

// int32AsFHRPGroupAssignmentRequestGroup is a convenience function that returns int32 wrapped in FHRPGroupAssignmentRequestGroup
func Int32AsFHRPGroupAssignmentRequestGroup(v *int32) FHRPGroupAssignmentRequestGroup {
	return FHRPGroupAssignmentRequestGroup{
		Int32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FHRPGroupAssignmentRequestGroup) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BriefFHRPGroupRequest
	err = newStrictDecoder(data).Decode(&dst.BriefFHRPGroupRequest)
	if err == nil {
		jsonBriefFHRPGroupRequest, _ := json.Marshal(dst.BriefFHRPGroupRequest)
		if string(jsonBriefFHRPGroupRequest) == "{}" { // empty struct
			dst.BriefFHRPGroupRequest = nil
		} else {
			if err = validator.Validate(dst.BriefFHRPGroupRequest); err != nil {
				dst.BriefFHRPGroupRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.BriefFHRPGroupRequest = nil
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
		dst.BriefFHRPGroupRequest = nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(FHRPGroupAssignmentRequestGroup)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(FHRPGroupAssignmentRequestGroup)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FHRPGroupAssignmentRequestGroup) MarshalJSON() ([]byte, error) {
	if src.BriefFHRPGroupRequest != nil {
		return json.Marshal(&src.BriefFHRPGroupRequest)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FHRPGroupAssignmentRequestGroup) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BriefFHRPGroupRequest != nil {
		return obj.BriefFHRPGroupRequest
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableFHRPGroupAssignmentRequestGroup struct {
	value *FHRPGroupAssignmentRequestGroup
	isSet bool
}

func (v NullableFHRPGroupAssignmentRequestGroup) Get() *FHRPGroupAssignmentRequestGroup {
	return v.value
}

func (v *NullableFHRPGroupAssignmentRequestGroup) Set(val *FHRPGroupAssignmentRequestGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableFHRPGroupAssignmentRequestGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableFHRPGroupAssignmentRequestGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFHRPGroupAssignmentRequestGroup(val *FHRPGroupAssignmentRequestGroup) *NullableFHRPGroupAssignmentRequestGroup {
	return &NullableFHRPGroupAssignmentRequestGroup{value: val, isSet: true}
}

func (v NullableFHRPGroupAssignmentRequestGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFHRPGroupAssignmentRequestGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
