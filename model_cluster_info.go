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

// ClusterInfo struct for ClusterInfo
type ClusterInfo struct {
	Name                       *string `json:"name,omitempty"`
	AppId                      *string `json:"appId,omitempty"`
	DataChangeCreatedBy        *string `json:"dataChangeCreatedBy,omitempty"`
	DataChangeLastModifiedBy   *string `json:"dataChangeLastModifiedBy,omitempty"`
	DataChangeCreatedTime      *string `json:"dataChangeCreatedTime,omitempty"`
	DataChangeLastModifiedTime *string `json:"dataChangeLastModifiedTime,omitempty"`
}

// NewClusterInfo instantiates a new ClusterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterInfo() *ClusterInfo {
	this := ClusterInfo{}
	return &this
}

// NewClusterInfoWithDefaults instantiates a new ClusterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterInfoWithDefaults() *ClusterInfo {
	this := ClusterInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterInfo) SetName(v string) {
	o.Name = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ClusterInfo) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInfo) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ClusterInfo) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *ClusterInfo) SetAppId(v string) {
	o.AppId = &v
}

// GetDataChangeCreatedBy returns the DataChangeCreatedBy field value if set, zero value otherwise.
func (o *ClusterInfo) GetDataChangeCreatedBy() string {
	if o == nil || o.DataChangeCreatedBy == nil {
		var ret string
		return ret
	}
	return *o.DataChangeCreatedBy
}

// GetDataChangeCreatedByOk returns a tuple with the DataChangeCreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInfo) GetDataChangeCreatedByOk() (*string, bool) {
	if o == nil || o.DataChangeCreatedBy == nil {
		return nil, false
	}
	return o.DataChangeCreatedBy, true
}

// HasDataChangeCreatedBy returns a boolean if a field has been set.
func (o *ClusterInfo) HasDataChangeCreatedBy() bool {
	if o != nil && o.DataChangeCreatedBy != nil {
		return true
	}

	return false
}

// SetDataChangeCreatedBy gets a reference to the given string and assigns it to the DataChangeCreatedBy field.
func (o *ClusterInfo) SetDataChangeCreatedBy(v string) {
	o.DataChangeCreatedBy = &v
}

// GetDataChangeLastModifiedBy returns the DataChangeLastModifiedBy field value if set, zero value otherwise.
func (o *ClusterInfo) GetDataChangeLastModifiedBy() string {
	if o == nil || o.DataChangeLastModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.DataChangeLastModifiedBy
}

// GetDataChangeLastModifiedByOk returns a tuple with the DataChangeLastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInfo) GetDataChangeLastModifiedByOk() (*string, bool) {
	if o == nil || o.DataChangeLastModifiedBy == nil {
		return nil, false
	}
	return o.DataChangeLastModifiedBy, true
}

// HasDataChangeLastModifiedBy returns a boolean if a field has been set.
func (o *ClusterInfo) HasDataChangeLastModifiedBy() bool {
	if o != nil && o.DataChangeLastModifiedBy != nil {
		return true
	}

	return false
}

// SetDataChangeLastModifiedBy gets a reference to the given string and assigns it to the DataChangeLastModifiedBy field.
func (o *ClusterInfo) SetDataChangeLastModifiedBy(v string) {
	o.DataChangeLastModifiedBy = &v
}

// GetDataChangeCreatedTime returns the DataChangeCreatedTime field value if set, zero value otherwise.
func (o *ClusterInfo) GetDataChangeCreatedTime() string {
	if o == nil || o.DataChangeCreatedTime == nil {
		var ret string
		return ret
	}
	return *o.DataChangeCreatedTime
}

// GetDataChangeCreatedTimeOk returns a tuple with the DataChangeCreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInfo) GetDataChangeCreatedTimeOk() (*string, bool) {
	if o == nil || o.DataChangeCreatedTime == nil {
		return nil, false
	}
	return o.DataChangeCreatedTime, true
}

// HasDataChangeCreatedTime returns a boolean if a field has been set.
func (o *ClusterInfo) HasDataChangeCreatedTime() bool {
	if o != nil && o.DataChangeCreatedTime != nil {
		return true
	}

	return false
}

// SetDataChangeCreatedTime gets a reference to the given string and assigns it to the DataChangeCreatedTime field.
func (o *ClusterInfo) SetDataChangeCreatedTime(v string) {
	o.DataChangeCreatedTime = &v
}

// GetDataChangeLastModifiedTime returns the DataChangeLastModifiedTime field value if set, zero value otherwise.
func (o *ClusterInfo) GetDataChangeLastModifiedTime() string {
	if o == nil || o.DataChangeLastModifiedTime == nil {
		var ret string
		return ret
	}
	return *o.DataChangeLastModifiedTime
}

// GetDataChangeLastModifiedTimeOk returns a tuple with the DataChangeLastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInfo) GetDataChangeLastModifiedTimeOk() (*string, bool) {
	if o == nil || o.DataChangeLastModifiedTime == nil {
		return nil, false
	}
	return o.DataChangeLastModifiedTime, true
}

// HasDataChangeLastModifiedTime returns a boolean if a field has been set.
func (o *ClusterInfo) HasDataChangeLastModifiedTime() bool {
	if o != nil && o.DataChangeLastModifiedTime != nil {
		return true
	}

	return false
}

// SetDataChangeLastModifiedTime gets a reference to the given string and assigns it to the DataChangeLastModifiedTime field.
func (o *ClusterInfo) SetDataChangeLastModifiedTime(v string) {
	o.DataChangeLastModifiedTime = &v
}

func (o ClusterInfo) MarshalJSON() ([]byte, error) {
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
	if o.DataChangeLastModifiedBy != nil {
		toSerialize["dataChangeLastModifiedBy"] = o.DataChangeLastModifiedBy
	}
	if o.DataChangeCreatedTime != nil {
		toSerialize["dataChangeCreatedTime"] = o.DataChangeCreatedTime
	}
	if o.DataChangeLastModifiedTime != nil {
		toSerialize["dataChangeLastModifiedTime"] = o.DataChangeLastModifiedTime
	}
	return json.Marshal(toSerialize)
}

type NullableClusterInfo struct {
	value *ClusterInfo
	isSet bool
}

func (v NullableClusterInfo) Get() *ClusterInfo {
	return v.value
}

func (v *NullableClusterInfo) Set(val *ClusterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterInfo(val *ClusterInfo) *NullableClusterInfo {
	return &NullableClusterInfo{value: val, isSet: true}
}

func (v NullableClusterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
