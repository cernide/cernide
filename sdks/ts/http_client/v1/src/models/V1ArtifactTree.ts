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
 * @interface V1ArtifactTree
 */
export interface V1ArtifactTree {
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof V1ArtifactTree
     */
    files?: { [key: string]: string; };
    /**
     * 
     * @type {Array<string>}
     * @memberof V1ArtifactTree
     */
    dirs?: Array<string>;
    /**
     * 
     * @type {boolean}
     * @memberof V1ArtifactTree
     */
    is_done?: boolean;
}

export function V1ArtifactTreeFromJSON(json: any): V1ArtifactTree {
    return V1ArtifactTreeFromJSONTyped(json, false);
}

export function V1ArtifactTreeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ArtifactTree {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'files': !exists(json, 'files') ? undefined : json['files'],
        'dirs': !exists(json, 'dirs') ? undefined : json['dirs'],
        'is_done': !exists(json, 'is_done') ? undefined : json['is_done'],
    };
}

export function V1ArtifactTreeToJSON(value?: V1ArtifactTree | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'files': value.files,
        'dirs': value.dirs,
        'is_done': value.is_done,
    };
}


