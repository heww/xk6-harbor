// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Permissions permissions
//
// swagger:model Permissions
type Permissions struct {

	// The project level permissions
	Project []*Permission `json:"project" js:"project"`

	// The system level permissions
	System []*Permission `json:"system" js:"system"`
}

// Validate validates this permissions
func (m *Permissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Permissions) validateProject(formats strfmt.Registry) error {
	if swag.IsZero(m.Project) { // not required
		return nil
	}

	for i := 0; i < len(m.Project); i++ {
		if swag.IsZero(m.Project[i]) { // not required
			continue
		}

		if m.Project[i] != nil {
			if err := m.Project[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("project" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("project" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Permissions) validateSystem(formats strfmt.Registry) error {
	if swag.IsZero(m.System) { // not required
		return nil
	}

	for i := 0; i < len(m.System); i++ {
		if swag.IsZero(m.System[i]) { // not required
			continue
		}

		if m.System[i] != nil {
			if err := m.System[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("system" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("system" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this permissions based on the context it is used
func (m *Permissions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Permissions) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Project); i++ {

		if m.Project[i] != nil {

			if swag.IsZero(m.Project[i]) { // not required
				return nil
			}

			if err := m.Project[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("project" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("project" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Permissions) contextValidateSystem(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.System); i++ {

		if m.System[i] != nil {

			if swag.IsZero(m.System[i]) { // not required
				return nil
			}

			if err := m.System[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("system" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("system" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Permissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Permissions) UnmarshalBinary(b []byte) error {
	var res Permissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
