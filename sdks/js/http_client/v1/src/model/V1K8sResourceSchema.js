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

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.10
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.V1K8sResourceSchema = factory(root.PolyaxonSdk.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';

  /**
   * The V1K8sResourceSchema model module.
   * @module model/V1K8sResourceSchema
   * @version 1.0.5
   */

  /**
   * Constructs a new <code>V1K8sResourceSchema</code>.
   * @alias module:model/V1K8sResourceSchema
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>V1K8sResourceSchema</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V1K8sResourceSchema} obj Optional instance to populate.
   * @return {module:model/V1K8sResourceSchema} The populated <code>V1K8sResourceSchema</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('name'))
        obj.name = ApiClient.convertToType(data['name'], 'String');
      if (data.hasOwnProperty('mount_path'))
        obj.mount_path = ApiClient.convertToType(data['mount_path'], 'String');
      if (data.hasOwnProperty('items'))
        obj.items = ApiClient.convertToType(data['items'], ['String']);
    }
    return obj;
  }

  /**
   * @member {String} name
   */
  exports.prototype.name = undefined;

  /**
   * @member {String} mount_path
   */
  exports.prototype.mount_path = undefined;

  /**
   * @member {Array.<String>} items
   */
  exports.prototype.items = undefined;

  return exports;

}));
