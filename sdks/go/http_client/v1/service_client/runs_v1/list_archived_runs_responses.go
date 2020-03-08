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

// ListArchivedRunsReader is a Reader for the ListArchivedRuns structure.
type ListArchivedRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListArchivedRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListArchivedRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListArchivedRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListArchivedRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListArchivedRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListArchivedRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListArchivedRunsOK creates a ListArchivedRunsOK with default headers values
func NewListArchivedRunsOK() *ListArchivedRunsOK {
	return &ListArchivedRunsOK{}
}

/*ListArchivedRunsOK handles this case with default header values.

A successful response.
*/
type ListArchivedRunsOK struct {
	Payload *service_model.V1ListRunsResponse
}

func (o *ListArchivedRunsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] listArchivedRunsOK  %+v", 200, o.Payload)
}

func (o *ListArchivedRunsOK) GetPayload() *service_model.V1ListRunsResponse {
	return o.Payload
}

func (o *ListArchivedRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedRunsNoContent creates a ListArchivedRunsNoContent with default headers values
func NewListArchivedRunsNoContent() *ListArchivedRunsNoContent {
	return &ListArchivedRunsNoContent{}
}

/*ListArchivedRunsNoContent handles this case with default header values.

No content.
*/
type ListArchivedRunsNoContent struct {
	Payload interface{}
}

func (o *ListArchivedRunsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] listArchivedRunsNoContent  %+v", 204, o.Payload)
}

func (o *ListArchivedRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListArchivedRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedRunsForbidden creates a ListArchivedRunsForbidden with default headers values
func NewListArchivedRunsForbidden() *ListArchivedRunsForbidden {
	return &ListArchivedRunsForbidden{}
}

/*ListArchivedRunsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListArchivedRunsForbidden struct {
	Payload interface{}
}

func (o *ListArchivedRunsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] listArchivedRunsForbidden  %+v", 403, o.Payload)
}

func (o *ListArchivedRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListArchivedRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedRunsNotFound creates a ListArchivedRunsNotFound with default headers values
func NewListArchivedRunsNotFound() *ListArchivedRunsNotFound {
	return &ListArchivedRunsNotFound{}
}

/*ListArchivedRunsNotFound handles this case with default header values.

Resource does not exist.
*/
type ListArchivedRunsNotFound struct {
	Payload interface{}
}

func (o *ListArchivedRunsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] listArchivedRunsNotFound  %+v", 404, o.Payload)
}

func (o *ListArchivedRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListArchivedRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedRunsDefault creates a ListArchivedRunsDefault with default headers values
func NewListArchivedRunsDefault(code int) *ListArchivedRunsDefault {
	return &ListArchivedRunsDefault{
		_statusCode: code,
	}
}

/*ListArchivedRunsDefault handles this case with default header values.

An unexpected error response
*/
type ListArchivedRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list archived runs default response
func (o *ListArchivedRunsDefault) Code() int {
	return o._statusCode
}

func (o *ListArchivedRunsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] ListArchivedRuns default  %+v", o._statusCode, o.Payload)
}

func (o *ListArchivedRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListArchivedRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
