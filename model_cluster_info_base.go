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

// ClusterInfoBase struct for ClusterInfoBase
type ClusterInfoBase struct {
	Name                *string `json:"name,omitempty"`
	AppId               *string `json:"appId,omitempty"`
	DataChangeCreatedBy *string `json:"dataChangeCreatedBy,omitempty"`
}

// NewClusterInfoBase instantiates a new ClusterInfoBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterInfoBase() *ClusterInfoBase {
	this := ClusterInfoBase{}
	return &this
}

// NewClusterInfoBaseWithDefaults instantiates a new ClusterInfoBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterInfoBaseWithDefaults() *ClusterInfoBase {
	this := ClusterInfoBase{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterInfoBase) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInfoBase) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterInfoBase) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterInfoBase) SetName(v string) {
	o.Name = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ClusterInfoBase) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInfoBase) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ClusterInfoBase) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *ClusterInfoBase) SetAppId(v string) {
	o.AppId = &v
}

// GetDataChangeCreatedBy returns the DataChangeCreatedBy field value if set, zero value otherwise.
func (o *ClusterInfoBase) GetDataChangeCreatedBy() string {
	if o == nil || o.DataChangeCreatedBy == nil {
		var ret string
		return ret
	}
	return *o.DataChangeCreatedBy
}

// GetDataChangeCreatedByOk returns a tuple with the DataChangeCreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInfoBase) GetDataChangeCreatedByOk() (*string, bool) {
	if o == nil || o.DataChangeCreatedBy == nil {
		return nil, false
	}
	return o.DataChangeCreatedBy, true
}

// HasDataChangeCreatedBy returns a boolean if a field has been set.
func (o *ClusterInfoBase) HasDataChangeCreatedBy() bool {
	if o != nil && o.DataChangeCreatedBy != nil {
		return true
	}

	return false
}

// SetDataChangeCreatedBy gets a reference to the given string and assigns it to the DataChangeCreatedBy field.
func (o *ClusterInfoBase) SetDataChangeCreatedBy(v string) {
	o.DataChangeCreatedBy = &v
}

func (o ClusterInfoBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.DataChangeCreatedBy != nil {
		toSerialize["dataChangeCreatedBy"] = o.DataChangeCreatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableClusterInfoBase struct {
	value *ClusterInfoBase
	isSet bool
}

func (v NullableClusterInfoBase) Get() *ClusterInfoBase {
	return v.value
}

func (v *NullableClusterInfoBase) Set(val *ClusterInfoBase) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterInfoBase) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterInfoBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterInfoBase(val *ClusterInfoBase) *NullableClusterInfoBase {
	return &NullableClusterInfoBase{value: val, isSet: true}
}

func (v NullableClusterInfoBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterInfoBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
