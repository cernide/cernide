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

package model_registry_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// DeleteModelRegistryReader is a Reader for the DeleteModelRegistry structure.
type DeleteModelRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteModelRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteModelRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteModelRegistryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteModelRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteModelRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteModelRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteModelRegistryOK creates a DeleteModelRegistryOK with default headers values
func NewDeleteModelRegistryOK() *DeleteModelRegistryOK {
	return &DeleteModelRegistryOK{}
}

/*DeleteModelRegistryOK handles this case with default header values.

A successful response.
*/
type DeleteModelRegistryOK struct {
}

func (o *DeleteModelRegistryOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/models/{uuid}][%d] deleteModelRegistryOK ", 200)
}

func (o *DeleteModelRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteModelRegistryNoContent creates a DeleteModelRegistryNoContent with default headers values
func NewDeleteModelRegistryNoContent() *DeleteModelRegistryNoContent {
	return &DeleteModelRegistryNoContent{}
}

/*DeleteModelRegistryNoContent handles this case with default header values.

No content.
*/
type DeleteModelRegistryNoContent struct {
	Payload interface{}
}

func (o *DeleteModelRegistryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/models/{uuid}][%d] deleteModelRegistryNoContent  %+v", 204, o.Payload)
}

func (o *DeleteModelRegistryNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteModelRegistryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModelRegistryForbidden creates a DeleteModelRegistryForbidden with default headers values
func NewDeleteModelRegistryForbidden() *DeleteModelRegistryForbidden {
	return &DeleteModelRegistryForbidden{}
}

/*DeleteModelRegistryForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DeleteModelRegistryForbidden struct {
	Payload interface{}
}

func (o *DeleteModelRegistryForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/models/{uuid}][%d] deleteModelRegistryForbidden  %+v", 403, o.Payload)
}

func (o *DeleteModelRegistryForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteModelRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModelRegistryNotFound creates a DeleteModelRegistryNotFound with default headers values
func NewDeleteModelRegistryNotFound() *DeleteModelRegistryNotFound {
	return &DeleteModelRegistryNotFound{}
}

/*DeleteModelRegistryNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteModelRegistryNotFound struct {
	Payload interface{}
}

func (o *DeleteModelRegistryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/models/{uuid}][%d] deleteModelRegistryNotFound  %+v", 404, o.Payload)
}

func (o *DeleteModelRegistryNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteModelRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModelRegistryDefault creates a DeleteModelRegistryDefault with default headers values
func NewDeleteModelRegistryDefault(code int) *DeleteModelRegistryDefault {
	return &DeleteModelRegistryDefault{
		_statusCode: code,
	}
}

/*DeleteModelRegistryDefault handles this case with default header values.

An unexpected error response
*/
type DeleteModelRegistryDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the delete model registry default response
func (o *DeleteModelRegistryDefault) Code() int {
	return o._statusCode
}

func (o *DeleteModelRegistryDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/models/{uuid}][%d] DeleteModelRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteModelRegistryDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteModelRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
