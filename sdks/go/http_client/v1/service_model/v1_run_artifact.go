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

// V1RunArtifact Run Artifact spec definition
//
// swagger:model v1RunArtifact
type V1RunArtifact struct {

	// Connection
	Connection string `json:"connection,omitempty"`

	// Optional flag to check the use of the artifact in a context
	IsInput bool `json:"is_input,omitempty"`

	// Artifact type
	Kind V1ArtifactKind `json:"kind,omitempty"`

	// Artifact name
	Name string `json:"name,omitempty"`

	// Artifact path
	Path string `json:"path,omitempty"`

	// State
	State string `json:"state,omitempty"`

	// Artifact schema
	Summary interface{} `json:"summary,omitempty"`
}

// Validate validates this v1 run artifact
func (m *V1RunArtifact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RunArtifact) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if err := m.Kind.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("kind")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RunArtifact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RunArtifact) UnmarshalBinary(b []byte) error {
	var res V1RunArtifact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
