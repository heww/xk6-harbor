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

// Task task
//
// swagger:model Task
type Task struct {

	// The creation time of task
	CreationTime string `json:"creation_time,omitempty" js:"creationTime"`

	// The end time of task
	EndTime string `json:"end_time,omitempty" js:"endTime"`

	// The ID of task execution
	ExecutionID int64 `json:"execution_id,omitempty" js:"executionID"`

	// extra attrs
	ExtraAttrs ExtraAttrs `json:"extra_attrs,omitempty" js:"extraAttrs"`

	// The ID of task
	ID int64 `json:"id,omitempty" js:"id"`

	// The count of task run
	RunCount int32 `json:"run_count,omitempty" js:"runCount"`

	// The start time of task
	StartTime string `json:"start_time,omitempty" js:"startTime"`

	// The status of task
	Status string `json:"status,omitempty" js:"status"`

	// The status message of task
	StatusMessage string `json:"status_message,omitempty" js:"statusMessage"`

	// The update time of task
	UpdateTime string `json:"update_time,omitempty" js:"updateTime"`
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtraAttrs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) validateExtraAttrs(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtraAttrs) { // not required
		return nil
	}

	if m.ExtraAttrs != nil {
		if err := m.ExtraAttrs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extra_attrs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extra_attrs")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this task based on the context it is used
func (m *Task) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExtraAttrs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) contextValidateExtraAttrs(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.ExtraAttrs) { // not required
		return nil
	}

	if err := m.ExtraAttrs.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("extra_attrs")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("extra_attrs")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Task) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Task) UnmarshalBinary(b []byte) error {
	var res Task
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
