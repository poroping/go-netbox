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

// checks if the PatchedWritableWirelessLANRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableWirelessLANRequest{}

// PatchedWritableWirelessLANRequest Adds support for custom fields and tags.
type PatchedWritableWirelessLANRequest struct {
	Ssid                 *string                                        `json:"ssid,omitempty"`
	Description          *string                                        `json:"description,omitempty"`
	Group                NullablePatchedWritableWirelessLANRequestGroup `json:"group,omitempty"`
	Status               *PatchedWritableWirelessLANRequestStatus       `json:"status,omitempty"`
	Vlan                 NullableInterfaceRequestUntaggedVlan           `json:"vlan,omitempty"`
	ScopeType            NullableString                                 `json:"scope_type,omitempty"`
	ScopeId              NullableInt32                                  `json:"scope_id,omitempty"`
	Tenant               NullableASNRangeRequestTenant                  `json:"tenant,omitempty"`
	AuthType             NullableAuthenticationType1                    `json:"auth_type,omitempty"`
	AuthCipher           NullableAuthenticationCipher                   `json:"auth_cipher,omitempty"`
	AuthPsk              *string                                        `json:"auth_psk,omitempty"`
	Comments             *string                                        `json:"comments,omitempty"`
	Tags                 []NestedTagRequest                             `json:"tags,omitempty"`
	CustomFields         map[string]interface{}                         `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableWirelessLANRequest PatchedWritableWirelessLANRequest

// NewPatchedWritableWirelessLANRequest instantiates a new PatchedWritableWirelessLANRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableWirelessLANRequest() *PatchedWritableWirelessLANRequest {
	this := PatchedWritableWirelessLANRequest{}
	return &this
}

// NewPatchedWritableWirelessLANRequestWithDefaults instantiates a new PatchedWritableWirelessLANRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableWirelessLANRequestWithDefaults() *PatchedWritableWirelessLANRequest {
	this := PatchedWritableWirelessLANRequest{}
	return &this
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLANRequest) GetSsid() string {
	if o == nil || IsNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLANRequest) GetSsidOk() (*string, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *PatchedWritableWirelessLANRequest) SetSsid(v string) {
	o.Ssid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLANRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLANRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableWirelessLANRequest) SetDescription(v string) {
	o.Description = &v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableWirelessLANRequest) GetGroup() PatchedWritableWirelessLANRequestGroup {
	if o == nil || IsNil(o.Group.Get()) {
		var ret PatchedWritableWirelessLANRequestGroup
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableWirelessLANRequest) GetGroupOk() (*PatchedWritableWirelessLANRequestGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullablePatchedWritableWirelessLANRequestGroup and assigns it to the Group field.
func (o *PatchedWritableWirelessLANRequest) SetGroup(v PatchedWritableWirelessLANRequestGroup) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *PatchedWritableWirelessLANRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *PatchedWritableWirelessLANRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLANRequest) GetStatus() PatchedWritableWirelessLANRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret PatchedWritableWirelessLANRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLANRequest) GetStatusOk() (*PatchedWritableWirelessLANRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PatchedWritableWirelessLANRequestStatus and assigns it to the Status field.
func (o *PatchedWritableWirelessLANRequest) SetStatus(v PatchedWritableWirelessLANRequestStatus) {
	o.Status = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableWirelessLANRequest) GetVlan() InterfaceRequestUntaggedVlan {
	if o == nil || IsNil(o.Vlan.Get()) {
		var ret InterfaceRequestUntaggedVlan
		return ret
	}
	return *o.Vlan.Get()
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableWirelessLANRequest) GetVlanOk() (*InterfaceRequestUntaggedVlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vlan.Get(), o.Vlan.IsSet()
}

// HasVlan returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasVlan() bool {
	if o != nil && o.Vlan.IsSet() {
		return true
	}

	return false
}

// SetVlan gets a reference to the given NullableInterfaceRequestUntaggedVlan and assigns it to the Vlan field.
func (o *PatchedWritableWirelessLANRequest) SetVlan(v InterfaceRequestUntaggedVlan) {
	o.Vlan.Set(&v)
}

// SetVlanNil sets the value for Vlan to be an explicit nil
func (o *PatchedWritableWirelessLANRequest) SetVlanNil() {
	o.Vlan.Set(nil)
}

// UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
func (o *PatchedWritableWirelessLANRequest) UnsetVlan() {
	o.Vlan.Unset()
}

// GetScopeType returns the ScopeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableWirelessLANRequest) GetScopeType() string {
	if o == nil || IsNil(o.ScopeType.Get()) {
		var ret string
		return ret
	}
	return *o.ScopeType.Get()
}

// GetScopeTypeOk returns a tuple with the ScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableWirelessLANRequest) GetScopeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeType.Get(), o.ScopeType.IsSet()
}

// HasScopeType returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasScopeType() bool {
	if o != nil && o.ScopeType.IsSet() {
		return true
	}

	return false
}

// SetScopeType gets a reference to the given NullableString and assigns it to the ScopeType field.
func (o *PatchedWritableWirelessLANRequest) SetScopeType(v string) {
	o.ScopeType.Set(&v)
}

// SetScopeTypeNil sets the value for ScopeType to be an explicit nil
func (o *PatchedWritableWirelessLANRequest) SetScopeTypeNil() {
	o.ScopeType.Set(nil)
}

// UnsetScopeType ensures that no value is present for ScopeType, not even an explicit nil
func (o *PatchedWritableWirelessLANRequest) UnsetScopeType() {
	o.ScopeType.Unset()
}

// GetScopeId returns the ScopeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableWirelessLANRequest) GetScopeId() int32 {
	if o == nil || IsNil(o.ScopeId.Get()) {
		var ret int32
		return ret
	}
	return *o.ScopeId.Get()
}

// GetScopeIdOk returns a tuple with the ScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableWirelessLANRequest) GetScopeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeId.Get(), o.ScopeId.IsSet()
}

// HasScopeId returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasScopeId() bool {
	if o != nil && o.ScopeId.IsSet() {
		return true
	}

	return false
}

// SetScopeId gets a reference to the given NullableInt32 and assigns it to the ScopeId field.
func (o *PatchedWritableWirelessLANRequest) SetScopeId(v int32) {
	o.ScopeId.Set(&v)
}

// SetScopeIdNil sets the value for ScopeId to be an explicit nil
func (o *PatchedWritableWirelessLANRequest) SetScopeIdNil() {
	o.ScopeId.Set(nil)
}

// UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
func (o *PatchedWritableWirelessLANRequest) UnsetScopeId() {
	o.ScopeId.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableWirelessLANRequest) GetTenant() ASNRangeRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret ASNRangeRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableWirelessLANRequest) GetTenantOk() (*ASNRangeRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableASNRangeRequestTenant and assigns it to the Tenant field.
func (o *PatchedWritableWirelessLANRequest) SetTenant(v ASNRangeRequestTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableWirelessLANRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableWirelessLANRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetAuthType returns the AuthType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableWirelessLANRequest) GetAuthType() AuthenticationType1 {
	if o == nil || IsNil(o.AuthType.Get()) {
		var ret AuthenticationType1
		return ret
	}
	return *o.AuthType.Get()
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableWirelessLANRequest) GetAuthTypeOk() (*AuthenticationType1, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthType.Get(), o.AuthType.IsSet()
}

// HasAuthType returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasAuthType() bool {
	if o != nil && o.AuthType.IsSet() {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given NullableAuthenticationType1 and assigns it to the AuthType field.
func (o *PatchedWritableWirelessLANRequest) SetAuthType(v AuthenticationType1) {
	o.AuthType.Set(&v)
}

// SetAuthTypeNil sets the value for AuthType to be an explicit nil
func (o *PatchedWritableWirelessLANRequest) SetAuthTypeNil() {
	o.AuthType.Set(nil)
}

// UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
func (o *PatchedWritableWirelessLANRequest) UnsetAuthType() {
	o.AuthType.Unset()
}

// GetAuthCipher returns the AuthCipher field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableWirelessLANRequest) GetAuthCipher() AuthenticationCipher {
	if o == nil || IsNil(o.AuthCipher.Get()) {
		var ret AuthenticationCipher
		return ret
	}
	return *o.AuthCipher.Get()
}

// GetAuthCipherOk returns a tuple with the AuthCipher field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableWirelessLANRequest) GetAuthCipherOk() (*AuthenticationCipher, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthCipher.Get(), o.AuthCipher.IsSet()
}

// HasAuthCipher returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasAuthCipher() bool {
	if o != nil && o.AuthCipher.IsSet() {
		return true
	}

	return false
}

// SetAuthCipher gets a reference to the given NullableAuthenticationCipher and assigns it to the AuthCipher field.
func (o *PatchedWritableWirelessLANRequest) SetAuthCipher(v AuthenticationCipher) {
	o.AuthCipher.Set(&v)
}

// SetAuthCipherNil sets the value for AuthCipher to be an explicit nil
func (o *PatchedWritableWirelessLANRequest) SetAuthCipherNil() {
	o.AuthCipher.Set(nil)
}

// UnsetAuthCipher ensures that no value is present for AuthCipher, not even an explicit nil
func (o *PatchedWritableWirelessLANRequest) UnsetAuthCipher() {
	o.AuthCipher.Unset()
}

// GetAuthPsk returns the AuthPsk field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLANRequest) GetAuthPsk() string {
	if o == nil || IsNil(o.AuthPsk) {
		var ret string
		return ret
	}
	return *o.AuthPsk
}

// GetAuthPskOk returns a tuple with the AuthPsk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLANRequest) GetAuthPskOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPsk) {
		return nil, false
	}
	return o.AuthPsk, true
}

// HasAuthPsk returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasAuthPsk() bool {
	if o != nil && !IsNil(o.AuthPsk) {
		return true
	}

	return false
}

// SetAuthPsk gets a reference to the given string and assigns it to the AuthPsk field.
func (o *PatchedWritableWirelessLANRequest) SetAuthPsk(v string) {
	o.AuthPsk = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLANRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLANRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableWirelessLANRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLANRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLANRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableWirelessLANRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLANRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLANRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLANRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableWirelessLANRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableWirelessLANRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableWirelessLANRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Vlan.IsSet() {
		toSerialize["vlan"] = o.Vlan.Get()
	}
	if o.ScopeType.IsSet() {
		toSerialize["scope_type"] = o.ScopeType.Get()
	}
	if o.ScopeId.IsSet() {
		toSerialize["scope_id"] = o.ScopeId.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.AuthType.IsSet() {
		toSerialize["auth_type"] = o.AuthType.Get()
	}
	if o.AuthCipher.IsSet() {
		toSerialize["auth_cipher"] = o.AuthCipher.Get()
	}
	if !IsNil(o.AuthPsk) {
		toSerialize["auth_psk"] = o.AuthPsk
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
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

func (o *PatchedWritableWirelessLANRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableWirelessLANRequest := _PatchedWritableWirelessLANRequest{}

	err = json.Unmarshal(data, &varPatchedWritableWirelessLANRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableWirelessLANRequest(varPatchedWritableWirelessLANRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "description")
		delete(additionalProperties, "group")
		delete(additionalProperties, "status")
		delete(additionalProperties, "vlan")
		delete(additionalProperties, "scope_type")
		delete(additionalProperties, "scope_id")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "auth_type")
		delete(additionalProperties, "auth_cipher")
		delete(additionalProperties, "auth_psk")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableWirelessLANRequest struct {
	value *PatchedWritableWirelessLANRequest
	isSet bool
}

func (v NullablePatchedWritableWirelessLANRequest) Get() *PatchedWritableWirelessLANRequest {
	return v.value
}

func (v *NullablePatchedWritableWirelessLANRequest) Set(val *PatchedWritableWirelessLANRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableWirelessLANRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableWirelessLANRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableWirelessLANRequest(val *PatchedWritableWirelessLANRequest) *NullablePatchedWritableWirelessLANRequest {
	return &NullablePatchedWritableWirelessLANRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableWirelessLANRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableWirelessLANRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
