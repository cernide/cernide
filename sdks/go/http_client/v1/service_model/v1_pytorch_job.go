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

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PytorchJob Pytorch Job specification
//
// swagger:model v1PytorchJob
type V1PytorchJob struct {

	// optional clean pod policy section
	CleanPodPolicy V1CleanPodPolicy `json:"cleanPodPolicy,omitempty"`

	// Optional component kind, should be equal to 'pytorch_job'
	Kind *string `json:"kind,omitempty"`

	// Optional master replica definition
	Master *V1KFReplica `json:"master,omitempty"`

	// Optional worker replica definition
	Worker *V1KFReplica `json:"worker,omitempty"`
}

// Validate validates this v1 pytorch job
func (m *V1PytorchJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCleanPodPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorker(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PytorchJob) validateCleanPodPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.CleanPodPolicy) { // not required
		return nil
	}

	if err := m.CleanPodPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cleanPodPolicy")
		}
		return err
	}

	return nil
}

func (m *V1PytorchJob) validateMaster(formats strfmt.Registry) error {

	if swag.IsZero(m.Master) { // not required
		return nil
	}

	if m.Master != nil {
		if err := m.Master.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master")
			}
			return err
		}
	}

	return nil
}

func (m *V1PytorchJob) validateWorker(formats strfmt.Registry) error {

	if swag.IsZero(m.Worker) { // not required
		return nil
	}

	if m.Worker != nil {
		if err := m.Worker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("worker")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PytorchJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PytorchJob) UnmarshalBinary(b []byte) error {
	var res V1PytorchJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
