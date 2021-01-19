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

// TerseSummary terse summary
//
// swagger:model TerseSummary
type TerseSummary struct {

	// original URL
	OriginalURL string `json:"originalURL,omitempty"`

	// redirect type
	RedirectType RedirectType `json:"redirectType,omitempty"`

	// shortened URL
	ShortenedURL string `json:"shortenedURL,omitempty"`

	// visit count
	VisitCount int64 `json:"visitCount,omitempty"`
}

// Validate validates this terse summary
func (m *TerseSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRedirectType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerseSummary) validateRedirectType(formats strfmt.Registry) error {
	if swag.IsZero(m.RedirectType) { // not required
		return nil
	}

	if err := m.RedirectType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("redirectType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this terse summary based on the context it is used
func (m *TerseSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRedirectType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerseSummary) contextValidateRedirectType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RedirectType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("redirectType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerseSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerseSummary) UnmarshalBinary(b []byte) error {
	var res TerseSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
