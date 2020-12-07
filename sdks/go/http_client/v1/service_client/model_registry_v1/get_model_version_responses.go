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

// GetModelVersionReader is a Reader for the GetModelVersion structure.
type GetModelVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModelVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetModelVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetModelVersionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetModelVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetModelVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetModelVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetModelVersionOK creates a GetModelVersionOK with default headers values
func NewGetModelVersionOK() *GetModelVersionOK {
	return &GetModelVersionOK{}
}

/*GetModelVersionOK handles this case with default header values.

A successful response.
*/
type GetModelVersionOK struct {
	Payload *service_model.V1ModelVersion
}

func (o *GetModelVersionOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry/{entity}/versions/{name}][%d] getModelVersionOK  %+v", 200, o.Payload)
}

func (o *GetModelVersionOK) GetPayload() *service_model.V1ModelVersion {
	return o.Payload
}

func (o *GetModelVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ModelVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelVersionNoContent creates a GetModelVersionNoContent with default headers values
func NewGetModelVersionNoContent() *GetModelVersionNoContent {
	return &GetModelVersionNoContent{}
}

/*GetModelVersionNoContent handles this case with default header values.

No content.
*/
type GetModelVersionNoContent struct {
	Payload interface{}
}

func (o *GetModelVersionNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry/{entity}/versions/{name}][%d] getModelVersionNoContent  %+v", 204, o.Payload)
}

func (o *GetModelVersionNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetModelVersionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelVersionForbidden creates a GetModelVersionForbidden with default headers values
func NewGetModelVersionForbidden() *GetModelVersionForbidden {
	return &GetModelVersionForbidden{}
}

/*GetModelVersionForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetModelVersionForbidden struct {
	Payload interface{}
}

func (o *GetModelVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry/{entity}/versions/{name}][%d] getModelVersionForbidden  %+v", 403, o.Payload)
}

func (o *GetModelVersionForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetModelVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelVersionNotFound creates a GetModelVersionNotFound with default headers values
func NewGetModelVersionNotFound() *GetModelVersionNotFound {
	return &GetModelVersionNotFound{}
}

/*GetModelVersionNotFound handles this case with default header values.

Resource does not exist.
*/
type GetModelVersionNotFound struct {
	Payload interface{}
}

func (o *GetModelVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry/{entity}/versions/{name}][%d] getModelVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetModelVersionNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetModelVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelVersionDefault creates a GetModelVersionDefault with default headers values
func NewGetModelVersionDefault(code int) *GetModelVersionDefault {
	return &GetModelVersionDefault{
		_statusCode: code,
	}
}

/*GetModelVersionDefault handles this case with default header values.

An unexpected error response.
*/
type GetModelVersionDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get model version default response
func (o *GetModelVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetModelVersionDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry/{entity}/versions/{name}][%d] GetModelVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetModelVersionDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetModelVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
