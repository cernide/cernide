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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListOrganizationQueuesReader is a Reader for the ListOrganizationQueues structure.
type ListOrganizationQueuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrganizationQueuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOrganizationQueuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListOrganizationQueuesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListOrganizationQueuesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListOrganizationQueuesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListOrganizationQueuesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOrganizationQueuesOK creates a ListOrganizationQueuesOK with default headers values
func NewListOrganizationQueuesOK() *ListOrganizationQueuesOK {
	return &ListOrganizationQueuesOK{}
}

/*ListOrganizationQueuesOK handles this case with default header values.

A successful response.
*/
type ListOrganizationQueuesOK struct {
	Payload *service_model.V1ListQueuesResponse
}

func (o *ListOrganizationQueuesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/queues][%d] listOrganizationQueuesOK  %+v", 200, o.Payload)
}

func (o *ListOrganizationQueuesOK) GetPayload() *service_model.V1ListQueuesResponse {
	return o.Payload
}

func (o *ListOrganizationQueuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListQueuesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationQueuesNoContent creates a ListOrganizationQueuesNoContent with default headers values
func NewListOrganizationQueuesNoContent() *ListOrganizationQueuesNoContent {
	return &ListOrganizationQueuesNoContent{}
}

/*ListOrganizationQueuesNoContent handles this case with default header values.

No content.
*/
type ListOrganizationQueuesNoContent struct {
	Payload interface{}
}

func (o *ListOrganizationQueuesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/queues][%d] listOrganizationQueuesNoContent  %+v", 204, o.Payload)
}

func (o *ListOrganizationQueuesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationQueuesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationQueuesForbidden creates a ListOrganizationQueuesForbidden with default headers values
func NewListOrganizationQueuesForbidden() *ListOrganizationQueuesForbidden {
	return &ListOrganizationQueuesForbidden{}
}

/*ListOrganizationQueuesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListOrganizationQueuesForbidden struct {
	Payload interface{}
}

func (o *ListOrganizationQueuesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/queues][%d] listOrganizationQueuesForbidden  %+v", 403, o.Payload)
}

func (o *ListOrganizationQueuesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationQueuesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationQueuesNotFound creates a ListOrganizationQueuesNotFound with default headers values
func NewListOrganizationQueuesNotFound() *ListOrganizationQueuesNotFound {
	return &ListOrganizationQueuesNotFound{}
}

/*ListOrganizationQueuesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListOrganizationQueuesNotFound struct {
	Payload interface{}
}

func (o *ListOrganizationQueuesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/queues][%d] listOrganizationQueuesNotFound  %+v", 404, o.Payload)
}

func (o *ListOrganizationQueuesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationQueuesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationQueuesDefault creates a ListOrganizationQueuesDefault with default headers values
func NewListOrganizationQueuesDefault(code int) *ListOrganizationQueuesDefault {
	return &ListOrganizationQueuesDefault{
		_statusCode: code,
	}
}

/*ListOrganizationQueuesDefault handles this case with default header values.

An unexpected error response
*/
type ListOrganizationQueuesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list organization queues default response
func (o *ListOrganizationQueuesDefault) Code() int {
	return o._statusCode
}

func (o *ListOrganizationQueuesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/queues][%d] ListOrganizationQueues default  %+v", o._statusCode, o.Payload)
}

func (o *ListOrganizationQueuesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListOrganizationQueuesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
