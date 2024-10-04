/*
Discovery Configuration API V2

The Discovery configuration service is a BloxOne Service that provides configuration for accessing and syncing the Cloud assets   Base Paths:  1. provider: **_/api/cloud_discovery/v2/_**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clouddiscovery

import (
	"encoding/json"
)

// checks if the ProviderCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderCreateResponse{}

// ProviderCreateResponse The Provider object create response format.
type ProviderCreateResponse struct {
	Result               *DiscoveryConfig `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProviderCreateResponse ProviderCreateResponse

// NewProviderCreateResponse instantiates a new ProviderCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderCreateResponse() *ProviderCreateResponse {
	this := ProviderCreateResponse{}
	return &this
}

// NewProviderCreateResponseWithDefaults instantiates a new ProviderCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderCreateResponseWithDefaults() *ProviderCreateResponse {
	this := ProviderCreateResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ProviderCreateResponse) GetResult() DiscoveryConfig {
	if o == nil || IsNil(o.Result) {
		var ret DiscoveryConfig
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderCreateResponse) GetResultOk() (*DiscoveryConfig, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ProviderCreateResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DiscoveryConfig and assigns it to the Result field.
func (o *ProviderCreateResponse) SetResult(v DiscoveryConfig) {
	o.Result = &v
}

func (o ProviderCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProviderCreateResponse) UnmarshalJSON(data []byte) (err error) {
	varProviderCreateResponse := _ProviderCreateResponse{}

	err = json.Unmarshal(data, &varProviderCreateResponse)

	if err != nil {
		return err
	}

	*o = ProviderCreateResponse(varProviderCreateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProviderCreateResponse struct {
	value *ProviderCreateResponse
	isSet bool
}

func (v NullableProviderCreateResponse) Get() *ProviderCreateResponse {
	return v.value
}

func (v *NullableProviderCreateResponse) Set(val *ProviderCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderCreateResponse(val *ProviderCreateResponse) *NullableProviderCreateResponse {
	return &NullableProviderCreateResponse{value: val, isSet: true}
}

func (v NullableProviderCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}