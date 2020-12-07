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
import {
    V1ModelVersion,
    V1ModelVersionFromJSON,
    V1ModelVersionFromJSONTyped,
    V1ModelVersionToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1ListModelVersionsResponse
 */
export interface V1ListModelVersionsResponse {
    /**
     * 
     * @type {number}
     * @memberof V1ListModelVersionsResponse
     */
    count?: number;
    /**
     * 
     * @type {Array<V1ModelVersion>}
     * @memberof V1ListModelVersionsResponse
     */
    results?: Array<V1ModelVersion>;
    /**
     * 
     * @type {string}
     * @memberof V1ListModelVersionsResponse
     */
    previous?: string;
    /**
     * 
     * @type {string}
     * @memberof V1ListModelVersionsResponse
     */
    next?: string;
}

export function V1ListModelVersionsResponseFromJSON(json: any): V1ListModelVersionsResponse {
    return V1ListModelVersionsResponseFromJSONTyped(json, false);
}

export function V1ListModelVersionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ListModelVersionsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'count': !exists(json, 'count') ? undefined : json['count'],
        'results': !exists(json, 'results') ? undefined : ((json['results'] as Array<any>).map(V1ModelVersionFromJSON)),
        'previous': !exists(json, 'previous') ? undefined : json['previous'],
        'next': !exists(json, 'next') ? undefined : json['next'],
    };
}

export function V1ListModelVersionsResponseToJSON(value?: V1ListModelVersionsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'count': value.count,
        'results': value.results === undefined ? undefined : ((value.results as Array<any>).map(V1ModelVersionToJSON)),
        'previous': value.previous,
        'next': value.next,
    };
}


