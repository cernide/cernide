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

// ListRunsReader is a Reader for the ListRuns structure.
type ListRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRunsOK creates a ListRunsOK with default headers values
func NewListRunsOK() *ListRunsOK {
	return &ListRunsOK{}
}

/*ListRunsOK handles this case with default header values.

A successful response.
*/
type ListRunsOK struct {
	Payload *service_model.V1ListRunsResponse
}

func (o *ListRunsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs][%d] listRunsOK  %+v", 200, o.Payload)
}

func (o *ListRunsOK) GetPayload() *service_model.V1ListRunsResponse {
	return o.Payload
}

func (o *ListRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRunsNoContent creates a ListRunsNoContent with default headers values
func NewListRunsNoContent() *ListRunsNoContent {
	return &ListRunsNoContent{}
}

/*ListRunsNoContent handles this case with default header values.

No content.
*/
type ListRunsNoContent struct {
	Payload interface{}
}

func (o *ListRunsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs][%d] listRunsNoContent  %+v", 204, o.Payload)
}

func (o *ListRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRunsForbidden creates a ListRunsForbidden with default headers values
func NewListRunsForbidden() *ListRunsForbidden {
	return &ListRunsForbidden{}
}

/*ListRunsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListRunsForbidden struct {
	Payload interface{}
}

func (o *ListRunsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs][%d] listRunsForbidden  %+v", 403, o.Payload)
}

func (o *ListRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRunsNotFound creates a ListRunsNotFound with default headers values
func NewListRunsNotFound() *ListRunsNotFound {
	return &ListRunsNotFound{}
}

/*ListRunsNotFound handles this case with default header values.

Resource does not exist.
*/
type ListRunsNotFound struct {
	Payload interface{}
}

func (o *ListRunsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs][%d] listRunsNotFound  %+v", 404, o.Payload)
}

func (o *ListRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRunsDefault creates a ListRunsDefault with default headers values
func NewListRunsDefault(code int) *ListRunsDefault {
	return &ListRunsDefault{
		_statusCode: code,
	}
}

/*ListRunsDefault handles this case with default header values.

An unexpected error response
*/
type ListRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list runs default response
func (o *ListRunsDefault) Code() int {
	return o._statusCode
}

func (o *ListRunsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs][%d] ListRuns default  %+v", o._statusCode, o.Payload)
}

func (o *ListRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
