# LinkPutModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **NullableString** | Link name. | [optional] 
**Url** | **string** | Address can be specified without protocol, but necessarily with the domain. | 
**Description** | Pointer to **NullableString** | Link description. | [optional] 
**Type** | Pointer to [**LinkType**](LinkType.md) |  | [optional] 
**HasInfo** | Pointer to **bool** |  | [optional] 

## Methods

### NewLinkPutModel

`func NewLinkPutModel(url string, ) *LinkPutModel`

NewLinkPutModel instantiates a new LinkPutModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkPutModelWithDefaults

`func NewLinkPutModelWithDefaults() *LinkPutModel`

NewLinkPutModelWithDefaults instantiates a new LinkPutModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkPutModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkPutModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkPutModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LinkPutModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *LinkPutModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LinkPutModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LinkPutModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LinkPutModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *LinkPutModel) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *LinkPutModel) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrl

`func (o *LinkPutModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LinkPutModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LinkPutModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *LinkPutModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LinkPutModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LinkPutModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LinkPutModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LinkPutModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LinkPutModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *LinkPutModel) GetType() LinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkPutModel) GetTypeOk() (*LinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkPutModel) SetType(v LinkType)`

SetType sets Type field to given value.

### HasType

`func (o *LinkPutModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHasInfo

`func (o *LinkPutModel) GetHasInfo() bool`

GetHasInfo returns the HasInfo field if non-nil, zero value otherwise.

### GetHasInfoOk

`func (o *LinkPutModel) GetHasInfoOk() (*bool, bool)`

GetHasInfoOk returns a tuple with the HasInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInfo

`func (o *LinkPutModel) SetHasInfo(v bool)`

SetHasInfo sets HasInfo field to given value.

### HasHasInfo

`func (o *LinkPutModel) HasHasInfo() bool`

HasHasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


