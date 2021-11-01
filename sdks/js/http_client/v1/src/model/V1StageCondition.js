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
import V1Stages from './V1Stages';

/**
 * The V1StageCondition model module.
 * @module model/V1StageCondition
 * @version 1.12.0
 */
class V1StageCondition {
    /**
     * Constructs a new <code>V1StageCondition</code>.
     * @alias module:model/V1StageCondition
     */
    constructor() { 
        
        V1StageCondition.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1StageCondition</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1StageCondition} obj Optional instance to populate.
     * @return {module:model/V1StageCondition} The populated <code>V1StageCondition</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1StageCondition();

            if (data.hasOwnProperty('type')) {
                obj['type'] = V1Stages.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('reason')) {
                obj['reason'] = ApiClient.convertToType(data['reason'], 'String');
            }
            if (data.hasOwnProperty('message')) {
                obj['message'] = ApiClient.convertToType(data['message'], 'String');
            }
            if (data.hasOwnProperty('last_update_time')) {
                obj['last_update_time'] = ApiClient.convertToType(data['last_update_time'], 'Date');
            }
            if (data.hasOwnProperty('last_transition_time')) {
                obj['last_transition_time'] = ApiClient.convertToType(data['last_transition_time'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/V1Stages} type
 */
V1StageCondition.prototype['type'] = undefined;

/**
 * @member {String} status
 */
V1StageCondition.prototype['status'] = undefined;

/**
 * @member {String} reason
 */
V1StageCondition.prototype['reason'] = undefined;

/**
 * @member {String} message
 */
V1StageCondition.prototype['message'] = undefined;

/**
 * @member {Date} last_update_time
 */
V1StageCondition.prototype['last_update_time'] = undefined;

/**
 * @member {Date} last_transition_time
 */
V1StageCondition.prototype['last_transition_time'] = undefined;






export default V1StageCondition;

