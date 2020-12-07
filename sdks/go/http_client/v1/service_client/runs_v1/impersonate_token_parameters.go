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

package runs_v1

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

// NewImpersonateTokenParams creates a new ImpersonateTokenParams object
// with the default values initialized.
func NewImpersonateTokenParams() *ImpersonateTokenParams {
	var ()
	return &ImpersonateTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImpersonateTokenParamsWithTimeout creates a new ImpersonateTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImpersonateTokenParamsWithTimeout(timeout time.Duration) *ImpersonateTokenParams {
	var ()
	return &ImpersonateTokenParams{

		timeout: timeout,
	}
}

// NewImpersonateTokenParamsWithContext creates a new ImpersonateTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewImpersonateTokenParamsWithContext(ctx context.Context) *ImpersonateTokenParams {
	var ()
	return &ImpersonateTokenParams{

		Context: ctx,
	}
}

// NewImpersonateTokenParamsWithHTTPClient creates a new ImpersonateTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImpersonateTokenParamsWithHTTPClient(client *http.Client) *ImpersonateTokenParams {
	var ()
	return &ImpersonateTokenParams{
		HTTPClient: client,
	}
}

/*ImpersonateTokenParams contains all the parameters to send to the API endpoint
for the impersonate token operation typically these are written to a http.Request
*/
type ImpersonateTokenParams struct {

	/*Entity
	  Owner of the namespace

	*/
	Entity string
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

// WithTimeout adds the timeout to the impersonate token params
func (o *ImpersonateTokenParams) WithTimeout(timeout time.Duration) *ImpersonateTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the impersonate token params
func (o *ImpersonateTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the impersonate token params
func (o *ImpersonateTokenParams) WithContext(ctx context.Context) *ImpersonateTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the impersonate token params
func (o *ImpersonateTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the impersonate token params
func (o *ImpersonateTokenParams) WithHTTPClient(client *http.Client) *ImpersonateTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the impersonate token params
func (o *ImpersonateTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the impersonate token params
func (o *ImpersonateTokenParams) WithEntity(entity string) *ImpersonateTokenParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the impersonate token params
func (o *ImpersonateTokenParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithOwner adds the owner to the impersonate token params
func (o *ImpersonateTokenParams) WithOwner(owner string) *ImpersonateTokenParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the impersonate token params
func (o *ImpersonateTokenParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the impersonate token params
func (o *ImpersonateTokenParams) WithUUID(uuid string) *ImpersonateTokenParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the impersonate token params
func (o *ImpersonateTokenParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ImpersonateTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
		return err
	}

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
