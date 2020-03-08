# HubComponentsV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createHubComponent**](HubComponentsV1Api.md#createHubComponent) | **POST** /api/v1/orgs/{owner}/components | Create hub component
[**deleteHubComponent**](HubComponentsV1Api.md#deleteHubComponent) | **DELETE** /api/v1/orgs/{owner}/components/{uuid} | Delete hub component
[**getHubComponent**](HubComponentsV1Api.md#getHubComponent) | **GET** /api/v1/orgs/{owner}/components/{uuid} | Get hub component
[**listHubComponebtNames**](HubComponentsV1Api.md#listHubComponebtNames) | **GET** /api/v1/orgs/{owner}/components/names | List hub component names
[**listHubComponents**](HubComponentsV1Api.md#listHubComponents) | **GET** /api/v1/orgs/{owner}/components | List hub components
[**patchHubComponent**](HubComponentsV1Api.md#patchHubComponent) | **PATCH** /api/v1/orgs/{owner}/components/{component.uuid} | Patch hub component
[**updateHubComponent**](HubComponentsV1Api.md#updateHubComponent) | **PUT** /api/v1/orgs/{owner}/components/{component.uuid} | Update hub component


<a name="createHubComponent"></a>
# **createHubComponent**
> V1HubComponent createHubComponent(owner, body)

Create hub component

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.HubComponentsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

HubComponentsV1Api apiInstance = new HubComponentsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
V1HubComponent body = new V1HubComponent(); // V1HubComponent | Component body
try {
    V1HubComponent result = apiInstance.createHubComponent(owner, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling HubComponentsV1Api#createHubComponent");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **body** | [**V1HubComponent**](V1HubComponent.md)| Component body |

### Return type

[**V1HubComponent**](V1HubComponent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteHubComponent"></a>
# **deleteHubComponent**
> deleteHubComponent(owner, uuid)

Delete hub component

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.HubComponentsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

HubComponentsV1Api apiInstance = new HubComponentsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String uuid = "uuid_example"; // String | Uuid identifier of the entity
try {
    apiInstance.deleteHubComponent(owner, uuid);
} catch (ApiException e) {
    System.err.println("Exception when calling HubComponentsV1Api#deleteHubComponent");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **uuid** | **String**| Uuid identifier of the entity |

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getHubComponent"></a>
# **getHubComponent**
> V1HubComponent getHubComponent(owner, uuid)

Get hub component

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.HubComponentsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

HubComponentsV1Api apiInstance = new HubComponentsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String uuid = "uuid_example"; // String | Uuid identifier of the entity
try {
    V1HubComponent result = apiInstance.getHubComponent(owner, uuid);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling HubComponentsV1Api#getHubComponent");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **uuid** | **String**| Uuid identifier of the entity |

### Return type

[**V1HubComponent**](V1HubComponent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listHubComponebtNames"></a>
# **listHubComponebtNames**
> V1ListHubComponentsResponse listHubComponebtNames(owner, offset, limit, sort, query)

List hub component names

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.HubComponentsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

HubComponentsV1Api apiInstance = new HubComponentsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
Integer offset = 56; // Integer | Pagination offset.
Integer limit = 56; // Integer | Limit size.
String sort = "sort_example"; // String | Sort to order the search.
String query = "query_example"; // String | Query filter the search search.
try {
    V1ListHubComponentsResponse result = apiInstance.listHubComponebtNames(owner, offset, limit, sort, query);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling HubComponentsV1Api#listHubComponebtNames");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **offset** | **Integer**| Pagination offset. | [optional]
 **limit** | **Integer**| Limit size. | [optional]
 **sort** | **String**| Sort to order the search. | [optional]
 **query** | **String**| Query filter the search search. | [optional]

### Return type

[**V1ListHubComponentsResponse**](V1ListHubComponentsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listHubComponents"></a>
# **listHubComponents**
> V1ListHubComponentsResponse listHubComponents(owner, offset, limit, sort, query)

List hub components

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.HubComponentsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

HubComponentsV1Api apiInstance = new HubComponentsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
Integer offset = 56; // Integer | Pagination offset.
Integer limit = 56; // Integer | Limit size.
String sort = "sort_example"; // String | Sort to order the search.
String query = "query_example"; // String | Query filter the search search.
try {
    V1ListHubComponentsResponse result = apiInstance.listHubComponents(owner, offset, limit, sort, query);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling HubComponentsV1Api#listHubComponents");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **offset** | **Integer**| Pagination offset. | [optional]
 **limit** | **Integer**| Limit size. | [optional]
 **sort** | **String**| Sort to order the search. | [optional]
 **query** | **String**| Query filter the search search. | [optional]

### Return type

[**V1ListHubComponentsResponse**](V1ListHubComponentsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="patchHubComponent"></a>
# **patchHubComponent**
> V1HubComponent patchHubComponent(owner, componentUuid, body)

Patch hub component

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.HubComponentsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

HubComponentsV1Api apiInstance = new HubComponentsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String componentUuid = "componentUuid_example"; // String | UUID
V1HubComponent body = new V1HubComponent(); // V1HubComponent | Component body
try {
    V1HubComponent result = apiInstance.patchHubComponent(owner, componentUuid, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling HubComponentsV1Api#patchHubComponent");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **componentUuid** | **String**| UUID |
 **body** | [**V1HubComponent**](V1HubComponent.md)| Component body |

### Return type

[**V1HubComponent**](V1HubComponent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="updateHubComponent"></a>
# **updateHubComponent**
> V1HubComponent updateHubComponent(owner, componentUuid, body)

Update hub component

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.HubComponentsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

HubComponentsV1Api apiInstance = new HubComponentsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String componentUuid = "componentUuid_example"; // String | UUID
V1HubComponent body = new V1HubComponent(); // V1HubComponent | Component body
try {
    V1HubComponent result = apiInstance.updateHubComponent(owner, componentUuid, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling HubComponentsV1Api#updateHubComponent");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **componentUuid** | **String**| UUID |
 **body** | [**V1HubComponent**](V1HubComponent.md)| Component body |

### Return type

[**V1HubComponent**](V1HubComponent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

