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

// checks if the PatchedObjectPermissionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedObjectPermissionRequest{}

// PatchedObjectPermissionRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedObjectPermissionRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Enabled     *bool    `json:"enabled,omitempty"`
	ObjectTypes []string `json:"object_types,omitempty"`
	// The list of actions granted by this permission
	Actions []string `json:"actions,omitempty"`
	// Queryset filter matching the applicable objects of the selected type(s)
	Constraints          interface{} `json:"constraints,omitempty"`
	Groups               []int32     `json:"groups,omitempty"`
	Users                []int32     `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedObjectPermissionRequest PatchedObjectPermissionRequest

// NewPatchedObjectPermissionRequest instantiates a new PatchedObjectPermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedObjectPermissionRequest() *PatchedObjectPermissionRequest {
	this := PatchedObjectPermissionRequest{}
	return &this
}

// NewPatchedObjectPermissionRequestWithDefaults instantiates a new PatchedObjectPermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedObjectPermissionRequestWithDefaults() *PatchedObjectPermissionRequest {
	this := PatchedObjectPermissionRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedObjectPermissionRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedObjectPermissionRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedObjectPermissionRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedObjectPermissionRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedObjectPermissionRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedObjectPermissionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedObjectPermissionRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedObjectPermissionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedObjectPermissionRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedObjectPermissionRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedObjectPermissionRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedObjectPermissionRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetObjectTypes returns the ObjectTypes field value if set, zero value otherwise.
func (o *PatchedObjectPermissionRequest) GetObjectTypes() []string {
	if o == nil || IsNil(o.ObjectTypes) {
		var ret []string
		return ret
	}
	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedObjectPermissionRequest) GetObjectTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ObjectTypes) {
		return nil, false
	}
	return o.ObjectTypes, true
}

// HasObjectTypes returns a boolean if a field has been set.
func (o *PatchedObjectPermissionRequest) HasObjectTypes() bool {
	if o != nil && !IsNil(o.ObjectTypes) {
		return true
	}

	return false
}

// SetObjectTypes gets a reference to the given []string and assigns it to the ObjectTypes field.
func (o *PatchedObjectPermissionRequest) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *PatchedObjectPermissionRequest) GetActions() []string {
	if o == nil || IsNil(o.Actions) {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedObjectPermissionRequest) GetActionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *PatchedObjectPermissionRequest) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []string and assigns it to the Actions field.
func (o *PatchedObjectPermissionRequest) SetActions(v []string) {
	o.Actions = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedObjectPermissionRequest) GetConstraints() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedObjectPermissionRequest) GetConstraintsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return &o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *PatchedObjectPermissionRequest) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given interface{} and assigns it to the Constraints field.
func (o *PatchedObjectPermissionRequest) SetConstraints(v interface{}) {
	o.Constraints = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *PatchedObjectPermissionRequest) GetGroups() []int32 {
	if o == nil || IsNil(o.Groups) {
		var ret []int32
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedObjectPermissionRequest) GetGroupsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *PatchedObjectPermissionRequest) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []int32 and assigns it to the Groups field.
func (o *PatchedObjectPermissionRequest) SetGroups(v []int32) {
	o.Groups = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *PatchedObjectPermissionRequest) GetUsers() []int32 {
	if o == nil || IsNil(o.Users) {
		var ret []int32
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedObjectPermissionRequest) GetUsersOk() ([]int32, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *PatchedObjectPermissionRequest) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []int32 and assigns it to the Users field.
func (o *PatchedObjectPermissionRequest) SetUsers(v []int32) {
	o.Users = v
}

func (o PatchedObjectPermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedObjectPermissionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ObjectTypes) {
		toSerialize["object_types"] = o.ObjectTypes
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedObjectPermissionRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedObjectPermissionRequest := _PatchedObjectPermissionRequest{}

	err = json.Unmarshal(data, &varPatchedObjectPermissionRequest)

	if err != nil {
		return err
	}

	*o = PatchedObjectPermissionRequest(varPatchedObjectPermissionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "object_types")
		delete(additionalProperties, "actions")
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedObjectPermissionRequest struct {
	value *PatchedObjectPermissionRequest
	isSet bool
}

func (v NullablePatchedObjectPermissionRequest) Get() *PatchedObjectPermissionRequest {
	return v.value
}

func (v *NullablePatchedObjectPermissionRequest) Set(val *PatchedObjectPermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedObjectPermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedObjectPermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedObjectPermissionRequest(val *PatchedObjectPermissionRequest) *NullablePatchedObjectPermissionRequest {
	return &NullablePatchedObjectPermissionRequest{value: val, isSet: true}
}

func (v NullablePatchedObjectPermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedObjectPermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
