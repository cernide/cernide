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
 * The version of the OpenAPI document: 1.0.99
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1DockerfileType model module.
 * @module model/V1DockerfileType
 * @version 1.0.99
 */
class V1DockerfileType {
    /**
     * Constructs a new <code>V1DockerfileType</code>.
     * @alias module:model/V1DockerfileType
     */
    constructor() { 
        
        V1DockerfileType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1DockerfileType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1DockerfileType} obj Optional instance to populate.
     * @return {module:model/V1DockerfileType} The populated <code>V1DockerfileType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1DockerfileType();

            if (data.hasOwnProperty('image')) {
                obj['image'] = ApiClient.convertToType(data['image'], 'String');
            }
            if (data.hasOwnProperty('env')) {
                obj['env'] = ApiClient.convertToType(data['env'], {'String': 'String'});
            }
            if (data.hasOwnProperty('path')) {
                obj['path'] = ApiClient.convertToType(data['path'], {'String': 'String'});
            }
            if (data.hasOwnProperty('copy')) {
                obj['copy'] = ApiClient.convertToType(data['copy'], {'String': 'String'});
            }
            if (data.hasOwnProperty('run')) {
                obj['run'] = ApiClient.convertToType(data['run'], ['String']);
            }
            if (data.hasOwnProperty('lang_env')) {
                obj['lang_env'] = ApiClient.convertToType(data['lang_env'], 'String');
            }
            if (data.hasOwnProperty('uid')) {
                obj['uid'] = ApiClient.convertToType(data['uid'], 'Number');
            }
            if (data.hasOwnProperty('gid')) {
                obj['gid'] = ApiClient.convertToType(data['gid'], 'Number');
            }
            if (data.hasOwnProperty('filename')) {
                obj['filename'] = ApiClient.convertToType(data['filename'], 'String');
            }
            if (data.hasOwnProperty('workdir')) {
                obj['workdir'] = ApiClient.convertToType(data['workdir'], 'String');
            }
            if (data.hasOwnProperty('workdir_path')) {
                obj['workdir_path'] = ApiClient.convertToType(data['workdir_path'], 'String');
            }
            if (data.hasOwnProperty('shell')) {
                obj['shell'] = ApiClient.convertToType(data['shell'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} image
 */
V1DockerfileType.prototype['image'] = undefined;

/**
 * @member {Object.<String, String>} env
 */
V1DockerfileType.prototype['env'] = undefined;

/**
 * @member {Object.<String, String>} path
 */
V1DockerfileType.prototype['path'] = undefined;

/**
 * @member {Object.<String, String>} copy
 */
V1DockerfileType.prototype['copy'] = undefined;

/**
 * @member {Array.<String>} run
 */
V1DockerfileType.prototype['run'] = undefined;

/**
 * @member {String} lang_env
 */
V1DockerfileType.prototype['lang_env'] = undefined;

/**
 * @member {Number} uid
 */
V1DockerfileType.prototype['uid'] = undefined;

/**
 * @member {Number} gid
 */
V1DockerfileType.prototype['gid'] = undefined;

/**
 * @member {String} filename
 */
V1DockerfileType.prototype['filename'] = undefined;

/**
 * @member {String} workdir
 */
V1DockerfileType.prototype['workdir'] = undefined;

/**
 * @member {String} workdir_path
 */
V1DockerfileType.prototype['workdir_path'] = undefined;

/**
 * @member {String} shell
 */
V1DockerfileType.prototype['shell'] = undefined;






export default V1DockerfileType;

