// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FunctionalCapabilityResult functional capability result
// swagger:model FunctionalCapabilityResult
type FunctionalCapabilityResult struct {

	// response
	Response *FunctionalCapabilityResultResponse `json:"response,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this functional capability result
func (m *FunctionalCapabilityResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FunctionalCapabilityResult) validateResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FunctionalCapabilityResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionalCapabilityResult) UnmarshalBinary(b []byte) error {
	var res FunctionalCapabilityResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FunctionalCapabilityResultResponse functional capability result response
// swagger:model FunctionalCapabilityResultResponse
type FunctionalCapabilityResultResponse struct {

	// attribute info
	AttributeInfo interface{} `json:"attributeInfo,omitempty"`

	// function details
	FunctionDetails []*FunctionalCapabilityResultResponseFunctionDetailsItems0 `json:"functionDetails"`

	// function name
	FunctionName string `json:"functionName,omitempty"`

	// function op state
	// Enum: [UNKNOWN NOT_APPLICABLE DISABLED ENABLED]
	FunctionOpState string `json:"functionOpState,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this functional capability result response
func (m *FunctionalCapabilityResultResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunctionDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunctionOpState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FunctionalCapabilityResultResponse) validateFunctionDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.FunctionDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.FunctionDetails); i++ {
		if swag.IsZero(m.FunctionDetails[i]) { // not required
			continue
		}

		if m.FunctionDetails[i] != nil {
			if err := m.FunctionDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("response" + "." + "functionDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var functionalCapabilityResultResponseTypeFunctionOpStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","NOT_APPLICABLE","DISABLED","ENABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		functionalCapabilityResultResponseTypeFunctionOpStatePropEnum = append(functionalCapabilityResultResponseTypeFunctionOpStatePropEnum, v)
	}
}

const (

	// FunctionalCapabilityResultResponseFunctionOpStateUNKNOWN captures enum value "UNKNOWN"
	FunctionalCapabilityResultResponseFunctionOpStateUNKNOWN string = "UNKNOWN"

	// FunctionalCapabilityResultResponseFunctionOpStateNOTAPPLICABLE captures enum value "NOT_APPLICABLE"
	FunctionalCapabilityResultResponseFunctionOpStateNOTAPPLICABLE string = "NOT_APPLICABLE"

	// FunctionalCapabilityResultResponseFunctionOpStateDISABLED captures enum value "DISABLED"
	FunctionalCapabilityResultResponseFunctionOpStateDISABLED string = "DISABLED"

	// FunctionalCapabilityResultResponseFunctionOpStateENABLED captures enum value "ENABLED"
	FunctionalCapabilityResultResponseFunctionOpStateENABLED string = "ENABLED"
)

// prop value enum
func (m *FunctionalCapabilityResultResponse) validateFunctionOpStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, functionalCapabilityResultResponseTypeFunctionOpStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FunctionalCapabilityResultResponse) validateFunctionOpState(formats strfmt.Registry) error {

	if swag.IsZero(m.FunctionOpState) { // not required
		return nil
	}

	// value enum
	if err := m.validateFunctionOpStateEnum("response"+"."+"functionOpState", "body", m.FunctionOpState); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FunctionalCapabilityResultResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionalCapabilityResultResponse) UnmarshalBinary(b []byte) error {
	var res FunctionalCapabilityResultResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FunctionalCapabilityResultResponseFunctionDetailsItems0 functional capability result response function details items0
// swagger:model FunctionalCapabilityResultResponseFunctionDetailsItems0
type FunctionalCapabilityResultResponseFunctionDetailsItems0 struct {

	// attribute info
	AttributeInfo interface{} `json:"attributeInfo,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// property name
	PropertyName string `json:"propertyName,omitempty"`

	// string value
	StringValue string `json:"stringValue,omitempty"`
}

// Validate validates this functional capability result response function details items0
func (m *FunctionalCapabilityResultResponseFunctionDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FunctionalCapabilityResultResponseFunctionDetailsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionalCapabilityResultResponseFunctionDetailsItems0) UnmarshalBinary(b []byte) error {
	var res FunctionalCapabilityResultResponseFunctionDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
