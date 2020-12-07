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

package projects_v1

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

// NewBookmarkProjectParams creates a new BookmarkProjectParams object
// with the default values initialized.
func NewBookmarkProjectParams() *BookmarkProjectParams {
	var ()
	return &BookmarkProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBookmarkProjectParamsWithTimeout creates a new BookmarkProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBookmarkProjectParamsWithTimeout(timeout time.Duration) *BookmarkProjectParams {
	var ()
	return &BookmarkProjectParams{

		timeout: timeout,
	}
}

// NewBookmarkProjectParamsWithContext creates a new BookmarkProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewBookmarkProjectParamsWithContext(ctx context.Context) *BookmarkProjectParams {
	var ()
	return &BookmarkProjectParams{

		Context: ctx,
	}
}

// NewBookmarkProjectParamsWithHTTPClient creates a new BookmarkProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBookmarkProjectParamsWithHTTPClient(client *http.Client) *BookmarkProjectParams {
	var ()
	return &BookmarkProjectParams{
		HTTPClient: client,
	}
}

/*BookmarkProjectParams contains all the parameters to send to the API endpoint
for the bookmark project operation typically these are written to a http.Request
*/
type BookmarkProjectParams struct {

	/*Name
	  Component under namesapce

	*/
	Name string
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bookmark project params
func (o *BookmarkProjectParams) WithTimeout(timeout time.Duration) *BookmarkProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bookmark project params
func (o *BookmarkProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bookmark project params
func (o *BookmarkProjectParams) WithContext(ctx context.Context) *BookmarkProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bookmark project params
func (o *BookmarkProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bookmark project params
func (o *BookmarkProjectParams) WithHTTPClient(client *http.Client) *BookmarkProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bookmark project params
func (o *BookmarkProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the bookmark project params
func (o *BookmarkProjectParams) WithName(name string) *BookmarkProjectParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the bookmark project params
func (o *BookmarkProjectParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the bookmark project params
func (o *BookmarkProjectParams) WithOwner(owner string) *BookmarkProjectParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the bookmark project params
func (o *BookmarkProjectParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *BookmarkProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
