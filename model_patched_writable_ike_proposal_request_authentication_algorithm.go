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

// PatchedWritableIKEProposalRequestAuthenticationAlgorithm * `hmac-sha1` - SHA-1 HMAC * `hmac-sha256` - SHA-256 HMAC * `hmac-sha384` - SHA-384 HMAC * `hmac-sha512` - SHA-512 HMAC * `hmac-md5` - MD5 HMAC
type PatchedWritableIKEProposalRequestAuthenticationAlgorithm string

// List of PatchedWritableIKEProposalRequest_authentication_algorithm
const (
	PATCHEDWRITABLEIKEPROPOSALREQUESTAUTHENTICATIONALGORITHM_HMAC_SHA1   PatchedWritableIKEProposalRequestAuthenticationAlgorithm = "hmac-sha1"
	PATCHEDWRITABLEIKEPROPOSALREQUESTAUTHENTICATIONALGORITHM_HMAC_SHA256 PatchedWritableIKEProposalRequestAuthenticationAlgorithm = "hmac-sha256"
	PATCHEDWRITABLEIKEPROPOSALREQUESTAUTHENTICATIONALGORITHM_HMAC_SHA384 PatchedWritableIKEProposalRequestAuthenticationAlgorithm = "hmac-sha384"
	PATCHEDWRITABLEIKEPROPOSALREQUESTAUTHENTICATIONALGORITHM_HMAC_SHA512 PatchedWritableIKEProposalRequestAuthenticationAlgorithm = "hmac-sha512"
	PATCHEDWRITABLEIKEPROPOSALREQUESTAUTHENTICATIONALGORITHM_HMAC_MD5    PatchedWritableIKEProposalRequestAuthenticationAlgorithm = "hmac-md5"
	PATCHEDWRITABLEIKEPROPOSALREQUESTAUTHENTICATIONALGORITHM_EMPTY       PatchedWritableIKEProposalRequestAuthenticationAlgorithm = ""
)

// All allowed values of PatchedWritableIKEProposalRequestAuthenticationAlgorithm enum
var AllowedPatchedWritableIKEProposalRequestAuthenticationAlgorithmEnumValues = []PatchedWritableIKEProposalRequestAuthenticationAlgorithm{
	"hmac-sha1",
	"hmac-sha256",
	"hmac-sha384",
	"hmac-sha512",
	"hmac-md5",
	"",
}

func (v *PatchedWritableIKEProposalRequestAuthenticationAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableIKEProposalRequestAuthenticationAlgorithm(value)
	for _, existing := range AllowedPatchedWritableIKEProposalRequestAuthenticationAlgorithmEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableIKEProposalRequestAuthenticationAlgorithm", value)
}

// NewPatchedWritableIKEProposalRequestAuthenticationAlgorithmFromValue returns a pointer to a valid PatchedWritableIKEProposalRequestAuthenticationAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableIKEProposalRequestAuthenticationAlgorithmFromValue(v string) (*PatchedWritableIKEProposalRequestAuthenticationAlgorithm, error) {
	ev := PatchedWritableIKEProposalRequestAuthenticationAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableIKEProposalRequestAuthenticationAlgorithm: valid values are %v", v, AllowedPatchedWritableIKEProposalRequestAuthenticationAlgorithmEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableIKEProposalRequestAuthenticationAlgorithm) IsValid() bool {
	for _, existing := range AllowedPatchedWritableIKEProposalRequestAuthenticationAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableIKEProposalRequest_authentication_algorithm value
func (v PatchedWritableIKEProposalRequestAuthenticationAlgorithm) Ptr() *PatchedWritableIKEProposalRequestAuthenticationAlgorithm {
	return &v
}

type NullablePatchedWritableIKEProposalRequestAuthenticationAlgorithm struct {
	value *PatchedWritableIKEProposalRequestAuthenticationAlgorithm
	isSet bool
}

func (v NullablePatchedWritableIKEProposalRequestAuthenticationAlgorithm) Get() *PatchedWritableIKEProposalRequestAuthenticationAlgorithm {
	return v.value
}

func (v *NullablePatchedWritableIKEProposalRequestAuthenticationAlgorithm) Set(val *PatchedWritableIKEProposalRequestAuthenticationAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableIKEProposalRequestAuthenticationAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableIKEProposalRequestAuthenticationAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableIKEProposalRequestAuthenticationAlgorithm(val *PatchedWritableIKEProposalRequestAuthenticationAlgorithm) *NullablePatchedWritableIKEProposalRequestAuthenticationAlgorithm {
	return &NullablePatchedWritableIKEProposalRequestAuthenticationAlgorithm{value: val, isSet: true}
}

func (v NullablePatchedWritableIKEProposalRequestAuthenticationAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableIKEProposalRequestAuthenticationAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
