/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
)

// checks if the AttachmentSubGetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentSubGetModel{}

// AttachmentSubGetModel struct for AttachmentSubGetModel
type AttachmentSubGetModel struct {
	Id *string `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewAttachmentSubGetModel instantiates a new AttachmentSubGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentSubGetModel() *AttachmentSubGetModel {
	this := AttachmentSubGetModel{}
	return &this
}

// NewAttachmentSubGetModelWithDefaults instantiates a new AttachmentSubGetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentSubGetModelWithDefaults() *AttachmentSubGetModel {
	this := AttachmentSubGetModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AttachmentSubGetModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentSubGetModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AttachmentSubGetModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AttachmentSubGetModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentSubGetModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentSubGetModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AttachmentSubGetModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AttachmentSubGetModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AttachmentSubGetModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AttachmentSubGetModel) UnsetName() {
	o.Name.Unset()
}

func (o AttachmentSubGetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentSubGetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableAttachmentSubGetModel struct {
	value *AttachmentSubGetModel
	isSet bool
}

func (v NullableAttachmentSubGetModel) Get() *AttachmentSubGetModel {
	return v.value
}

func (v *NullableAttachmentSubGetModel) Set(val *AttachmentSubGetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentSubGetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentSubGetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentSubGetModel(val *AttachmentSubGetModel) *NullableAttachmentSubGetModel {
	return &NullableAttachmentSubGetModel{value: val, isSet: true}
}

func (v NullableAttachmentSubGetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentSubGetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


