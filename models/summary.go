// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Summary summary
//
// swagger:model Summary
type Summary struct {

	// terse
	Terse *TerseSummary `json:"terse,omitempty"`

	// visits
	Visits *VisitsSummary `json:"visits,omitempty"`
}

// Validate validates this summary
func (m *Summary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTerse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Summary) validateTerse(formats strfmt.Registry) error {
	if swag.IsZero(m.Terse) { // not required
		return nil
	}

	if m.Terse != nil {
		if err := m.Terse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terse")
			}
			return err
		}
	}

	return nil
}

func (m *Summary) validateVisits(formats strfmt.Registry) error {
	if swag.IsZero(m.Visits) { // not required
		return nil
	}

	if m.Visits != nil {
		if err := m.Visits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visits")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this summary based on the context it is used
func (m *Summary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTerse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVisits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Summary) contextValidateTerse(ctx context.Context, formats strfmt.Registry) error {

	if m.Terse != nil {
		if err := m.Terse.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terse")
			}
			return err
		}
	}

	return nil
}

func (m *Summary) contextValidateVisits(ctx context.Context, formats strfmt.Registry) error {

	if m.Visits != nil {
		if err := m.Visits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Summary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Summary) UnmarshalBinary(b []byte) error {
	var res Summary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
