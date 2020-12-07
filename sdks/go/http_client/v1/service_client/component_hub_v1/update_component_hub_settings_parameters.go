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

package component_hub_v1

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

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewUpdateComponentHubSettingsParams creates a new UpdateComponentHubSettingsParams object
// with the default values initialized.
func NewUpdateComponentHubSettingsParams() *UpdateComponentHubSettingsParams {
	var ()
	return &UpdateComponentHubSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateComponentHubSettingsParamsWithTimeout creates a new UpdateComponentHubSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateComponentHubSettingsParamsWithTimeout(timeout time.Duration) *UpdateComponentHubSettingsParams {
	var ()
	return &UpdateComponentHubSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateComponentHubSettingsParamsWithContext creates a new UpdateComponentHubSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateComponentHubSettingsParamsWithContext(ctx context.Context) *UpdateComponentHubSettingsParams {
	var ()
	return &UpdateComponentHubSettingsParams{

		Context: ctx,
	}
}

// NewUpdateComponentHubSettingsParamsWithHTTPClient creates a new UpdateComponentHubSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateComponentHubSettingsParamsWithHTTPClient(client *http.Client) *UpdateComponentHubSettingsParams {
	var ()
	return &UpdateComponentHubSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateComponentHubSettingsParams contains all the parameters to send to the API endpoint
for the update component hub settings operation typically these are written to a http.Request
*/
type UpdateComponentHubSettingsParams struct {

	/*Body
	  Hub settings body

	*/
	Body *service_model.V1ComponentHubSettings
	/*Component
	  Hub name

	*/
	Component string
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update component hub settings params
func (o *UpdateComponentHubSettingsParams) WithTimeout(timeout time.Duration) *UpdateComponentHubSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update component hub settings params
func (o *UpdateComponentHubSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update component hub settings params
func (o *UpdateComponentHubSettingsParams) WithContext(ctx context.Context) *UpdateComponentHubSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update component hub settings params
func (o *UpdateComponentHubSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update component hub settings params
func (o *UpdateComponentHubSettingsParams) WithHTTPClient(client *http.Client) *UpdateComponentHubSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update component hub settings params
func (o *UpdateComponentHubSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update component hub settings params
func (o *UpdateComponentHubSettingsParams) WithBody(body *service_model.V1ComponentHubSettings) *UpdateComponentHubSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update component hub settings params
func (o *UpdateComponentHubSettingsParams) SetBody(body *service_model.V1ComponentHubSettings) {
	o.Body = body
}

// WithComponent adds the component to the update component hub settings params
func (o *UpdateComponentHubSettingsParams) WithComponent(component string) *UpdateComponentHubSettingsParams {
	o.SetComponent(component)
	return o
}

// SetComponent adds the component to the update component hub settings params
func (o *UpdateComponentHubSettingsParams) SetComponent(component string) {
	o.Component = component
}

// WithOwner adds the owner to the update component hub settings params
func (o *UpdateComponentHubSettingsParams) WithOwner(owner string) *UpdateComponentHubSettingsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update component hub settings params
func (o *UpdateComponentHubSettingsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateComponentHubSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param component
	if err := r.SetPathParam("component", o.Component); err != nil {
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
