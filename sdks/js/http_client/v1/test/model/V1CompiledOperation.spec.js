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
 * The version of the OpenAPI document: 1.0.89
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
    instance = new PolyaxonSdk.V1CompiledOperation();
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

  describe('V1CompiledOperation', function() {
    it('should create an instance of V1CompiledOperation', function() {
      // uncomment below and update the code to test V1CompiledOperation
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be.a(PolyaxonSdk.V1CompiledOperation);
    });

    it('should have the property version (base name: "version")', function() {
      // uncomment below and update the code to test the property version
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property kind (base name: "kind")', function() {
      // uncomment below and update the code to test the property kind
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property tags (base name: "tags")', function() {
      // uncomment below and update the code to test the property tags
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property profile (base name: "profile")', function() {
      // uncomment below and update the code to test the property profile
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property queue (base name: "queue")', function() {
      // uncomment below and update the code to test the property queue
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property cache (base name: "cache")', function() {
      // uncomment below and update the code to test the property cache
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property schedule (base name: "schedule")', function() {
      // uncomment below and update the code to test the property schedule
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property events (base name: "events")', function() {
      // uncomment below and update the code to test the property events
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property parallel (base name: "parallel")', function() {
      // uncomment below and update the code to test the property parallel
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property dependencies (base name: "dependencies")', function() {
      // uncomment below and update the code to test the property dependencies
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property trigger (base name: "trigger")', function() {
      // uncomment below and update the code to test the property trigger
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property conditions (base name: "conditions")', function() {
      // uncomment below and update the code to test the property conditions
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property skip_on_upstream_skip (base name: "skip_on_upstream_skip")', function() {
      // uncomment below and update the code to test the property skip_on_upstream_skip
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property termination (base name: "termination")', function() {
      // uncomment below and update the code to test the property termination
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property plugins (base name: "plugins")', function() {
      // uncomment below and update the code to test the property plugins
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property inputs (base name: "inputs")', function() {
      // uncomment below and update the code to test the property inputs
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property outputs (base name: "outputs")', function() {
      // uncomment below and update the code to test the property outputs
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

    it('should have the property run (base name: "run")', function() {
      // uncomment below and update the code to test the property run
      //var instane = new PolyaxonSdk.V1CompiledOperation();
      //expect(instance).to.be();
    });

  });

}));
