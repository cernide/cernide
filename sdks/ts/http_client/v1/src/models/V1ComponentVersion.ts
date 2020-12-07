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
    V1StageCondition,
    V1StageConditionFromJSON,
    V1StageConditionFromJSONTyped,
    V1StageConditionToJSON,
    V1Stages,
    V1StagesFromJSON,
    V1StagesFromJSONTyped,
    V1StagesToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1ComponentVersion
 */
export interface V1ComponentVersion {
    /**
     * 
     * @type {string}
     * @memberof V1ComponentVersion
     */
    uuid?: string;
    /**
     * 
     * @type {string}
     * @memberof V1ComponentVersion
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof V1ComponentVersion
     */
    description?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1ComponentVersion
     */
    tags?: Array<string>;
    /**
     * 
     * @type {number}
     * @memberof V1ComponentVersion
     */
    live_state?: number;
    /**
     * 
     * @type {Date}
     * @memberof V1ComponentVersion
     */
    created_at?: Date;
    /**
     * 
     * @type {Date}
     * @memberof V1ComponentVersion
     */
    updated_at?: Date;
    /**
     * 
     * @type {V1Stages}
     * @memberof V1ComponentVersion
     */
    stage?: V1Stages;
    /**
     * 
     * @type {Array<V1StageCondition>}
     * @memberof V1ComponentVersion
     */
    stage_conditions?: Array<V1StageCondition>;
    /**
     * 
     * @type {string}
     * @memberof V1ComponentVersion
     */
    content?: string;
}

export function V1ComponentVersionFromJSON(json: any): V1ComponentVersion {
    return V1ComponentVersionFromJSONTyped(json, false);
}

export function V1ComponentVersionFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ComponentVersion {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'uuid': !exists(json, 'uuid') ? undefined : json['uuid'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'live_state': !exists(json, 'live_state') ? undefined : json['live_state'],
        'created_at': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'updated_at': !exists(json, 'updated_at') ? undefined : (new Date(json['updated_at'])),
        'stage': !exists(json, 'stage') ? undefined : V1StagesFromJSON(json['stage']),
        'stage_conditions': !exists(json, 'stage_conditions') ? undefined : ((json['stage_conditions'] as Array<any>).map(V1StageConditionFromJSON)),
        'content': !exists(json, 'content') ? undefined : json['content'],
    };
}

export function V1ComponentVersionToJSON(value?: V1ComponentVersion | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'uuid': value.uuid,
        'name': value.name,
        'description': value.description,
        'tags': value.tags,
        'live_state': value.live_state,
        'created_at': value.created_at === undefined ? undefined : (value.created_at.toISOString()),
        'updated_at': value.updated_at === undefined ? undefined : (value.updated_at.toISOString()),
        'stage': V1StagesToJSON(value.stage),
        'stage_conditions': value.stage_conditions === undefined ? undefined : ((value.stage_conditions as Array<any>).map(V1StageConditionToJSON)),
        'content': value.content,
    };
}


