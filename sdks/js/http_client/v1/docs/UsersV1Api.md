# PolyaxonSdk.UsersV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getUser**](UsersV1Api.md#getUser) | **GET** /api/v1/users | Get current user


<a name="getUser"></a>
# **getUser**
> V1User getUser()

Get current user

### Example
```javascript
var PolyaxonSdk = require('polyaxon-sdk');
var defaultClient = PolyaxonSdk.ApiClient.instance;

// Configure API key authorization: ApiKey
var ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

var apiInstance = new PolyaxonSdk.UsersV1Api();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getUser(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**V1User**](V1User.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

