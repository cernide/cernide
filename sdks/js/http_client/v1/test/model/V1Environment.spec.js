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
 * The version of the OpenAPI document: 1.4.0-rc8
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
    instance = new PolyaxonSdk.V1Environment();
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

  describe('V1Environment', function() {
    it('should create an instance of V1Environment', function() {
      // uncomment below and update the code to test V1Environment
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be.a(PolyaxonSdk.V1Environment);
    });

    it('should have the property labels (base name: "labels")', function() {
      // uncomment below and update the code to test the property labels
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property annotations (base name: "annotations")', function() {
      // uncomment below and update the code to test the property annotations
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property node_selector (base name: "node_selector")', function() {
      // uncomment below and update the code to test the property node_selector
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property affinity (base name: "affinity")', function() {
      // uncomment below and update the code to test the property affinity
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property tolerations (base name: "tolerations")', function() {
      // uncomment below and update the code to test the property tolerations
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property node_name (base name: "node_name")', function() {
      // uncomment below and update the code to test the property node_name
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property service_account_name (base name: "service_account_name")', function() {
      // uncomment below and update the code to test the property service_account_name
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property host_aliases (base name: "host_aliases")', function() {
      // uncomment below and update the code to test the property host_aliases
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property security_context (base name: "security_context")', function() {
      // uncomment below and update the code to test the property security_context
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property image_pull_secrets (base name: "image_pull_secrets")', function() {
      // uncomment below and update the code to test the property image_pull_secrets
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property host_network (base name: "host_network")', function() {
      // uncomment below and update the code to test the property host_network
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property hostPID (base name: "hostPID")', function() {
      // uncomment below and update the code to test the property hostPID
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property dns_policy (base name: "dns_policy")', function() {
      // uncomment below and update the code to test the property dns_policy
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property dns_config (base name: "dns_config")', function() {
      // uncomment below and update the code to test the property dns_config
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property scheduler_name (base name: "scheduler_name")', function() {
      // uncomment below and update the code to test the property scheduler_name
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property priority_class_name (base name: "priority_class_name")', function() {
      // uncomment below and update the code to test the property priority_class_name
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property priority (base name: "priority")', function() {
      // uncomment below and update the code to test the property priority
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property restart_policy (base name: "restart_policy")', function() {
      // uncomment below and update the code to test the property restart_policy
      //var instane = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

  });

}));
