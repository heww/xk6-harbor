// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetentionSelectorMetadata retention selector
//
// swagger:model RetentionSelectorMetadata
type RetentionSelectorMetadata struct {

	// decorations
	Decorations []string `json:"decorations" js:"decorations"`

	// display text
	DisplayText string `json:"display_text,omitempty" js:"displayText"`

	// kind
	Kind string `json:"kind,omitempty" js:"kind"`
}

// Validate validates this retention selector metadata
func (m *RetentionSelectorMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this retention selector metadata based on context it is used
func (m *RetentionSelectorMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RetentionSelectorMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionSelectorMetadata) UnmarshalBinary(b []byte) error {
	var res RetentionSelectorMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
