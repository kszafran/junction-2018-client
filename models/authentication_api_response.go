// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AuthenticationAPIResponse authentication API response
// swagger:model AuthenticationAPIResponse
type AuthenticationAPIResponse struct {

	// token
	Token string `json:"Token,omitempty"`
}

// Validate validates this authentication API response
func (m *AuthenticationAPIResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthenticationAPIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthenticationAPIResponse) UnmarshalBinary(b []byte) error {
	var res AuthenticationAPIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
