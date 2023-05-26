# TestPointAnalyticResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountGroupByStatus** | Pointer to [**[]TestPlanGroupByStatus**](TestPlanGroupByStatus.md) |  | [optional] 
**SumGroupByTester** | Pointer to [**[]TestPlanGroupByTester**](TestPlanGroupByTester.md) |  | [optional] 
**CountGroupByTester** | Pointer to [**[]TestPlanGroupByTester**](TestPlanGroupByTester.md) |  | [optional] 
**CountGroupByTestSuite** | Pointer to [**[]TestPlanGroupByTestSuite**](TestPlanGroupByTestSuite.md) |  | [optional] 
**CountGroupByTesterAndStatus** | Pointer to [**[]TestPlanGroupByTesterAndStatus**](TestPlanGroupByTesterAndStatus.md) |  | [optional] 

## Methods

### NewTestPointAnalyticResult

`func NewTestPointAnalyticResult() *TestPointAnalyticResult`

NewTestPointAnalyticResult instantiates a new TestPointAnalyticResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointAnalyticResultWithDefaults

`func NewTestPointAnalyticResultWithDefaults() *TestPointAnalyticResult`

NewTestPointAnalyticResultWithDefaults instantiates a new TestPointAnalyticResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountGroupByStatus

`func (o *TestPointAnalyticResult) GetCountGroupByStatus() []TestPlanGroupByStatus`

GetCountGroupByStatus returns the CountGroupByStatus field if non-nil, zero value otherwise.

### GetCountGroupByStatusOk

`func (o *TestPointAnalyticResult) GetCountGroupByStatusOk() (*[]TestPlanGroupByStatus, bool)`

GetCountGroupByStatusOk returns a tuple with the CountGroupByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountGroupByStatus

`func (o *TestPointAnalyticResult) SetCountGroupByStatus(v []TestPlanGroupByStatus)`

SetCountGroupByStatus sets CountGroupByStatus field to given value.

### HasCountGroupByStatus

`func (o *TestPointAnalyticResult) HasCountGroupByStatus() bool`

HasCountGroupByStatus returns a boolean if a field has been set.

### SetCountGroupByStatusNil

`func (o *TestPointAnalyticResult) SetCountGroupByStatusNil(b bool)`

 SetCountGroupByStatusNil sets the value for CountGroupByStatus to be an explicit nil

### UnsetCountGroupByStatus
`func (o *TestPointAnalyticResult) UnsetCountGroupByStatus()`

UnsetCountGroupByStatus ensures that no value is present for CountGroupByStatus, not even an explicit nil
### GetSumGroupByTester

`func (o *TestPointAnalyticResult) GetSumGroupByTester() []TestPlanGroupByTester`

GetSumGroupByTester returns the SumGroupByTester field if non-nil, zero value otherwise.

### GetSumGroupByTesterOk

`func (o *TestPointAnalyticResult) GetSumGroupByTesterOk() (*[]TestPlanGroupByTester, bool)`

GetSumGroupByTesterOk returns a tuple with the SumGroupByTester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGroupByTester

`func (o *TestPointAnalyticResult) SetSumGroupByTester(v []TestPlanGroupByTester)`

SetSumGroupByTester sets SumGroupByTester field to given value.

### HasSumGroupByTester

`func (o *TestPointAnalyticResult) HasSumGroupByTester() bool`

HasSumGroupByTester returns a boolean if a field has been set.

### SetSumGroupByTesterNil

`func (o *TestPointAnalyticResult) SetSumGroupByTesterNil(b bool)`

 SetSumGroupByTesterNil sets the value for SumGroupByTester to be an explicit nil

### UnsetSumGroupByTester
`func (o *TestPointAnalyticResult) UnsetSumGroupByTester()`

UnsetSumGroupByTester ensures that no value is present for SumGroupByTester, not even an explicit nil
### GetCountGroupByTester

`func (o *TestPointAnalyticResult) GetCountGroupByTester() []TestPlanGroupByTester`

GetCountGroupByTester returns the CountGroupByTester field if non-nil, zero value otherwise.

### GetCountGroupByTesterOk

