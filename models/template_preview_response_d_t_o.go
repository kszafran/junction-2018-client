// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TemplatePreviewResponseDTO template preview response d t o
// swagger:model TemplatePreviewResponseDTO
type TemplatePreviewResponseDTO struct {

	// cli preview
	CliPreview string `json:"cliPreview,omitempty"`

	// template Id
	TemplateID string `json:"templateId,omitempty"`

	// validation errors
	ValidationErrors interface{} `json:"validationErrors,omitempty"`
}

// Validate validates this template preview response d t o
func (m *TemplatePreviewResponseDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TemplatePreviewResponseDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplatePreviewResponseDTO) UnmarshalBinary(b []byte) error {
	var res TemplatePreviewResponseDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
