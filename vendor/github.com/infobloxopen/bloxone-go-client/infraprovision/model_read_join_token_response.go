/*
Host Activation Service

Host activation service provides a RESTful interface to manage cert and join token object. Join tokens are essentially a password that allows on-prem hosts to auto-associate themselves to a customer's account and receive a signed cert.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infraprovision

import (
	"encoding/json"
)

// checks if the ReadJoinTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadJoinTokenResponse{}

// ReadJoinTokenResponse struct for ReadJoinTokenResponse
type ReadJoinTokenResponse struct {
	Result *JoinToken `json:"result,omitempty"`
}

// NewReadJoinTokenResponse instantiates a new ReadJoinTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadJoinTokenResponse() *ReadJoinTokenResponse {
	this := ReadJoinTokenResponse{}
	return &this
}

// NewReadJoinTokenResponseWithDefaults instantiates a new ReadJoinTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadJoinTokenResponseWithDefaults() *ReadJoinTokenResponse {
	this := ReadJoinTokenResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReadJoinTokenResponse) GetResult() JoinToken {
	if o == nil || IsNil(o.Result) {
		var ret JoinToken
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadJoinTokenResponse) GetResultOk() (*JoinToken, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReadJoinTokenResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given JoinToken and assigns it to the Result field.
func (o *ReadJoinTokenResponse) SetResult(v JoinToken) {
	o.Result = &v
}

func (o ReadJoinTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadJoinTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableReadJoinTokenResponse struct {
	value *ReadJoinTokenResponse
	isSet bool
}

func (v NullableReadJoinTokenResponse) Get() *ReadJoinTokenResponse {
	return v.value
}

func (v *NullableReadJoinTokenResponse) Set(val *ReadJoinTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadJoinTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadJoinTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadJoinTokenResponse(val *ReadJoinTokenResponse) *NullableReadJoinTokenResponse {
	return &NullableReadJoinTokenResponse{value: val, isSet: true}
}

func (v NullableReadJoinTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadJoinTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}