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

import ApiClient from '../ApiClient';

/**
 * The V1HpQLogUniform model module.
 * @module model/V1HpQLogUniform
 * @version 1.3.2
 */
class V1HpQLogUniform {
    /**
     * Constructs a new <code>V1HpQLogUniform</code>.
     * @alias module:model/V1HpQLogUniform
     */
    constructor() { 
        
        V1HpQLogUniform.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1HpQLogUniform</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1HpQLogUniform} obj Optional instance to populate.
     * @return {module:model/V1HpQLogUniform} The populated <code>V1HpQLogUniform</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1HpQLogUniform();

            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('value')) {
                obj['value'] = ApiClient.convertToType(data['value'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {String} kind
 * @default 'qloguniform'
 */
V1HpQLogUniform.prototype['kind'] = 'qloguniform';

/**
 * @member {Object} value
 */
V1HpQLogUniform.prototype['value'] = undefined;






export default V1HpQLogUniform;

