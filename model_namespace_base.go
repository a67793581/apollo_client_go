/*
 * 阿波罗客户端
 *
 * 阿波罗客户端
 *
 * API version: v0.0.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// NamespaceBase struct for NamespaceBase
type NamespaceBase struct {
	Name                *string `json:"name,omitempty"`
	AppId               *string `json:"appId,omitempty"`
	Comment             *string `json:"comment,omitempty"`
	Format              *string `json:"format,omitempty"`
	IsPublic            *bool   `json:"isPublic,omitempty"`
	DataChangeCreatedBy *string `json:"dataChangeCreatedBy,omitempty"`
}

// NewNamespaceBase instantiates a new NamespaceBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespaceBase() *NamespaceBase {
	this := NamespaceBase{}
	return &this
}

// NewNamespaceBaseWithDefaults instantiates a new NamespaceBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceBaseWithDefaults() *NamespaceBase {
	this := NamespaceBase{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NamespaceBase) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceBase) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NamespaceBase) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NamespaceBase) SetName(v string) {
	o.Name = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *NamespaceBase) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceBase) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *NamespaceBase) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *NamespaceBase) SetAppId(v string) {
	o.AppId = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NamespaceBase) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceBase) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NamespaceBase) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NamespaceBase) SetComment(v string) {
	o.Comment = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *NamespaceBase) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceBase) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *NamespaceBase) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *NamespaceBase) SetFormat(v string) {
	o.Format = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *NamespaceBase) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceBase) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *NamespaceBase) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *NamespaceBase) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetDataChangeCreatedBy returns the DataChangeCreatedBy field value if set, zero value otherwise.
func (o *NamespaceBase) GetDataChangeCreatedBy() string {
	if o == nil || o.DataChangeCreatedBy == nil {
		var ret string
		return ret
	}
	return *o.DataChangeCreatedBy
}

// GetDataChangeCreatedByOk returns a tuple with the DataChangeCreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceBase) GetDataChangeCreatedByOk() (*string, bool) {
	if o == nil || o.DataChangeCreatedBy == nil {
		return nil, false
	}
	return o.DataChangeCreatedBy, true
}

// HasDataChangeCreatedBy returns a boolean if a field has been set.
func (o *NamespaceBase) HasDataChangeCreatedBy() bool {
	if o != nil && o.DataChangeCreatedBy != nil {
		return true
	}

	return false
}

// SetDataChangeCreatedBy gets a reference to the given string and assigns it to the DataChangeCreatedBy field.
func (o *NamespaceBase) SetDataChangeCreatedBy(v string) {
	o.DataChangeCreatedBy = &v
}

func (o NamespaceBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.IsPublic != nil {
		toSerialize["isPublic"] = o.IsPublic
	}
	if o.DataChangeCreatedBy != nil {
		toSerialize["dataChangeCreatedBy"] = o.DataChangeCreatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableNamespaceBase struct {
	value *NamespaceBase
	isSet bool
}

func (v NullableNamespaceBase) Get() *NamespaceBase {
	return v.value
}

func (v *NullableNamespaceBase) Set(val *NamespaceBase) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespaceBase) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespaceBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespaceBase(val *NamespaceBase) *NullableNamespaceBase {
	return &NullableNamespaceBase{value: val, isSet: true}
}

func (v NullableNamespaceBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespaceBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
