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

package project_dashboards_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListProjectDashboardNamesReader is a Reader for the ListProjectDashboardNames structure.
type ListProjectDashboardNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectDashboardNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectDashboardNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListProjectDashboardNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListProjectDashboardNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectDashboardNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListProjectDashboardNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectDashboardNamesOK creates a ListProjectDashboardNamesOK with default headers values
func NewListProjectDashboardNamesOK() *ListProjectDashboardNamesOK {
	return &ListProjectDashboardNamesOK{}
}

/*ListProjectDashboardNamesOK handles this case with default header values.

A successful response.
*/
type ListProjectDashboardNamesOK struct {
	Payload *service_model.V1ListDashboardsResponse
}

func (o *ListProjectDashboardNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] listProjectDashboardNamesOK  %+v", 200, o.Payload)
}

func (o *ListProjectDashboardNamesOK) GetPayload() *service_model.V1ListDashboardsResponse {
	return o.Payload
}

func (o *ListProjectDashboardNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListDashboardsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDashboardNamesNoContent creates a ListProjectDashboardNamesNoContent with default headers values
func NewListProjectDashboardNamesNoContent() *ListProjectDashboardNamesNoContent {
	return &ListProjectDashboardNamesNoContent{}
}

/*ListProjectDashboardNamesNoContent handles this case with default header values.

No content.
*/
type ListProjectDashboardNamesNoContent struct {
	Payload interface{}
}

func (o *ListProjectDashboardNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] listProjectDashboardNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListProjectDashboardNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectDashboardNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDashboardNamesForbidden creates a ListProjectDashboardNamesForbidden with default headers values
func NewListProjectDashboardNamesForbidden() *ListProjectDashboardNamesForbidden {
	return &ListProjectDashboardNamesForbidden{}
}

/*ListProjectDashboardNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListProjectDashboardNamesForbidden struct {
	Payload interface{}
}

func (o *ListProjectDashboardNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] listProjectDashboardNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectDashboardNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectDashboardNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDashboardNamesNotFound creates a ListProjectDashboardNamesNotFound with default headers values
func NewListProjectDashboardNamesNotFound() *ListProjectDashboardNamesNotFound {
	return &ListProjectDashboardNamesNotFound{}
}

/*ListProjectDashboardNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListProjectDashboardNamesNotFound struct {
	Payload interface{}
}

func (o *ListProjectDashboardNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] listProjectDashboardNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectDashboardNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectDashboardNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDashboardNamesDefault creates a ListProjectDashboardNamesDefault with default headers values
func NewListProjectDashboardNamesDefault(code int) *ListProjectDashboardNamesDefault {
	return &ListProjectDashboardNamesDefault{
		_statusCode: code,
	}
}

/*ListProjectDashboardNamesDefault handles this case with default header values.

An unexpected error response.
*/
type ListProjectDashboardNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list project dashboard names default response
func (o *ListProjectDashboardNamesDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectDashboardNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] ListProjectDashboardNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectDashboardNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListProjectDashboardNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
