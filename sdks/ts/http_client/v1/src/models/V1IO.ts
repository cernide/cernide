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

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.4.0-rc8
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface V1IO
 */
export interface V1IO {
    /**
     * 
     * @type {string}
     * @memberof V1IO
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof V1IO
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof V1IO
     */
    iotype?: string;
    /**
     * 
     * @type {object}
     * @memberof V1IO
     */
    value?: object;
    /**
     * 
     * @type {boolean}
     * @memberof V1IO
     */
    is_optional?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof V1IO
     */
    is_list?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof V1IO
     */
    is_flag?: boolean;
    /**
     * 
     * @type {string}
     * @memberof V1IO
     */
    arg_format?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V1IO
     */
    delay_validation?: boolean;
    /**
     * 
     * @type {Array<object>}
     * @memberof V1IO
     */
    options?: Array<object>;
    /**
     * 
     * @type {string}
     * @memberof V1IO
     */
    connection?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V1IO
     */
    to_init?: boolean;
}

export function V1IOFromJSON(json: any): V1IO {
    return V1IOFromJSONTyped(json, false);
}

export function V1IOFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1IO {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'iotype': !exists(json, 'iotype') ? undefined : json['iotype'],
        'value': !exists(json, 'value') ? undefined : json['value'],
        'is_optional': !exists(json, 'is_optional') ? undefined : json['is_optional'],
        'is_list': !exists(json, 'is_list') ? undefined : json['is_list'],
        'is_flag': !exists(json, 'is_flag') ? undefined : json['is_flag'],
        'arg_format': !exists(json, 'arg_format') ? undefined : json['arg_format'],
        'delay_validation': !exists(json, 'delay_validation') ? undefined : json['delay_validation'],
        'options': !exists(json, 'options') ? undefined : json['options'],
        'connection': !exists(json, 'connection') ? undefined : json['connection'],
        'to_init': !exists(json, 'to_init') ? undefined : json['to_init'],
    };
}

export function V1IOToJSON(value?: V1IO | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'description': value.description,
        'iotype': value.iotype,
        'value': value.value,
        'is_optional': value.is_optional,
        'is_list': value.is_list,
        'is_flag': value.is_flag,
        'arg_format': value.arg_format,
        'delay_validation': value.delay_validation,
        'options': value.options,
        'connection': value.connection,
        'to_init': value.to_init,
    };
}


