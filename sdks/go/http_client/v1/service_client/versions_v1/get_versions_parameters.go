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

package versions_v1

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

// NewGetVersionsParams creates a new GetVersionsParams object
// with the default values initialized.
func NewGetVersionsParams() *GetVersionsParams {

	return &GetVersionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionsParamsWithTimeout creates a new GetVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVersionsParamsWithTimeout(timeout time.Duration) *GetVersionsParams {

	return &GetVersionsParams{

		timeout: timeout,
	}
}

// NewGetVersionsParamsWithContext creates a new GetVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVersionsParamsWithContext(ctx context.Context) *GetVersionsParams {

	return &GetVersionsParams{

		Context: ctx,
	}
}

// NewGetVersionsParamsWithHTTPClient creates a new GetVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVersionsParamsWithHTTPClient(client *http.Client) *GetVersionsParams {

	return &GetVersionsParams{
		HTTPClient: client,
	}
}

/*GetVersionsParams contains all the parameters to send to the API endpoint
for the get versions operation typically these are written to a http.Request
*/
type GetVersionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get versions params
func (o *GetVersionsParams) WithTimeout(timeout time.Duration) *GetVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get versions params
func (o *GetVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get versions params
func (o *GetVersionsParams) WithContext(ctx context.Context) *GetVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get versions params
func (o *GetVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get versions params
func (o *GetVersionsParams) WithHTTPClient(client *http.Client) *GetVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get versions params
func (o *GetVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
