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
 * The version of the OpenAPI document: 1.0.99
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';

export interface UploadArtifactRequest {
    owner: string;
    uuid: string;
    uploadfile: Blob;
    path?: string;
    overwrite?: boolean;
}

/**
 * Polyaxon sdk
 */
export class ArtifactsStoresV1Api extends runtime.BaseAPI {

    /**
     * Upload artifact to a store
     */
    async uploadArtifactRaw(requestParameters: UploadArtifactRequest): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling uploadArtifact.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling uploadArtifact.');
        }

        if (requestParameters.uploadfile === null || requestParameters.uploadfile === undefined) {
            throw new runtime.RequiredError('uploadfile','Required parameter requestParameters.uploadfile was null or undefined when calling uploadArtifact.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.path !== undefined) {
            queryParameters['path'] = requestParameters.path;
        }

        if (requestParameters.overwrite !== undefined) {
            queryParameters['overwrite'] = requestParameters.overwrite;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const consumes: runtime.Consume[] = [
            { contentType: 'multipart/form-data' },
        ];
        // @ts-ignore: canConsumeForm may be unused
        const canConsumeForm = runtime.canConsumeForm(consumes);

        let formParams: { append(param: string, value: any): any };
        let useForm = false;
        // use FormData to transmit files using content-type "multipart/form-data"
        useForm = canConsumeForm;
        if (useForm) {
            formParams = new FormData();
        } else {
            formParams = new URLSearchParams();
        }

        if (requestParameters.uploadfile !== undefined) {
            formParams.append('uploadfile', requestParameters.uploadfile as any);
        }

        const response = await this.request({
            path: `/api/v1/catalogs/{owner}/artifacts/{uuid}/upload`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        });

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Upload artifact to a store
     */
    async uploadArtifact(requestParameters: UploadArtifactRequest): Promise<void> {
        await this.uploadArtifactRaw(requestParameters);
    }

}
