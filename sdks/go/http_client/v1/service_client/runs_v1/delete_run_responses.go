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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// DeleteRunReader is a Reader for the DeleteRun structure.
type DeleteRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRunOK creates a DeleteRunOK with default headers values
func NewDeleteRunOK() *DeleteRunOK {
	return &DeleteRunOK{}
}

/*DeleteRunOK handles this case with default header values.

A successful response.
*/
type DeleteRunOK struct {
}

func (o *DeleteRunOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/{uuid}][%d] deleteRunOK ", 200)
}

func (o *DeleteRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRunNoContent creates a DeleteRunNoContent with default headers values
func NewDeleteRunNoContent() *DeleteRunNoContent {
	return &DeleteRunNoContent{}
}

/*DeleteRunNoContent handles this case with default header values.

No content.
*/
type DeleteRunNoContent struct {
	Payload interface{}
}

func (o *DeleteRunNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/{uuid}][%d] deleteRunNoContent  %+v", 204, o.Payload)
}

func (o *DeleteRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunForbidden creates a DeleteRunForbidden with default headers values
func NewDeleteRunForbidden() *DeleteRunForbidden {
	return &DeleteRunForbidden{}
}

/*DeleteRunForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DeleteRunForbidden struct {
	Payload interface{}
}

func (o *DeleteRunForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/{uuid}][%d] deleteRunForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunNotFound creates a DeleteRunNotFound with default headers values
func NewDeleteRunNotFound() *DeleteRunNotFound {
	return &DeleteRunNotFound{}
}

/*DeleteRunNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteRunNotFound struct {
	Payload interface{}
}

func (o *DeleteRunNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/{uuid}][%d] deleteRunNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunDefault creates a DeleteRunDefault with default headers values
func NewDeleteRunDefault(code int) *DeleteRunDefault {
	return &DeleteRunDefault{
		_statusCode: code,
	}
}

/*DeleteRunDefault handles this case with default header values.

An unexpected error response
*/
type DeleteRunDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the delete run default response
func (o *DeleteRunDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRunDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/{uuid}][%d] DeleteRun default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRunDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
