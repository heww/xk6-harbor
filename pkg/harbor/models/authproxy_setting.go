// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthproxySetting authproxy setting
//
// swagger:model AuthproxySetting
type AuthproxySetting struct {

	// The fully qualified URI of login endpoint of authproxy, such as 'https://192.168.1.2:8443/login'
	Endpoint string `json:"endpoint,omitempty" js:"endpoint"`

	// The certificate to be pinned when connecting auth proxy.
	ServerCertificate string `json:"server_certificate,omitempty" js:"serverCertificate"`

	// The flag to determine whether Harbor can skip search the user/group when adding him as a member.
	SkipSearch bool `json:"skip_search,omitempty" js:"skipSearch"`

	// The fully qualified URI of token review endpoint of authproxy, such as 'https://192.168.1.2:8443/tokenreview'
	TokenreivewEndpoint string `json:"tokenreivew_endpoint,omitempty" js:"tokenreivewEndpoint"`

	// The flag to determine whether Harbor should verify the certificate when connecting to the auth proxy.
	VerifyCert bool `json:"verify_cert,omitempty" js:"verifyCert"`
}

// Validate validates this authproxy setting
func (m *AuthproxySetting) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this authproxy setting based on context it is used
func (m *AuthproxySetting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthproxySetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthproxySetting) UnmarshalBinary(b []byte) error {
	var res AuthproxySetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
