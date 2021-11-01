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

/* tslint:disable */
/* eslint-disable */
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
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface V1Mapping
 */
export interface V1Mapping {
    /**
     * 
     * @type {string}
     * @memberof V1Mapping
     */
    kind?: string;
    /**
     * 
     * @type {Array<object>}
     * @memberof V1Mapping
     */
    values?: Array<object>;
    /**
     * 
     * @type {number}
     * @memberof V1Mapping
     */
    concurrency?: number;
    /**
     * 
     * @type {Array<object>}
     * @memberof V1Mapping
     */
    earlyStopping?: Array<object>;
}

export function V1MappingFromJSON(json: any): V1Mapping {
    return V1MappingFromJSONTyped(json, false);
}

export function V1MappingFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Mapping {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'values': !exists(json, 'values') ? undefined : json['values'],
        'concurrency': !exists(json, 'concurrency') ? undefined : json['concurrency'],
        'earlyStopping': !exists(json, 'earlyStopping') ? undefined : json['earlyStopping'],
    };
}

export function V1MappingToJSON(value?: V1Mapping | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'kind': value.kind,
        'values': value.values,
        'concurrency': value.concurrency,
        'earlyStopping': value.earlyStopping,
    };
}


