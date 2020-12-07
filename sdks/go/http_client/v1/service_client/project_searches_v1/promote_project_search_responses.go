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

package project_searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// PromoteProjectSearchReader is a Reader for the PromoteProjectSearch structure.
type PromoteProjectSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PromoteProjectSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPromoteProjectSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPromoteProjectSearchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPromoteProjectSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPromoteProjectSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPromoteProjectSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPromoteProjectSearchOK creates a PromoteProjectSearchOK with default headers values
func NewPromoteProjectSearchOK() *PromoteProjectSearchOK {
	return &PromoteProjectSearchOK{}
}

/*PromoteProjectSearchOK handles this case with default header values.

A successful response.
*/
type PromoteProjectSearchOK struct {
}

func (o *PromoteProjectSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/searches/{uuid}/promote][%d] promoteProjectSearchOK ", 200)
}

func (o *PromoteProjectSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPromoteProjectSearchNoContent creates a PromoteProjectSearchNoContent with default headers values
func NewPromoteProjectSearchNoContent() *PromoteProjectSearchNoContent {
	return &PromoteProjectSearchNoContent{}
}

/*PromoteProjectSearchNoContent handles this case with default header values.

No content.
*/
type PromoteProjectSearchNoContent struct {
	Payload interface{}
}

func (o *PromoteProjectSearchNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/searches/{uuid}/promote][%d] promoteProjectSearchNoContent  %+v", 204, o.Payload)
}

func (o *PromoteProjectSearchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PromoteProjectSearchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteProjectSearchForbidden creates a PromoteProjectSearchForbidden with default headers values
func NewPromoteProjectSearchForbidden() *PromoteProjectSearchForbidden {
	return &PromoteProjectSearchForbidden{}
}

/*PromoteProjectSearchForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PromoteProjectSearchForbidden struct {
	Payload interface{}
}

func (o *PromoteProjectSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/searches/{uuid}/promote][%d] promoteProjectSearchForbidden  %+v", 403, o.Payload)
}

func (o *PromoteProjectSearchForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PromoteProjectSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteProjectSearchNotFound creates a PromoteProjectSearchNotFound with default headers values
func NewPromoteProjectSearchNotFound() *PromoteProjectSearchNotFound {
	return &PromoteProjectSearchNotFound{}
}

/*PromoteProjectSearchNotFound handles this case with default header values.

Resource does not exist.
*/
type PromoteProjectSearchNotFound struct {
	Payload interface{}
}

func (o *PromoteProjectSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/searches/{uuid}/promote][%d] promoteProjectSearchNotFound  %+v", 404, o.Payload)
}

func (o *PromoteProjectSearchNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PromoteProjectSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteProjectSearchDefault creates a PromoteProjectSearchDefault with default headers values
func NewPromoteProjectSearchDefault(code int) *PromoteProjectSearchDefault {
	return &PromoteProjectSearchDefault{
		_statusCode: code,
	}
}

/*PromoteProjectSearchDefault handles this case with default header values.

An unexpected error response.
*/
type PromoteProjectSearchDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the promote project search default response
func (o *PromoteProjectSearchDefault) Code() int {
	return o._statusCode
}

func (o *PromoteProjectSearchDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/searches/{uuid}/promote][%d] PromoteProjectSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PromoteProjectSearchDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PromoteProjectSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
