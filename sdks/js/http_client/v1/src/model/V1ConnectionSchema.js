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
import V1BucketConnection from './V1BucketConnection';
import V1ClaimConnection from './V1ClaimConnection';
import V1GitConnection from './V1GitConnection';
import V1HostConnection from './V1HostConnection';
import V1HostPathConnection from './V1HostPathConnection';

/**
 * The V1ConnectionSchema model module.
 * @module model/V1ConnectionSchema
 * @version 1.12.0
 */
class V1ConnectionSchema {
    /**
     * Constructs a new <code>V1ConnectionSchema</code>.
     * @alias module:model/V1ConnectionSchema
     */
    constructor() { 
        
        V1ConnectionSchema.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1ConnectionSchema</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1ConnectionSchema} obj Optional instance to populate.
     * @return {module:model/V1ConnectionSchema} The populated <code>V1ConnectionSchema</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1ConnectionSchema();

            if (data.hasOwnProperty('bucketConnection')) {
                obj['bucketConnection'] = V1BucketConnection.constructFromObject(data['bucketConnection']);
            }
            if (data.hasOwnProperty('hostPathConnection')) {
                obj['hostPathConnection'] = V1HostPathConnection.constructFromObject(data['hostPathConnection']);
            }
            if (data.hasOwnProperty('claimConnection')) {
                obj['claimConnection'] = V1ClaimConnection.constructFromObject(data['claimConnection']);
            }
            if (data.hasOwnProperty('hostConnection')) {
                obj['hostConnection'] = V1HostConnection.constructFromObject(data['hostConnection']);
            }
            if (data.hasOwnProperty('gitConnection')) {
                obj['gitConnection'] = V1GitConnection.constructFromObject(data['gitConnection']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/V1BucketConnection} bucketConnection
 */
V1ConnectionSchema.prototype['bucketConnection'] = undefined;

/**
 * @member {module:model/V1HostPathConnection} hostPathConnection
 */
V1ConnectionSchema.prototype['hostPathConnection'] = undefined;

/**
 * @member {module:model/V1ClaimConnection} claimConnection
 */
V1ConnectionSchema.prototype['claimConnection'] = undefined;

/**
 * @member {module:model/V1HostConnection} hostConnection
 */
V1ConnectionSchema.prototype['hostConnection'] = undefined;

/**
 * @member {module:model/V1GitConnection} gitConnection
 */
V1ConnectionSchema.prototype['gitConnection'] = undefined;






export default V1ConnectionSchema;

