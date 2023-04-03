# EnvCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to **[]string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvCluster

`func NewEnvCluster() *EnvCluster`

NewEnvCluster instantiates a new EnvCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvClusterWithDefaults

`func NewEnvClusterWithDefaults() *EnvCluster`

NewEnvClusterWithDefaults instantiates a new EnvCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *EnvCluster) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *EnvCluster) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *EnvCluster) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *EnvCluster) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetEnv

`func (o *EnvCluster) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *EnvCluster) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *EnvCluster) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *EnvCluster) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


