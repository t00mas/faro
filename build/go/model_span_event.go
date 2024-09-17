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

// checks if the SpanEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpanEvent{}

// SpanEvent struct for SpanEvent
type SpanEvent struct {
	Attributes []string `json:"attributes,omitempty"`
	DroppedAttributesCount *int32 `json:"droppedAttributesCount,omitempty"`
	Name string `json:"name"`
	TimeUnixNano *string `json:"timeUnixNano,omitempty"`
}

type _SpanEvent SpanEvent

// NewSpanEvent instantiates a new SpanEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanEvent(name string) *SpanEvent {
	this := SpanEvent{}
	this.Name = name
	return &this
}

// NewSpanEventWithDefaults instantiates a new SpanEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanEventWithDefaults() *SpanEvent {
	this := SpanEvent{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SpanEvent) GetAttributes() []string {
	if o == nil || IsNil(o.Attributes) {
		var ret []string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanEvent) GetAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SpanEvent) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *SpanEvent) SetAttributes(v []string) {
	o.Attributes = v
}

// GetDroppedAttributesCount returns the DroppedAttributesCount field value if set, zero value otherwise.
func (o *SpanEvent) GetDroppedAttributesCount() int32 {
	if o == nil || IsNil(o.DroppedAttributesCount) {
		var ret int32
		return ret
	}
	return *o.DroppedAttributesCount
}

// GetDroppedAttributesCountOk returns a tuple with the DroppedAttributesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanEvent) GetDroppedAttributesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DroppedAttributesCount) {
		return nil, false
	}
	return o.DroppedAttributesCount, true
}

// HasDroppedAttributesCount returns a boolean if a field has been set.
func (o *SpanEvent) HasDroppedAttributesCount() bool {
	if o != nil && !IsNil(o.DroppedAttributesCount) {
		return true
	}

	return false
}

// SetDroppedAttributesCount gets a reference to the given int32 and assigns it to the DroppedAttributesCount field.
func (o *SpanEvent) SetDroppedAttributesCount(v int32) {
	o.DroppedAttributesCount = &v
}

// GetName returns the Name field value
func (o *SpanEvent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SpanEvent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SpanEvent) SetName(v string) {
	o.Name = v
}

// GetTimeUnixNano returns the TimeUnixNano field value if set, zero value otherwise.
func (o *SpanEvent) GetTimeUnixNano() string {
	if o == nil || IsNil(o.TimeUnixNano) {
		var ret string
		return ret
	}
	return *o.TimeUnixNano
}

// GetTimeUnixNanoOk returns a tuple with the TimeUnixNano field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanEvent) GetTimeUnixNanoOk() (*string, bool) {
	if o == nil || IsNil(o.TimeUnixNano) {
		return nil, false
	}
	return o.TimeUnixNano, true
}

// HasTimeUnixNano returns a boolean if a field has been set.
func (o *SpanEvent) HasTimeUnixNano() bool {
	if o != nil && !IsNil(o.TimeUnixNano) {
		return true
	}

	return false
}

// SetTimeUnixNano gets a reference to the given string and assigns it to the TimeUnixNano field.
func (o *SpanEvent) SetTimeUnixNano(v string) {
	o.TimeUnixNano = &v
}

func (o SpanEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpanEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.DroppedAttributesCount) {
		toSerialize["droppedAttributesCount"] = o.DroppedAttributesCount
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.TimeUnixNano) {
		toSerialize["timeUnixNano"] = o.TimeUnixNano
	}
	return toSerialize, nil
}

func (o *SpanEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varSpanEvent := _SpanEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpanEvent)

	if err != nil {
		return err
	}

	*o = SpanEvent(varSpanEvent)

	return err
}

type NullableSpanEvent struct {
	value *SpanEvent
	isSet bool
}

func (v NullableSpanEvent) Get() *SpanEvent {
	return v.value
}

func (v *NullableSpanEvent) Set(val *SpanEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanEvent(val *SpanEvent) *NullableSpanEvent {
	return &NullableSpanEvent{value: val, isSet: true}
}

func (v NullableSpanEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


