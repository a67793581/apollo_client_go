# Namespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**NamespaceName** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Items** | Pointer to [**[]Field**](Field.md) |  | [optional] 
**DataChangeCreatedBy** | Pointer to **string** |  | [optional] 
**DataChangeLastModifiedBy** | Pointer to **string** |  | [optional] 
**DataChangeCreatedTime** | Pointer to **string** |  | [optional] 
**DataChangeLastModifiedTime** | Pointer to **string** |  | [optional] 

## Methods

### NewNamespace

`func NewNamespace() *Namespace`

NewNamespace instantiates a new Namespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceWithDefaults

`func NewNamespaceWithDefaults() *Namespace`

NewNamespaceWithDefaults instantiates a new Namespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *Namespace) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Namespace) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Namespace) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Namespace) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetClusterName

`func (o *Namespace) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Namespace) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Namespace) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *Namespace) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetNamespaceName

`func (o *Namespace) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *Namespace) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *Namespace) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *Namespace) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.

### GetComment

`func (o *Namespace) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Namespace) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Namespace) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Namespace) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetFormat

`func (o *Namespace) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Namespace) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Namespace) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Namespace) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetIsPublic

`func (o *Namespace) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Namespace) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Namespace) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *Namespace) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetItems

`func (o *Namespace) GetItems() []Field`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Namespace) GetItemsOk() (*[]Field, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Namespace) SetItems(v []Field)`

SetItems sets Items field to given value.

### HasItems

`func (o *Namespace) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetDataChangeCreatedBy

`func (o *Namespace) GetDataChangeCreatedBy() string`

GetDataChangeCreatedBy returns the DataChangeCreatedBy field if non-nil, zero value otherwise.

### GetDataChangeCreatedByOk

`func (o *Namespace) GetDataChangeCreatedByOk() (*string, bool)`

GetDataChangeCreatedByOk returns a tuple with the DataChangeCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChangeCreatedBy

`func (o *Namespace) SetDataChangeCreatedBy(v string)`

SetDataChangeCreatedBy sets DataChangeCreatedBy field to given value.

### HasDataChangeCreatedBy

`func (o *Namespace) HasDataChangeCreatedBy() bool`

HasDataChangeCreatedBy returns a boolean if a field has been set.

### GetDataChangeLastModifiedBy

`func (o *Namespace) GetDataChangeLastModifiedBy() string`

GetDataChangeLastModifiedBy returns the DataChangeLastModifiedBy field if non-nil, zero value otherwise.

### GetDataChangeLastModifiedByOk

`func (o *Namespace) GetDataChangeLastModifiedByOk() (*string, bool)`

GetDataChangeLastModifiedByOk returns a tuple with the DataChangeLastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChangeLastModifiedBy

`func (o *Namespace) SetDataChangeLastModifiedBy(v string)`

SetDataChangeLastModifiedBy sets DataChangeLastModifiedBy field to given value.

### HasDataChangeLastModifiedBy

`func (o *Namespace) HasDataChangeLastModifiedBy() bool`

HasDataChangeLastModifiedBy returns a boolean if a field has been set.

### GetDataChangeCreatedTime

`func (o *Namespace) GetDataChangeCreatedTime() string`

GetDataChangeCreatedTime returns the DataChangeCreatedTime field if non-nil, zero value otherwise.

### GetDataChangeCreatedTimeOk

`func (o *Namespace) GetDataChangeCreatedTimeOk() (*string, bool)`

GetDataChangeCreatedTimeOk returns a tuple with the DataChangeCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChangeCreatedTime

`func (o *Namespace) SetDataChangeCreatedTime(v string)`

SetDataChangeCreatedTime sets DataChangeCreatedTime field to given value.

### HasDataChangeCreatedTime

`func (o *Namespace) HasDataChangeCreatedTime() bool`

HasDataChangeCreatedTime returns a boolean if a field has been set.

### GetDataChangeLastModifiedTime

`func (o *Namespace) GetDataChangeLastModifiedTime() string`

GetDataChangeLastModifiedTime returns the DataChangeLastModifiedTime field if non-nil, zero value otherwise.

### GetDataChangeLastModifiedTimeOk

`func (o *Namespace) GetDataChangeLastModifiedTimeOk() (*string, bool)`

GetDataChangeLastModifiedTimeOk returns a tuple with the DataChangeLastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChangeLastModifiedTime

`func (o *Namespace) SetDataChangeLastModifiedTime(v string)`

SetDataChangeLastModifiedTime sets DataChangeLastModifiedTime field to given value.

### HasDataChangeLastModifiedTime

`func (o *Namespace) HasDataChangeLastModifiedTime() bool`

HasDataChangeLastModifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


