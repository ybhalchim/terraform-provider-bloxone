/*
BloxOne Anycast API

Anycast capability enables HA (High Availability) configuration of BloxOne applications that run on equipment located on customer's premises (on-prem hosts). Anycast supports DNS, as well as DNS-forwarding services.  Anycast-enabled application setups use multiple on-premises installations for one particular application type. Multiple application instances are configured to use the same endpoint address. Anycast capability is collocated with such application instance, monitoring the local application instance and advertising to the upstream router (a customer equipment) a per-instance, local route to the common application endpoint address, as long as the local application instance is available. Depending on the type of the upstream router, the customer may configure local route advertisement via either BGP (Boarder Gateway Protocol) or OSPF (Open Shortest Path First) routing protocols. Both protocols may be enabled as well. Multiple routes to the common application service address provide redundancy without the need to reconfigure application clients.  Should an application instance become unavailable, the local route advertisements stop, resulting in withdrawal of the route (in the upstream router) to the application instance that has gone out of service and ensuring that subsequent application requests thus get routed to the remaining available application instances.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package anycast

import (
	"encoding/json"
)

// checks if the AnycastConfigResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnycastConfigResponse{}

// AnycastConfigResponse struct for AnycastConfigResponse
type AnycastConfigResponse struct {
	Results *AnycastConfig `json:"results,omitempty"`
}

// NewAnycastConfigResponse instantiates a new AnycastConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnycastConfigResponse() *AnycastConfigResponse {
	this := AnycastConfigResponse{}
	return &this
}

// NewAnycastConfigResponseWithDefaults instantiates a new AnycastConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnycastConfigResponseWithDefaults() *AnycastConfigResponse {
	this := AnycastConfigResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AnycastConfigResponse) GetResults() AnycastConfig {
	if o == nil || IsNil(o.Results) {
		var ret AnycastConfig
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnycastConfigResponse) GetResultsOk() (*AnycastConfig, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AnycastConfigResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given AnycastConfig and assigns it to the Results field.
func (o *AnycastConfigResponse) SetResults(v AnycastConfig) {
	o.Results = &v
}

func (o AnycastConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnycastConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableAnycastConfigResponse struct {
	value *AnycastConfigResponse
	isSet bool
}

func (v NullableAnycastConfigResponse) Get() *AnycastConfigResponse {
	return v.value
}

func (v *NullableAnycastConfigResponse) Set(val *AnycastConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnycastConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnycastConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnycastConfigResponse(val *AnycastConfigResponse) *NullableAnycastConfigResponse {
	return &NullableAnycastConfigResponse{value: val, isSet: true}
}

func (v NullableAnycastConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnycastConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}