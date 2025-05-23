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

// DeviceWithConfigContextRequestVirtualChassis - struct for DeviceWithConfigContextRequestVirtualChassis
type DeviceWithConfigContextRequestVirtualChassis struct {
	BriefVirtualChassisRequest *BriefVirtualChassisRequest
	Int32                      *int32
}

// BriefVirtualChassisRequestAsDeviceWithConfigContextRequestVirtualChassis is a convenience function that returns BriefVirtualChassisRequest wrapped in DeviceWithConfigContextRequestVirtualChassis
func BriefVirtualChassisRequestAsDeviceWithConfigContextRequestVirtualChassis(v *BriefVirtualChassisRequest) DeviceWithConfigContextRequestVirtualChassis {
	return DeviceWithConfigContextRequestVirtualChassis{
		BriefVirtualChassisRequest: v,
	}
}

// int32AsDeviceWithConfigContextRequestVirtualChassis is a convenience function that returns int32 wrapped in DeviceWithConfigContextRequestVirtualChassis
func Int32AsDeviceWithConfigContextRequestVirtualChassis(v *int32) DeviceWithConfigContextRequestVirtualChassis {
	return DeviceWithConfigContextRequestVirtualChassis{
		Int32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DeviceWithConfigContextRequestVirtualChassis) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into BriefVirtualChassisRequest
	err = newStrictDecoder(data).Decode(&dst.BriefVirtualChassisRequest)
	if err == nil {
		jsonBriefVirtualChassisRequest, _ := json.Marshal(dst.BriefVirtualChassisRequest)
		if string(jsonBriefVirtualChassisRequest) == "{}" { // empty struct
			dst.BriefVirtualChassisRequest = nil
		} else {
			if err = validator.Validate(dst.BriefVirtualChassisRequest); err != nil {
				dst.BriefVirtualChassisRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.BriefVirtualChassisRequest = nil
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
		dst.BriefVirtualChassisRequest = nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DeviceWithConfigContextRequestVirtualChassis)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DeviceWithConfigContextRequestVirtualChassis)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DeviceWithConfigContextRequestVirtualChassis) MarshalJSON() ([]byte, error) {
	if src.BriefVirtualChassisRequest != nil {
		return json.Marshal(&src.BriefVirtualChassisRequest)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DeviceWithConfigContextRequestVirtualChassis) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BriefVirtualChassisRequest != nil {
		return obj.BriefVirtualChassisRequest
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableDeviceWithConfigContextRequestVirtualChassis struct {
	value *DeviceWithConfigContextRequestVirtualChassis
	isSet bool
}

func (v NullableDeviceWithConfigContextRequestVirtualChassis) Get() *DeviceWithConfigContextRequestVirtualChassis {
	return v.value
}

func (v *NullableDeviceWithConfigContextRequestVirtualChassis) Set(val *DeviceWithConfigContextRequestVirtualChassis) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceWithConfigContextRequestVirtualChassis) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceWithConfigContextRequestVirtualChassis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceWithConfigContextRequestVirtualChassis(val *DeviceWithConfigContextRequestVirtualChassis) *NullableDeviceWithConfigContextRequestVirtualChassis {
	return &NullableDeviceWithConfigContextRequestVirtualChassis{value: val, isSet: true}
}

func (v NullableDeviceWithConfigContextRequestVirtualChassis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceWithConfigContextRequestVirtualChassis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
