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

// DeviceWithConfigContextRequestCluster - struct for DeviceWithConfigContextRequestCluster
type DeviceWithConfigContextRequestCluster struct {
	BriefClusterRequest *BriefClusterRequest
	Int32               *int32
}

// BriefClusterRequestAsDeviceWithConfigContextRequestCluster is a convenience function that returns BriefClusterRequest wrapped in DeviceWithConfigContextRequestCluster
func BriefClusterRequestAsDeviceWithConfigContextRequestCluster(v *BriefClusterRequest) DeviceWithConfigContextRequestCluster {
	return DeviceWithConfigContextRequestCluster{
		BriefClusterRequest: v,
	}
}

// int32AsDeviceWithConfigContextRequestCluster is a convenience function that returns int32 wrapped in DeviceWithConfigContextRequestCluster
func Int32AsDeviceWithConfigContextRequestCluster(v *int32) DeviceWithConfigContextRequestCluster {
	return DeviceWithConfigContextRequestCluster{
		Int32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DeviceWithConfigContextRequestCluster) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into BriefClusterRequest
	err = newStrictDecoder(data).Decode(&dst.BriefClusterRequest)
	if err == nil {
		jsonBriefClusterRequest, _ := json.Marshal(dst.BriefClusterRequest)
		if string(jsonBriefClusterRequest) == "{}" { // empty struct
			dst.BriefClusterRequest = nil
		} else {
			if err = validator.Validate(dst.BriefClusterRequest); err != nil {
				dst.BriefClusterRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.BriefClusterRequest = nil
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
		dst.BriefClusterRequest = nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DeviceWithConfigContextRequestCluster)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DeviceWithConfigContextRequestCluster)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DeviceWithConfigContextRequestCluster) MarshalJSON() ([]byte, error) {
	if src.BriefClusterRequest != nil {
		return json.Marshal(&src.BriefClusterRequest)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DeviceWithConfigContextRequestCluster) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BriefClusterRequest != nil {
		return obj.BriefClusterRequest
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableDeviceWithConfigContextRequestCluster struct {
	value *DeviceWithConfigContextRequestCluster
	isSet bool
}

func (v NullableDeviceWithConfigContextRequestCluster) Get() *DeviceWithConfigContextRequestCluster {
	return v.value
}

func (v *NullableDeviceWithConfigContextRequestCluster) Set(val *DeviceWithConfigContextRequestCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceWithConfigContextRequestCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceWithConfigContextRequestCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceWithConfigContextRequestCluster(val *DeviceWithConfigContextRequestCluster) *NullableDeviceWithConfigContextRequestCluster {
	return &NullableDeviceWithConfigContextRequestCluster{value: val, isSet: true}
}

func (v NullableDeviceWithConfigContextRequestCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceWithConfigContextRequestCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
