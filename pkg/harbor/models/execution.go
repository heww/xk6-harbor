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

// Execution execution
//
// swagger:model Execution
type Execution struct {

	// The end time of execution
	EndTime string `json:"end_time,omitempty" js:"endTime"`

	// extra attrs
	ExtraAttrs ExtraAttrs `json:"extra_attrs,omitempty" js:"extraAttrs"`

	// The ID of execution
	ID int64 `json:"id,omitempty" js:"id"`

	// metrics
	Metrics *Metrics `json:"metrics,omitempty" js:"metrics"`

	// The start time of execution
	StartTime string `json:"start_time,omitempty" js:"startTime"`

	// The status of execution
	Status string `json:"status,omitempty" js:"status"`

	// The status message of execution
	StatusMessage string `json:"status_message,omitempty" js:"statusMessage"`

	// The trigger of execution
	Trigger string `json:"trigger,omitempty" js:"trigger"`

	// The vendor id of execution
	VendorID int64 `json:"vendor_id,omitempty" js:"vendorID"`

	// The vendor type of execution
	VendorType string `json:"vendor_type,omitempty" js:"vendorType"`
}

// Validate validates this execution
func (m *Execution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtraAttrs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Execution) validateExtraAttrs(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtraAttrs) { // not required
		return nil
	}

	if m.ExtraAttrs != nil {
		if err := m.ExtraAttrs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extra_attrs")
			}
			return err
		}
	}

	return nil
}

func (m *Execution) validateMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	if m.Metrics != nil {
		if err := m.Metrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this execution based on the context it is used
func (m *Execution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExtraAttrs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Execution) contextValidateExtraAttrs(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ExtraAttrs.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("extra_attrs")
		}
		return err
	}

	return nil
}

func (m *Execution) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	if m.Metrics != nil {
		if err := m.Metrics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Execution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Execution) UnmarshalBinary(b []byte) error {
	var res Execution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
