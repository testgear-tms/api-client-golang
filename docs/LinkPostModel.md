# LinkPostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | Link name. | [optional] 
**Url** | **string** | Address can be specified without protocol, but necessarily with the domain. | 
**Description** | Pointer to **NullableString** | Link description. | [optional] 
**Type** | Pointer to [**LinkType**](LinkType.md) |  | [optional] 
**HasInfo** | Pointer to **bool** |  | [optional] 

## Methods

### NewLinkPostModel

`func NewLinkPostModel(url string, ) *LinkPostModel`

NewLinkPostModel instantiates a new LinkPostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkPostModelWithDefaults

`func NewLinkPostModelWithDefaults() *LinkPostModel`

NewLinkPostModelWithDefaults instantiates a new LinkPostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *LinkPostModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LinkPostModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LinkPostModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LinkPostModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *LinkPostModel) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *LinkPostModel) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrl

`func (o *LinkPostModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LinkPostModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LinkPostModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *LinkPostModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LinkPostModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LinkPostModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LinkPostModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LinkPostModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LinkPostModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *LinkPostModel) GetType() LinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkPostModel) GetTypeOk() (*LinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkPostModel) SetType(v LinkType)`

SetType sets Type field to given value.

### HasType

`func (o *LinkPostModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHasInfo

`func (o *LinkPostModel) GetHasInfo() bool`

GetHasInfo returns the HasInfo field if non-nil, zero value otherwise.

### GetHasInfoOk

`func (o *LinkPostModel) GetHasInfoOk() (*bool, bool)`

GetHasInfoOk returns a tuple with the HasInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInfo

`func (o *LinkPostModel) SetHasInfo(v bool)`

SetHasInfo sets HasInfo field to given value.

### HasHasInfo

`func (o *LinkPostModel) HasHasInfo() bool`

HasHasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


