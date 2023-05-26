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

// checks if the CustomAttributeTemplatePostModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomAttributeTemplatePostModel{}

// CustomAttributeTemplatePostModel struct for CustomAttributeTemplatePostModel
type CustomAttributeTemplatePostModel struct {
	// Collection of attribute IDs
	CustomAttributeIds []string `json:"customAttributeIds,omitempty"`
	// Custom attributes template name
	Name string `json:"name"`
}

// NewCustomAttributeTemplatePostModel instantiates a new CustomAttributeTemplatePostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAttributeTemplatePostModel(name string) *CustomAttributeTemplatePostModel {
	this := CustomAttributeTemplatePostModel{}
	this.Name = name
	return &this
}

// NewCustomAttributeTemplatePostModelWithDefaults instantiates a new CustomAttributeTemplatePostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAttributeTemplatePostModelWithDefaults() *CustomAttributeTemplatePostModel {
	this := CustomAttributeTemplatePostModel{}
	return &this
}

// GetCustomAttributeIds returns the CustomAttributeIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomAttributeTemplatePostModel) GetCustomAttributeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CustomAttributeIds
}

// GetCustomAttributeIdsOk returns a tuple with the CustomAttributeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomAttributeTemplatePostModel) GetCustomAttributeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomAttributeIds) {
		return nil, false
	}
	return o.CustomAttributeIds, true
}

// HasCustomAttributeIds returns a boolean if a field has been set.
func (o *CustomAttributeTemplatePostModel) HasCustomAttributeIds() bool {
	if o != nil && IsNil(o.CustomAttributeIds) {
		return true
	}

	return false
}

// SetCustomAttributeIds gets a reference to the given []string and assigns it to the CustomAttributeIds field.
func (o *CustomAttributeTemplatePostModel) SetCustomAttributeIds(v []string) {
	o.CustomAttributeIds = v
}

// GetName returns the Name field value
func (o *CustomAttributeTemplatePostModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeTemplatePostModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomAttributeTemplatePostModel) SetName(v string) {
	o.Name = v
}

func (o CustomAttributeTemplatePostModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomAttributeTemplatePostModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomAttributeIds != nil {
		toSerialize["customAttributeIds"] = o.CustomAttributeIds
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableCustomAttributeTemplatePostModel struct {
	value *CustomAttributeTemplatePostModel
	isSet bool
}

func (v NullableCustomAttributeTemplatePostModel) Get() *CustomAttributeTemplatePostModel {
	return v.value
}

func (v *NullableCustomAttributeTemplatePostModel) Set(val *CustomAttributeTemplatePostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAttributeTemplatePostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAttributeTemplatePostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAttributeTemplatePostModel(val *CustomAttributeTemplatePostModel) *NullableCustomAttributeTemplatePostModel {
	return &NullableCustomAttributeTemplatePostModel{value: val, isSet: true}
}

func (v NullableCustomAttributeTemplatePostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAttributeTemplatePostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

