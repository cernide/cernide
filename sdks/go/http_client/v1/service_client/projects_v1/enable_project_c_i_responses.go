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

// EnableProjectCIReader is a Reader for the EnableProjectCI structure.
type EnableProjectCIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableProjectCIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableProjectCIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewEnableProjectCINoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewEnableProjectCIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnableProjectCINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEnableProjectCIDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnableProjectCIOK creates a EnableProjectCIOK with default headers values
func NewEnableProjectCIOK() *EnableProjectCIOK {
	return &EnableProjectCIOK{}
}

/*EnableProjectCIOK handles this case with default header values.

A successful response.
*/
type EnableProjectCIOK struct {
}

func (o *EnableProjectCIOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/ci][%d] enableProjectCIOK ", 200)
}

func (o *EnableProjectCIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableProjectCINoContent creates a EnableProjectCINoContent with default headers values
func NewEnableProjectCINoContent() *EnableProjectCINoContent {
	return &EnableProjectCINoContent{}
}

/*EnableProjectCINoContent handles this case with default header values.

No content.
*/
type EnableProjectCINoContent struct {
	Payload interface{}
}

func (o *EnableProjectCINoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/ci][%d] enableProjectCINoContent  %+v", 204, o.Payload)
}

func (o *EnableProjectCINoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *EnableProjectCINoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableProjectCIForbidden creates a EnableProjectCIForbidden with default headers values
func NewEnableProjectCIForbidden() *EnableProjectCIForbidden {
	return &EnableProjectCIForbidden{}
}

/*EnableProjectCIForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type EnableProjectCIForbidden struct {
	Payload interface{}
}

func (o *EnableProjectCIForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/ci][%d] enableProjectCIForbidden  %+v", 403, o.Payload)
}

func (o *EnableProjectCIForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *EnableProjectCIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableProjectCINotFound creates a EnableProjectCINotFound with default headers values
func NewEnableProjectCINotFound() *EnableProjectCINotFound {
	return &EnableProjectCINotFound{}
}

/*EnableProjectCINotFound handles this case with default header values.

Resource does not exist.
*/
type EnableProjectCINotFound struct {
	Payload interface{}
}

func (o *EnableProjectCINotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/ci][%d] enableProjectCINotFound  %+v", 404, o.Payload)
}

func (o *EnableProjectCINotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *EnableProjectCINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableProjectCIDefault creates a EnableProjectCIDefault with default headers values
func NewEnableProjectCIDefault(code int) *EnableProjectCIDefault {
	return &EnableProjectCIDefault{
		_statusCode: code,
	}
}

/*EnableProjectCIDefault handles this case with default header values.

An unexpected error response.
*/
type EnableProjectCIDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the enable project c i default response
func (o *EnableProjectCIDefault) Code() int {
	return o._statusCode
}

func (o *EnableProjectCIDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/ci][%d] EnableProjectCI default  %+v", o._statusCode, o.Payload)
}

func (o *EnableProjectCIDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *EnableProjectCIDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
