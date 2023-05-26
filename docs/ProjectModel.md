# ProjectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the project | [optional] 
**Description** | Pointer to **NullableString** | Description of the project | [optional] 
**Name** | Pointer to **NullableString** | Name of the project | [optional] 
**IsFavorite** | Pointer to **bool** | Indicates if the project is marked as favorite | [optional] 
**AttributesScheme** | Pointer to [**[]CustomAttributeModel**](CustomAttributeModel.md) | Collection of the project attributes | [optional] 
**TestPlansAttributesScheme** | Pointer to [**[]CustomAttributeModel**](CustomAttributeModel.md) | Collection of the project test plans attributes | [optional] 
**TestCasesCount** | Pointer to **NullableInt32** | Number of test cases in the project | [optional] 
**SharedStepsCount** | Pointer to **NullableInt32** | Number of shared steps in the project | [optional] 
**CheckListsCount** | Pointer to **NullableInt32** | Number of checklists in the project | [optional] 
**AutoTestsCount** | Pointer to **NullableInt32** | Number of autotests in the project | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates if the project is deleted | [optional] 
**CreatedDate** | Pointer to **time.Time** | Creation date of the project | [optional] 
**ModifiedDate** | Pointer to **NullableTime** | Last modification date of the project | [optional] 
**CreatedById** | Pointer to **string** | Unique ID of the project creator | [optional] 
**ModifiedById** | Pointer to **NullableString** | Unique ID of the project last editor | [optional] 
**GlobalId** | Pointer to **int64** | Global ID of the project | [optional] 

## Methods

### NewProjectModel

`func NewProjectModel() *ProjectModel`

NewProjectModel instantiates a new ProjectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectModelWithDefaults

`func NewProjectModelWithDefaults() *ProjectModel`

NewProjectModelWithDefaults instantiates a new ProjectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *ProjectModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsFavorite

