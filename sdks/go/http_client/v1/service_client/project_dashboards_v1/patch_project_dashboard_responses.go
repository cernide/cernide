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

// PatchProjectDashboardReader is a Reader for the PatchProjectDashboard structure.
type PatchProjectDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchProjectDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchProjectDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchProjectDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchProjectDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchProjectDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchProjectDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchProjectDashboardOK creates a PatchProjectDashboardOK with default headers values
func NewPatchProjectDashboardOK() *PatchProjectDashboardOK {
	return &PatchProjectDashboardOK{}
}

/*PatchProjectDashboardOK handles this case with default header values.

A successful response.
*/
type PatchProjectDashboardOK struct {
	Payload *service_model.V1Dashboard
}

func (o *PatchProjectDashboardOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] patchProjectDashboardOK  %+v", 200, o.Payload)
}

func (o *PatchProjectDashboardOK) GetPayload() *service_model.V1Dashboard {
	return o.Payload
}

func (o *PatchProjectDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProjectDashboardNoContent creates a PatchProjectDashboardNoContent with default headers values
func NewPatchProjectDashboardNoContent() *PatchProjectDashboardNoContent {
	return &PatchProjectDashboardNoContent{}
}

/*PatchProjectDashboardNoContent handles this case with default header values.

No content.
*/
type PatchProjectDashboardNoContent struct {
	Payload interface{}
}

func (o *PatchProjectDashboardNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] patchProjectDashboardNoContent  %+v", 204, o.Payload)
}

func (o *PatchProjectDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchProjectDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProjectDashboardForbidden creates a PatchProjectDashboardForbidden with default headers values
func NewPatchProjectDashboardForbidden() *PatchProjectDashboardForbidden {
	return &PatchProjectDashboardForbidden{}
}

/*PatchProjectDashboardForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PatchProjectDashboardForbidden struct {
	Payload interface{}
}

func (o *PatchProjectDashboardForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] patchProjectDashboardForbidden  %+v", 403, o.Payload)
}

func (o *PatchProjectDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchProjectDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProjectDashboardNotFound creates a PatchProjectDashboardNotFound with default headers values
func NewPatchProjectDashboardNotFound() *PatchProjectDashboardNotFound {
	return &PatchProjectDashboardNotFound{}
}

/*PatchProjectDashboardNotFound handles this case with default header values.

Resource does not exist.
*/
type PatchProjectDashboardNotFound struct {
	Payload interface{}
}

func (o *PatchProjectDashboardNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] patchProjectDashboardNotFound  %+v", 404, o.Payload)
}

func (o *PatchProjectDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchProjectDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProjectDashboardDefault creates a PatchProjectDashboardDefault with default headers values
func NewPatchProjectDashboardDefault(code int) *PatchProjectDashboardDefault {
	return &PatchProjectDashboardDefault{
		_statusCode: code,
	}
}

/*PatchProjectDashboardDefault handles this case with default header values.

An unexpected error response
*/
type PatchProjectDashboardDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the patch project dashboard default response
func (o *PatchProjectDashboardDefault) Code() int {
	return o._statusCode
}

func (o *PatchProjectDashboardDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] PatchProjectDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *PatchProjectDashboardDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchProjectDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
