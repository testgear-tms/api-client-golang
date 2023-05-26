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

// LinkType the model 'LinkType'
type LinkType string

// List of LinkType
const (
	LINKTYPE_RELATED LinkType = "Related"
	LINKTYPE_BLOCKED_BY LinkType = "BlockedBy"
	LINKTYPE_DEFECT LinkType = "Defect"
	LINKTYPE_ISSUE LinkType = "Issue"
	LINKTYPE_REQUIREMENT LinkType = "Requirement"
	LINKTYPE_REPOSITORY LinkType = "Repository"
)

// All allowed values of LinkType enum
var AllowedLinkTypeEnumValues = []LinkType{
	"Related",
	"BlockedBy",
	"Defect",
	"Issue",
	"Requirement",
	"Repository",
}

func (v *LinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LinkType(value)
	for _, existing := range AllowedLinkTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LinkType", value)
}

// NewLinkTypeFromValue returns a pointer to a valid LinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLinkTypeFromValue(v string) (*LinkType, error) {
	ev := LinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LinkType: valid values are %v", v, AllowedLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LinkType) IsValid() bool {
	for _, existing := range AllowedLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LinkType value
func (v LinkType) Ptr() *LinkType {
	return &v
}

type NullableLinkType struct {
	value *LinkType
	isSet bool
}

func (v NullableLinkType) Get() *LinkType {
	return v.value
}

func (v *NullableLinkType) Set(val *LinkType) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkType) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkType(val *LinkType) *NullableLinkType {
	return &NullableLinkType{value: val, isSet: true}
}

func (v NullableLinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
