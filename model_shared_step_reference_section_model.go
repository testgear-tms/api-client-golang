/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"time"
)

// checks if the SharedStepReferenceSectionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedStepReferenceSectionModel{}

// SharedStepReferenceSectionModel struct for SharedStepReferenceSectionModel
type SharedStepReferenceSectionModel struct {
	Id *string `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	HasThisSharedStepAsPrecondition *bool `json:"hasThisSharedStepAsPrecondition,omitempty"`
	HasThisSharedStepAsPostcondition *bool `json:"hasThisSharedStepAsPostcondition,omitempty"`
	CreatedById *string `json:"createdById,omitempty"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	CreatedDate NullableTime `json:"createdDate,omitempty"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// NewSharedStepReferenceSectionModel instantiates a new SharedStepReferenceSectionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedStepReferenceSectionModel() *SharedStepReferenceSectionModel {
	this := SharedStepReferenceSectionModel{}
	return &this
}

// NewSharedStepReferenceSectionModelWithDefaults instantiates a new SharedStepReferenceSectionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedStepReferenceSectionModelWithDefaults() *SharedStepReferenceSectionModel {
	this := SharedStepReferenceSectionModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SharedStepReferenceSectionModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepReferenceSectionModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SharedStepReferenceSectionModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SharedStepReferenceSectionModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedStepReferenceSectionModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharedStepReferenceSectionModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SharedStepReferenceSectionModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SharedStepReferenceSectionModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SharedStepReferenceSectionModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SharedStepReferenceSectionModel) UnsetName() {
	o.Name.Unset()
}

// GetHasThisSharedStepAsPrecondition returns the HasThisSharedStepAsPrecondition field value if set, zero value otherwise.
func (o *SharedStepReferenceSectionModel) GetHasThisSharedStepAsPrecondition() bool {
	if o == nil || IsNil(o.HasThisSharedStepAsPrecondition) {
		var ret bool
		return ret
	}
	return *o.HasThisSharedStepAsPrecondition
}

// GetHasThisSharedStepAsPreconditionOk returns a tuple with the HasThisSharedStepAsPrecondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepReferenceSectionModel) GetHasThisSharedStepAsPreconditionOk() (*bool, bool) {
	if o == nil || IsNil(o.HasThisSharedStepAsPrecondition) {
		return nil, false
	}
	return o.HasThisSharedStepAsPrecondition, true
}

// HasHasThisSharedStepAsPrecondition returns a boolean if a field has been set.
func (o *SharedStepReferenceSectionModel) HasHasThisSharedStepAsPrecondition() bool {
	if o != nil && !IsNil(o.HasThisSharedStepAsPrecondition) {
		return true
	}

	return false
}

// SetHasThisSharedStepAsPrecondition gets a reference to the given bool and assigns it to the HasThisSharedStepAsPrecondition field.
func (o *SharedStepReferenceSectionModel) SetHasThisSharedStepAsPrecondition(v bool) {
	o.HasThisSharedStepAsPrecondition = &v
}

// GetHasThisSharedStepAsPostcondition returns the HasThisSharedStepAsPostcondition field value if set, zero value otherwise.
func (o *SharedStepReferenceSectionModel) GetHasThisSharedStepAsPostcondition() bool {
	if o == nil || IsNil(o.HasThisSharedStepAsPostcondition) {
		var ret bool
		return ret
	}
	return *o.HasThisSharedStepAsPostcondition
}

// GetHasThisSharedStepAsPostconditionOk returns a tuple with the HasThisSharedStepAsPostcondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepReferenceSectionModel) GetHasThisSharedStepAsPostconditionOk() (*bool, bool) {
	if o == nil || IsNil(o.HasThisSharedStepAsPostcondition) {
		return nil, false
	}
	return o.HasThisSharedStepAsPostcondition, true
}

// HasHasThisSharedStepAsPostcondition returns a boolean if a field has been set.
func (o *SharedStepReferenceSectionModel) HasHasThisSharedStepAsPostcondition() bool {
	if o != nil && !IsNil(o.HasThisSharedStepAsPostcondition) {
		return true
	}

	return false
}

// SetHasThisSharedStepAsPostcondition gets a reference to the given bool and assigns it to the HasThisSharedStepAsPostcondition field.
func (o *SharedStepReferenceSectionModel) SetHasThisSharedStepAsPostcondition(v bool) {
	o.HasThisSharedStepAsPostcondition = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *SharedStepReferenceSectionModel) GetCreatedById() string {
	if o == nil || IsNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepReferenceSectionModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *SharedStepReferenceSectionModel) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *SharedStepReferenceSectionModel) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedStepReferenceSectionModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharedStepReferenceSectionModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *SharedStepReferenceSectionModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *SharedStepReferenceSectionModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *SharedStepReferenceSectionModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *SharedStepReferenceSectionModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedStepReferenceSectionModel) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharedStepReferenceSectionModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *SharedStepReferenceSectionModel) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTime and assigns it to the CreatedDate field.
func (o *SharedStepReferenceSectionModel) SetCreatedDate(v time.Time) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *SharedStepReferenceSectionModel) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *SharedStepReferenceSectionModel) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedStepReferenceSectionModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharedStepReferenceSectionModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *SharedStepReferenceSectionModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *SharedStepReferenceSectionModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *SharedStepReferenceSectionModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *SharedStepReferenceSectionModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *SharedStepReferenceSectionModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepReferenceSectionModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *SharedStepReferenceSectionModel) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *SharedStepReferenceSectionModel) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

func (o SharedStepReferenceSectionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedStepReferenceSectionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.HasThisSharedStepAsPrecondition) {
		toSerialize["hasThisSharedStepAsPrecondition"] = o.HasThisSharedStepAsPrecondition
	}
	if !IsNil(o.HasThisSharedStepAsPostcondition) {
		toSerialize["hasThisSharedStepAsPostcondition"] = o.HasThisSharedStepAsPostcondition
	}
	if !IsNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	return toSerialize, nil
}

type NullableSharedStepReferenceSectionModel struct {
	value *SharedStepReferenceSectionModel
	isSet bool
}

func (v NullableSharedStepReferenceSectionModel) Get() *SharedStepReferenceSectionModel {
	return v.value
}

func (v *NullableSharedStepReferenceSectionModel) Set(val *SharedStepReferenceSectionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedStepReferenceSectionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedStepReferenceSectionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedStepReferenceSectionModel(val *SharedStepReferenceSectionModel) *NullableSharedStepReferenceSectionModel {
	return &NullableSharedStepReferenceSectionModel{value: val, isSet: true}
}

func (v NullableSharedStepReferenceSectionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedStepReferenceSectionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


