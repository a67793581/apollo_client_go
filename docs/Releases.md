# Releases

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**NamespaceName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Configurations** | Pointer to [**ReleasesConfigurations**](releases_configurations.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**DataChangeCreatedBy** | Pointer to **string** |  | [optional] 
**DataChangeLastModifiedBy** | Pointer to **string** |  | [optional] 
**DataChangeCreatedTime** | Pointer to **string** |  | [optional] 
**DataChangeLastModifiedTime** | Pointer to **string** |  | [optional] 

## Methods

### NewReleases

`func NewReleases() *Releases`

NewReleases instantiates a new Releases object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleasesWithDefaults

`func NewReleasesWithDefaults() *Releases`

NewReleasesWithDefaults instantiates a new Releases object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *Releases) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Releases) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Releases) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Releases) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetClusterName

`func (o *Releases) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Releases) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Releases) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *Releases) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetNamespaceName

`func (o *Releases) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *Releases) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *Releases) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *Releases) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.

### GetName

`func (o *Releases) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Releases) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Releases) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Releases) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfigurations

`func (o *Releases) GetConfigurations() ReleasesConfigurations`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *Releases) GetConfigurationsOk() (*ReleasesConfigurations, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *Releases) SetConfigurations(v ReleasesConfigurations)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *Releases) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetComment

`func (o *Releases) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Releases) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Releases) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Releases) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDataChangeCreatedBy

`func (o *Releases) GetDataChangeCreatedBy() string`

GetDataChangeCreatedBy returns the DataChangeCreatedBy field if non-nil, zero value otherwise.

### GetDataChangeCreatedByOk

`func (o *Releases) GetDataChangeCreatedByOk() (*string, bool)`

GetDataChangeCreatedByOk returns a tuple with the DataChangeCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChangeCreatedBy

`func (o *Releases) SetDataChangeCreatedBy(v string)`

SetDataChangeCreatedBy sets DataChangeCreatedBy field to given value.

### HasDataChangeCreatedBy

`func (o *Releases) HasDataChangeCreatedBy() bool`

HasDataChangeCreatedBy returns a boolean if a field has been set.

### GetDataChangeLastModifiedBy

`func (o *Releases) GetDataChangeLastModifiedBy() string`

GetDataChangeLastModifiedBy returns the DataChangeLastModifiedBy field if non-nil, zero value otherwise.

### GetDataChangeLastModifiedByOk

`func (o *Releases) GetDataChangeLastModifiedByOk() (*string, bool)`

GetDataChangeLastModifiedByOk returns a tuple with the DataChangeLastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChangeLastModifiedBy

`func (o *Releases) SetDataChangeLastModifiedBy(v string)`

SetDataChangeLastModifiedBy sets DataChangeLastModifiedBy field to given value.

### HasDataChangeLastModifiedBy

`func (o *Releases) HasDataChangeLastModifiedBy() bool`

HasDataChangeLastModifiedBy returns a boolean if a field has been set.

### GetDataChangeCreatedTime

`func (o *Releases) GetDataChangeCreatedTime() string`

GetDataChangeCreatedTime returns the DataChangeCreatedTime field if non-nil, zero value otherwise.

### GetDataChangeCreatedTimeOk

`func (o *Releases) GetDataChangeCreatedTimeOk() (*string, bool)`

GetDataChangeCreatedTimeOk returns a tuple with the DataChangeCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChangeCreatedTime

`func (o *Releases) SetDataChangeCreatedTime(v string)`

SetDataChangeCreatedTime sets DataChangeCreatedTime field to given value.

### HasDataChangeCreatedTime

`func (o *Releases) HasDataChangeCreatedTime() bool`

HasDataChangeCreatedTime returns a boolean if a field has been set.

### GetDataChangeLastModifiedTime

`func (o *Releases) GetDataChangeLastModifiedTime() string`

GetDataChangeLastModifiedTime returns the DataChangeLastModifiedTime field if non-nil, zero value otherwise.

### GetDataChangeLastModifiedTimeOk

`func (o *Releases) GetDataChangeLastModifiedTimeOk() (*string, bool)`

GetDataChangeLastModifiedTimeOk returns a tuple with the DataChangeLastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChangeLastModifiedTime

`func (o *Releases) SetDataChangeLastModifiedTime(v string)`

SetDataChangeLastModifiedTime sets DataChangeLastModifiedTime field to given value.

### HasDataChangeLastModifiedTime

`func (o *Releases) HasDataChangeLastModifiedTime() bool`

HasDataChangeLastModifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


