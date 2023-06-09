/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"fmt"
)

// WorkItemPriorityModel the model 'WorkItemPriorityModel'
type WorkItemPriorityModel string

// List of WorkItemPriorityModel
const (
	WORKITEMPRIORITYMODEL_LOWEST WorkItemPriorityModel = "Lowest"
	WORKITEMPRIORITYMODEL_LOW WorkItemPriorityModel = "Low"
	WORKITEMPRIORITYMODEL_MEDIUM WorkItemPriorityModel = "Medium"
	WORKITEMPRIORITYMODEL_HIGH WorkItemPriorityModel = "High"
	WORKITEMPRIORITYMODEL_HIGHEST WorkItemPriorityModel = "Highest"
)

// All allowed values of WorkItemPriorityModel enum
var AllowedWorkItemPriorityModelEnumValues = []WorkItemPriorityModel{
	"Lowest",
	"Low",
	"Medium",
	"High",
	"Highest",
}

func (v *WorkItemPriorityModel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WorkItemPriorityModel(value)
	for _, existing := range AllowedWorkItemPriorityModelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WorkItemPriorityModel", value)
}

// NewWorkItemPriorityModelFromValue returns a pointer to a valid WorkItemPriorityModel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWorkItemPriorityModelFromValue(v string) (*WorkItemPriorityModel, error) {
	ev := WorkItemPriorityModel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WorkItemPriorityModel: valid values are %v", v, AllowedWorkItemPriorityModelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WorkItemPriorityModel) IsValid() bool {
	for _, existing := range AllowedWorkItemPriorityModelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkItemPriorityModel value
func (v WorkItemPriorityModel) Ptr() *WorkItemPriorityModel {
	return &v
}

type NullableWorkItemPriorityModel struct {
	value *WorkItemPriorityModel
	isSet bool
}

func (v NullableWorkItemPriorityModel) Get() *WorkItemPriorityModel {
	return v.value
}

func (v *NullableWorkItemPriorityModel) Set(val *WorkItemPriorityModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemPriorityModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemPriorityModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemPriorityModel(val *WorkItemPriorityModel) *NullableWorkItemPriorityModel {
	return &NullableWorkItemPriorityModel{value: val, isSet: true}
}

func (v NullableWorkItemPriorityModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemPriorityModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

