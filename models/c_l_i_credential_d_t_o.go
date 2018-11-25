// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CLICredentialDTO c l i credential d t o
// swagger:model CLICredentialDTO
type CLICredentialDTO struct {

	// comments
	Comments string `json:"comments,omitempty"`

	// credential type
	// Enum: [GLOBAL APP]
	CredentialType string `json:"credentialType,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enable password
	EnablePassword string `json:"enablePassword,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// instance tenant Id
	InstanceTenantID string `json:"instanceTenantId,omitempty"`

	// instance Uuid
	InstanceUUID string `json:"instanceUuid,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this c l i credential d t o
func (m *CLICredentialDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cLICredentialDTOTypeCredentialTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GLOBAL","APP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cLICredentialDTOTypeCredentialTypePropEnum = append(cLICredentialDTOTypeCredentialTypePropEnum, v)
	}
}

const (

	// CLICredentialDTOCredentialTypeGLOBAL captures enum value "GLOBAL"
	CLICredentialDTOCredentialTypeGLOBAL string = "GLOBAL"

	// CLICredentialDTOCredentialTypeAPP captures enum value "APP"
	CLICredentialDTOCredentialTypeAPP string = "APP"
)

// prop value enum
func (m *CLICredentialDTO) validateCredentialTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cLICredentialDTOTypeCredentialTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CLICredentialDTO) validateCredentialType(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCredentialTypeEnum("credentialType", "body", m.CredentialType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CLICredentialDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CLICredentialDTO) UnmarshalBinary(b []byte) error {
	var res CLICredentialDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
