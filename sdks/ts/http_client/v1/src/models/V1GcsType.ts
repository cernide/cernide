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
 * The version of the OpenAPI document: 1.3.2
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
 * @interface V1GcsType
 */
export interface V1GcsType {
    /**
     * 
     * @type {string}
     * @memberof V1GcsType
     */
    bucket?: string;
    /**
     * 
     * @type {string}
     * @memberof V1GcsType
     */
    blob?: string;
}

export function V1GcsTypeFromJSON(json: any): V1GcsType {
    return V1GcsTypeFromJSONTyped(json, false);
}

export function V1GcsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1GcsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bucket': !exists(json, 'bucket') ? undefined : json['bucket'],
        'blob': !exists(json, 'blob') ? undefined : json['blob'],
    };
}

export function V1GcsTypeToJSON(value?: V1GcsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bucket': value.bucket,
        'blob': value.blob,
    };
}


