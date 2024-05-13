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

// NativeReportSummary The summary for the native report
//
// swagger:model NativeReportSummary
type NativeReportSummary struct {

	// The complete percent of the scanning which value is between 0 and 100
	// Example: 100
	CompletePercent int64 `json:"complete_percent,omitempty" js:"completePercent"`

	// The seconds spent for generating the report
	// Example: 300
	Duration int64 `json:"duration,omitempty" js:"duration"`

	// The end time of the scan process that generating report
	// Example: 2006-01-02T15:04:05Z
	// Format: date-time
	EndTime strfmt.DateTime `json:"end_time,omitempty" js:"endTime"`

	// id of the native scan report
	// Example: 5f62c830-f996-11e9-957f-0242c0a89008
	ReportID string `json:"report_id,omitempty" js:"reportID"`

	// The status of the report generating process
	// Example: Success
	ScanStatus string `json:"scan_status,omitempty" js:"scanStatus"`

	// scanner
	Scanner *Scanner `json:"scanner,omitempty" js:"scanner"`

	// The overall severity
	// Example: High
	Severity string `json:"severity,omitempty" js:"severity"`

	// The start time of the scan process that generating report
	// Example: 2006-01-02T14:04:05Z
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty" js:"startTime"`

	// summary
	Summary *VulnerabilitySummary `json:"summary,omitempty" js:"summary"`
}

// Validate validates this native report summary
func (m *NativeReportSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NativeReportSummary) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NativeReportSummary) validateScanner(formats strfmt.Registry) error {
	if swag.IsZero(m.Scanner) { // not required
		return nil
	}

	if m.Scanner != nil {
		if err := m.Scanner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scanner")
			}
			return err
		}
	}

	return nil
}

func (m *NativeReportSummary) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NativeReportSummary) validateSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.Summary) { // not required
		return nil
	}

	if m.Summary != nil {
		if err := m.Summary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this native report summary based on the context it is used
func (m *NativeReportSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScanner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NativeReportSummary) contextValidateScanner(ctx context.Context, formats strfmt.Registry) error {

	if m.Scanner != nil {

		if swag.IsZero(m.Scanner) { // not required
			return nil
		}

		if err := m.Scanner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scanner")
			}
			return err
		}
	}

	return nil
}

func (m *NativeReportSummary) contextValidateSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.Summary != nil {

		if swag.IsZero(m.Summary) { // not required
			return nil
		}

		if err := m.Summary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NativeReportSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NativeReportSummary) UnmarshalBinary(b []byte) error {
	var res NativeReportSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
