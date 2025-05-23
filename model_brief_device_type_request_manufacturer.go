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

// BriefDeviceTypeRequestManufacturer - struct for BriefDeviceTypeRequestManufacturer
type BriefDeviceTypeRequestManufacturer struct {
	BriefManufacturerRequest *BriefManufacturerRequest
	Int32                    *int32
}

// BriefManufacturerRequestAsBriefDeviceTypeRequestManufacturer is a convenience function that returns BriefManufacturerRequest wrapped in BriefDeviceTypeRequestManufacturer
func BriefManufacturerRequestAsBriefDeviceTypeRequestManufacturer(v *BriefManufacturerRequest) BriefDeviceTypeRequestManufacturer {
	return BriefDeviceTypeRequestManufacturer{
		BriefManufacturerRequest: v,
	}
}

// int32AsBriefDeviceTypeRequestManufacturer is a convenience function that returns int32 wrapped in BriefDeviceTypeRequestManufacturer
func Int32AsBriefDeviceTypeRequestManufacturer(v *int32) BriefDeviceTypeRequestManufacturer {
	return BriefDeviceTypeRequestManufacturer{
		Int32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BriefDeviceTypeRequestManufacturer) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BriefManufacturerRequest
	err = newStrictDecoder(data).Decode(&dst.BriefManufacturerRequest)
	if err == nil {
		jsonBriefManufacturerRequest, _ := json.Marshal(dst.BriefManufacturerRequest)
		if string(jsonBriefManufacturerRequest) == "{}" { // empty struct
			dst.BriefManufacturerRequest = nil
		} else {
			if err = validator.Validate(dst.BriefManufacturerRequest); err != nil {
				dst.BriefManufacturerRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.BriefManufacturerRequest = nil
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
		dst.BriefManufacturerRequest = nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BriefDeviceTypeRequestManufacturer)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BriefDeviceTypeRequestManufacturer)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BriefDeviceTypeRequestManufacturer) MarshalJSON() ([]byte, error) {
	if src.BriefManufacturerRequest != nil {
		return json.Marshal(&src.BriefManufacturerRequest)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BriefDeviceTypeRequestManufacturer) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BriefManufacturerRequest != nil {
		return obj.BriefManufacturerRequest
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableBriefDeviceTypeRequestManufacturer struct {
	value *BriefDeviceTypeRequestManufacturer
	isSet bool
}

func (v NullableBriefDeviceTypeRequestManufacturer) Get() *BriefDeviceTypeRequestManufacturer {
	return v.value
}

func (v *NullableBriefDeviceTypeRequestManufacturer) Set(val *BriefDeviceTypeRequestManufacturer) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefDeviceTypeRequestManufacturer) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefDeviceTypeRequestManufacturer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefDeviceTypeRequestManufacturer(val *BriefDeviceTypeRequestManufacturer) *NullableBriefDeviceTypeRequestManufacturer {
	return &NullableBriefDeviceTypeRequestManufacturer{value: val, isSet: true}
}

func (v NullableBriefDeviceTypeRequestManufacturer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefDeviceTypeRequestManufacturer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
