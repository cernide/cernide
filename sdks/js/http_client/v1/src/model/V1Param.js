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
 * The version of the OpenAPI document: 1.2.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1ParamSearch from './V1ParamSearch';

/**
 * The V1Param model module.
 * @module model/V1Param
 * @version 1.2.0
 */
class V1Param {
    /**
     * Constructs a new <code>V1Param</code>.
     * @alias module:model/V1Param
     */
    constructor() { 
        
        V1Param.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Param</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Param} obj Optional instance to populate.
     * @return {module:model/V1Param} The populated <code>V1Param</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Param();

            if (data.hasOwnProperty('value')) {
                obj['value'] = ApiClient.convertToType(data['value'], Object);
            }
            if (data.hasOwnProperty('ref')) {
                obj['ref'] = ApiClient.convertToType(data['ref'], 'String');
            }
            if (data.hasOwnProperty('search')) {
                obj['search'] = V1ParamSearch.constructFromObject(data['search']);
            }
            if (data.hasOwnProperty('connection')) {
                obj['connection'] = ApiClient.convertToType(data['connection'], 'String');
            }
            if (data.hasOwnProperty('to_init')) {
                obj['to_init'] = ApiClient.convertToType(data['to_init'], 'Boolean');
            }
            if (data.hasOwnProperty('context_only')) {
                obj['context_only'] = ApiClient.convertToType(data['context_only'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {Object} value
 */
V1Param.prototype['value'] = undefined;

/**
 * @member {String} ref
 */
V1Param.prototype['ref'] = undefined;

/**
 * @member {module:model/V1ParamSearch} search
 */
V1Param.prototype['search'] = undefined;

/**
 * @member {String} connection
 */
V1Param.prototype['connection'] = undefined;

/**
 * @member {Boolean} to_init
 */
V1Param.prototype['to_init'] = undefined;

/**
 * @member {Boolean} context_only
 */
V1Param.prototype['context_only'] = undefined;






export default V1Param;