`func (o *TestPointAnalyticResult) GetCountGroupByTesterOk() (*[]TestPlanGroupByTester, bool)`

GetCountGroupByTesterOk returns a tuple with the CountGroupByTester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountGroupByTester

`func (o *TestPointAnalyticResult) SetCountGroupByTester(v []TestPlanGroupByTester)`

SetCountGroupByTester sets CountGroupByTester field to given value.

### HasCountGroupByTester

`func (o *TestPointAnalyticResult) HasCountGroupByTester() bool`

HasCountGroupByTester returns a boolean if a field has been set.

### SetCountGroupByTesterNil

`func (o *TestPointAnalyticResult) SetCountGroupByTesterNil(b bool)`

 SetCountGroupByTesterNil sets the value for CountGroupByTester to be an explicit nil

### UnsetCountGroupByTester
`func (o *TestPointAnalyticResult) UnsetCountGroupByTester()`

UnsetCountGroupByTester ensures that no value is present for CountGroupByTester, not even an explicit nil
### GetCountGroupByTestSuite

`func (o *TestPointAnalyticResult) GetCountGroupByTestSuite() []TestPlanGroupByTestSuite`

GetCountGroupByTestSuite returns the CountGroupByTestSuite field if non-nil, zero value otherwise.

### GetCountGroupByTestSuiteOk

`func (o *TestPointAnalyticResult) GetCountGroupByTestSuiteOk() (*[]TestPlanGroupByTestSuite, bool)`

GetCountGroupByTestSuiteOk returns a tuple with the CountGroupByTestSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountGroupByTestSuite

`func (o *TestPointAnalyticResult) SetCountGroupByTestSuite(v []TestPlanGroupByTestSuite)`

SetCountGroupByTestSuite sets CountGroupByTestSuite field to given value.

### HasCountGroupByTestSuite

`func (o *TestPointAnalyticResult) HasCountGroupByTestSuite() bool`

HasCountGroupByTestSuite returns a boolean if a field has been set.

### SetCountGroupByTestSuiteNil

`func (o *TestPointAnalyticResult) SetCountGroupByTestSuiteNil(b bool)`

 SetCountGroupByTestSuiteNil sets the value for CountGroupByTestSuite to be an explicit nil

### UnsetCountGroupByTestSuite
`func (o *TestPointAnalyticResult) UnsetCountGroupByTestSuite()`

UnsetCountGroupByTestSuite ensures that no value is present for CountGroupByTestSuite, not even an explicit nil
### GetCountGroupByTesterAndStatus

`func (o *TestPointAnalyticResult) GetCountGroupByTesterAndStatus() []TestPlanGroupByTesterAndStatus`

GetCountGroupByTesterAndStatus returns the CountGroupByTesterAndStatus field if non-nil, zero value otherwise.

### GetCountGroupByTesterAndStatusOk

`func (o *TestPointAnalyticResult) GetCountGroupByTesterAndStatusOk() (*[]TestPlanGroupByTesterAndStatus, bool)`

GetCountGroupByTesterAndStatusOk returns a tuple with the CountGroupByTesterAndStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountGroupByTesterAndStatus

`func (o *TestPointAnalyticResult) SetCountGroupByTesterAndStatus(v []TestPlanGroupByTesterAndStatus)`

SetCountGroupByTesterAndStatus sets CountGroupByTesterAndStatus field to given value.

### HasCountGroupByTesterAndStatus

`func (o *TestPointAnalyticResult) HasCountGroupByTesterAndStatus() bool`

HasCountGroupByTesterAndStatus returns a boolean if a field has been set.

### SetCountGroupByTesterAndStatusNil

`func (o *TestPointAnalyticResult) SetCountGroupByTesterAndStatusNil(b bool)`

 SetCountGroupByTesterAndStatusNil sets the value for CountGroupByTesterAndStatus to be an explicit nil

### UnsetCountGroupByTesterAndStatus
`func (o *TestPointAnalyticResult) UnsetCountGroupByTesterAndStatus()`

UnsetCountGroupByTesterAndStatus ensures that no value is present for CountGroupByTesterAndStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


