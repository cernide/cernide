// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.3.2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import RuntimeError from '../model/RuntimeError';
import V1ListModelRegistryResponse from '../model/V1ListModelRegistryResponse';
import V1ModelRegistry from '../model/V1ModelRegistry';

/**
* ModelRegistryV1 service.
* @module api/ModelRegistryV1Api
* @version 1.3.2
*/
export default class ModelRegistryV1Api {

    /**
    * Constructs a new ModelRegistryV1Api. 
    * Polyaxon sdk
    * @alias module:api/ModelRegistryV1Api
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the createModelRegistry operation.
     * @callback module:api/ModelRegistryV1Api~createModelRegistryCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ModelRegistry} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create hub model
     * @param {String} owner Owner of the namespace
     * @param {module:model/V1ModelRegistry} body Model body
     * @param {module:api/ModelRegistryV1Api~createModelRegistryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ModelRegistry}
     */
    createModelRegistry(owner, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling createModelRegistry");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createModelRegistry");
      }

      let pathParams = {
        'owner': owner
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1ModelRegistry;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteModelRegistry operation.
     * @callback module:api/ModelRegistryV1Api~deleteModelRegistryCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete hub model
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/ModelRegistryV1Api~deleteModelRegistryCallback} callback The callback function, accepting three arguments: error, data, response
     */
    deleteModelRegistry(owner, uuid, callback) {
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling deleteModelRegistry");
      }
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling deleteModelRegistry");
      }

      let pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/{uuid}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getModelRegistry operation.
     * @callback module:api/ModelRegistryV1Api~getModelRegistryCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ModelRegistry} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get hub model
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/ModelRegistryV1Api~getModelRegistryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ModelRegistry}
     */
    getModelRegistry(owner, uuid, callback) {
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getModelRegistry");
      }
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling getModelRegistry");
      }

      let pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1ModelRegistry;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/{uuid}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listModelRegistry operation.
     * @callback module:api/ModelRegistryV1Api~listModelRegistryCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListModelRegistryResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List hub models
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/ModelRegistryV1Api~listModelRegistryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListModelRegistryResponse}
     */
    listModelRegistry(owner, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listModelRegistry");
      }

      let pathParams = {
        'owner': owner
      };
      let queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1ListModelRegistryResponse;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listModelRegistryNames operation.
     * @callback module:api/ModelRegistryV1Api~listModelRegistryNamesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListModelRegistryResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List hub model names
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/ModelRegistryV1Api~listModelRegistryNamesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListModelRegistryResponse}
     */
    listModelRegistryNames(owner, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listModelRegistryNames");
      }

      let pathParams = {
        'owner': owner
      };
      let queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1ListModelRegistryResponse;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/names', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the patchModelRegistry operation.
     * @callback module:api/ModelRegistryV1Api~patchModelRegistryCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ModelRegistry} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch hub model
     * @param {String} owner Owner of the namespace
     * @param {String} model_uuid UUID
     * @param {module:model/V1ModelRegistry} body Model body
     * @param {module:api/ModelRegistryV1Api~patchModelRegistryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ModelRegistry}
     */
    patchModelRegistry(owner, model_uuid, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling patchModelRegistry");
      }
      // verify the required parameter 'model_uuid' is set
      if (model_uuid === undefined || model_uuid === null) {
        throw new Error("Missing the required parameter 'model_uuid' when calling patchModelRegistry");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling patchModelRegistry");
      }

      let pathParams = {
        'owner': owner,
        'model.uuid': model_uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1ModelRegistry;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/{model.uuid}', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateModelRegistry operation.
     * @callback module:api/ModelRegistryV1Api~updateModelRegistryCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ModelRegistry} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update hub model
     * @param {String} owner Owner of the namespace
     * @param {String} model_uuid UUID
     * @param {module:model/V1ModelRegistry} body Model body
     * @param {module:api/ModelRegistryV1Api~updateModelRegistryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ModelRegistry}
     */
    updateModelRegistry(owner, model_uuid, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling updateModelRegistry");
      }
      // verify the required parameter 'model_uuid' is set
      if (model_uuid === undefined || model_uuid === null) {
        throw new Error("Missing the required parameter 'model_uuid' when calling updateModelRegistry");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateModelRegistry");
      }

      let pathParams = {
        'owner': owner,
        'model.uuid': model_uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1ModelRegistry;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/{model.uuid}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
