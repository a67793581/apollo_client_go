# Lock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamespaceName** | Pointer to **string** |  | [optional] 
**LockedBy** | Pointer to **string** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 

## Methods

### NewLock

`func NewLock() *Lock`

NewLock instantiates a new Lock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockWithDefaults

`func NewLockWithDefaults() *Lock`

NewLockWithDefaults instantiates a new Lock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespaceName

`func (o *Lock) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *Lock) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *Lock) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *Lock) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.

### GetLockedBy

`func (o *Lock) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *Lock) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *Lock) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *Lock) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### GetIsLocked

`func (o *Lock) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *Lock) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *Lock) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *Lock) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


