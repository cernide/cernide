# polyaxon_sdk.RunsV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**archive_run**](RunsV1Api.md#archive_run) | **POST** /api/v1/{owner}/{project}/runs/{uuid}/archive | Archive run
[**bookmark_run**](RunsV1Api.md#bookmark_run) | **POST** /api/v1/{owner}/{project}/runs/{uuid}/bookmark | Bookmark run
[**collect_run_logs**](RunsV1Api.md#collect_run_logs) | **POST** /streams/v1/{namespace}/_internal/{owner}/{project}/runs/{uuid}/logs | Collect run logs
[**copy_run**](RunsV1Api.md#copy_run) | **POST** /api/v1/{entity.owner}/{entity.project}/runs/{entity.uuid}/copy | Restart run with copy
[**create_run**](RunsV1Api.md#create_run) | **POST** /api/v1/{owner}/{project}/runs | Create new run
[**create_run_artifacts_lineage**](RunsV1Api.md#create_run_artifacts_lineage) | **POST** /api/v1/{owner}/{project}/runs/{uuid}/artifacts_lineage | Create bulk run run artifacts lineage
[**create_run_status**](RunsV1Api.md#create_run_status) | **POST** /api/v1/{owner}/{project}/runs/{uuid}/statuses | Create new run status
[**delete_run**](RunsV1Api.md#delete_run) | **DELETE** /api/v1/{owner}/{project}/runs/{uuid} | Delete run
[**delete_run_artifact_lineage**](RunsV1Api.md#delete_run_artifact_lineage) | **DELETE** /api/v1/{owner}/{project}/runs/{uuid}/artifacts_lineage/{name} | Delete run artifact lineage
[**delete_runs**](RunsV1Api.md#delete_runs) | **DELETE** /api/v1/{owner}/{project}/runs/delete | Delete runs
[**get_multi_run_events**](RunsV1Api.md#get_multi_run_events) | **GET** /streams/v1/{namespace}/{owner}/{project}/runs/multi/events/{kind} | Get multi runs events
[**get_run**](RunsV1Api.md#get_run) | **GET** /api/v1/{owner}/{project}/runs/{uuid} | Get run
[**get_run_artifact**](RunsV1Api.md#get_run_artifact) | **GET** /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact | Get run artifact
[**get_run_artifact_lineage**](RunsV1Api.md#get_run_artifact_lineage) | **GET** /api/v1/{owner}/{project}/runs/{uuid}/artifacts_lineage/{name} | Get run artifacts lineage
[**get_run_artifacts**](RunsV1Api.md#get_run_artifacts) | **GET** /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifacts | Get run artifacts
[**get_run_artifacts_lineage**](RunsV1Api.md#get_run_artifacts_lineage) | **GET** /api/v1/{owner}/{project}/runs/{uuid}/artifacts_lineage | Get run artifacts lineage
[**get_run_artifacts_lineage_names**](RunsV1Api.md#get_run_artifacts_lineage_names) | **GET** /api/v1/{owner}/{project}/runs/{uuid}/artifacts_lineage/names | Get run artifacts lineage names
[**get_run_artifacts_tree**](RunsV1Api.md#get_run_artifacts_tree) | **GET** /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifacts/tree | Get run artifacts tree
[**get_run_events**](RunsV1Api.md#get_run_events) | **GET** /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/events/{kind} | Get run events
[**get_run_logs**](RunsV1Api.md#get_run_logs) | **GET** /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/logs | Get run logs
[**get_run_namespace**](RunsV1Api.md#get_run_namespace) | **GET** /api/v1/{owner}/{project}/runs/{uuid}/namespace | Get Run namespace
[**get_run_settings**](RunsV1Api.md#get_run_settings) | **GET** /api/v1/{owner}/{project}/runs/{uuid}/settings | Get Run settings
[**get_run_statuses**](RunsV1Api.md#get_run_statuses) | **GET** /api/v1/{owner}/{project}/runs/{uuid}/statuses | Get run status
[**impersonate_token**](RunsV1Api.md#impersonate_token) | **POST** /api/v1/{owner}/{project}/runs/{uuid}/impersonate | Impersonate run token
[**invalidate_run**](RunsV1Api.md#invalidate_run) | **POST** /api/v1/{owner}/{project}/runs/{uuid}/invalidate | Invalidate run
[**invalidate_runs**](RunsV1Api.md#invalidate_runs) | **POST** /api/v1/{owner}/{project}/runs/invalidate | Invalidate runs
[**list_archived_runs**](RunsV1Api.md#list_archived_runs) | **GET** /api/v1/archives/{user}/runs | List archived runs for user
[**list_bookmarked_runs**](RunsV1Api.md#list_bookmarked_runs) | **GET** /api/v1/bookmarks/{user}/runs | List bookmarked runs for user
[**list_runs**](RunsV1Api.md#list_runs) | **GET** /api/v1/{owner}/{project}/runs | List runs
[**list_runs_io**](RunsV1Api.md#list_runs_io) | **GET** /api/v1/{owner}/{project}/runs/io | List runs
[**notify_run_status**](RunsV1Api.md#notify_run_status) | **POST** /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/notify | Notify run status
[**patch_run**](RunsV1Api.md#patch_run) | **PATCH** /api/v1/{owner}/{project}/runs/{run.uuid} | Patch run
[**restart_run**](RunsV1Api.md#restart_run) | **POST** /api/v1/{entity.owner}/{entity.project}/runs/{entity.uuid}/restart | Restart run
[**restore_run**](RunsV1Api.md#restore_run) | **POST** /api/v1/{owner}/{project}/runs/{uuid}/restore | Restore run
[**resume_run**](RunsV1Api.md#resume_run) | **POST** /api/v1/{entity.owner}/{entity.project}/runs/{entity.uuid}/resume | Resume run
[**start_run_tensorboard**](RunsV1Api.md#start_run_tensorboard) | **POST** /api/v1/{owner}/{project}/runs/{uuid}/tensorboard/start | Start run tensorboard
[**stop_run**](RunsV1Api.md#stop_run) | **POST** /api/v1/{owner}/{project}/runs/{uuid}/stop | Stop run
[**stop_run_tensorboard**](RunsV1Api.md#stop_run_tensorboard) | **POST** /api/v1/{owner}/{project}/runs/{uuid}/tensorboard/stop | Stop run tensorboard
[**stop_runs**](RunsV1Api.md#stop_runs) | **POST** /api/v1/{owner}/{project}/runs/stop | Stop runs
[**unbookmark_run**](RunsV1Api.md#unbookmark_run) | **DELETE** /api/v1/{owner}/{project}/runs/{uuid}/unbookmark | Unbookmark run
[**update_run**](RunsV1Api.md#update_run) | **PUT** /api/v1/{owner}/{project}/runs/{run.uuid} | Update run
[**upload_run_artifact**](RunsV1Api.md#upload_run_artifact) | **POST** /api/v1/{owner}/{project}/runs/{uuid}/artifacts/upload | Upload an artifact file to a store via run access
[**upload_run_logs**](RunsV1Api.md#upload_run_logs) | **POST** /api/v1/{owner}/{project}/runs/{uuid}/logs/upload | Upload a logs file to a store via run access


# **archive_run**
> archive_run(owner, project, uuid)

Archive run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Archive run
    api_instance.archive_run(owner, project, uuid)
except ApiException as e:
    print("Exception when calling RunsV1Api->archive_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **bookmark_run**
> bookmark_run(owner, project, uuid)

Bookmark run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Bookmark run
    api_instance.bookmark_run(owner, project, uuid)
except ApiException as e:
    print("Exception when calling RunsV1Api->bookmark_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **collect_run_logs**
> collect_run_logs(namespace, owner, project, uuid)

Collect run logs

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
namespace = 'namespace_example' # str | 
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Collect run logs
    api_instance.collect_run_logs(namespace, owner, project, uuid)
except ApiException as e:
    print("Exception when calling RunsV1Api->collect_run_logs: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**|  | 
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **copy_run**
> V1Run copy_run(entity_owner, entity_project, entity_uuid, body)

Restart run with copy

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
entity_owner = 'entity_owner_example' # str | Owner of the namespace
entity_project = 'entity_project_example' # str | Project where the notification will be assigned
entity_uuid = 'entity_uuid_example' # str | Uuid identifier of the entity
body = polyaxon_sdk.V1Run() # V1Run | Run object

try:
    # Restart run with copy
    api_response = api_instance.copy_run(entity_owner, entity_project, entity_uuid, body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->copy_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity_owner** | **str**| Owner of the namespace | 
 **entity_project** | **str**| Project where the notification will be assigned | 
 **entity_uuid** | **str**| Uuid identifier of the entity | 
 **body** | [**V1Run**](V1Run.md)| Run object | 

### Return type

[**V1Run**](V1Run.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_run**
> V1Run create_run(owner, project, body)

Create new run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
body = polyaxon_sdk.V1OperationBody() # V1OperationBody | operation object

try:
    # Create new run
    api_response = api_instance.create_run(owner, project, body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->create_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **body** | [**V1OperationBody**](V1OperationBody.md)| operation object | 

### Return type

[**V1Run**](V1Run.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_run_artifacts_lineage**
> create_run_artifacts_lineage(owner, project, uuid, body)

Create bulk run run artifacts lineage

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity
body = polyaxon_sdk.V1RunArtifacts() # V1RunArtifacts | Run Artifacts

try:
    # Create bulk run run artifacts lineage
    api_instance.create_run_artifacts_lineage(owner, project, uuid, body)
except ApiException as e:
    print("Exception when calling RunsV1Api->create_run_artifacts_lineage: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 
 **body** | [**V1RunArtifacts**](V1RunArtifacts.md)| Run Artifacts | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_run_status**
> V1Status create_run_status(owner, project, uuid, body)

Create new run status

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity
body = polyaxon_sdk.V1EntityStatusBodyRequest() # V1EntityStatusBodyRequest | 

try:
    # Create new run status
    api_response = api_instance.create_run_status(owner, project, uuid, body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->create_run_status: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 
 **body** | [**V1EntityStatusBodyRequest**](V1EntityStatusBodyRequest.md)|  | 

### Return type

[**V1Status**](V1Status.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_run**
> delete_run(owner, project, uuid)

Delete run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Delete run
    api_instance.delete_run(owner, project, uuid)
except ApiException as e:
    print("Exception when calling RunsV1Api->delete_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_run_artifact_lineage**
> delete_run_artifact_lineage(owner, project, uuid, name, namespace=namespace, kind=kind, names=names, runs=runs, orient=orient, path=path)

Delete run artifact lineage

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity
name = 'name_example' # str | Artifact name
namespace = 'namespace_example' # str | namespace. (optional)
kind = 'model' # str | The artifact kind.   - model: model  - audio: audio  - video: vidio  - histogram: histogram  - image: image  - tensor: tensor  - dataframe: dataframe  - chart: plotly/bokeh chart  - csv: Comma  - tsv: Tab  - psv: Pipe  - ssv: Space  - metric: Metric  - env: Env  - html: HTML  - text: Text  - file: File  - dir: Dir  - dockerfile: Dockerfile  - docker_image: docker image  - data: data  - coderef: coderef  - table: table (optional) (default to model)
names = 'names_example' # str | Names query param. (optional)
runs = 'runs_example' # str | Runs query param. (optional)
orient = 'orient_example' # str | Orient query param. (optional)
path = 'path_example' # str | Path query param. (optional)

try:
    # Delete run artifact lineage
    api_instance.delete_run_artifact_lineage(owner, project, uuid, name, namespace=namespace, kind=kind, names=names, runs=runs, orient=orient, path=path)
except ApiException as e:
    print("Exception when calling RunsV1Api->delete_run_artifact_lineage: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 
 **name** | **str**| Artifact name | 
 **namespace** | **str**| namespace. | [optional] 
 **kind** | **str**| The artifact kind.   - model: model  - audio: audio  - video: vidio  - histogram: histogram  - image: image  - tensor: tensor  - dataframe: dataframe  - chart: plotly/bokeh chart  - csv: Comma  - tsv: Tab  - psv: Pipe  - ssv: Space  - metric: Metric  - env: Env  - html: HTML  - text: Text  - file: File  - dir: Dir  - dockerfile: Dockerfile  - docker_image: docker image  - data: data  - coderef: coderef  - table: table | [optional] [default to model]
 **names** | **str**| Names query param. | [optional] 
 **runs** | **str**| Runs query param. | [optional] 
 **orient** | **str**| Orient query param. | [optional] 
 **path** | **str**| Path query param. | [optional] 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_runs**
> delete_runs(owner, project, body)

Delete runs

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project under namesapce
body = polyaxon_sdk.V1Uuids() # V1Uuids | Uuids of the entities

try:
    # Delete runs
    api_instance.delete_runs(owner, project, body)
except ApiException as e:
    print("Exception when calling RunsV1Api->delete_runs: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project under namesapce | 
 **body** | [**V1Uuids**](V1Uuids.md)| Uuids of the entities | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_multi_run_events**
> V1EventsResponse get_multi_run_events(namespace, owner, project, kind, uuid=uuid, name=name, names=names, runs=runs, orient=orient, path=path)

Get multi runs events

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
namespace = 'namespace_example' # str | namespace
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
kind = 'kind_example' # str | The artifact kind
uuid = 'uuid_example' # str | Uuid identifier of the entity. (optional)
name = 'name_example' # str | Artifact name. (optional)
names = 'names_example' # str | Names query param. (optional)
runs = 'runs_example' # str | Runs query param. (optional)
orient = 'orient_example' # str | Orient query param. (optional)
path = 'path_example' # str | Path query param. (optional)

try:
    # Get multi runs events
    api_response = api_instance.get_multi_run_events(namespace, owner, project, kind, uuid=uuid, name=name, names=names, runs=runs, orient=orient, path=path)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_multi_run_events: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| namespace | 
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **kind** | **str**| The artifact kind | 
 **uuid** | **str**| Uuid identifier of the entity. | [optional] 
 **name** | **str**| Artifact name. | [optional] 
 **names** | **str**| Names query param. | [optional] 
 **runs** | **str**| Runs query param. | [optional] 
 **orient** | **str**| Orient query param. | [optional] 
 **path** | **str**| Path query param. | [optional] 

### Return type

[**V1EventsResponse**](V1EventsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_run**
> V1Run get_run(owner, project, uuid)

Get run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Get run
    api_response = api_instance.get_run(owner, project, uuid)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

[**V1Run**](V1Run.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_run_artifact**
> str get_run_artifact(namespace, owner, project, uuid, path=path, stream=stream)

Get run artifact

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
namespace = 'namespace_example' # str | namespace
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the experiement will be assigned
uuid = 'uuid_example' # str | Unique integer identifier of the entity
path = 'path_example' # str | Artifact filepath. (optional)
stream = true # bool | Whether to stream the file. (optional)

try:
    # Get run artifact
    api_response = api_instance.get_run_artifact(namespace, owner, project, uuid, path=path, stream=stream)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_run_artifact: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| namespace | 
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the experiement will be assigned | 
 **uuid** | **str**| Unique integer identifier of the entity | 
 **path** | **str**| Artifact filepath. | [optional] 
 **stream** | **bool**| Whether to stream the file. | [optional] 

### Return type

**str**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_run_artifact_lineage**
> V1RunArtifact get_run_artifact_lineage(owner, project, uuid, name, namespace=namespace, kind=kind, names=names, runs=runs, orient=orient, path=path)

Get run artifacts lineage

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity
name = 'name_example' # str | Artifact name
namespace = 'namespace_example' # str | namespace. (optional)
kind = 'model' # str | The artifact kind.   - model: model  - audio: audio  - video: vidio  - histogram: histogram  - image: image  - tensor: tensor  - dataframe: dataframe  - chart: plotly/bokeh chart  - csv: Comma  - tsv: Tab  - psv: Pipe  - ssv: Space  - metric: Metric  - env: Env  - html: HTML  - text: Text  - file: File  - dir: Dir  - dockerfile: Dockerfile  - docker_image: docker image  - data: data  - coderef: coderef  - table: table (optional) (default to model)
names = 'names_example' # str | Names query param. (optional)
runs = 'runs_example' # str | Runs query param. (optional)
orient = 'orient_example' # str | Orient query param. (optional)
path = 'path_example' # str | Path query param. (optional)

try:
    # Get run artifacts lineage
    api_response = api_instance.get_run_artifact_lineage(owner, project, uuid, name, namespace=namespace, kind=kind, names=names, runs=runs, orient=orient, path=path)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_run_artifact_lineage: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 
 **name** | **str**| Artifact name | 
 **namespace** | **str**| namespace. | [optional] 
 **kind** | **str**| The artifact kind.   - model: model  - audio: audio  - video: vidio  - histogram: histogram  - image: image  - tensor: tensor  - dataframe: dataframe  - chart: plotly/bokeh chart  - csv: Comma  - tsv: Tab  - psv: Pipe  - ssv: Space  - metric: Metric  - env: Env  - html: HTML  - text: Text  - file: File  - dir: Dir  - dockerfile: Dockerfile  - docker_image: docker image  - data: data  - coderef: coderef  - table: table | [optional] [default to model]
 **names** | **str**| Names query param. | [optional] 
 **runs** | **str**| Runs query param. | [optional] 
 **orient** | **str**| Orient query param. | [optional] 
 **path** | **str**| Path query param. | [optional] 

### Return type

[**V1RunArtifact**](V1RunArtifact.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_run_artifacts**
> str get_run_artifacts(namespace, owner, project, uuid, path=path)

Get run artifacts

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
namespace = 'namespace_example' # str | namespace
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the experiement will be assigned
uuid = 'uuid_example' # str | Unique integer identifier of the entity
path = 'path_example' # str | Artifact filepath. (optional)

try:
    # Get run artifacts
    api_response = api_instance.get_run_artifacts(namespace, owner, project, uuid, path=path)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_run_artifacts: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| namespace | 
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the experiement will be assigned | 
 **uuid** | **str**| Unique integer identifier of the entity | 
 **path** | **str**| Artifact filepath. | [optional] 

### Return type

**str**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_run_artifacts_lineage**
> V1ListRunArtifactsResponse get_run_artifacts_lineage(owner, project, uuid, limit=limit, sort=sort, query=query)

Get run artifacts lineage

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity
limit = 56 # int | Limit size. (optional)
sort = 'sort_example' # str | Sort to order the search. (optional)
query = 'query_example' # str | Query filter the search search. (optional)

try:
    # Get run artifacts lineage
    api_response = api_instance.get_run_artifacts_lineage(owner, project, uuid, limit=limit, sort=sort, query=query)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_run_artifacts_lineage: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 
 **limit** | **int**| Limit size. | [optional] 
 **sort** | **str**| Sort to order the search. | [optional] 
 **query** | **str**| Query filter the search search. | [optional] 

### Return type

[**V1ListRunArtifactsResponse**](V1ListRunArtifactsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_run_artifacts_lineage_names**
> V1ListRunArtifactsResponse get_run_artifacts_lineage_names(owner, project, uuid, limit=limit, sort=sort, query=query)

Get run artifacts lineage names

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity
limit = 56 # int | Limit size. (optional)
sort = 'sort_example' # str | Sort to order the search. (optional)
query = 'query_example' # str | Query filter the search search. (optional)

try:
    # Get run artifacts lineage names
    api_response = api_instance.get_run_artifacts_lineage_names(owner, project, uuid, limit=limit, sort=sort, query=query)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_run_artifacts_lineage_names: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 
 **limit** | **int**| Limit size. | [optional] 
 **sort** | **str**| Sort to order the search. | [optional] 
 **query** | **str**| Query filter the search search. | [optional] 

### Return type

[**V1ListRunArtifactsResponse**](V1ListRunArtifactsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_run_artifacts_tree**
> V1ArtifactTreeResponse get_run_artifacts_tree(namespace, owner, project, uuid, name=name, kind=kind, names=names, runs=runs, orient=orient, path=path)

Get run artifacts tree

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
namespace = 'namespace_example' # str | namespace
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity
name = 'name_example' # str | Artifact name. (optional)
kind = 'model' # str | The artifact kind.   - model: model  - audio: audio  - video: vidio  - histogram: histogram  - image: image  - tensor: tensor  - dataframe: dataframe  - chart: plotly/bokeh chart  - csv: Comma  - tsv: Tab  - psv: Pipe  - ssv: Space  - metric: Metric  - env: Env  - html: HTML  - text: Text  - file: File  - dir: Dir  - dockerfile: Dockerfile  - docker_image: docker image  - data: data  - coderef: coderef  - table: table (optional) (default to model)
names = 'names_example' # str | Names query param. (optional)
runs = 'runs_example' # str | Runs query param. (optional)
orient = 'orient_example' # str | Orient query param. (optional)
path = 'path_example' # str | Path query param. (optional)

try:
    # Get run artifacts tree
    api_response = api_instance.get_run_artifacts_tree(namespace, owner, project, uuid, name=name, kind=kind, names=names, runs=runs, orient=orient, path=path)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_run_artifacts_tree: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| namespace | 
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 
 **name** | **str**| Artifact name. | [optional] 
 **kind** | **str**| The artifact kind.   - model: model  - audio: audio  - video: vidio  - histogram: histogram  - image: image  - tensor: tensor  - dataframe: dataframe  - chart: plotly/bokeh chart  - csv: Comma  - tsv: Tab  - psv: Pipe  - ssv: Space  - metric: Metric  - env: Env  - html: HTML  - text: Text  - file: File  - dir: Dir  - dockerfile: Dockerfile  - docker_image: docker image  - data: data  - coderef: coderef  - table: table | [optional] [default to model]
 **names** | **str**| Names query param. | [optional] 
 **runs** | **str**| Runs query param. | [optional] 
 **orient** | **str**| Orient query param. | [optional] 
 **path** | **str**| Path query param. | [optional] 

### Return type

[**V1ArtifactTreeResponse**](V1ArtifactTreeResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_run_events**
> V1EventsResponse get_run_events(namespace, owner, project, uuid, kind, name=name, names=names, runs=runs, orient=orient, path=path)

Get run events

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
namespace = 'namespace_example' # str | namespace
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity
kind = 'kind_example' # str | The artifact kind
name = 'name_example' # str | Artifact name. (optional)
names = 'names_example' # str | Names query param. (optional)
runs = 'runs_example' # str | Runs query param. (optional)
orient = 'orient_example' # str | Orient query param. (optional)
path = 'path_example' # str | Path query param. (optional)

try:
    # Get run events
    api_response = api_instance.get_run_events(namespace, owner, project, uuid, kind, name=name, names=names, runs=runs, orient=orient, path=path)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_run_events: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| namespace | 
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 
 **kind** | **str**| The artifact kind | 
 **name** | **str**| Artifact name. | [optional] 
 **names** | **str**| Names query param. | [optional] 
 **runs** | **str**| Runs query param. | [optional] 
 **orient** | **str**| Orient query param. | [optional] 
 **path** | **str**| Path query param. | [optional] 

### Return type

[**V1EventsResponse**](V1EventsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_run_logs**
> V1Logs get_run_logs(namespace, owner, project, uuid, last_time=last_time, last_file=last_file)

Get run logs

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
namespace = 'namespace_example' # str | 
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity
last_time = '2013-10-20T19:20:30+01:00' # datetime | last time. (optional)
last_file = 'last_file_example' # str | last file. (optional)

try:
    # Get run logs
    api_response = api_instance.get_run_logs(namespace, owner, project, uuid, last_time=last_time, last_file=last_file)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_run_logs: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**|  | 
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 
 **last_time** | **datetime**| last time. | [optional] 
 **last_file** | **str**| last file. | [optional] 

### Return type

[**V1Logs**](V1Logs.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_run_namespace**
> V1RunSettings get_run_namespace(owner, project, uuid)

Get Run namespace

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Get Run namespace
    api_response = api_instance.get_run_namespace(owner, project, uuid)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_run_namespace: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

[**V1RunSettings**](V1RunSettings.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_run_settings**
> V1RunSettings get_run_settings(owner, project, uuid)

Get Run settings

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Get Run settings
    api_response = api_instance.get_run_settings(owner, project, uuid)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_run_settings: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

[**V1RunSettings**](V1RunSettings.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_run_statuses**
> V1Status get_run_statuses(owner, project, uuid)

Get run status

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Get run status
    api_response = api_instance.get_run_statuses(owner, project, uuid)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->get_run_statuses: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

[**V1Status**](V1Status.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **impersonate_token**
> V1Auth impersonate_token(owner, project, uuid)

Impersonate run token

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Impersonate run token
    api_response = api_instance.impersonate_token(owner, project, uuid)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->impersonate_token: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

[**V1Auth**](V1Auth.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **invalidate_run**
> invalidate_run(owner, project, uuid, body)

Invalidate run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity
body = polyaxon_sdk.V1ProjectEntityResourceRequest() # V1ProjectEntityResourceRequest | 

try:
    # Invalidate run
    api_instance.invalidate_run(owner, project, uuid, body)
except ApiException as e:
    print("Exception when calling RunsV1Api->invalidate_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 
 **body** | [**V1ProjectEntityResourceRequest**](V1ProjectEntityResourceRequest.md)|  | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **invalidate_runs**
> invalidate_runs(owner, project, body)

Invalidate runs

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project under namesapce
body = polyaxon_sdk.V1Uuids() # V1Uuids | Uuids of the entities

try:
    # Invalidate runs
    api_instance.invalidate_runs(owner, project, body)
except ApiException as e:
    print("Exception when calling RunsV1Api->invalidate_runs: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project under namesapce | 
 **body** | [**V1Uuids**](V1Uuids.md)| Uuids of the entities | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_archived_runs**
> V1ListRunsResponse list_archived_runs(user, offset=offset, limit=limit, sort=sort, query=query)

List archived runs for user

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
user = 'user_example' # str | User
offset = 56 # int | Pagination offset. (optional)
limit = 56 # int | Limit size. (optional)
sort = 'sort_example' # str | Sort to order the search. (optional)
query = 'query_example' # str | Query filter the search search. (optional)

try:
    # List archived runs for user
    api_response = api_instance.list_archived_runs(user, offset=offset, limit=limit, sort=sort, query=query)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->list_archived_runs: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **str**| User | 
 **offset** | **int**| Pagination offset. | [optional] 
 **limit** | **int**| Limit size. | [optional] 
 **sort** | **str**| Sort to order the search. | [optional] 
 **query** | **str**| Query filter the search search. | [optional] 

### Return type

[**V1ListRunsResponse**](V1ListRunsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_bookmarked_runs**
> V1ListRunsResponse list_bookmarked_runs(user, offset=offset, limit=limit, sort=sort, query=query)

List bookmarked runs for user

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
user = 'user_example' # str | User
offset = 56 # int | Pagination offset. (optional)
limit = 56 # int | Limit size. (optional)
sort = 'sort_example' # str | Sort to order the search. (optional)
query = 'query_example' # str | Query filter the search search. (optional)

try:
    # List bookmarked runs for user
    api_response = api_instance.list_bookmarked_runs(user, offset=offset, limit=limit, sort=sort, query=query)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->list_bookmarked_runs: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **str**| User | 
 **offset** | **int**| Pagination offset. | [optional] 
 **limit** | **int**| Limit size. | [optional] 
 **sort** | **str**| Sort to order the search. | [optional] 
 **query** | **str**| Query filter the search search. | [optional] 

### Return type

[**V1ListRunsResponse**](V1ListRunsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_runs**
> V1ListRunsResponse list_runs(owner, project, offset=offset, limit=limit, sort=sort, query=query)

List runs

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project under namesapce
offset = 56 # int | Pagination offset. (optional)
limit = 56 # int | Limit size. (optional)
sort = 'sort_example' # str | Sort to order the search. (optional)
query = 'query_example' # str | Query filter the search search. (optional)

try:
    # List runs
    api_response = api_instance.list_runs(owner, project, offset=offset, limit=limit, sort=sort, query=query)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->list_runs: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project under namesapce | 
 **offset** | **int**| Pagination offset. | [optional] 
 **limit** | **int**| Limit size. | [optional] 
 **sort** | **str**| Sort to order the search. | [optional] 
 **query** | **str**| Query filter the search search. | [optional] 

### Return type

[**V1ListRunsResponse**](V1ListRunsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_runs_io**
> V1ListRunsResponse list_runs_io(owner, project, offset=offset, limit=limit, sort=sort, query=query)

List runs

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project under namesapce
offset = 56 # int | Pagination offset. (optional)
limit = 56 # int | Limit size. (optional)
sort = 'sort_example' # str | Sort to order the search. (optional)
query = 'query_example' # str | Query filter the search search. (optional)

try:
    # List runs
    api_response = api_instance.list_runs_io(owner, project, offset=offset, limit=limit, sort=sort, query=query)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->list_runs_io: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project under namesapce | 
 **offset** | **int**| Pagination offset. | [optional] 
 **limit** | **int**| Limit size. | [optional] 
 **sort** | **str**| Sort to order the search. | [optional] 
 **query** | **str**| Query filter the search search. | [optional] 

### Return type

[**V1ListRunsResponse**](V1ListRunsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **notify_run_status**
> notify_run_status(namespace, owner, project, uuid, body)

Notify run status

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
namespace = 'namespace_example' # str | Na,espace
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity
body = polyaxon_sdk.V1EntityNotificationBody() # V1EntityNotificationBody | 

try:
    # Notify run status
    api_instance.notify_run_status(namespace, owner, project, uuid, body)
except ApiException as e:
    print("Exception when calling RunsV1Api->notify_run_status: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| Na,espace | 
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 
 **body** | [**V1EntityNotificationBody**](V1EntityNotificationBody.md)|  | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **patch_run**
> V1Run patch_run(owner, project, run_uuid, body)

Patch run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
run_uuid = 'run_uuid_example' # str | UUID
body = polyaxon_sdk.V1Run() # V1Run | Run object

try:
    # Patch run
    api_response = api_instance.patch_run(owner, project, run_uuid, body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->patch_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **run_uuid** | **str**| UUID | 
 **body** | [**V1Run**](V1Run.md)| Run object | 

### Return type

[**V1Run**](V1Run.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **restart_run**
> V1Run restart_run(entity_owner, entity_project, entity_uuid, body)

Restart run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
entity_owner = 'entity_owner_example' # str | Owner of the namespace
entity_project = 'entity_project_example' # str | Project where the notification will be assigned
entity_uuid = 'entity_uuid_example' # str | Uuid identifier of the entity
body = polyaxon_sdk.V1Run() # V1Run | Run object

try:
    # Restart run
    api_response = api_instance.restart_run(entity_owner, entity_project, entity_uuid, body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->restart_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity_owner** | **str**| Owner of the namespace | 
 **entity_project** | **str**| Project where the notification will be assigned | 
 **entity_uuid** | **str**| Uuid identifier of the entity | 
 **body** | [**V1Run**](V1Run.md)| Run object | 

### Return type

[**V1Run**](V1Run.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **restore_run**
> restore_run(owner, project, uuid)

Restore run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Restore run
    api_instance.restore_run(owner, project, uuid)
except ApiException as e:
    print("Exception when calling RunsV1Api->restore_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **resume_run**
> V1Run resume_run(entity_owner, entity_project, entity_uuid, body)

Resume run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
entity_owner = 'entity_owner_example' # str | Owner of the namespace
entity_project = 'entity_project_example' # str | Project where the notification will be assigned
entity_uuid = 'entity_uuid_example' # str | Uuid identifier of the entity
body = polyaxon_sdk.V1Run() # V1Run | Run object

try:
    # Resume run
    api_response = api_instance.resume_run(entity_owner, entity_project, entity_uuid, body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->resume_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity_owner** | **str**| Owner of the namespace | 
 **entity_project** | **str**| Project where the notification will be assigned | 
 **entity_uuid** | **str**| Uuid identifier of the entity | 
 **body** | [**V1Run**](V1Run.md)| Run object | 

### Return type

[**V1Run**](V1Run.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **start_run_tensorboard**
> start_run_tensorboard(owner, project, uuid, body)

Start run tensorboard

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity
body = polyaxon_sdk.V1ProjectEntityResourceRequest() # V1ProjectEntityResourceRequest | 

try:
    # Start run tensorboard
    api_instance.start_run_tensorboard(owner, project, uuid, body)
except ApiException as e:
    print("Exception when calling RunsV1Api->start_run_tensorboard: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 
 **body** | [**V1ProjectEntityResourceRequest**](V1ProjectEntityResourceRequest.md)|  | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **stop_run**
> stop_run(owner, project, uuid)

Stop run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Stop run
    api_instance.stop_run(owner, project, uuid)
except ApiException as e:
    print("Exception when calling RunsV1Api->stop_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **stop_run_tensorboard**
> stop_run_tensorboard(owner, project, uuid)

Stop run tensorboard

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Stop run tensorboard
    api_instance.stop_run_tensorboard(owner, project, uuid)
except ApiException as e:
    print("Exception when calling RunsV1Api->stop_run_tensorboard: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **stop_runs**
> stop_runs(owner, project, body)

Stop runs

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project under namesapce
body = polyaxon_sdk.V1Uuids() # V1Uuids | Uuids of the entities

try:
    # Stop runs
    api_instance.stop_runs(owner, project, body)
except ApiException as e:
    print("Exception when calling RunsV1Api->stop_runs: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project under namesapce | 
 **body** | [**V1Uuids**](V1Uuids.md)| Uuids of the entities | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **unbookmark_run**
> unbookmark_run(owner, project, uuid)

Unbookmark run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the notification will be assigned
uuid = 'uuid_example' # str | Uuid identifier of the entity

try:
    # Unbookmark run
    api_instance.unbookmark_run(owner, project, uuid)
except ApiException as e:
    print("Exception when calling RunsV1Api->unbookmark_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the notification will be assigned | 
 **uuid** | **str**| Uuid identifier of the entity | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_run**
> V1Run update_run(owner, project, run_uuid, body)

Update run

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project where the run will be assigned
run_uuid = 'run_uuid_example' # str | UUID
body = polyaxon_sdk.V1Run() # V1Run | Run object

try:
    # Update run
    api_response = api_instance.update_run(owner, project, run_uuid, body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RunsV1Api->update_run: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project where the run will be assigned | 
 **run_uuid** | **str**| UUID | 
 **body** | [**V1Run**](V1Run.md)| Run object | 

### Return type

[**V1Run**](V1Run.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **upload_run_artifact**
> upload_run_artifact(owner, project, uuid, uploadfile, path=path, overwrite=overwrite)

Upload an artifact file to a store via run access

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project having access to the store
uuid = 'uuid_example' # str | Unique integer identifier of the entity
uploadfile = '/path/to/file.txt' # file | The file to upload.
path = 'path_example' # str | File path query params. (optional)
overwrite = true # bool | File path query params. (optional)

try:
    # Upload an artifact file to a store via run access
    api_instance.upload_run_artifact(owner, project, uuid, uploadfile, path=path, overwrite=overwrite)
except ApiException as e:
    print("Exception when calling RunsV1Api->upload_run_artifact: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project having access to the store | 
 **uuid** | **str**| Unique integer identifier of the entity | 
 **uploadfile** | **file**| The file to upload. | 
 **path** | **str**| File path query params. | [optional] 
 **overwrite** | **bool**| File path query params. | [optional] 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **upload_run_logs**
> upload_run_logs(owner, project, uuid, uploadfile, path=path, overwrite=overwrite)

Upload a logs file to a store via run access

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.RunsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
project = 'project_example' # str | Project having access to the store
uuid = 'uuid_example' # str | Unique integer identifier of the entity
uploadfile = '/path/to/file.txt' # file | The file to upload.
path = 'path_example' # str | File path query params. (optional)
overwrite = true # bool | File path query params. (optional)

try:
    # Upload a logs file to a store via run access
    api_instance.upload_run_logs(owner, project, uuid, uploadfile, path=path, overwrite=overwrite)
except ApiException as e:
    print("Exception when calling RunsV1Api->upload_run_logs: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **project** | **str**| Project having access to the store | 
 **uuid** | **str**| Unique integer identifier of the entity | 
 **uploadfile** | **file**| The file to upload. | 
 **path** | **str**| File path query params. | [optional] 
 **overwrite** | **bool**| File path query params. | [optional] 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

