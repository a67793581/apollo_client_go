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

// EnvCluster struct for EnvCluster
type EnvCluster struct {
	Clusters *[]string `json:"clusters,omitempty"`
	Env      *string   `json:"env,omitempty"`
}

// NewEnvCluster instantiates a new EnvCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvCluster() *EnvCluster {
	this := EnvCluster{}
	return &this
}

// NewEnvClusterWithDefaults instantiates a new EnvCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvClusterWithDefaults() *EnvCluster {
	this := EnvCluster{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *EnvCluster) GetClusters() []string {
	if o == nil || o.Clusters == nil {
		var ret []string
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvCluster) GetClustersOk() (*[]string, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *EnvCluster) HasClusters() bool {
	if o != nil && o.Clusters != nil {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *EnvCluster) SetClusters(v []string) {
	o.Clusters = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *EnvCluster) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvCluster) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *EnvCluster) HasEnv() bool {
	if o != nil && o.Env != nil {
		return true
	}

	return false
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *EnvCluster) SetEnv(v string) {
	o.Env = &v
}

func (o EnvCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	return json.Marshal(toSerialize)
}

type NullableEnvCluster struct {
	value *EnvCluster
	isSet bool
}

func (v NullableEnvCluster) Get() *EnvCluster {
	return v.value
}

func (v *NullableEnvCluster) Set(val *EnvCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvCluster(val *EnvCluster) *NullableEnvCluster {
	return &NullableEnvCluster{value: val, isSet: true}
}

func (v NullableEnvCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}