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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreateProjectReader is a Reader for the CreateProject structure.
type CreateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateProjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProjectOK creates a CreateProjectOK with default headers values
func NewCreateProjectOK() *CreateProjectOK {
	return &CreateProjectOK{}
}

/*CreateProjectOK handles this case with default header values.

A successful response.
*/
type CreateProjectOK struct {
	Payload *service_model.V1Project
}

func (o *CreateProjectOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/projects/create][%d] createProjectOK  %+v", 200, o.Payload)
}

func (o *CreateProjectOK) GetPayload() *service_model.V1Project {
	return o.Payload
}

func (o *CreateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectNoContent creates a CreateProjectNoContent with default headers values
func NewCreateProjectNoContent() *CreateProjectNoContent {
	return &CreateProjectNoContent{}
}

/*CreateProjectNoContent handles this case with default header values.

No content.
*/
type CreateProjectNoContent struct {
	Payload interface{}
}

func (o *CreateProjectNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/projects/create][%d] createProjectNoContent  %+v", 204, o.Payload)
}

func (o *CreateProjectNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateProjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectForbidden creates a CreateProjectForbidden with default headers values
func NewCreateProjectForbidden() *CreateProjectForbidden {
	return &CreateProjectForbidden{}
}

/*CreateProjectForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type CreateProjectForbidden struct {
	Payload interface{}
}

func (o *CreateProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/projects/create][%d] createProjectForbidden  %+v", 403, o.Payload)
}

func (o *CreateProjectForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectNotFound creates a CreateProjectNotFound with default headers values
func NewCreateProjectNotFound() *CreateProjectNotFound {
	return &CreateProjectNotFound{}
}

/*CreateProjectNotFound handles this case with default header values.

Resource does not exist.
*/
type CreateProjectNotFound struct {
	Payload interface{}
}

func (o *CreateProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/projects/create][%d] createProjectNotFound  %+v", 404, o.Payload)
}

func (o *CreateProjectNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectDefault creates a CreateProjectDefault with default headers values
func NewCreateProjectDefault(code int) *CreateProjectDefault {
	return &CreateProjectDefault{
		_statusCode: code,
	}
}

/*CreateProjectDefault handles this case with default header values.

An unexpected error response
*/
type CreateProjectDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the create project default response
func (o *CreateProjectDefault) Code() int {
	return o._statusCode
}

func (o *CreateProjectDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/projects/create][%d] CreateProject default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProjectDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
