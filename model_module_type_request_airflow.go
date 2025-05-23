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

// ModuleTypeRequestAirflow * `front-to-rear` - Front to rear * `rear-to-front` - Rear to front * `left-to-right` - Left to right * `right-to-left` - Right to left * `side-to-rear` - Side to rear * `passive` - Passive
type ModuleTypeRequestAirflow string

// List of ModuleTypeRequest_airflow
const (
	MODULETYPEREQUESTAIRFLOW_FRONT_TO_REAR ModuleTypeRequestAirflow = "front-to-rear"
	MODULETYPEREQUESTAIRFLOW_REAR_TO_FRONT ModuleTypeRequestAirflow = "rear-to-front"
	MODULETYPEREQUESTAIRFLOW_LEFT_TO_RIGHT ModuleTypeRequestAirflow = "left-to-right"
	MODULETYPEREQUESTAIRFLOW_RIGHT_TO_LEFT ModuleTypeRequestAirflow = "right-to-left"
	MODULETYPEREQUESTAIRFLOW_SIDE_TO_REAR  ModuleTypeRequestAirflow = "side-to-rear"
	MODULETYPEREQUESTAIRFLOW_PASSIVE       ModuleTypeRequestAirflow = "passive"
	MODULETYPEREQUESTAIRFLOW_EMPTY         ModuleTypeRequestAirflow = ""
)

// All allowed values of ModuleTypeRequestAirflow enum
var AllowedModuleTypeRequestAirflowEnumValues = []ModuleTypeRequestAirflow{
	"front-to-rear",
	"rear-to-front",
	"left-to-right",
	"right-to-left",
	"side-to-rear",
	"passive",
	"",
}

func (v *ModuleTypeRequestAirflow) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModuleTypeRequestAirflow(value)
	for _, existing := range AllowedModuleTypeRequestAirflowEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModuleTypeRequestAirflow", value)
}

// NewModuleTypeRequestAirflowFromValue returns a pointer to a valid ModuleTypeRequestAirflow
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModuleTypeRequestAirflowFromValue(v string) (*ModuleTypeRequestAirflow, error) {
	ev := ModuleTypeRequestAirflow(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModuleTypeRequestAirflow: valid values are %v", v, AllowedModuleTypeRequestAirflowEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModuleTypeRequestAirflow) IsValid() bool {
	for _, existing := range AllowedModuleTypeRequestAirflowEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModuleTypeRequest_airflow value
func (v ModuleTypeRequestAirflow) Ptr() *ModuleTypeRequestAirflow {
	return &v
}

type NullableModuleTypeRequestAirflow struct {
	value *ModuleTypeRequestAirflow
	isSet bool
}

func (v NullableModuleTypeRequestAirflow) Get() *ModuleTypeRequestAirflow {
	return v.value
}

func (v *NullableModuleTypeRequestAirflow) Set(val *ModuleTypeRequestAirflow) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleTypeRequestAirflow) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleTypeRequestAirflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleTypeRequestAirflow(val *ModuleTypeRequestAirflow) *NullableModuleTypeRequestAirflow {
	return &NullableModuleTypeRequestAirflow{value: val, isSet: true}
}

func (v NullableModuleTypeRequestAirflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleTypeRequestAirflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
