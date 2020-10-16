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
import V1Environment from './V1Environment';
import V1Init from './V1Init';

/**
 * The V1Dask model module.
 * @module model/V1Dask
 * @version 1.2.0
 */
class V1Dask {
    /**
     * Constructs a new <code>V1Dask</code>.
     * @alias module:model/V1Dask
     */
    constructor() { 
        
        V1Dask.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Dask</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Dask} obj Optional instance to populate.
     * @return {module:model/V1Dask} The populated <code>V1Dask</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Dask();

            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('threads')) {
                obj['threads'] = ApiClient.convertToType(data['threads'], 'Number');
            }
            if (data.hasOwnProperty('scale')) {
                obj['scale'] = ApiClient.convertToType(data['scale'], 'Number');
            }
            if (data.hasOwnProperty('adapt_min')) {
                obj['adapt_min'] = ApiClient.convertToType(data['adapt_min'], 'Number');
            }
            if (data.hasOwnProperty('adapt_max')) {
                obj['adapt_max'] = ApiClient.convertToType(data['adapt_max'], 'Number');
            }
            if (data.hasOwnProperty('adapt_interval')) {
                obj['adapt_interval'] = ApiClient.convertToType(data['adapt_interval'], 'String');
            }
            if (data.hasOwnProperty('environment')) {
                obj['environment'] = V1Environment.constructFromObject(data['environment']);
            }
            if (data.hasOwnProperty('connections')) {
                obj['connections'] = ApiClient.convertToType(data['connections'], ['String']);
            }
            if (data.hasOwnProperty('volumes')) {
                obj['volumes'] = ApiClient.convertToType(data['volumes'], [Object]);
            }
            if (data.hasOwnProperty('init')) {
                obj['init'] = ApiClient.convertToType(data['init'], [V1Init]);
            }
            if (data.hasOwnProperty('sidecars')) {
                obj['sidecars'] = ApiClient.convertToType(data['sidecars'], [Object]);
            }
            if (data.hasOwnProperty('container')) {
                obj['container'] = ApiClient.convertToType(data['container'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {String} kind
 * @default 'dask'
 */
V1Dask.prototype['kind'] = 'dask';

/**
 * @member {Number} threads
 */
V1Dask.prototype['threads'] = undefined;

/**
 * @member {Number} scale
 */
V1Dask.prototype['scale'] = undefined;

/**
 * @member {Number} adapt_min
 */
V1Dask.prototype['adapt_min'] = undefined;

/**
 * @member {Number} adapt_max
 */
V1Dask.prototype['adapt_max'] = undefined;

/**
 * @member {String} adapt_interval
 */
V1Dask.prototype['adapt_interval'] = undefined;

/**
 * @member {module:model/V1Environment} environment
 */
V1Dask.prototype['environment'] = undefined;

/**
 * @member {Array.<String>} connections
 */
V1Dask.prototype['connections'] = undefined;

/**
 * Volumes is a list of volumes that can be mounted.
 * @member {Array.<Object>} volumes
 */
V1Dask.prototype['volumes'] = undefined;

/**
 * @member {Array.<module:model/V1Init>} init
 */
V1Dask.prototype['init'] = undefined;

/**
 * @member {Array.<Object>} sidecars
 */
V1Dask.prototype['sidecars'] = undefined;

/**
 * @member {Object} container
 */
V1Dask.prototype['container'] = undefined;






export default V1Dask;

