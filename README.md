# Go API client for openapi

阿波罗客户端

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v0.0.2
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://config-admin.test.huajiao.com/openapi/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**DeleteField**](docs/DefaultApi.md#deletefield) | **Delete** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items/{key} | 删除配置接口
*DefaultApi* | [**GetAppInfo**](docs/DefaultApi.md#getappinfo) | **Get** /apps | 获取App信息
*DefaultApi* | [**GetClusters**](docs/DefaultApi.md#getclusters) | **Get** /envs/{env}/apps/{appId}/clusters/{clusterName} | 获取集群接口
*DefaultApi* | [**GetEnvClustersById**](docs/DefaultApi.md#getenvclustersbyid) | **Get** /apps/{appId}/envclusters | 获取App的环境，集群信息
*DefaultApi* | [**GetField**](docs/DefaultApi.md#getfield) | **Get** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items/{key} | 读取配置接口
*DefaultApi* | [**GetLock**](docs/DefaultApi.md#getlock) | **Get** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/lock | 获取某个Namespace当前编辑人接口
*DefaultApi* | [**GetNamespace**](docs/DefaultApi.md#getnamespace) | **Get** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName} | 获取某个Namespace信息接口
*DefaultApi* | [**GetNamespaces**](docs/DefaultApi.md#getnamespaces) | **Get** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces | 获取集群下所有Namespace信息接口
*DefaultApi* | [**GetReleases**](docs/DefaultApi.md#getreleases) | **Get** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/releases/latest | 获取生效配置接口
*DefaultApi* | [**PostClusters**](docs/DefaultApi.md#postclusters) | **Post** /envs/{env}/apps/{appId}/clusters | 创建集群接口
*DefaultApi* | [**PostField**](docs/DefaultApi.md#postfield) | **Post** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items | 新增配置接口
*DefaultApi* | [**PostNamespace**](docs/DefaultApi.md#postnamespace) | **Post** /apps/{appId}/appnamespaces | 创建Namespace
*DefaultApi* | [**PostReleases**](docs/DefaultApi.md#postreleases) | **Post** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/releases | 发布配置接口
*DefaultApi* | [**PutField**](docs/DefaultApi.md#putfield) | **Put** /envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items/{key} | 修改配置接口
*DefaultApi* | [**Rollback**](docs/DefaultApi.md#rollback) | **Put** /envs/{env}/releases/{releaseId}/rollback | 回滚已发布配置接口


## Documentation For Models

 - [AppInfo](docs/AppInfo.md)
 - [ClusterInfo](docs/ClusterInfo.md)
 - [ClusterInfoBase](docs/ClusterInfoBase.md)
 - [EnvCluster](docs/EnvCluster.md)
 - [Field](docs/Field.md)
 - [FieldBase](docs/FieldBase.md)
 - [FieldPut](docs/FieldPut.md)
 - [Lock](docs/Lock.md)
 - [Namespace](docs/Namespace.md)
 - [NamespaceBase](docs/NamespaceBase.md)
 - [NamespaceInfo](docs/NamespaceInfo.md)
 - [Releases](docs/Releases.md)
 - [ReleasesConfigurations](docs/ReleasesConfigurations.md)
 - [ReleasesPost](docs/ReleasesPost.md)


## Documentation For Authorization



### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



