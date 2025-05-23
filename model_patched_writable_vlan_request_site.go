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

// PatchedWritableVLANRequestSite - struct for PatchedWritableVLANRequestSite
type PatchedWritableVLANRequestSite struct {
	BriefSiteRequest *BriefSiteRequest
	Int32            *int32
}

// BriefSiteRequestAsPatchedWritableVLANRequestSite is a convenience function that returns BriefSiteRequest wrapped in PatchedWritableVLANRequestSite
func BriefSiteRequestAsPatchedWritableVLANRequestSite(v *BriefSiteRequest) PatchedWritableVLANRequestSite {
	return PatchedWritableVLANRequestSite{
		BriefSiteRequest: v,
	}
}

// int32AsPatchedWritableVLANRequestSite is a convenience function that returns int32 wrapped in PatchedWritableVLANRequestSite
func Int32AsPatchedWritableVLANRequestSite(v *int32) PatchedWritableVLANRequestSite {
	return PatchedWritableVLANRequestSite{
		Int32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchedWritableVLANRequestSite) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into BriefSiteRequest
	err = newStrictDecoder(data).Decode(&dst.BriefSiteRequest)
	if err == nil {
		jsonBriefSiteRequest, _ := json.Marshal(dst.BriefSiteRequest)
		if string(jsonBriefSiteRequest) == "{}" { // empty struct
			dst.BriefSiteRequest = nil
		} else {
			if err = validator.Validate(dst.BriefSiteRequest); err != nil {
				dst.BriefSiteRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.BriefSiteRequest = nil
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
		dst.BriefSiteRequest = nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PatchedWritableVLANRequestSite)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PatchedWritableVLANRequestSite)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchedWritableVLANRequestSite) MarshalJSON() ([]byte, error) {
	if src.BriefSiteRequest != nil {
		return json.Marshal(&src.BriefSiteRequest)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchedWritableVLANRequestSite) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BriefSiteRequest != nil {
		return obj.BriefSiteRequest
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullablePatchedWritableVLANRequestSite struct {
	value *PatchedWritableVLANRequestSite
	isSet bool
}

func (v NullablePatchedWritableVLANRequestSite) Get() *PatchedWritableVLANRequestSite {
	return v.value
}

func (v *NullablePatchedWritableVLANRequestSite) Set(val *PatchedWritableVLANRequestSite) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableVLANRequestSite) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableVLANRequestSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableVLANRequestSite(val *PatchedWritableVLANRequestSite) *NullablePatchedWritableVLANRequestSite {
	return &NullablePatchedWritableVLANRequestSite{value: val, isSet: true}
}

func (v NullablePatchedWritableVLANRequestSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableVLANRequestSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
