// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IdentityRange Status of identity range of the cluster
//
// swagger:model IdentityRange
type IdentityRange struct {

	// Maximum identity of the cluster
	MaxIdentity int64 `json:"max-identity,omitempty"`

	// Minimum identity of the cluster
	MinIdentity int64 `json:"min-identity,omitempty"`
}

// Validate validates this identity range
func (m *IdentityRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this identity range based on context it is used
func (m *IdentityRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IdentityRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityRange) UnmarshalBinary(b []byte) error {
	var res IdentityRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
