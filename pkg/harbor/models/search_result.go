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

// SearchResult The chart search result item
//
// swagger:model SearchResult
type SearchResult struct {

	// chart
	Chart *ChartVersion `json:"Chart,omitempty" js:"chart"`

	// The chart name with repo name
	Name string `json:"Name,omitempty" js:"name"`

	// The matched level
	Score int64 `json:"Score,omitempty" js:"score"`
}

// Validate validates this search result
func (m *SearchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchResult) validateChart(formats strfmt.Registry) error {
	if swag.IsZero(m.Chart) { // not required
		return nil
	}

	if m.Chart != nil {
		if err := m.Chart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Chart")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this search result based on the context it is used
func (m *SearchResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchResult) contextValidateChart(ctx context.Context, formats strfmt.Registry) error {

	if m.Chart != nil {
		if err := m.Chart.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Chart")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchResult) UnmarshalBinary(b []byte) error {
	var res SearchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
