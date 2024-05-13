// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserGroupSearchItem user group search item
//
// swagger:model UserGroupSearchItem
type UserGroupSearchItem struct {

	// The name of the user group
	GroupName string `json:"group_name,omitempty" js:"groupName"`

	// The group type, 1 for LDAP group, 2 for HTTP group, 3 for OIDC group.
	GroupType int64 `json:"group_type,omitempty" js:"groupType"`

	// The ID of the user group
	ID int64 `json:"id,omitempty" js:"id"`
}

// Validate validates this user group search item
func (m *UserGroupSearchItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user group search item based on context it is used
func (m *UserGroupSearchItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserGroupSearchItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGroupSearchItem) UnmarshalBinary(b []byte) error {
	var res UserGroupSearchItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
