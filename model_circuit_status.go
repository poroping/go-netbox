/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.3.0-Docker-3.3.0 (4.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the CircuitStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CircuitStatus{}

// CircuitStatus struct for CircuitStatus
type CircuitStatus struct {
	Value                *CircuitStatusValue `json:"value,omitempty"`
	Label                *CircuitStatusLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CircuitStatus CircuitStatus

// NewCircuitStatus instantiates a new CircuitStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuitStatus() *CircuitStatus {
	this := CircuitStatus{}
	return &this
}

// NewCircuitStatusWithDefaults instantiates a new CircuitStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitStatusWithDefaults() *CircuitStatus {
	this := CircuitStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CircuitStatus) GetValue() CircuitStatusValue {
	if o == nil || IsNil(o.Value) {
		var ret CircuitStatusValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitStatus) GetValueOk() (*CircuitStatusValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CircuitStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CircuitStatusValue and assigns it to the Value field.
func (o *CircuitStatus) SetValue(v CircuitStatusValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CircuitStatus) GetLabel() CircuitStatusLabel {
	if o == nil || IsNil(o.Label) {
		var ret CircuitStatusLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitStatus) GetLabelOk() (*CircuitStatusLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CircuitStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given CircuitStatusLabel and assigns it to the Label field.
func (o *CircuitStatus) SetLabel(v CircuitStatusLabel) {
	o.Label = &v
}

func (o CircuitStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CircuitStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CircuitStatus) UnmarshalJSON(data []byte) (err error) {
	varCircuitStatus := _CircuitStatus{}

	err = json.Unmarshal(data, &varCircuitStatus)

	if err != nil {
		return err
	}

	*o = CircuitStatus(varCircuitStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCircuitStatus struct {
	value *CircuitStatus
	isSet bool
}

func (v NullableCircuitStatus) Get() *CircuitStatus {
	return v.value
}

func (v *NullableCircuitStatus) Set(val *CircuitStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitStatus(val *CircuitStatus) *NullableCircuitStatus {
	return &NullableCircuitStatus{value: val, isSet: true}
}

func (v NullableCircuitStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
