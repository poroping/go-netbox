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

// checks if the PatchedWritablePowerPortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritablePowerPortRequest{}

// PatchedWritablePowerPortRequest Adds support for custom fields and tags.
type PatchedWritablePowerPortRequest struct {
	Device *BriefInterfaceRequestDevice     `json:"device,omitempty"`
	Module NullableConsolePortRequestModule `json:"module,omitempty"`
	Name   *string                          `json:"name,omitempty"`
	// Physical label
	Label *string                                     `json:"label,omitempty"`
	Type  NullablePatchedWritablePowerPortRequestType `json:"type,omitempty"`
	// Maximum power draw (watts)
	MaximumDraw NullableInt32 `json:"maximum_draw,omitempty"`
	// Allocated power draw (watts)
	AllocatedDraw NullableInt32 `json:"allocated_draw,omitempty"`
	Description   *string       `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected        *bool                  `json:"mark_connected,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritablePowerPortRequest PatchedWritablePowerPortRequest

// NewPatchedWritablePowerPortRequest instantiates a new PatchedWritablePowerPortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritablePowerPortRequest() *PatchedWritablePowerPortRequest {
	this := PatchedWritablePowerPortRequest{}
	return &this
}

// NewPatchedWritablePowerPortRequestWithDefaults instantiates a new PatchedWritablePowerPortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritablePowerPortRequestWithDefaults() *PatchedWritablePowerPortRequest {
	this := PatchedWritablePowerPortRequest{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortRequest) GetDevice() BriefInterfaceRequestDevice {
	if o == nil || IsNil(o.Device) {
		var ret BriefInterfaceRequestDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortRequest) GetDeviceOk() (*BriefInterfaceRequestDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortRequest) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given BriefInterfaceRequestDevice and assigns it to the Device field.
func (o *PatchedWritablePowerPortRequest) SetDevice(v BriefInterfaceRequestDevice) {
	o.Device = &v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerPortRequest) GetModule() ConsolePortRequestModule {
	if o == nil || IsNil(o.Module.Get()) {
		var ret ConsolePortRequestModule
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerPortRequest) GetModuleOk() (*ConsolePortRequestModule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortRequest) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableConsolePortRequestModule and assigns it to the Module field.
func (o *PatchedWritablePowerPortRequest) SetModule(v ConsolePortRequestModule) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil
func (o *PatchedWritablePowerPortRequest) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *PatchedWritablePowerPortRequest) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritablePowerPortRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritablePowerPortRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerPortRequest) GetType() PatchedWritablePowerPortRequestType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret PatchedWritablePowerPortRequestType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerPortRequest) GetTypeOk() (*PatchedWritablePowerPortRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortRequest) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullablePatchedWritablePowerPortRequestType and assigns it to the Type field.
func (o *PatchedWritablePowerPortRequest) SetType(v PatchedWritablePowerPortRequestType) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *PatchedWritablePowerPortRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *PatchedWritablePowerPortRequest) UnsetType() {
	o.Type.Unset()
}

// GetMaximumDraw returns the MaximumDraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerPortRequest) GetMaximumDraw() int32 {
	if o == nil || IsNil(o.MaximumDraw.Get()) {
		var ret int32
		return ret
	}
	return *o.MaximumDraw.Get()
}

// GetMaximumDrawOk returns a tuple with the MaximumDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerPortRequest) GetMaximumDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaximumDraw.Get(), o.MaximumDraw.IsSet()
}

// HasMaximumDraw returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortRequest) HasMaximumDraw() bool {
	if o != nil && o.MaximumDraw.IsSet() {
		return true
	}

	return false
}

// SetMaximumDraw gets a reference to the given NullableInt32 and assigns it to the MaximumDraw field.
func (o *PatchedWritablePowerPortRequest) SetMaximumDraw(v int32) {
	o.MaximumDraw.Set(&v)
}

// SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil
func (o *PatchedWritablePowerPortRequest) SetMaximumDrawNil() {
	o.MaximumDraw.Set(nil)
}

// UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
func (o *PatchedWritablePowerPortRequest) UnsetMaximumDraw() {
	o.MaximumDraw.Unset()
}

// GetAllocatedDraw returns the AllocatedDraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerPortRequest) GetAllocatedDraw() int32 {
	if o == nil || IsNil(o.AllocatedDraw.Get()) {
		var ret int32
		return ret
	}
	return *o.AllocatedDraw.Get()
}

// GetAllocatedDrawOk returns a tuple with the AllocatedDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerPortRequest) GetAllocatedDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllocatedDraw.Get(), o.AllocatedDraw.IsSet()
}

// HasAllocatedDraw returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortRequest) HasAllocatedDraw() bool {
	if o != nil && o.AllocatedDraw.IsSet() {
		return true
	}

	return false
}

// SetAllocatedDraw gets a reference to the given NullableInt32 and assigns it to the AllocatedDraw field.
func (o *PatchedWritablePowerPortRequest) SetAllocatedDraw(v int32) {
	o.AllocatedDraw.Set(&v)
}

// SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil
func (o *PatchedWritablePowerPortRequest) SetAllocatedDrawNil() {
	o.AllocatedDraw.Set(nil)
}

// UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
func (o *PatchedWritablePowerPortRequest) UnsetAllocatedDraw() {
	o.AllocatedDraw.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritablePowerPortRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *PatchedWritablePowerPortRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritablePowerPortRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritablePowerPortRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritablePowerPortRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritablePowerPortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.MaximumDraw.IsSet() {
		toSerialize["maximum_draw"] = o.MaximumDraw.Get()
	}
	if o.AllocatedDraw.IsSet() {
		toSerialize["allocated_draw"] = o.AllocatedDraw.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritablePowerPortRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritablePowerPortRequest := _PatchedWritablePowerPortRequest{}

	err = json.Unmarshal(data, &varPatchedWritablePowerPortRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritablePowerPortRequest(varPatchedWritablePowerPortRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "maximum_draw")
		delete(additionalProperties, "allocated_draw")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritablePowerPortRequest struct {
	value *PatchedWritablePowerPortRequest
	isSet bool
}

func (v NullablePatchedWritablePowerPortRequest) Get() *PatchedWritablePowerPortRequest {
	return v.value
}

func (v *NullablePatchedWritablePowerPortRequest) Set(val *PatchedWritablePowerPortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerPortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerPortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerPortRequest(val *PatchedWritablePowerPortRequest) *NullablePatchedWritablePowerPortRequest {
	return &NullablePatchedWritablePowerPortRequest{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerPortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerPortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
