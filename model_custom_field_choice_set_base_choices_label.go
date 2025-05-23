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

// CustomFieldChoiceSetBaseChoicesLabel the model 'CustomFieldChoiceSetBaseChoicesLabel'
type CustomFieldChoiceSetBaseChoicesLabel string

// List of CustomFieldChoiceSet_base_choices_label
const (
	CUSTOMFIELDCHOICESETBASECHOICESLABEL_IATA__AIRPORT_CODES       CustomFieldChoiceSetBaseChoicesLabel = "IATA (Airport codes)"
	CUSTOMFIELDCHOICESETBASECHOICESLABEL_ISO_3166__COUNTRY_CODES   CustomFieldChoiceSetBaseChoicesLabel = "ISO 3166 (Country codes)"
	CUSTOMFIELDCHOICESETBASECHOICESLABEL_UN_LOCODE__LOCATION_CODES CustomFieldChoiceSetBaseChoicesLabel = "UN/LOCODE (Location codes)"
)

// All allowed values of CustomFieldChoiceSetBaseChoicesLabel enum
var AllowedCustomFieldChoiceSetBaseChoicesLabelEnumValues = []CustomFieldChoiceSetBaseChoicesLabel{
	"IATA (Airport codes)",
	"ISO 3166 (Country codes)",
	"UN/LOCODE (Location codes)",
}

func (v *CustomFieldChoiceSetBaseChoicesLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomFieldChoiceSetBaseChoicesLabel(value)
	for _, existing := range AllowedCustomFieldChoiceSetBaseChoicesLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomFieldChoiceSetBaseChoicesLabel", value)
}

// NewCustomFieldChoiceSetBaseChoicesLabelFromValue returns a pointer to a valid CustomFieldChoiceSetBaseChoicesLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomFieldChoiceSetBaseChoicesLabelFromValue(v string) (*CustomFieldChoiceSetBaseChoicesLabel, error) {
	ev := CustomFieldChoiceSetBaseChoicesLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomFieldChoiceSetBaseChoicesLabel: valid values are %v", v, AllowedCustomFieldChoiceSetBaseChoicesLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomFieldChoiceSetBaseChoicesLabel) IsValid() bool {
	for _, existing := range AllowedCustomFieldChoiceSetBaseChoicesLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomFieldChoiceSet_base_choices_label value
func (v CustomFieldChoiceSetBaseChoicesLabel) Ptr() *CustomFieldChoiceSetBaseChoicesLabel {
	return &v
}

type NullableCustomFieldChoiceSetBaseChoicesLabel struct {
	value *CustomFieldChoiceSetBaseChoicesLabel
	isSet bool
}

func (v NullableCustomFieldChoiceSetBaseChoicesLabel) Get() *CustomFieldChoiceSetBaseChoicesLabel {
	return v.value
}

func (v *NullableCustomFieldChoiceSetBaseChoicesLabel) Set(val *CustomFieldChoiceSetBaseChoicesLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldChoiceSetBaseChoicesLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldChoiceSetBaseChoicesLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldChoiceSetBaseChoicesLabel(val *CustomFieldChoiceSetBaseChoicesLabel) *NullableCustomFieldChoiceSetBaseChoicesLabel {
	return &NullableCustomFieldChoiceSetBaseChoicesLabel{value: val, isSet: true}
}

func (v NullableCustomFieldChoiceSetBaseChoicesLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldChoiceSetBaseChoicesLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
