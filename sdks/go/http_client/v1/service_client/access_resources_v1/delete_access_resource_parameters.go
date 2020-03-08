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

// Code generated by go-swagger; DO NOT EDIT.

package access_resources_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteAccessResourceParams creates a new DeleteAccessResourceParams object
// with the default values initialized.
func NewDeleteAccessResourceParams() *DeleteAccessResourceParams {
	var ()
	return &DeleteAccessResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAccessResourceParamsWithTimeout creates a new DeleteAccessResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAccessResourceParamsWithTimeout(timeout time.Duration) *DeleteAccessResourceParams {
	var ()
	return &DeleteAccessResourceParams{

		timeout: timeout,
	}
}

// NewDeleteAccessResourceParamsWithContext creates a new DeleteAccessResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAccessResourceParamsWithContext(ctx context.Context) *DeleteAccessResourceParams {
	var ()
	return &DeleteAccessResourceParams{

		Context: ctx,
	}
}

// NewDeleteAccessResourceParamsWithHTTPClient creates a new DeleteAccessResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAccessResourceParamsWithHTTPClient(client *http.Client) *DeleteAccessResourceParams {
	var ()
	return &DeleteAccessResourceParams{
		HTTPClient: client,
	}
}

/*DeleteAccessResourceParams contains all the parameters to send to the API endpoint
for the delete access resource operation typically these are written to a http.Request
*/
type DeleteAccessResourceParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*UUID
	  Uuid identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete access resource params
func (o *DeleteAccessResourceParams) WithTimeout(timeout time.Duration) *DeleteAccessResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete access resource params
func (o *DeleteAccessResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete access resource params
func (o *DeleteAccessResourceParams) WithContext(ctx context.Context) *DeleteAccessResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete access resource params
func (o *DeleteAccessResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete access resource params
func (o *DeleteAccessResourceParams) WithHTTPClient(client *http.Client) *DeleteAccessResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete access resource params
func (o *DeleteAccessResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the delete access resource params
func (o *DeleteAccessResourceParams) WithOwner(owner string) *DeleteAccessResourceParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete access resource params
func (o *DeleteAccessResourceParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the delete access resource params
func (o *DeleteAccessResourceParams) WithUUID(uuid string) *DeleteAccessResourceParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete access resource params
func (o *DeleteAccessResourceParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAccessResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
