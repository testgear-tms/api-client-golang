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

// checks if the FailureClassModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailureClassModel{}

// FailureClassModel struct for FailureClassModel
type FailureClassModel struct {
	Name NullableString `json:"name,omitempty"`
	FailureCategory FailureCategoryModel `json:"failureCategory"`
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	CreatedById *string `json:"createdById,omitempty"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	FailureClassRegexes []FailureClassRegexModel `json:"failureClassRegexes,omitempty"`
	// Unique ID of the entity
	Id *string `json:"id,omitempty"`
	// Indicates if the entity is deleted
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// NewFailureClassModel instantiates a new FailureClassModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailureClassModel(failureCategory FailureCategoryModel) *FailureClassModel {
	this := FailureClassModel{}
	this.FailureCategory = failureCategory
	return &this
}

// NewFailureClassModelWithDefaults instantiates a new FailureClassModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailureClassModelWithDefaults() *FailureClassModel {
	this := FailureClassModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FailureClassModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailureClassModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FailureClassModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FailureClassModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *FailureClassModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FailureClassModel) UnsetName() {
	o.Name.Unset()
}

// GetFailureCategory returns the FailureCategory field value
func (o *FailureClassModel) GetFailureCategory() FailureCategoryModel {
	if o == nil {
		var ret FailureCategoryModel
		return ret
	}

	return o.FailureCategory
}

// GetFailureCategoryOk returns a tuple with the FailureCategory field value
// and a boolean to check if the value has been set.
func (o *FailureClassModel) GetFailureCategoryOk() (*FailureCategoryModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureCategory, true
}

// SetFailureCategory sets field value
func (o *FailureClassModel) SetFailureCategory(v FailureCategoryModel) {
	o.FailureCategory = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *FailureClassModel) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureClassModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *FailureClassModel) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *FailureClassModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FailureClassModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailureClassModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *FailureClassModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *FailureClassModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *FailureClassModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *FailureClassModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *FailureClassModel) GetCreatedById() string {
	if o == nil || IsNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureClassModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *FailureClassModel) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *FailureClassModel) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FailureClassModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailureClassModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *FailureClassModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *FailureClassModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *FailureClassModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *FailureClassModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetFailureClassRegexes returns the FailureClassRegexes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FailureClassModel) GetFailureClassRegexes() []FailureClassRegexModel {
	if o == nil {
		var ret []FailureClassRegexModel
		return ret
	}
	return o.FailureClassRegexes
}

// GetFailureClassRegexesOk returns a tuple with the FailureClassRegexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailureClassModel) GetFailureClassRegexesOk() ([]FailureClassRegexModel, bool) {
	if o == nil || IsNil(o.FailureClassRegexes) {
		return nil, false
	}
	return o.FailureClassRegexes, true
}

// HasFailureClassRegexes returns a boolean if a field has been set.
func (o *FailureClassModel) HasFailureClassRegexes() bool {
	if o != nil && IsNil(o.FailureClassRegexes) {
		return true
	}

	return false
}

// SetFailureClassRegexes gets a reference to the given []FailureClassRegexModel and assigns it to the FailureClassRegexes field.
func (o *FailureClassModel) SetFailureClassRegexes(v []FailureClassRegexModel) {
	o.FailureClassRegexes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FailureClassModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureClassModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FailureClassModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FailureClassModel) SetId(v string) {
	o.Id = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *FailureClassModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureClassModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *FailureClassModel) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *FailureClassModel) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

func (o FailureClassModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailureClassModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["failureCategory"] = o.FailureCategory
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	if !IsNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	if o.FailureClassRegexes != nil {
		toSerialize["failureClassRegexes"] = o.FailureClassRegexes
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	return toSerialize, nil
}

type NullableFailureClassModel struct {
	value *FailureClassModel
	isSet bool
}

func (v NullableFailureClassModel) Get() *FailureClassModel {
	return v.value
}

func (v *NullableFailureClassModel) Set(val *FailureClassModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureClassModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureClassModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureClassModel(val *FailureClassModel) *NullableFailureClassModel {
	return &NullableFailureClassModel{value: val, isSet: true}
}

func (v NullableFailureClassModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureClassModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

