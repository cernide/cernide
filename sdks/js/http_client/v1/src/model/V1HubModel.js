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
 * The version of the OpenAPI document: 1.0.89
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1HubModel model module.
 * @module model/V1HubModel
 * @version 1.0.89
 */
class V1HubModel {
    /**
     * Constructs a new <code>V1HubModel</code>.
     * @alias module:model/V1HubModel
     */
    constructor() { 
        
        V1HubModel.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1HubModel</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1HubModel} obj Optional instance to populate.
     * @return {module:model/V1HubModel} The populated <code>V1HubModel</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1HubModel();

            if (data.hasOwnProperty('uuid')) {
                obj['uuid'] = ApiClient.convertToType(data['uuid'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('framework')) {
                obj['framework'] = ApiClient.convertToType(data['framework'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], ['String']);
            }
            if (data.hasOwnProperty('disabled')) {
                obj['disabled'] = ApiClient.convertToType(data['disabled'], 'Boolean');
            }
            if (data.hasOwnProperty('deleted')) {
                obj['deleted'] = ApiClient.convertToType(data['deleted'], 'Boolean');
            }
            if (data.hasOwnProperty('created_at')) {
                obj['created_at'] = ApiClient.convertToType(data['created_at'], 'Date');
            }
            if (data.hasOwnProperty('updated_at')) {
                obj['updated_at'] = ApiClient.convertToType(data['updated_at'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {String} uuid
 */
V1HubModel.prototype['uuid'] = undefined;

/**
 * @member {String} name
 */
V1HubModel.prototype['name'] = undefined;

/**
 * @member {String} framework
 */
V1HubModel.prototype['framework'] = undefined;

/**
 * @member {String} description
 */
V1HubModel.prototype['description'] = undefined;

/**
 * @member {Array.<String>} tags
 */
V1HubModel.prototype['tags'] = undefined;

/**
 * @member {Boolean} disabled
 */
V1HubModel.prototype['disabled'] = undefined;

/**
 * @member {Boolean} deleted
 */
V1HubModel.prototype['deleted'] = undefined;

/**
 * @member {Date} created_at
 */
V1HubModel.prototype['created_at'] = undefined;

/**
 * @member {Date} updated_at
 */
V1HubModel.prototype['updated_at'] = undefined;






export default V1HubModel;

