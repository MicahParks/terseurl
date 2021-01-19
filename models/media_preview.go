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

// MediaPreview media preview
//
// swagger:model MediaPreview
type MediaPreview struct {

	// og
	Og OpenGraph `json:"og,omitempty"`

	// redirect type
	RedirectType RedirectType `json:"redirectType,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// twitter
	Twitter Twitter `json:"twitter,omitempty"`
}

// Validate validates this media preview
func (m *MediaPreview) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTwitter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediaPreview) validateOg(formats strfmt.Registry) error {
	if swag.IsZero(m.Og) { // not required
		return nil
	}

	if m.Og != nil {
		if err := m.Og.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("og")
			}
			return err
		}
	}

	return nil
}

func (m *MediaPreview) validateRedirectType(formats strfmt.Registry) error {
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

func (m *MediaPreview) validateTwitter(formats strfmt.Registry) error {
	if swag.IsZero(m.Twitter) { // not required
		return nil
	}

	if m.Twitter != nil {
		if err := m.Twitter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("twitter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this media preview based on the context it is used
func (m *MediaPreview) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOg(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRedirectType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTwitter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediaPreview) contextValidateOg(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Og.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("og")
		}
		return err
	}

	return nil
}

func (m *MediaPreview) contextValidateRedirectType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RedirectType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("redirectType")
		}
		return err
	}

	return nil
}

func (m *MediaPreview) contextValidateTwitter(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Twitter.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("twitter")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediaPreview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediaPreview) UnmarshalBinary(b []byte) error {
	var res MediaPreview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
