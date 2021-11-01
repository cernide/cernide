// Copyright 2018-2021 Polyaxon, Inc.
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
 * The version of the OpenAPI document: 1.12.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1Token model module.
 * @module model/V1Token
 * @version 1.12.0
 */
class V1Token {
    /**
     * Constructs a new <code>V1Token</code>.
     * @alias module:model/V1Token
     */
    constructor() { 
        
        V1Token.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Token</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Token} obj Optional instance to populate.
     * @return {module:model/V1Token} The populated <code>V1Token</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Token();

            if (data.hasOwnProperty('uuid')) {
                obj['uuid'] = ApiClient.convertToType(data['uuid'], 'String');
            }
            if (data.hasOwnProperty('key')) {
                obj['key'] = ApiClient.convertToType(data['key'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('scopes')) {
                obj['scopes'] = ApiClient.convertToType(data['scopes'], ['String']);
            }
            if (data.hasOwnProperty('services')) {
                obj['services'] = ApiClient.convertToType(data['services'], ['String']);
            }
            if (data.hasOwnProperty('started_at')) {
                obj['started_at'] = ApiClient.convertToType(data['started_at'], 'Date');
            }
            if (data.hasOwnProperty('expires_at')) {
                obj['expires_at'] = ApiClient.convertToType(data['expires_at'], 'Date');
            }
            if (data.hasOwnProperty('created_at')) {
                obj['created_at'] = ApiClient.convertToType(data['created_at'], 'Date');
            }
            if (data.hasOwnProperty('updated_at')) {
                obj['updated_at'] = ApiClient.convertToType(data['updated_at'], 'Date');
            }
            if (data.hasOwnProperty('expiration')) {
                obj['expiration'] = ApiClient.convertToType(data['expiration'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {String} uuid
 */
V1Token.prototype['uuid'] = undefined;

/**
 * @member {String} key
 */
V1Token.prototype['key'] = undefined;

/**
 * @member {String} name
 */
V1Token.prototype['name'] = undefined;

/**
 * @member {Array.<String>} scopes
 */
V1Token.prototype['scopes'] = undefined;

/**
 * @member {Array.<String>} services
 */
V1Token.prototype['services'] = undefined;

/**
 * @member {Date} started_at
 */
V1Token.prototype['started_at'] = undefined;

/**
 * @member {Date} expires_at
 */
V1Token.prototype['expires_at'] = undefined;

/**
 * @member {Date} created_at
 */
V1Token.prototype['created_at'] = undefined;

/**
 * @member {Date} updated_at
 */
V1Token.prototype['updated_at'] = undefined;

/**
 * @member {Number} expiration
 */
V1Token.prototype['expiration'] = undefined;






export default V1Token;

