/*
Faro Collector API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package faro

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PayloadTracesResourceSpansInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayloadTracesResourceSpansInner{}

// PayloadTracesResourceSpansInner struct for PayloadTracesResourceSpansInner
type PayloadTracesResourceSpansInner struct {
	Resource PayloadTracesResourceSpansInnerResource `json:"resource"`
	InstrumentationLibrarySpans []PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInner `json:"instrumentationLibrarySpans"`
}

type _PayloadTracesResourceSpansInner PayloadTracesResourceSpansInner

// NewPayloadTracesResourceSpansInner instantiates a new PayloadTracesResourceSpansInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayloadTracesResourceSpansInner(resource PayloadTracesResourceSpansInnerResource, instrumentationLibrarySpans []PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInner) *PayloadTracesResourceSpansInner {
	this := PayloadTracesResourceSpansInner{}
	this.Resource = resource
	this.InstrumentationLibrarySpans = instrumentationLibrarySpans
	return &this
}

// NewPayloadTracesResourceSpansInnerWithDefaults instantiates a new PayloadTracesResourceSpansInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayloadTracesResourceSpansInnerWithDefaults() *PayloadTracesResourceSpansInner {
	this := PayloadTracesResourceSpansInner{}
	return &this
}

// GetResource returns the Resource field value
func (o *PayloadTracesResourceSpansInner) GetResource() PayloadTracesResourceSpansInnerResource {
	if o == nil {
		var ret PayloadTracesResourceSpansInnerResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *PayloadTracesResourceSpansInner) GetResourceOk() (*PayloadTracesResourceSpansInnerResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *PayloadTracesResourceSpansInner) SetResource(v PayloadTracesResourceSpansInnerResource) {
	o.Resource = v
}

// GetInstrumentationLibrarySpans returns the InstrumentationLibrarySpans field value
func (o *PayloadTracesResourceSpansInner) GetInstrumentationLibrarySpans() []PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInner {
	if o == nil {
		var ret []PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInner
		return ret
	}

	return o.InstrumentationLibrarySpans
}

// GetInstrumentationLibrarySpansOk returns a tuple with the InstrumentationLibrarySpans field value
// and a boolean to check if the value has been set.
func (o *PayloadTracesResourceSpansInner) GetInstrumentationLibrarySpansOk() ([]PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstrumentationLibrarySpans, true
}

// SetInstrumentationLibrarySpans sets field value
func (o *PayloadTracesResourceSpansInner) SetInstrumentationLibrarySpans(v []PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInner) {
	o.InstrumentationLibrarySpans = v
}

func (o PayloadTracesResourceSpansInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayloadTracesResourceSpansInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource"] = o.Resource
	toSerialize["instrumentationLibrarySpans"] = o.InstrumentationLibrarySpans
	return toSerialize, nil
}

func (o *PayloadTracesResourceSpansInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource",
		"instrumentationLibrarySpans",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPayloadTracesResourceSpansInner := _PayloadTracesResourceSpansInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPayloadTracesResourceSpansInner)

	if err != nil {
		return err
	}

	*o = PayloadTracesResourceSpansInner(varPayloadTracesResourceSpansInner)

	return err
}

type NullablePayloadTracesResourceSpansInner struct {
	value *PayloadTracesResourceSpansInner
	isSet bool
}

func (v NullablePayloadTracesResourceSpansInner) Get() *PayloadTracesResourceSpansInner {
	return v.value
}

func (v *NullablePayloadTracesResourceSpansInner) Set(val *PayloadTracesResourceSpansInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePayloadTracesResourceSpansInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePayloadTracesResourceSpansInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayloadTracesResourceSpansInner(val *PayloadTracesResourceSpansInner) *NullablePayloadTracesResourceSpansInner {
	return &NullablePayloadTracesResourceSpansInner{value: val, isSet: true}
}

func (v NullablePayloadTracesResourceSpansInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayloadTracesResourceSpansInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


