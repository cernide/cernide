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

  beforeEach(function() {
    instance = new PolyaxonSdk.AccessResourcesV1Api();
  });

  describe('(package)', function() {
    describe('AccessResourcesV1Api', function() {
      describe('createAccessResource', function() {
        it('should call createAccessResource successfully', function(done) {
          // TODO: uncomment, update parameter values for createAccessResource call and complete the assertions
          /*
          var owner = "owner_example";
          var body = new PolyaxonSdk.V1AccessResource();
          body.uuid = "";
          body.name = "";
          body.description = "";
          body.tags = [""];
          body.created_at = new Date();
          body.updated_at = new Date();
          body.kind = "";
          body.frozen = false;
          body.disabled = false;
          body.deleted = false;

          instance.createAccessResource(owner, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1AccessResource);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());
            expect(data.kind).to.be.a('string');
            expect(data.kind).to.be("");
            expect(data.frozen).to.be.a('boolean');
            expect(data.frozen).to.be(false);
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('deleteAccessResource', function() {
        it('should call deleteAccessResource successfully', function(done) {
          // TODO: uncomment, update parameter values for deleteAccessResource call
          /*
          var owner = "owner_example";
          var uuid = "uuid_example";

          instance.deleteAccessResource(owner, uuid, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('getAccessResource', function() {
        it('should call getAccessResource successfully', function(done) {
          // TODO: uncomment, update parameter values for getAccessResource call and complete the assertions
          /*
          var owner = "owner_example";
          var uuid = "uuid_example";

          instance.getAccessResource(owner, uuid, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1AccessResource);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());
            expect(data.kind).to.be.a('string');
            expect(data.kind).to.be("");
            expect(data.frozen).to.be.a('boolean');
            expect(data.frozen).to.be(false);
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('listAccessResourceNames', function() {
        it('should call listAccessResourceNames successfully', function(done) {
          // TODO: uncomment, update parameter values for listAccessResourceNames call and complete the assertions
          /*
          var owner = "owner_example";
          var opts = {};
          opts.offset = 56;
          opts.limit = 56;
          opts.sort = "sort_example";
          opts.query = "query_example";

          instance.listAccessResourceNames(owner, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListAccessResourcesResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1AccessResource);
                expect(data.uuid).to.be.a('string');
                expect(data.uuid).to.be("");
                expect(data.name).to.be.a('string');
                expect(data.name).to.be("");
                expect(data.description).to.be.a('string');
                expect(data.description).to.be("");
                {
                  let dataCtr = data.tags;
                  expect(dataCtr).to.be.an(Array);
                  expect(dataCtr).to.not.be.empty();
                  for (let p in dataCtr) {
                    let data = dataCtr[p];
                    expect(data).to.be.a('string');
                    expect(data).to.be("");
                  }
                }
                expect(data.created_at).to.be.a(Date);
                expect(data.created_at).to.be(new Date());
                expect(data.updated_at).to.be.a(Date);
                expect(data.updated_at).to.be(new Date());
                expect(data.kind).to.be.a('string');
                expect(data.kind).to.be("");
                expect(data.frozen).to.be.a('boolean');
                expect(data.frozen).to.be(false);
                expect(data.disabled).to.be.a('boolean');
                expect(data.disabled).to.be(false);
                expect(data.deleted).to.be.a('boolean');
                expect(data.deleted).to.be(false);
              }
            }
            expect(data.previous).to.be.a('string');
            expect(data.previous).to.be("");
            expect(data.next).to.be.a('string');
            expect(data.next).to.be("");

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('listAccessResources', function() {
        it('should call listAccessResources successfully', function(done) {
          // TODO: uncomment, update parameter values for listAccessResources call and complete the assertions
          /*
          var owner = "owner_example";
          var opts = {};
          opts.offset = 56;
          opts.limit = 56;
          opts.sort = "sort_example";
          opts.query = "query_example";

          instance.listAccessResources(owner, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListAccessResourcesResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1AccessResource);
                expect(data.uuid).to.be.a('string');
                expect(data.uuid).to.be("");
                expect(data.name).to.be.a('string');
                expect(data.name).to.be("");
                expect(data.description).to.be.a('string');
                expect(data.description).to.be("");
                {
                  let dataCtr = data.tags;
                  expect(dataCtr).to.be.an(Array);
                  expect(dataCtr).to.not.be.empty();
                  for (let p in dataCtr) {
                    let data = dataCtr[p];
                    expect(data).to.be.a('string');
                    expect(data).to.be("");
                  }
                }
                expect(data.created_at).to.be.a(Date);
                expect(data.created_at).to.be(new Date());
                expect(data.updated_at).to.be.a(Date);
                expect(data.updated_at).to.be(new Date());
                expect(data.kind).to.be.a('string');
                expect(data.kind).to.be("");
                expect(data.frozen).to.be.a('boolean');
                expect(data.frozen).to.be(false);
                expect(data.disabled).to.be.a('boolean');
                expect(data.disabled).to.be(false);
                expect(data.deleted).to.be.a('boolean');
                expect(data.deleted).to.be(false);
              }
            }
            expect(data.previous).to.be.a('string');
            expect(data.previous).to.be("");
            expect(data.next).to.be.a('string');
            expect(data.next).to.be("");

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('patchAccessResource', function() {
        it('should call patchAccessResource successfully', function(done) {
          // TODO: uncomment, update parameter values for patchAccessResource call and complete the assertions
          /*
          var owner = "owner_example";
          var access_resource_uuid = "access_resource_uuid_example";
          var body = new PolyaxonSdk.V1AccessResource();
          body.uuid = "";
          body.name = "";
          body.description = "";
          body.tags = [""];
          body.created_at = new Date();
          body.updated_at = new Date();
          body.kind = "";
          body.frozen = false;
          body.disabled = false;
          body.deleted = false;

          instance.patchAccessResource(owner, access_resource_uuid, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1AccessResource);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());
            expect(data.kind).to.be.a('string');
            expect(data.kind).to.be("");
            expect(data.frozen).to.be.a('boolean');
            expect(data.frozen).to.be(false);
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('updateAccessResource', function() {
        it('should call updateAccessResource successfully', function(done) {
          // TODO: uncomment, update parameter values for updateAccessResource call and complete the assertions
          /*
          var owner = "owner_example";
          var access_resource_uuid = "access_resource_uuid_example";
          var body = new PolyaxonSdk.V1AccessResource();
          body.uuid = "";
          body.name = "";
          body.description = "";
          body.tags = [""];
          body.created_at = new Date();
          body.updated_at = new Date();
          body.kind = "";
          body.frozen = false;
          body.disabled = false;
          body.deleted = false;

          instance.updateAccessResource(owner, access_resource_uuid, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1AccessResource);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());
            expect(data.kind).to.be.a('string');
            expect(data.kind).to.be("");
            expect(data.frozen).to.be.a('boolean');
            expect(data.frozen).to.be(false);
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
    });
  });

}));
