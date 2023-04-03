# \DefaultApi

All URIs are relative to *http://config-admin.test.huajiao.com/openapi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteField**](DefaultApi.md#DeleteField) | **Delete** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items/{key} | 删除配置接口
[**GetAppInfo**](DefaultApi.md#GetAppInfo) | **Get** /apps | 获取App信息
[**GetClusters**](DefaultApi.md#GetClusters) | **Get** /envs/{env}/apps/{appId}/clusters/{clusterName} | 获取集群接口
[**GetEnvClustersById**](DefaultApi.md#GetEnvClustersById) | **Get** /apps/{appId}/envclusters | 获取App的环境，集群信息
[**GetField**](DefaultApi.md#GetField) | **Get** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items/{key} | 读取配置接口
[**GetLock**](DefaultApi.md#GetLock) | **Get** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/lock | 获取某个Namespace当前编辑人接口
[**GetNamespace**](DefaultApi.md#GetNamespace) | **Get** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName} | 获取某个Namespace信息接口
[**GetNamespaces**](DefaultApi.md#GetNamespaces) | **Get** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces | 获取集群下所有Namespace信息接口
[**GetReleases**](DefaultApi.md#GetReleases) | **Get** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/releases/latest | 获取生效配置接口
[**PostClusters**](DefaultApi.md#PostClusters) | **Post** /envs/{env}/apps/{appId}/clusters | 创建集群接口
[**PostField**](DefaultApi.md#PostField) | **Post** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items | 新增配置接口
[**PostNamespace**](DefaultApi.md#PostNamespace) | **Post** /apps/{appId}/appnamespaces | 创建Namespace
[**PostReleases**](DefaultApi.md#PostReleases) | **Post** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/releases | 发布配置接口
[**PutField**](DefaultApi.md#PutField) | **Put** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items/{key} | 修改配置接口
[**Rollback**](DefaultApi.md#Rollback) | **Put** /envs/{env}/releases/{releaseId}/rollback | 回滚已发布配置接口



## DeleteField

> DeleteField(ctx, appId, env, clusterName, namespaceName, key).Operator(operator).Execute()

删除配置接口



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id
    env := "env_example" // string | 环境
    clusterName := "clusterName_example" // string | 集群名称
    namespaceName := "namespaceName_example" // string | 命名空间名称
    key := "key_example" // string | 字段名
    operator := "operator_example" // string | 删除配置的操作者，域账号

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteField(context.Background(), appId, env, clusterName, namespaceName, key).Operator(operator).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 
**env** | **string** | 环境 | 
**clusterName** | **string** | 集群名称 | 
**namespaceName** | **string** | 命名空间名称 | 
**key** | **string** | 字段名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **operator** | **string** | 删除配置的操作者，域账号 | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppInfo

> []AppInfo GetAppInfo(ctx).Execute()

获取App信息



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAppInfo(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAppInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppInfo`: []AppInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAppInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppInfoRequest struct via the builder pattern


### Return type

[**[]AppInfo**](appInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusters

> ClusterInfo GetClusters(ctx, appId, env, clusterName).Execute()

获取集群接口



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id
    env := "env_example" // string | 环境
    clusterName := "clusterName_example" // string | 集群名称

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetClusters(context.Background(), appId, env, clusterName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusters`: ClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 
**env** | **string** | 环境 | 
**clusterName** | **string** | 集群名称 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ClusterInfo**](clusterInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvClustersById

> []EnvCluster GetEnvClustersById(ctx, appId).Execute()

获取App的环境，集群信息



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetEnvClustersById(context.Background(), appId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEnvClustersById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvClustersById`: []EnvCluster
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEnvClustersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvClustersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EnvCluster**](envCluster.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetField

> Field GetField(ctx, appId, env, clusterName, namespaceName, key).Execute()

读取配置接口



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id
    env := "env_example" // string | 环境
    clusterName := "clusterName_example" // string | 集群名称
    namespaceName := "namespaceName_example" // string | 命名空间名称
    key := "key_example" // string | 字段名

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetField(context.Background(), appId, env, clusterName, namespaceName, key).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetField`: Field
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 
**env** | **string** | 环境 | 
**clusterName** | **string** | 集群名称 | 
**namespaceName** | **string** | 命名空间名称 | 
**key** | **string** | 字段名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**Field**](field.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLock

> Lock GetLock(ctx, appId, env, clusterName, namespaceName).Execute()

获取某个Namespace当前编辑人接口



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id
    env := "env_example" // string | 环境
    clusterName := "clusterName_example" // string | 集群名称
    namespaceName := "namespaceName_example" // string | 命名空间名称

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetLock(context.Background(), appId, env, clusterName, namespaceName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLock`: Lock
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 
**env** | **string** | 环境 | 
**clusterName** | **string** | 集群名称 | 
**namespaceName** | **string** | 命名空间名称 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Lock**](lock.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespace

> Namespace GetNamespace(ctx, appId, env, clusterName, namespaceName).Execute()

获取某个Namespace信息接口



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id
    env := "env_example" // string | 环境
    clusterName := "clusterName_example" // string | 集群名称
    namespaceName := "namespaceName_example" // string | 命名空间名称

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetNamespace(context.Background(), appId, env, clusterName, namespaceName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespace`: Namespace
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 
**env** | **string** | 环境 | 
**clusterName** | **string** | 集群名称 | 
**namespaceName** | **string** | 命名空间名称 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Namespace**](namespace.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaces

> []Namespace GetNamespaces(ctx, appId, env, clusterName).Execute()

获取集群下所有Namespace信息接口



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id
    env := "env_example" // string | 环境
    clusterName := "clusterName_example" // string | 集群名称

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetNamespaces(context.Background(), appId, env, clusterName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespaces`: []Namespace
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 
**env** | **string** | 环境 | 
**clusterName** | **string** | 集群名称 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]Namespace**](namespace.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleases

> Releases GetReleases(ctx, appId, env, clusterName, namespaceName).Execute()

获取生效配置接口



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id
    env := "env_example" // string | 环境
    clusterName := "clusterName_example" // string | 集群名称
    namespaceName := "namespaceName_example" // string | 命名空间名称

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetReleases(context.Background(), appId, env, clusterName, namespaceName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetReleases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleases`: Releases
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 
**env** | **string** | 环境 | 
**clusterName** | **string** | 集群名称 | 
**namespaceName** | **string** | 命名空间名称 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Releases**](releases.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClusters

> ClusterInfo PostClusters(ctx, appId, env).ClusterInfoBase(clusterInfoBase).Execute()

创建集群接口



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id
    env := "env_example" // string | 环境
    clusterInfoBase := *openapiclient.NewClusterInfoBase() // ClusterInfoBase | post携带参数

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PostClusters(context.Background(), appId, env).ClusterInfoBase(clusterInfoBase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostClusters`: ClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 
**env** | **string** | 环境 | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterInfoBase** | [**ClusterInfoBase**](ClusterInfoBase.md) | post携带参数 | 

### Return type

[**ClusterInfo**](clusterInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostField

> Field PostField(ctx, appId, env, clusterName, namespaceName).FieldBase(fieldBase).Execute()

新增配置接口



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id
    env := "env_example" // string | 环境
    clusterName := "clusterName_example" // string | 集群名称
    namespaceName := "namespaceName_example" // string | 命名空间名称
    fieldBase := *openapiclient.NewFieldBase() // FieldBase | post携带参数

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PostField(context.Background(), appId, env, clusterName, namespaceName).FieldBase(fieldBase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostField`: Field
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 
**env** | **string** | 环境 | 
**clusterName** | **string** | 集群名称 | 
**namespaceName** | **string** | 命名空间名称 | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **fieldBase** | [**FieldBase**](FieldBase.md) | post携带参数 | 

### Return type

[**Field**](field.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNamespace

> NamespaceInfo PostNamespace(ctx, appId).NamespaceBase(namespaceBase).Execute()

创建Namespace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id
    namespaceBase := *openapiclient.NewNamespaceBase() // NamespaceBase | post携带参数

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PostNamespace(context.Background(), appId).NamespaceBase(namespaceBase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostNamespace`: NamespaceInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespaceBase** | [**NamespaceBase**](NamespaceBase.md) | post携带参数 | 

### Return type

[**NamespaceInfo**](namespaceInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReleases

> Releases PostReleases(ctx, appId, env, clusterName, namespaceName).ReleasesPost(releasesPost).Execute()

发布配置接口



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id
    env := "env_example" // string | 环境
    clusterName := "clusterName_example" // string | 集群名称
    namespaceName := "namespaceName_example" // string | 命名空间名称
    releasesPost := *openapiclient.NewReleasesPost() // ReleasesPost | post携带参数

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PostReleases(context.Background(), appId, env, clusterName, namespaceName).ReleasesPost(releasesPost).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostReleases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostReleases`: Releases
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 
**env** | **string** | 环境 | 
**clusterName** | **string** | 集群名称 | 
**namespaceName** | **string** | 命名空间名称 | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **releasesPost** | [**ReleasesPost**](ReleasesPost.md) | post携带参数 | 

### Return type

[**Releases**](releases.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutField

> PutField(ctx, appId, env, clusterName, namespaceName, key).FieldPut(fieldPut).CreateIfNotExists(createIfNotExists).Execute()

修改配置接口



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 应用id
    env := "env_example" // string | 环境
    clusterName := "clusterName_example" // string | 集群名称
    namespaceName := "namespaceName_example" // string | 命名空间名称
    key := "key_example" // string | 字段名
    fieldPut := *openapiclient.NewFieldPut() // FieldPut | 携带参数
    createIfNotExists := true // bool | 当配置不存在时是否自动创建 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PutField(context.Background(), appId, env, clusterName, namespaceName, key).FieldPut(fieldPut).CreateIfNotExists(createIfNotExists).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | 应用id | 
**env** | **string** | 环境 | 
**clusterName** | **string** | 集群名称 | 
**namespaceName** | **string** | 命名空间名称 | 
**key** | **string** | 字段名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **fieldPut** | [**FieldPut**](FieldPut.md) | 携带参数 | 
 **createIfNotExists** | **bool** | 当配置不存在时是否自动创建 | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rollback

> Releases Rollback(ctx, releaseId, env).Operator(operator).Execute()

回滚已发布配置接口



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    operator := "operator_example" // string | 删除配置的操作者，域账号
    releaseId := "releaseId_example" // string | 发布记录id
    env := "env_example" // string | 环境

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Rollback(context.Background(), releaseId, env).Operator(operator).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Rollback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Rollback`: Releases
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Rollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** | 发布记录id | 
**env** | **string** | 环境 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operator** | **string** | 删除配置的操作者，域账号 | 



### Return type

[**Releases**](releases.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

