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

// PatchModelRegistryReader is a Reader for the PatchModelRegistry structure.
type PatchModelRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchModelRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchModelRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchModelRegistryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchModelRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchModelRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchModelRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchModelRegistryOK creates a PatchModelRegistryOK with default headers values
func NewPatchModelRegistryOK() *PatchModelRegistryOK {
	return &PatchModelRegistryOK{}
}

/*PatchModelRegistryOK handles this case with default header values.

A successful response.
*/
type PatchModelRegistryOK struct {
	Payload *service_model.V1ModelRegistry
}

func (o *PatchModelRegistryOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry/{model.name}][%d] patchModelRegistryOK  %+v", 200, o.Payload)
}

func (o *PatchModelRegistryOK) GetPayload() *service_model.V1ModelRegistry {
	return o.Payload
}

func (o *PatchModelRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ModelRegistry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchModelRegistryNoContent creates a PatchModelRegistryNoContent with default headers values
func NewPatchModelRegistryNoContent() *PatchModelRegistryNoContent {
	return &PatchModelRegistryNoContent{}
}

/*PatchModelRegistryNoContent handles this case with default header values.

No content.
*/
type PatchModelRegistryNoContent struct {
	Payload interface{}
}

func (o *PatchModelRegistryNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry/{model.name}][%d] patchModelRegistryNoContent  %+v", 204, o.Payload)
}

func (o *PatchModelRegistryNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchModelRegistryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchModelRegistryForbidden creates a PatchModelRegistryForbidden with default headers values
func NewPatchModelRegistryForbidden() *PatchModelRegistryForbidden {
	return &PatchModelRegistryForbidden{}
}

/*PatchModelRegistryForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PatchModelRegistryForbidden struct {
	Payload interface{}
}

func (o *PatchModelRegistryForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry/{model.name}][%d] patchModelRegistryForbidden  %+v", 403, o.Payload)
}

func (o *PatchModelRegistryForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchModelRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchModelRegistryNotFound creates a PatchModelRegistryNotFound with default headers values
func NewPatchModelRegistryNotFound() *PatchModelRegistryNotFound {
	return &PatchModelRegistryNotFound{}
}

/*PatchModelRegistryNotFound handles this case with default header values.

Resource does not exist.
*/
type PatchModelRegistryNotFound struct {
	Payload interface{}
}

func (o *PatchModelRegistryNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry/{model.name}][%d] patchModelRegistryNotFound  %+v", 404, o.Payload)
}

func (o *PatchModelRegistryNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchModelRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchModelRegistryDefault creates a PatchModelRegistryDefault with default headers values
func NewPatchModelRegistryDefault(code int) *PatchModelRegistryDefault {
	return &PatchModelRegistryDefault{
		_statusCode: code,
	}
}

/*PatchModelRegistryDefault handles this case with default header values.

An unexpected error response.
*/
type PatchModelRegistryDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the patch model registry default response
func (o *PatchModelRegistryDefault) Code() int {
	return o._statusCode
}

func (o *PatchModelRegistryDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry/{model.name}][%d] PatchModelRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *PatchModelRegistryDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchModelRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
