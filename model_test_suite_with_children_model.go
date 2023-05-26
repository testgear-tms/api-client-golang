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

// checks if the TestSuiteWithChildrenModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSuiteWithChildrenModel{}

// TestSuiteWithChildrenModel struct for TestSuiteWithChildrenModel
type TestSuiteWithChildrenModel struct {
	Children []TestSuiteWithChildrenModel `json:"children,omitempty"`
	TesterId NullableString `json:"testerId,omitempty"`
	ParentId NullableString `json:"parentId,omitempty"`
	TestPlanId *string `json:"testPlanId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	// Unique ID of the entity
	Id *string `json:"id,omitempty"`
	// Indicates if the entity is deleted
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// NewTestSuiteWithChildrenModel instantiates a new TestSuiteWithChildrenModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSuiteWithChildrenModel() *TestSuiteWithChildrenModel {
	this := TestSuiteWithChildrenModel{}
	return &this
}

// NewTestSuiteWithChildrenModelWithDefaults instantiates a new TestSuiteWithChildrenModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSuiteWithChildrenModelWithDefaults() *TestSuiteWithChildrenModel {
	this := TestSuiteWithChildrenModel{}
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteWithChildrenModel) GetChildren() []TestSuiteWithChildrenModel {
	if o == nil {
		var ret []TestSuiteWithChildrenModel
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteWithChildrenModel) GetChildrenOk() ([]TestSuiteWithChildrenModel, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *TestSuiteWithChildrenModel) HasChildren() bool {
	if o != nil && IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []TestSuiteWithChildrenModel and assigns it to the Children field.
func (o *TestSuiteWithChildrenModel) SetChildren(v []TestSuiteWithChildrenModel) {
	o.Children = v
}

// GetTesterId returns the TesterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteWithChildrenModel) GetTesterId() string {
	if o == nil || IsNil(o.TesterId.Get()) {
		var ret string
		return ret
	}
	return *o.TesterId.Get()
}

// GetTesterIdOk returns a tuple with the TesterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteWithChildrenModel) GetTesterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TesterId.Get(), o.TesterId.IsSet()
}

// HasTesterId returns a boolean if a field has been set.
func (o *TestSuiteWithChildrenModel) HasTesterId() bool {
	if o != nil && o.TesterId.IsSet() {
		return true
	}

	return false
}

// SetTesterId gets a reference to the given NullableString and assigns it to the TesterId field.
func (o *TestSuiteWithChildrenModel) SetTesterId(v string) {
	o.TesterId.Set(&v)
}
// SetTesterIdNil sets the value for TesterId to be an explicit nil
func (o *TestSuiteWithChildrenModel) SetTesterIdNil() {
	o.TesterId.Set(nil)
}

// UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
func (o *TestSuiteWithChildrenModel) UnsetTesterId() {
	o.TesterId.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteWithChildrenModel) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteWithChildrenModel) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *TestSuiteWithChildrenModel) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *TestSuiteWithChildrenModel) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *TestSuiteWithChildrenModel) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *TestSuiteWithChildrenModel) UnsetParentId() {
	o.ParentId.Unset()
}

// GetTestPlanId returns the TestPlanId field value if set, zero value otherwise.
func (o *TestSuiteWithChildrenModel) GetTestPlanId() string {
	if o == nil || IsNil(o.TestPlanId) {
		var ret string
		return ret
	}
	return *o.TestPlanId
}

// GetTestPlanIdOk returns a tuple with the TestPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSuiteWithChildrenModel) GetTestPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.TestPlanId) {
		return nil, false
	}
	return o.TestPlanId, true
}

// HasTestPlanId returns a boolean if a field has been set.
func (o *TestSuiteWithChildrenModel) HasTestPlanId() bool {
	if o != nil && !IsNil(o.TestPlanId) {
		return true
	}

	return false
}

// SetTestPlanId gets a reference to the given string and assigns it to the TestPlanId field.
func (o *TestSuiteWithChildrenModel) SetTestPlanId(v string) {
	o.TestPlanId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteWithChildrenModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteWithChildrenModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TestSuiteWithChildrenModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TestSuiteWithChildrenModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TestSuiteWithChildrenModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TestSuiteWithChildrenModel) UnsetName() {
	o.Name.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestSuiteWithChildrenModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSuiteWithChildrenModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestSuiteWithChildrenModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestSuiteWithChildrenModel) SetId(v string) {
	o.Id = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *TestSuiteWithChildrenModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSuiteWithChildrenModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *TestSuiteWithChildrenModel) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *TestSuiteWithChildrenModel) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

func (o TestSuiteWithChildrenModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSuiteWithChildrenModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.TesterId.IsSet() {
		toSerialize["testerId"] = o.TesterId.Get()
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	if !IsNil(o.TestPlanId) {
		toSerialize["testPlanId"] = o.TestPlanId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	return toSerialize, nil
}

type NullableTestSuiteWithChildrenModel struct {
	value *TestSuiteWithChildrenModel
	isSet bool
}

func (v NullableTestSuiteWithChildrenModel) Get() *TestSuiteWithChildrenModel {
	return v.value
}

func (v *NullableTestSuiteWithChildrenModel) Set(val *TestSuiteWithChildrenModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSuiteWithChildrenModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSuiteWithChildrenModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSuiteWithChildrenModel(val *TestSuiteWithChildrenModel) *NullableTestSuiteWithChildrenModel {
	return &NullableTestSuiteWithChildrenModel{value: val, isSet: true}
}

func (v NullableTestSuiteWithChildrenModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSuiteWithChildrenModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


