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

// checks if the ProjectAttributesFilterModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectAttributesFilterModel{}

// ProjectAttributesFilterModel struct for ProjectAttributesFilterModel
type ProjectAttributesFilterModel struct {
	// Specifies an attribute name to search for
	Name NullableString `json:"name,omitempty"`
	// Specifies an attribute mandatory status to search for
	IsRequired NullableBool `json:"isRequired,omitempty"`
	// Specifies an attribute global status to search for
	IsGlobal NullableBool `json:"isGlobal,omitempty"`
	// Specifies an attribute types to search for
	Types []CustomAttributeTypesEnum `json:"types,omitempty"`
	// Specifies an attribute enabled status to search for
	IsEnabled NullableBool `json:"isEnabled,omitempty"`
}

// NewProjectAttributesFilterModel instantiates a new ProjectAttributesFilterModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAttributesFilterModel() *ProjectAttributesFilterModel {
	this := ProjectAttributesFilterModel{}
	return &this
}

// NewProjectAttributesFilterModelWithDefaults instantiates a new ProjectAttributesFilterModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAttributesFilterModelWithDefaults() *ProjectAttributesFilterModel {
	this := ProjectAttributesFilterModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAttributesFilterModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAttributesFilterModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProjectAttributesFilterModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProjectAttributesFilterModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProjectAttributesFilterModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProjectAttributesFilterModel) UnsetName() {
	o.Name.Unset()
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAttributesFilterModel) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRequired.Get()
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAttributesFilterModel) GetIsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRequired.Get(), o.IsRequired.IsSet()
}

// HasIsRequired returns a boolean if a field has been set.
func (o *ProjectAttributesFilterModel) HasIsRequired() bool {
	if o != nil && o.IsRequired.IsSet() {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given NullableBool and assigns it to the IsRequired field.
func (o *ProjectAttributesFilterModel) SetIsRequired(v bool) {
	o.IsRequired.Set(&v)
}
// SetIsRequiredNil sets the value for IsRequired to be an explicit nil
func (o *ProjectAttributesFilterModel) SetIsRequiredNil() {
	o.IsRequired.Set(nil)
}

// UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
func (o *ProjectAttributesFilterModel) UnsetIsRequired() {
	o.IsRequired.Unset()
}

// GetIsGlobal returns the IsGlobal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAttributesFilterModel) GetIsGlobal() bool {
	if o == nil || IsNil(o.IsGlobal.Get()) {
		var ret bool
		return ret
	}
	return *o.IsGlobal.Get()
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAttributesFilterModel) GetIsGlobalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsGlobal.Get(), o.IsGlobal.IsSet()
}

// HasIsGlobal returns a boolean if a field has been set.
func (o *ProjectAttributesFilterModel) HasIsGlobal() bool {
	if o != nil && o.IsGlobal.IsSet() {
		return true
	}

	return false
}

// SetIsGlobal gets a reference to the given NullableBool and assigns it to the IsGlobal field.
func (o *ProjectAttributesFilterModel) SetIsGlobal(v bool) {
	o.IsGlobal.Set(&v)
}
// SetIsGlobalNil sets the value for IsGlobal to be an explicit nil
func (o *ProjectAttributesFilterModel) SetIsGlobalNil() {
	o.IsGlobal.Set(nil)
}

// UnsetIsGlobal ensures that no value is present for IsGlobal, not even an explicit nil
func (o *ProjectAttributesFilterModel) UnsetIsGlobal() {
	o.IsGlobal.Unset()
}

// GetTypes returns the Types field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAttributesFilterModel) GetTypes() []CustomAttributeTypesEnum {
	if o == nil {
		var ret []CustomAttributeTypesEnum
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAttributesFilterModel) GetTypesOk() ([]CustomAttributeTypesEnum, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *ProjectAttributesFilterModel) HasTypes() bool {
	if o != nil && IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []CustomAttributeTypesEnum and assigns it to the Types field.
func (o *ProjectAttributesFilterModel) SetTypes(v []CustomAttributeTypesEnum) {
	o.Types = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAttributesFilterModel) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEnabled.Get()
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAttributesFilterModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEnabled.Get(), o.IsEnabled.IsSet()
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *ProjectAttributesFilterModel) HasIsEnabled() bool {
	if o != nil && o.IsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given NullableBool and assigns it to the IsEnabled field.
func (o *ProjectAttributesFilterModel) SetIsEnabled(v bool) {
	o.IsEnabled.Set(&v)
}
// SetIsEnabledNil sets the value for IsEnabled to be an explicit nil
func (o *ProjectAttributesFilterModel) SetIsEnabledNil() {
	o.IsEnabled.Set(nil)
}

// UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
func (o *ProjectAttributesFilterModel) UnsetIsEnabled() {
	o.IsEnabled.Unset()
}

func (o ProjectAttributesFilterModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectAttributesFilterModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IsRequired.IsSet() {
		toSerialize["isRequired"] = o.IsRequired.Get()
	}
	if o.IsGlobal.IsSet() {
		toSerialize["isGlobal"] = o.IsGlobal.Get()
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.IsEnabled.IsSet() {
		toSerialize["isEnabled"] = o.IsEnabled.Get()
	}
	return toSerialize, nil
}

type NullableProjectAttributesFilterModel struct {
	value *ProjectAttributesFilterModel
	isSet bool
}

func (v NullableProjectAttributesFilterModel) Get() *ProjectAttributesFilterModel {
	return v.value
}

func (v *NullableProjectAttributesFilterModel) Set(val *ProjectAttributesFilterModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAttributesFilterModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAttributesFilterModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAttributesFilterModel(val *ProjectAttributesFilterModel) *NullableProjectAttributesFilterModel {
	return &NullableProjectAttributesFilterModel{value: val, isSet: true}
}

func (v NullableProjectAttributesFilterModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAttributesFilterModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

