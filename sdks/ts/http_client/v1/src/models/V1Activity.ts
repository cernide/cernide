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
 * The version of the OpenAPI document: 1.8.0
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
 * @interface V1Activity
 */
export interface V1Activity {
    /**
     * 
     * @type {string}
     * @memberof V1Activity
     */
    actor?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Activity
     */
    owner?: string;
    /**
     * 
     * @type {Date}
     * @memberof V1Activity
     */
    created_at?: Date;
    /**
     * 
     * @type {string}
     * @memberof V1Activity
     */
    event_action?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Activity
     */
    event_subject?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Activity
     */
    object_name?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Activity
     */
    object_uuid?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Activity
     */
    object_parent?: string;
}

export function V1ActivityFromJSON(json: any): V1Activity {
    return V1ActivityFromJSONTyped(json, false);
}

export function V1ActivityFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Activity {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'actor': !exists(json, 'actor') ? undefined : json['actor'],
        'owner': !exists(json, 'owner') ? undefined : json['owner'],
        'created_at': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'event_action': !exists(json, 'event_action') ? undefined : json['event_action'],
        'event_subject': !exists(json, 'event_subject') ? undefined : json['event_subject'],
        'object_name': !exists(json, 'object_name') ? undefined : json['object_name'],
        'object_uuid': !exists(json, 'object_uuid') ? undefined : json['object_uuid'],
        'object_parent': !exists(json, 'object_parent') ? undefined : json['object_parent'],
    };
}

export function V1ActivityToJSON(value?: V1Activity | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'actor': value.actor,
        'owner': value.owner,
        'created_at': value.created_at === undefined ? undefined : (value.created_at.toISOString()),
        'event_action': value.event_action,
        'event_subject': value.event_subject,
        'object_name': value.object_name,
        'object_uuid': value.object_uuid,
        'object_parent': value.object_parent,
    };
}


