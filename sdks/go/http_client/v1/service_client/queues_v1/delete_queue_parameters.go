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

package queues_v1

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

// NewDeleteQueueParams creates a new DeleteQueueParams object
// with the default values initialized.
func NewDeleteQueueParams() *DeleteQueueParams {
	var ()
	return &DeleteQueueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteQueueParamsWithTimeout creates a new DeleteQueueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteQueueParamsWithTimeout(timeout time.Duration) *DeleteQueueParams {
	var ()
	return &DeleteQueueParams{

		timeout: timeout,
	}
}

// NewDeleteQueueParamsWithContext creates a new DeleteQueueParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteQueueParamsWithContext(ctx context.Context) *DeleteQueueParams {
	var ()
	return &DeleteQueueParams{

		Context: ctx,
	}
}

// NewDeleteQueueParamsWithHTTPClient creates a new DeleteQueueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteQueueParamsWithHTTPClient(client *http.Client) *DeleteQueueParams {
	var ()
	return &DeleteQueueParams{
		HTTPClient: client,
	}
}

/*DeleteQueueParams contains all the parameters to send to the API endpoint
for the delete queue operation typically these are written to a http.Request
*/
type DeleteQueueParams struct {

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

// WithTimeout adds the timeout to the delete queue params
func (o *DeleteQueueParams) WithTimeout(timeout time.Duration) *DeleteQueueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete queue params
func (o *DeleteQueueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete queue params
func (o *DeleteQueueParams) WithContext(ctx context.Context) *DeleteQueueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete queue params
func (o *DeleteQueueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete queue params
func (o *DeleteQueueParams) WithHTTPClient(client *http.Client) *DeleteQueueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete queue params
func (o *DeleteQueueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the delete queue params
func (o *DeleteQueueParams) WithEntity(entity string) *DeleteQueueParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the delete queue params
func (o *DeleteQueueParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithOwner adds the owner to the delete queue params
func (o *DeleteQueueParams) WithOwner(owner string) *DeleteQueueParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete queue params
func (o *DeleteQueueParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the delete queue params
func (o *DeleteQueueParams) WithUUID(uuid string) *DeleteQueueParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete queue params
func (o *DeleteQueueParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteQueueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
