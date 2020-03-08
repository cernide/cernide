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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('V1Schedule', function() {
      beforeEach(function() {
        instance = new PolyaxonSdk.V1Schedule();
      });

      it('should create an instance of V1Schedule', function() {
        // TODO: update the code to test V1Schedule
        expect(instance).to.be.a(PolyaxonSdk.V1Schedule);
      });

      it('should have the property cron (base name: "cron")', function() {
        // TODO: update the code to test the property cron
        expect(instance).to.have.property('cron');
        // expect(instance.cron).to.be(expectedValueLiteral);
      });

      it('should have the property exact_time (base name: "exact_time")', function() {
        // TODO: update the code to test the property exact_time
        expect(instance).to.have.property('exact_time');
        // expect(instance.exact_time).to.be(expectedValueLiteral);
      });

      it('should have the property interval (base name: "interval")', function() {
        // TODO: update the code to test the property interval
        expect(instance).to.have.property('interval');
        // expect(instance.interval).to.be(expectedValueLiteral);
      });

      it('should have the property repeatable (base name: "repeatable")', function() {
        // TODO: update the code to test the property repeatable
        expect(instance).to.have.property('repeatable');
        // expect(instance.repeatable).to.be(expectedValueLiteral);
      });

    });
  });

}));
