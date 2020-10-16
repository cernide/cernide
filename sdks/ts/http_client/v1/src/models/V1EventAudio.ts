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
 * The version of the OpenAPI document: 1.2.0
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
 * @interface V1EventAudio
 */
export interface V1EventAudio {
    /**
     * Sample rate of the audio in Hz.
     * @type {number}
     * @memberof V1EventAudio
     */
    sample_rate?: number;
    /**
     * Number of channels of audio.
     * @type {number}
     * @memberof V1EventAudio
     */
    num_channels?: number;
    /**
     * Length of the audio in frames (samples per channel).
     * @type {number}
     * @memberof V1EventAudio
     */
    length_frames?: number;
    /**
     * 
     * @type {string}
     * @memberof V1EventAudio
     */
    content_type?: string;
    /**
     * 
     * @type {string}
     * @memberof V1EventAudio
     */
    path?: string;
}

export function V1EventAudioFromJSON(json: any): V1EventAudio {
    return V1EventAudioFromJSONTyped(json, false);
}

export function V1EventAudioFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1EventAudio {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'sample_rate': !exists(json, 'sample_rate') ? undefined : json['sample_rate'],
        'num_channels': !exists(json, 'num_channels') ? undefined : json['num_channels'],
        'length_frames': !exists(json, 'length_frames') ? undefined : json['length_frames'],
        'content_type': !exists(json, 'content_type') ? undefined : json['content_type'],
        'path': !exists(json, 'path') ? undefined : json['path'],
    };
}

export function V1EventAudioToJSON(value?: V1EventAudio | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'sample_rate': value.sample_rate,
        'num_channels': value.num_channels,
        'length_frames': value.length_frames,
        'content_type': value.content_type,
        'path': value.path,
    };
}


