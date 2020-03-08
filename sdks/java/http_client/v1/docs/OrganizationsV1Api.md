# OrganizationsV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createOrganization**](OrganizationsV1Api.md#createOrganization) | **POST** /api/v1/orgs/create | Create organization
[**createOrganizationMember**](OrganizationsV1Api.md#createOrganizationMember) | **POST** /api/v1/orgs/{owner}/members | Create organization member
[**deleteOrganization**](OrganizationsV1Api.md#deleteOrganization) | **DELETE** /api/v1/orgs/{owner} | Delete organization
[**deleteOrganizationMember**](OrganizationsV1Api.md#deleteOrganizationMember) | **DELETE** /api/v1/orgs/{owner}/members/{user} | Delete organization member details
[**getOrganization**](OrganizationsV1Api.md#getOrganization) | **GET** /api/v1/orgs/{owner} | Get organization
[**getOrganizationMember**](OrganizationsV1Api.md#getOrganizationMember) | **GET** /api/v1/orgs/{owner}/members/{user} | Get organization member details
[**listOrganizationMembers**](OrganizationsV1Api.md#listOrganizationMembers) | **GET** /api/v1/orgs/{owner}/members | Get organization members
[**listOrganizationNames**](OrganizationsV1Api.md#listOrganizationNames) | **GET** /api/v1/orgs/names | List organizations names
[**listOrganizations**](OrganizationsV1Api.md#listOrganizations) | **GET** /api/v1/orgs/list | List organizations
[**patchOrganization**](OrganizationsV1Api.md#patchOrganization) | **PATCH** /api/v1/orgs/{owner} | Patch organization
[**patchOrganizationMember**](OrganizationsV1Api.md#patchOrganizationMember) | **PATCH** /api/v1/orgs/{owner}/members/{member.user} | Patch organization member
[**updateOrganization**](OrganizationsV1Api.md#updateOrganization) | **PUT** /api/v1/orgs/{owner} | Update organization
[**updateOrganizationMember**](OrganizationsV1Api.md#updateOrganizationMember) | **PUT** /api/v1/orgs/{owner}/members/{member.user} | Update organization member


<a name="createOrganization"></a>
# **createOrganization**
> V1Organization createOrganization(body)

Create organization

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
V1Organization body = new V1Organization(); // V1Organization | 
try {
    V1Organization result = apiInstance.createOrganization(body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#createOrganization");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1Organization**](V1Organization.md)|  |

### Return type

[**V1Organization**](V1Organization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="createOrganizationMember"></a>
# **createOrganizationMember**
> V1OrganizationMember createOrganizationMember(owner, body)

Create organization member

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
V1OrganizationMember body = new V1OrganizationMember(); // V1OrganizationMember | Organization body
try {
    V1OrganizationMember result = apiInstance.createOrganizationMember(owner, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#createOrganizationMember");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **body** | [**V1OrganizationMember**](V1OrganizationMember.md)| Organization body |

### Return type

[**V1OrganizationMember**](V1OrganizationMember.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteOrganization"></a>
# **deleteOrganization**
> deleteOrganization(owner)

Delete organization

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
try {
    apiInstance.deleteOrganization(owner);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#deleteOrganization");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteOrganizationMember"></a>
# **deleteOrganizationMember**
> deleteOrganizationMember(owner, user)

Delete organization member details

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String user = "user_example"; // String | Memeber under namesapce
try {
    apiInstance.deleteOrganizationMember(owner, user);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#deleteOrganizationMember");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **user** | **String**| Memeber under namesapce |

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getOrganization"></a>
# **getOrganization**
> V1Organization getOrganization(owner)

Get organization

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
try {
    V1Organization result = apiInstance.getOrganization(owner);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#getOrganization");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |

### Return type

[**V1Organization**](V1Organization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getOrganizationMember"></a>
# **getOrganizationMember**
> V1OrganizationMember getOrganizationMember(owner, user)

Get organization member details

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String user = "user_example"; // String | Memeber under namesapce
try {
    V1OrganizationMember result = apiInstance.getOrganizationMember(owner, user);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#getOrganizationMember");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **user** | **String**| Memeber under namesapce |

### Return type

[**V1OrganizationMember**](V1OrganizationMember.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listOrganizationMembers"></a>
# **listOrganizationMembers**
> V1ListOrganizationMembersResponse listOrganizationMembers(owner, offset, limit, sort, query)

Get organization members

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
Integer offset = 56; // Integer | Pagination offset.
Integer limit = 56; // Integer | Limit size.
String sort = "sort_example"; // String | Sort to order the search.
String query = "query_example"; // String | Query filter the search search.
try {
    V1ListOrganizationMembersResponse result = apiInstance.listOrganizationMembers(owner, offset, limit, sort, query);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#listOrganizationMembers");
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

[**V1ListOrganizationMembersResponse**](V1ListOrganizationMembersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listOrganizationNames"></a>
# **listOrganizationNames**
> V1ListOrganizationsResponse listOrganizationNames()

List organizations names

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
try {
    V1ListOrganizationsResponse result = apiInstance.listOrganizationNames();
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#listOrganizationNames");
    e.printStackTrace();
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**V1ListOrganizationsResponse**](V1ListOrganizationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listOrganizations"></a>
# **listOrganizations**
> V1ListOrganizationsResponse listOrganizations()

List organizations

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
try {
    V1ListOrganizationsResponse result = apiInstance.listOrganizations();
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#listOrganizations");
    e.printStackTrace();
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**V1ListOrganizationsResponse**](V1ListOrganizationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="patchOrganization"></a>
# **patchOrganization**
> V1Organization patchOrganization(owner, body)

Patch organization

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
V1Organization body = new V1Organization(); // V1Organization | Organization body
try {
    V1Organization result = apiInstance.patchOrganization(owner, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#patchOrganization");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **body** | [**V1Organization**](V1Organization.md)| Organization body |

### Return type

[**V1Organization**](V1Organization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="patchOrganizationMember"></a>
# **patchOrganizationMember**
> V1OrganizationMember patchOrganizationMember(owner, memberUser, body)

Patch organization member

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String memberUser = "memberUser_example"; // String | User
V1OrganizationMember body = new V1OrganizationMember(); // V1OrganizationMember | Organization body
try {
    V1OrganizationMember result = apiInstance.patchOrganizationMember(owner, memberUser, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#patchOrganizationMember");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **memberUser** | **String**| User |
 **body** | [**V1OrganizationMember**](V1OrganizationMember.md)| Organization body |

### Return type

[**V1OrganizationMember**](V1OrganizationMember.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="updateOrganization"></a>
# **updateOrganization**
> V1Organization updateOrganization(owner, body)

Update organization

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
V1Organization body = new V1Organization(); // V1Organization | Organization body
try {
    V1Organization result = apiInstance.updateOrganization(owner, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#updateOrganization");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **body** | [**V1Organization**](V1Organization.md)| Organization body |

### Return type

[**V1Organization**](V1Organization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="updateOrganizationMember"></a>
# **updateOrganizationMember**
> V1OrganizationMember updateOrganizationMember(owner, memberUser, body)

Update organization member

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.OrganizationsV1Api;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: ApiKey
ApiKeyAuth ApiKey = (ApiKeyAuth) defaultClient.getAuthentication("ApiKey");
ApiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.setApiKeyPrefix("Token");

OrganizationsV1Api apiInstance = new OrganizationsV1Api();
String owner = "owner_example"; // String | Owner of the namespace
String memberUser = "memberUser_example"; // String | User
V1OrganizationMember body = new V1OrganizationMember(); // V1OrganizationMember | Organization body
try {
    V1OrganizationMember result = apiInstance.updateOrganizationMember(owner, memberUser, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling OrganizationsV1Api#updateOrganizationMember");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace |
 **memberUser** | **String**| User |
 **body** | [**V1OrganizationMember**](V1OrganizationMember.md)| Organization body |

### Return type

[**V1OrganizationMember**](V1OrganizationMember.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

