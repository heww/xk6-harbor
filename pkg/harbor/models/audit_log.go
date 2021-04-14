// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuditLog audit log
//
// swagger:model AuditLog
type AuditLog struct {

	// The ID of the audit log entry.
	ID int64 `json:"id,omitempty" js:"id"`

	// The time when this operation is triggered.
	// Example: 2006-01-02T15:04:05
	// Format: date-time
	OpTime strfmt.DateTime `json:"op_time,omitempty" js:"opTime"`

	// The operation against the repository in this log entry.
	Operation string `json:"operation,omitempty" js:"operation"`

	// Name of the repository in this log entry.
	Resource string `json:"resource,omitempty" js:"resource"`

	// Tag of the repository in this log entry.
	ResourceType string `json:"resource_type,omitempty" js:"resourceType"`

	// Username of the user in this log entry.
	Username string `json:"username,omitempty" js:"username"`
}

// Validate validates this audit log
func (m *AuditLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOpTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLog) validateOpTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OpTime) { // not required
		return nil
	}

	if err := validate.FormatOf("op_time", "body", "date-time", m.OpTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this audit log based on context it is used
func (m *AuditLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuditLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLog) UnmarshalBinary(b []byte) error {
	var res AuditLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
