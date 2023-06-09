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

// checks if the AutoTestModelV2GetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTestModelV2GetModel{}

// AutoTestModelV2GetModel struct for AutoTestModelV2GetModel
type AutoTestModelV2GetModel struct {
	// This property is used to set autotest identifier from client system
	ExternalId NullableString `json:"externalId,omitempty"`
	Links []LinkModel `json:"links,omitempty"`
	// This property is used to link autotest with project
	ProjectId *string `json:"projectId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
	Classname NullableString `json:"classname,omitempty"`
	Steps []AutoTestStepModel `json:"steps,omitempty"`
	Setup []AutoTestStepModel `json:"setup,omitempty"`
	Teardown []AutoTestStepModel `json:"teardown,omitempty"`
	GlobalId *int64 `json:"globalId,omitempty"`
	CreatedDate NullableTime `json:"createdDate,omitempty"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	CreatedById *string `json:"createdById,omitempty"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	Labels []LabelShortModel `json:"labels,omitempty"`
	// Unique ID of the entity
	Id *string `json:"id,omitempty"`
	// Indicates if the entity is deleted
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// NewAutoTestModelV2GetModel instantiates a new AutoTestModelV2GetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTestModelV2GetModel() *AutoTestModelV2GetModel {
	this := AutoTestModelV2GetModel{}
	return &this
}

// NewAutoTestModelV2GetModelWithDefaults instantiates a new AutoTestModelV2GetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTestModelV2GetModelWithDefaults() *AutoTestModelV2GetModel {
	this := AutoTestModelV2GetModel{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModelV2GetModel) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModelV2GetModel) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *AutoTestModelV2GetModel) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *AutoTestModelV2GetModel) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *AutoTestModelV2GetModel) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModelV2GetModel) GetLinks() []LinkModel {
	if o == nil {
		var ret []LinkModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModelV2GetModel) GetLinksOk() ([]LinkModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasLinks() bool {
	if o != nil && IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkModel and assigns it to the Links field.
func (o *AutoTestModelV2GetModel) SetLinks(v []LinkModel) {
	o.Links = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AutoTestModelV2GetModel) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTestModelV2GetModel) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *AutoTestModelV2GetModel) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModelV2GetModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModelV2GetModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AutoTestModelV2GetModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AutoTestModelV2GetModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AutoTestModelV2GetModel) UnsetName() {
	o.Name.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModelV2GetModel) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModelV2GetModel) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *AutoTestModelV2GetModel) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *AutoTestModelV2GetModel) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *AutoTestModelV2GetModel) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetClassname returns the Classname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModelV2GetModel) GetClassname() string {
	if o == nil || IsNil(o.Classname.Get()) {
		var ret string
		return ret
	}
	return *o.Classname.Get()
}

// GetClassnameOk returns a tuple with the Classname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModelV2GetModel) GetClassnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Classname.Get(), o.Classname.IsSet()
}

// HasClassname returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasClassname() bool {
	if o != nil && o.Classname.IsSet() {
		return true
	}

	return false
}

// SetClassname gets a reference to the given NullableString and assigns it to the Classname field.
func (o *AutoTestModelV2GetModel) SetClassname(v string) {
	o.Classname.Set(&v)
}
// SetClassnameNil sets the value for Classname to be an explicit nil
func (o *AutoTestModelV2GetModel) SetClassnameNil() {
	o.Classname.Set(nil)
}

// UnsetClassname ensures that no value is present for Classname, not even an explicit nil
func (o *AutoTestModelV2GetModel) UnsetClassname() {
	o.Classname.Unset()
}

// GetSteps returns the Steps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModelV2GetModel) GetSteps() []AutoTestStepModel {
	if o == nil {
		var ret []AutoTestStepModel
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModelV2GetModel) GetStepsOk() ([]AutoTestStepModel, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasSteps() bool {
	if o != nil && IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []AutoTestStepModel and assigns it to the Steps field.
func (o *AutoTestModelV2GetModel) SetSteps(v []AutoTestStepModel) {
	o.Steps = v
}

// GetSetup returns the Setup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModelV2GetModel) GetSetup() []AutoTestStepModel {
	if o == nil {
		var ret []AutoTestStepModel
		return ret
	}
	return o.Setup
}

// GetSetupOk returns a tuple with the Setup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModelV2GetModel) GetSetupOk() ([]AutoTestStepModel, bool) {
	if o == nil || IsNil(o.Setup) {
		return nil, false
	}
	return o.Setup, true
}

// HasSetup returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasSetup() bool {
	if o != nil && IsNil(o.Setup) {
		return true
	}

	return false
}

// SetSetup gets a reference to the given []AutoTestStepModel and assigns it to the Setup field.
func (o *AutoTestModelV2GetModel) SetSetup(v []AutoTestStepModel) {
	o.Setup = v
}

// GetTeardown returns the Teardown field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModelV2GetModel) GetTeardown() []AutoTestStepModel {
	if o == nil {
		var ret []AutoTestStepModel
		return ret
	}
	return o.Teardown
}

// GetTeardownOk returns a tuple with the Teardown field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModelV2GetModel) GetTeardownOk() ([]AutoTestStepModel, bool) {
	if o == nil || IsNil(o.Teardown) {
		return nil, false
	}
	return o.Teardown, true
}

// HasTeardown returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasTeardown() bool {
	if o != nil && IsNil(o.Teardown) {
		return true
	}

	return false
}

// SetTeardown gets a reference to the given []AutoTestStepModel and assigns it to the Teardown field.
func (o *AutoTestModelV2GetModel) SetTeardown(v []AutoTestStepModel) {
	o.Teardown = v
}

// GetGlobalId returns the GlobalId field value if set, zero value otherwise.
func (o *AutoTestModelV2GetModel) GetGlobalId() int64 {
	if o == nil || IsNil(o.GlobalId) {
		var ret int64
		return ret
	}
	return *o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTestModelV2GetModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GlobalId) {
		return nil, false
	}
	return o.GlobalId, true
}

// HasGlobalId returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasGlobalId() bool {
	if o != nil && !IsNil(o.GlobalId) {
		return true
	}

	return false
}

// SetGlobalId gets a reference to the given int64 and assigns it to the GlobalId field.
func (o *AutoTestModelV2GetModel) SetGlobalId(v int64) {
	o.GlobalId = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModelV2GetModel) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModelV2GetModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTime and assigns it to the CreatedDate field.
func (o *AutoTestModelV2GetModel) SetCreatedDate(v time.Time) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *AutoTestModelV2GetModel) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *AutoTestModelV2GetModel) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModelV2GetModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModelV2GetModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *AutoTestModelV2GetModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *AutoTestModelV2GetModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *AutoTestModelV2GetModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *AutoTestModelV2GetModel) GetCreatedById() string {
	if o == nil || IsNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTestModelV2GetModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *AutoTestModelV2GetModel) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModelV2GetModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModelV2GetModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *AutoTestModelV2GetModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *AutoTestModelV2GetModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *AutoTestModelV2GetModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModelV2GetModel) GetLabels() []LabelShortModel {
	if o == nil {
		var ret []LabelShortModel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModelV2GetModel) GetLabelsOk() ([]LabelShortModel, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasLabels() bool {
	if o != nil && IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelShortModel and assigns it to the Labels field.
func (o *AutoTestModelV2GetModel) SetLabels(v []LabelShortModel) {
	o.Labels = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutoTestModelV2GetModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTestModelV2GetModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AutoTestModelV2GetModel) SetId(v string) {
	o.Id = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *AutoTestModelV2GetModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTestModelV2GetModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *AutoTestModelV2GetModel) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *AutoTestModelV2GetModel) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

func (o AutoTestModelV2GetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTestModelV2GetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Classname.IsSet() {
		toSerialize["classname"] = o.Classname.Get()
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.Setup != nil {
		toSerialize["setup"] = o.Setup
	}
	if o.Teardown != nil {
		toSerialize["teardown"] = o.Teardown
	}
	if !IsNil(o.GlobalId) {
		toSerialize["globalId"] = o.GlobalId
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
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
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	return toSerialize, nil
}

type NullableAutoTestModelV2GetModel struct {
	value *AutoTestModelV2GetModel
	isSet bool
}

func (v NullableAutoTestModelV2GetModel) Get() *AutoTestModelV2GetModel {
	return v.value
}

func (v *NullableAutoTestModelV2GetModel) Set(val *AutoTestModelV2GetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTestModelV2GetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTestModelV2GetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTestModelV2GetModel(val *AutoTestModelV2GetModel) *NullableAutoTestModelV2GetModel {
	return &NullableAutoTestModelV2GetModel{value: val, isSet: true}
}

func (v NullableAutoTestModelV2GetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTestModelV2GetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


