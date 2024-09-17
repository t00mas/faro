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

// checks if the SdkMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SdkMeta{}

// SdkMeta struct for SdkMeta
type SdkMeta struct {
	Name string `json:"name"`
	Version string `json:"version"`
}

type _SdkMeta SdkMeta

// NewSdkMeta instantiates a new SdkMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdkMeta(name string, version string) *SdkMeta {
	this := SdkMeta{}
	this.Name = name
	this.Version = version
	return &this
}

// NewSdkMetaWithDefaults instantiates a new SdkMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdkMetaWithDefaults() *SdkMeta {
	this := SdkMeta{}
	return &this
}

// GetName returns the Name field value
func (o *SdkMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SdkMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SdkMeta) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *SdkMeta) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SdkMeta) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SdkMeta) SetVersion(v string) {
	o.Version = v
}

func (o SdkMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SdkMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *SdkMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"version",
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

	varSdkMeta := _SdkMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSdkMeta)

	if err != nil {
		return err
	}

	*o = SdkMeta(varSdkMeta)

	return err
}

type NullableSdkMeta struct {
	value *SdkMeta
	isSet bool
}

func (v NullableSdkMeta) Get() *SdkMeta {
	return v.value
}

func (v *NullableSdkMeta) Set(val *SdkMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableSdkMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableSdkMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdkMeta(val *SdkMeta) *NullableSdkMeta {
	return &NullableSdkMeta{value: val, isSet: true}
}

func (v NullableSdkMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdkMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


