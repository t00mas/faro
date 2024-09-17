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

// checks if the AppMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppMeta{}

// AppMeta struct for AppMeta
type AppMeta struct {
	Name string `json:"name"`
	Release string `json:"release"`
	Version string `json:"version"`
	Environment string `json:"environment"`
}

type _AppMeta AppMeta

// NewAppMeta instantiates a new AppMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppMeta(name string, release string, version string, environment string) *AppMeta {
	this := AppMeta{}
	this.Name = name
	this.Release = release
	this.Version = version
	this.Environment = environment
	return &this
}

// NewAppMetaWithDefaults instantiates a new AppMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppMetaWithDefaults() *AppMeta {
	this := AppMeta{}
	return &this
}

// GetName returns the Name field value
func (o *AppMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppMeta) SetName(v string) {
	o.Name = v
}

// GetRelease returns the Release field value
func (o *AppMeta) GetRelease() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Release
}

// GetReleaseOk returns a tuple with the Release field value
// and a boolean to check if the value has been set.
func (o *AppMeta) GetReleaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Release, true
}

// SetRelease sets field value
func (o *AppMeta) SetRelease(v string) {
	o.Release = v
}

// GetVersion returns the Version field value
func (o *AppMeta) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AppMeta) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AppMeta) SetVersion(v string) {
	o.Version = v
}

// GetEnvironment returns the Environment field value
func (o *AppMeta) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *AppMeta) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *AppMeta) SetEnvironment(v string) {
	o.Environment = v
}

func (o AppMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["release"] = o.Release
	toSerialize["version"] = o.Version
	toSerialize["environment"] = o.Environment
	return toSerialize, nil
}

func (o *AppMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"release",
		"version",
		"environment",
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

	varAppMeta := _AppMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppMeta)

	if err != nil {
		return err
	}

	*o = AppMeta(varAppMeta)

	return err
}

type NullableAppMeta struct {
	value *AppMeta
	isSet bool
}

func (v NullableAppMeta) Get() *AppMeta {
	return v.value
}

func (v *NullableAppMeta) Set(val *AppMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableAppMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableAppMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppMeta(val *AppMeta) *NullableAppMeta {
	return &NullableAppMeta{value: val, isSet: true}
}

func (v NullableAppMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


