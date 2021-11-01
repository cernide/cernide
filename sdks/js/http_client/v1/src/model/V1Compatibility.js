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
import V1Version from './V1Version';

/**
 * The V1Compatibility model module.
 * @module model/V1Compatibility
 * @version 1.12.0
 */
class V1Compatibility {
    /**
     * Constructs a new <code>V1Compatibility</code>.
     * @alias module:model/V1Compatibility
     */
    constructor() { 
        
        V1Compatibility.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Compatibility</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Compatibility} obj Optional instance to populate.
     * @return {module:model/V1Compatibility} The populated <code>V1Compatibility</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Compatibility();

            if (data.hasOwnProperty('cli')) {
                obj['cli'] = V1Version.constructFromObject(data['cli']);
            }
            if (data.hasOwnProperty('platform')) {
                obj['platform'] = V1Version.constructFromObject(data['platform']);
            }
            if (data.hasOwnProperty('agent')) {
                obj['agent'] = V1Version.constructFromObject(data['agent']);
            }
            if (data.hasOwnProperty('ui')) {
                obj['ui'] = V1Version.constructFromObject(data['ui']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/V1Version} cli
 */
V1Compatibility.prototype['cli'] = undefined;

/**
 * @member {module:model/V1Version} platform
 */
V1Compatibility.prototype['platform'] = undefined;

/**
 * @member {module:model/V1Version} agent
 */
V1Compatibility.prototype['agent'] = undefined;

/**
 * @member {module:model/V1Version} ui
 */
V1Compatibility.prototype['ui'] = undefined;






export default V1Compatibility;

