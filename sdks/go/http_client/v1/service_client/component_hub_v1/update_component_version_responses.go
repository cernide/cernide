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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateComponentVersionReader is a Reader for the UpdateComponentVersion structure.
type UpdateComponentVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateComponentVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateComponentVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateComponentVersionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateComponentVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateComponentVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateComponentVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateComponentVersionOK creates a UpdateComponentVersionOK with default headers values
func NewUpdateComponentVersionOK() *UpdateComponentVersionOK {
	return &UpdateComponentVersionOK{}
}

/*UpdateComponentVersionOK handles this case with default header values.

A successful response.
*/
type UpdateComponentVersionOK struct {
	Payload *service_model.V1ComponentVersion
}

func (o *UpdateComponentVersionOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/hub/{component}/versions/{version.name}][%d] updateComponentVersionOK  %+v", 200, o.Payload)
}

func (o *UpdateComponentVersionOK) GetPayload() *service_model.V1ComponentVersion {
	return o.Payload
}

func (o *UpdateComponentVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ComponentVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateComponentVersionNoContent creates a UpdateComponentVersionNoContent with default headers values
func NewUpdateComponentVersionNoContent() *UpdateComponentVersionNoContent {
	return &UpdateComponentVersionNoContent{}
}

/*UpdateComponentVersionNoContent handles this case with default header values.

No content.
*/
type UpdateComponentVersionNoContent struct {
	Payload interface{}
}

func (o *UpdateComponentVersionNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/hub/{component}/versions/{version.name}][%d] updateComponentVersionNoContent  %+v", 204, o.Payload)
}

func (o *UpdateComponentVersionNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateComponentVersionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateComponentVersionForbidden creates a UpdateComponentVersionForbidden with default headers values
func NewUpdateComponentVersionForbidden() *UpdateComponentVersionForbidden {
	return &UpdateComponentVersionForbidden{}
}

/*UpdateComponentVersionForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UpdateComponentVersionForbidden struct {
	Payload interface{}
}

func (o *UpdateComponentVersionForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/hub/{component}/versions/{version.name}][%d] updateComponentVersionForbidden  %+v", 403, o.Payload)
}

func (o *UpdateComponentVersionForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateComponentVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateComponentVersionNotFound creates a UpdateComponentVersionNotFound with default headers values
func NewUpdateComponentVersionNotFound() *UpdateComponentVersionNotFound {
	return &UpdateComponentVersionNotFound{}
}

/*UpdateComponentVersionNotFound handles this case with default header values.

Resource does not exist.
*/
type UpdateComponentVersionNotFound struct {
	Payload interface{}
}

func (o *UpdateComponentVersionNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/hub/{component}/versions/{version.name}][%d] updateComponentVersionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateComponentVersionNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateComponentVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateComponentVersionDefault creates a UpdateComponentVersionDefault with default headers values
func NewUpdateComponentVersionDefault(code int) *UpdateComponentVersionDefault {
	return &UpdateComponentVersionDefault{
		_statusCode: code,
	}
}

/*UpdateComponentVersionDefault handles this case with default header values.

An unexpected error response.
*/
type UpdateComponentVersionDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the update component version default response
func (o *UpdateComponentVersionDefault) Code() int {
	return o._statusCode
}

func (o *UpdateComponentVersionDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/hub/{component}/versions/{version.name}][%d] UpdateComponentVersion default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateComponentVersionDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateComponentVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
