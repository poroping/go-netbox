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

// DataSourceTypeLabel the model 'DataSourceTypeLabel'
type DataSourceTypeLabel string

// List of DataSource_type_label
const (
	DATASOURCETYPELABEL________   DataSourceTypeLabel = "---------"
	DATASOURCETYPELABEL_LOCAL     DataSourceTypeLabel = "Local"
	DATASOURCETYPELABEL_GIT       DataSourceTypeLabel = "Git"
	DATASOURCETYPELABEL_AMAZON_S3 DataSourceTypeLabel = "Amazon S3"
)

// All allowed values of DataSourceTypeLabel enum
var AllowedDataSourceTypeLabelEnumValues = []DataSourceTypeLabel{
	"---------",
	"Local",
	"Git",
	"Amazon S3",
}

func (v *DataSourceTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSourceTypeLabel(value)
	for _, existing := range AllowedDataSourceTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSourceTypeLabel", value)
}

// NewDataSourceTypeLabelFromValue returns a pointer to a valid DataSourceTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSourceTypeLabelFromValue(v string) (*DataSourceTypeLabel, error) {
	ev := DataSourceTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSourceTypeLabel: valid values are %v", v, AllowedDataSourceTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSourceTypeLabel) IsValid() bool {
	for _, existing := range AllowedDataSourceTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSource_type_label value
func (v DataSourceTypeLabel) Ptr() *DataSourceTypeLabel {
	return &v
}

type NullableDataSourceTypeLabel struct {
	value *DataSourceTypeLabel
	isSet bool
}

func (v NullableDataSourceTypeLabel) Get() *DataSourceTypeLabel {
	return v.value
}

func (v *NullableDataSourceTypeLabel) Set(val *DataSourceTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceTypeLabel(val *DataSourceTypeLabel) *NullableDataSourceTypeLabel {
	return &NullableDataSourceTypeLabel{value: val, isSet: true}
}

func (v NullableDataSourceTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
