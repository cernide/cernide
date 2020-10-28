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

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.2.1-rc5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new PolyaxonSdk.V1ListDashboardsResponse();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('V1ListDashboardsResponse', function() {
    it('should create an instance of V1ListDashboardsResponse', function() {
      // uncomment below and update the code to test V1ListDashboardsResponse
      //var instane = new PolyaxonSdk.V1ListDashboardsResponse();
      //expect(instance).to.be.a(PolyaxonSdk.V1ListDashboardsResponse);
    });

    it('should have the property count (base name: "count")', function() {
      // uncomment below and update the code to test the property count
      //var instane = new PolyaxonSdk.V1ListDashboardsResponse();
      //expect(instance).to.be();
    });

    it('should have the property results (base name: "results")', function() {
      // uncomment below and update the code to test the property results
      //var instane = new PolyaxonSdk.V1ListDashboardsResponse();
      //expect(instance).to.be();
    });

    it('should have the property previous (base name: "previous")', function() {
      // uncomment below and update the code to test the property previous
      //var instane = new PolyaxonSdk.V1ListDashboardsResponse();
      //expect(instance).to.be();
    });

    it('should have the property next (base name: "next")', function() {
      // uncomment below and update the code to test the property next
      //var instane = new PolyaxonSdk.V1ListDashboardsResponse();
      //expect(instance).to.be();
    });

  });

}));
