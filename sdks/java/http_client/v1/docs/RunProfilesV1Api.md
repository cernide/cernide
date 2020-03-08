# RunProfilesV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createRunProfile**](RunProfilesV1Api.md#createRunProfile) | **POST** /api/v1/orgs/{owner}/run_profiles | Create run profile
[**deleteRunProfile**](RunProfilesV1Api.md#deleteRunProfile) | **DELETE** /api/v1/orgs/{owner}/run_profiles/{uuid} | Delete run profile
[**getRunProfile**](RunProfilesV1Api.md#getRunProfile) | **GET** /api/v1/orgs/{owner}/run_profiles/{uuid} | Get run profile
[**listRunProfileNames**](RunProfilesV1Api.md#listRunProfileNames) | **GET** /api/v1/orgs/{owner}/run_profiles/names | List run profiles names
[**listRunProfiles**](RunProfilesV1Api.md#listRunProfiles) | **GET** /api/v1/orgs/{owner}/run_profiles | List run profiles
[**patchRunProfile**](RunProfilesV1Api.md#patchRunProfile) | **PATCH** /api/v1/orgs/{owner}/run_profiles/{run_profile.uuid} | Patch run profile
[**updateRunProfile**](RunProfilesV1Api.md#updateRunProfile) | **PUT** /api/v1/orgs/{owner}/run_profiles/{run_profile.uuid} | Update run profile


<a name="createRunProfile"></a>
# **createRunProfile**
> V1RunProfile createRunProfile(owner, body)

Create run profile

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.RunProfilesV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

RunProfilesV1Api apiInstance = new RunProfilesV1Api();
String owner = "owner_example"; // String | Owner of the namespace
V1RunProfile body = new V1RunProfile(); // V1RunProfile | Artifact store body
try {
    V1RunProfile result = apiInstance.createRunProfile(owner, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling RunProfilesV1Api#createRunProfile");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **body** | [**V1RunProfile**](V1RunProfile.md)| Artifact store body |

### Return type

[**V1RunProfile**](V1RunProfile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteRunProfile"></a>
# **deleteRunProfile**
> deleteRunProfile(owner, uuid)

Delete run profile

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.RunProfilesV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

RunProfilesV1Api apiInstance = new RunProfilesV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String uuid = "uuid_example"; // String | Uuid identifier of the entity
try {
    apiInstance.deleteRunProfile(owner, uuid);
} catch (ApiException e) {
    System.err.println("Exception when calling RunProfilesV1Api#deleteRunProfile");
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

<a name="getRunProfile"></a>
# **getRunProfile**
> V1RunProfile getRunProfile(owner, uuid)

Get run profile

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.RunProfilesV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

RunProfilesV1Api apiInstance = new RunProfilesV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String uuid = "uuid_example"; // String | Uuid identifier of the entity
try {
    V1RunProfile result = apiInstance.getRunProfile(owner, uuid);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling RunProfilesV1Api#getRunProfile");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **uuid** | **String**| Uuid identifier of the entity |

### Return type

[**V1RunProfile**](V1RunProfile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listRunProfileNames"></a>
# **listRunProfileNames**
> V1ListRunProfilesResponse listRunProfileNames(owner, offset, limit, sort, query)

List run profiles names

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.RunProfilesV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

RunProfilesV1Api apiInstance = new RunProfilesV1Api();
String owner = "owner_example"; // String | Owner of the namespace
Integer offset = 56; // Integer | Pagination offset.
Integer limit = 56; // Integer | Limit size.
String sort = "sort_example"; // String | Sort to order the search.
String query = "query_example"; // String | Query filter the search search.
try {
    V1ListRunProfilesResponse result = apiInstance.listRunProfileNames(owner, offset, limit, sort, query);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling RunProfilesV1Api#listRunProfileNames");
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

[**V1ListRunProfilesResponse**](V1ListRunProfilesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listRunProfiles"></a>
# **listRunProfiles**
> V1ListRunProfilesResponse listRunProfiles(owner, offset, limit, sort, query)

List run profiles

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.RunProfilesV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

RunProfilesV1Api apiInstance = new RunProfilesV1Api();
String owner = "owner_example"; // String | Owner of the namespace
Integer offset = 56; // Integer | Pagination offset.
Integer limit = 56; // Integer | Limit size.
String sort = "sort_example"; // String | Sort to order the search.
String query = "query_example"; // String | Query filter the search search.
try {
    V1ListRunProfilesResponse result = apiInstance.listRunProfiles(owner, offset, limit, sort, query);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling RunProfilesV1Api#listRunProfiles");
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

[**V1ListRunProfilesResponse**](V1ListRunProfilesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="patchRunProfile"></a>
# **patchRunProfile**
> V1RunProfile patchRunProfile(owner, runProfileUuid, body)

Patch run profile

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.RunProfilesV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

RunProfilesV1Api apiInstance = new RunProfilesV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String runProfileUuid = "runProfileUuid_example"; // String | UUID
V1RunProfile body = new V1RunProfile(); // V1RunProfile | Artifact store body
try {
    V1RunProfile result = apiInstance.patchRunProfile(owner, runProfileUuid, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling RunProfilesV1Api#patchRunProfile");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **runProfileUuid** | **String**| UUID |
 **body** | [**V1RunProfile**](V1RunProfile.md)| Artifact store body |

### Return type

[**V1RunProfile**](V1RunProfile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="updateRunProfile"></a>
# **updateRunProfile**
> V1RunProfile updateRunProfile(owner, runProfileUuid, body)

Update run profile

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.RunProfilesV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

RunProfilesV1Api apiInstance = new RunProfilesV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String runProfileUuid = "runProfileUuid_example"; // String | UUID
V1RunProfile body = new V1RunProfile(); // V1RunProfile | Artifact store body
try {
    V1RunProfile result = apiInstance.updateRunProfile(owner, runProfileUuid, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling RunProfilesV1Api#updateRunProfile");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **runProfileUuid** | **String**| UUID |
 **body** | [**V1RunProfile**](V1RunProfile.md)| Artifact store body |

### Return type

[**V1RunProfile**](V1RunProfile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

