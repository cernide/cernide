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

package organizations_v1

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

// NewGetOrganizationParams creates a new GetOrganizationParams object
// with the default values initialized.
func NewGetOrganizationParams() *GetOrganizationParams {
	var ()
	return &GetOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationParamsWithTimeout creates a new GetOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationParamsWithTimeout(timeout time.Duration) *GetOrganizationParams {
	var ()
	return &GetOrganizationParams{

		timeout: timeout,
	}
}

// NewGetOrganizationParamsWithContext creates a new GetOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationParamsWithContext(ctx context.Context) *GetOrganizationParams {
	var ()
	return &GetOrganizationParams{

		Context: ctx,
	}
}

// NewGetOrganizationParamsWithHTTPClient creates a new GetOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationParamsWithHTTPClient(client *http.Client) *GetOrganizationParams {
	var ()
	return &GetOrganizationParams{
		HTTPClient: client,
	}
}

/*GetOrganizationParams contains all the parameters to send to the API endpoint
for the get organization operation typically these are written to a http.Request
*/
type GetOrganizationParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organization params
func (o *GetOrganizationParams) WithTimeout(timeout time.Duration) *GetOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization params
func (o *GetOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization params
func (o *GetOrganizationParams) WithContext(ctx context.Context) *GetOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization params
func (o *GetOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization params
func (o *GetOrganizationParams) WithHTTPClient(client *http.Client) *GetOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization params
func (o *GetOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get organization params
func (o *GetOrganizationParams) WithOwner(owner string) *GetOrganizationParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get organization params
func (o *GetOrganizationParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
