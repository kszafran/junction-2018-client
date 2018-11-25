// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GlobalCredentialSubTypeResult global credential sub type result
// swagger:model GlobalCredentialSubTypeResult
type GlobalCredentialSubTypeResult struct {

	// response
	Response string `json:"response,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this global credential sub type result
func (m *GlobalCredentialSubTypeResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GlobalCredentialSubTypeResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalCredentialSubTypeResult) UnmarshalBinary(b []byte) error {
	var res GlobalCredentialSubTypeResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