`func (o *ProjectModel) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ProjectModel) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ProjectModel) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *ProjectModel) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetAttributesScheme

`func (o *ProjectModel) GetAttributesScheme() []CustomAttributeModel`

GetAttributesScheme returns the AttributesScheme field if non-nil, zero value otherwise.

### GetAttributesSchemeOk

`func (o *ProjectModel) GetAttributesSchemeOk() (*[]CustomAttributeModel, bool)`

GetAttributesSchemeOk returns a tuple with the AttributesScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesScheme

`func (o *ProjectModel) SetAttributesScheme(v []CustomAttributeModel)`

SetAttributesScheme sets AttributesScheme field to given value.

### HasAttributesScheme

`func (o *ProjectModel) HasAttributesScheme() bool`

HasAttributesScheme returns a boolean if a field has been set.

### SetAttributesSchemeNil

`func (o *ProjectModel) SetAttributesSchemeNil(b bool)`

 SetAttributesSchemeNil sets the value for AttributesScheme to be an explicit nil

### UnsetAttributesScheme
`func (o *ProjectModel) UnsetAttributesScheme()`

UnsetAttributesScheme ensures that no value is present for AttributesScheme, not even an explicit nil
### GetTestPlansAttributesScheme

`func (o *ProjectModel) GetTestPlansAttributesScheme() []CustomAttributeModel`

GetTestPlansAttributesScheme returns the TestPlansAttributesScheme field if non-nil, zero value otherwise.

### GetTestPlansAttributesSchemeOk

`func (o *ProjectModel) GetTestPlansAttributesSchemeOk() (*[]CustomAttributeModel, bool)`

GetTestPlansAttributesSchemeOk returns a tuple with the TestPlansAttributesScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlansAttributesScheme

`func (o *ProjectModel) SetTestPlansAttributesScheme(v []CustomAttributeModel)`

SetTestPlansAttributesScheme sets TestPlansAttributesScheme field to given value.

### HasTestPlansAttributesScheme

`func (o *ProjectModel) HasTestPlansAttributesScheme() bool`

HasTestPlansAttributesScheme returns a boolean if a field has been set.

### SetTestPlansAttributesSchemeNil

`func (o *ProjectModel) SetTestPlansAttributesSchemeNil(b bool)`

 SetTestPlansAttributesSchemeNil sets the value for TestPlansAttributesScheme to be an explicit nil

### UnsetTestPlansAttributesScheme
`func (o *ProjectModel) UnsetTestPlansAttributesScheme()`

UnsetTestPlansAttributesScheme ensures that no value is present for TestPlansAttributesScheme, not even an explicit nil
### GetTestCasesCount

`func (o *ProjectModel) GetTestCasesCount() int32`

GetTestCasesCount returns the TestCasesCount field if non-nil, zero value otherwise.

### GetTestCasesCountOk

`func (o *ProjectModel) GetTestCasesCountOk() (*int32, bool)`

GetTestCasesCountOk returns a tuple with the TestCasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCasesCount

`func (o *ProjectModel) SetTestCasesCount(v int32)`

SetTestCasesCount sets TestCasesCount field to given value.

### HasTestCasesCount

`func (o *ProjectModel) HasTestCasesCount() bool`

HasTestCasesCount returns a boolean if a field has been set.

### SetTestCasesCountNil

`func (o *ProjectModel) SetTestCasesCountNil(b bool)`

 SetTestCasesCountNil sets the value for TestCasesCount to be an explicit nil

### UnsetTestCasesCount
`func (o *ProjectModel) UnsetTestCasesCount()`

UnsetTestCasesCount ensures that no value is present for TestCasesCount, not even an explicit nil
### GetSharedStepsCount

`func (o *ProjectModel) GetSharedStepsCount() int32`

GetSharedStepsCount returns the SharedStepsCount field if non-nil, zero value otherwise.

### GetSharedStepsCountOk

`func (o *ProjectModel) GetSharedStepsCountOk() (*int32, bool)`

GetSharedStepsCountOk returns a tuple with the SharedStepsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepsCount

`func (o *ProjectModel) SetSharedStepsCount(v int32)`

SetSharedStepsCount sets SharedStepsCount field to given value.

### HasSharedStepsCount

`func (o *ProjectModel) HasSharedStepsCount() bool`

HasSharedStepsCount returns a boolean if a field has been set.

### SetSharedStepsCountNil

`func (o *ProjectModel) SetSharedStepsCountNil(b bool)`

 SetSharedStepsCountNil sets the value for SharedStepsCount to be an explicit nil

### UnsetSharedStepsCount
`func (o *ProjectModel) UnsetSharedStepsCount()`

UnsetSharedStepsCount ensures that no value is present for SharedStepsCount, not even an explicit nil
### GetCheckListsCount

`func (o *ProjectModel) GetCheckListsCount() int32`

GetCheckListsCount returns the CheckListsCount field if non-nil, zero value otherwise.

### GetCheckListsCountOk

`func (o *ProjectModel) GetCheckListsCountOk() (*int32, bool)`

GetCheckListsCountOk returns a tuple with the CheckListsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckListsCount

`func (o *ProjectModel) SetCheckListsCount(v int32)`

SetCheckListsCount sets CheckListsCount field to given value.

### HasCheckListsCount

`func (o *ProjectModel) HasCheckListsCount() bool`

HasCheckListsCount returns a boolean if a field has been set.

### SetCheckListsCountNil

`func (o *ProjectModel) SetCheckListsCountNil(b bool)`

 SetCheckListsCountNil sets the value for CheckListsCount to be an explicit nil

### UnsetCheckListsCount
`func (o *ProjectModel) UnsetCheckListsCount()`

UnsetCheckListsCount ensures that no value is present for CheckListsCount, not even an explicit nil
### GetAutoTestsCount

`func (o *ProjectModel) GetAutoTestsCount() int32`

GetAutoTestsCount returns the AutoTestsCount field if non-nil, zero value otherwise.

### GetAutoTestsCountOk

`func (o *ProjectModel) GetAutoTestsCountOk() (*int32, bool)`

GetAutoTestsCountOk returns a tuple with the AutoTestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestsCount

`func (o *ProjectModel) SetAutoTestsCount(v int32)`

SetAutoTestsCount sets AutoTestsCount field to given value.

### HasAutoTestsCount

`func (o *ProjectModel) HasAutoTestsCount() bool`

HasAutoTestsCount returns a boolean if a field has been set.

### SetAutoTestsCountNil

`func (o *ProjectModel) SetAutoTestsCountNil(b bool)`

 SetAutoTestsCountNil sets the value for AutoTestsCount to be an explicit nil

### UnsetAutoTestsCount
`func (o *ProjectModel) UnsetAutoTestsCount()`

UnsetAutoTestsCount ensures that no value is present for AutoTestsCount, not even an explicit nil
### GetIsDeleted

`func (o *ProjectModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProjectModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProjectModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ProjectModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ProjectModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ProjectModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *ProjectModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ProjectModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ProjectModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ProjectModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *ProjectModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ProjectModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *ProjectModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ProjectModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ProjectModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *ProjectModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *ProjectModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *ProjectModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *ProjectModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *ProjectModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *ProjectModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *ProjectModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetGlobalId

`func (o *ProjectModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ProjectModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ProjectModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *ProjectModel) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


