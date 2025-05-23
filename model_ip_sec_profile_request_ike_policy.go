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

// IPSecProfileRequestIkePolicy - struct for IPSecProfileRequestIkePolicy
type IPSecProfileRequestIkePolicy struct {
	BriefIKEPolicyRequest *BriefIKEPolicyRequest
	Int32                 *int32
}

// BriefIKEPolicyRequestAsIPSecProfileRequestIkePolicy is a convenience function that returns BriefIKEPolicyRequest wrapped in IPSecProfileRequestIkePolicy
func BriefIKEPolicyRequestAsIPSecProfileRequestIkePolicy(v *BriefIKEPolicyRequest) IPSecProfileRequestIkePolicy {
	return IPSecProfileRequestIkePolicy{
		BriefIKEPolicyRequest: v,
	}
}

// int32AsIPSecProfileRequestIkePolicy is a convenience function that returns int32 wrapped in IPSecProfileRequestIkePolicy
func Int32AsIPSecProfileRequestIkePolicy(v *int32) IPSecProfileRequestIkePolicy {
	return IPSecProfileRequestIkePolicy{
		Int32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IPSecProfileRequestIkePolicy) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BriefIKEPolicyRequest
	err = newStrictDecoder(data).Decode(&dst.BriefIKEPolicyRequest)
	if err == nil {
		jsonBriefIKEPolicyRequest, _ := json.Marshal(dst.BriefIKEPolicyRequest)
		if string(jsonBriefIKEPolicyRequest) == "{}" { // empty struct
			dst.BriefIKEPolicyRequest = nil
		} else {
			if err = validator.Validate(dst.BriefIKEPolicyRequest); err != nil {
				dst.BriefIKEPolicyRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.BriefIKEPolicyRequest = nil
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
		dst.BriefIKEPolicyRequest = nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IPSecProfileRequestIkePolicy)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IPSecProfileRequestIkePolicy)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IPSecProfileRequestIkePolicy) MarshalJSON() ([]byte, error) {
	if src.BriefIKEPolicyRequest != nil {
		return json.Marshal(&src.BriefIKEPolicyRequest)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IPSecProfileRequestIkePolicy) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BriefIKEPolicyRequest != nil {
		return obj.BriefIKEPolicyRequest
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableIPSecProfileRequestIkePolicy struct {
	value *IPSecProfileRequestIkePolicy
	isSet bool
}

func (v NullableIPSecProfileRequestIkePolicy) Get() *IPSecProfileRequestIkePolicy {
	return v.value
}

func (v *NullableIPSecProfileRequestIkePolicy) Set(val *IPSecProfileRequestIkePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecProfileRequestIkePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecProfileRequestIkePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecProfileRequestIkePolicy(val *IPSecProfileRequestIkePolicy) *NullableIPSecProfileRequestIkePolicy {
	return &NullableIPSecProfileRequestIkePolicy{value: val, isSet: true}
}

func (v NullableIPSecProfileRequestIkePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecProfileRequestIkePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
