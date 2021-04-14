// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// ImportLdapUserReader is a Reader for the ImportLdapUser structure.
type ImportLdapUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportLdapUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportLdapUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportLdapUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportLdapUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImportLdapUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImportLdapUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImportLdapUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImportLdapUserOK creates a ImportLdapUserOK with default headers values
func NewImportLdapUserOK() *ImportLdapUserOK {
	return &ImportLdapUserOK{}
}

/* ImportLdapUserOK describes a response with status code 200, with default header values.

Add ldap users successfully.
*/
type ImportLdapUserOK struct {
}

func (o *ImportLdapUserOK) Error() string {
	return fmt.Sprintf("[POST /ldap/users/import][%d] importLdapUserOK ", 200)
}

func (o *ImportLdapUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportLdapUserBadRequest creates a ImportLdapUserBadRequest with default headers values
func NewImportLdapUserBadRequest() *ImportLdapUserBadRequest {
	return &ImportLdapUserBadRequest{}
}

/* ImportLdapUserBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ImportLdapUserBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ImportLdapUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /ldap/users/import][%d] importLdapUserBadRequest  %+v", 400, o.Payload)
}
func (o *ImportLdapUserBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ImportLdapUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportLdapUserUnauthorized creates a ImportLdapUserUnauthorized with default headers values
func NewImportLdapUserUnauthorized() *ImportLdapUserUnauthorized {
	return &ImportLdapUserUnauthorized{}
}

/* ImportLdapUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImportLdapUserUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ImportLdapUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ldap/users/import][%d] importLdapUserUnauthorized  %+v", 401, o.Payload)
}
func (o *ImportLdapUserUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ImportLdapUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportLdapUserForbidden creates a ImportLdapUserForbidden with default headers values
func NewImportLdapUserForbidden() *ImportLdapUserForbidden {
	return &ImportLdapUserForbidden{}
}

/* ImportLdapUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImportLdapUserForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ImportLdapUserForbidden) Error() string {
	return fmt.Sprintf("[POST /ldap/users/import][%d] importLdapUserForbidden  %+v", 403, o.Payload)
}
func (o *ImportLdapUserForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ImportLdapUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportLdapUserNotFound creates a ImportLdapUserNotFound with default headers values
func NewImportLdapUserNotFound() *ImportLdapUserNotFound {
	return &ImportLdapUserNotFound{}
}

/* ImportLdapUserNotFound describes a response with status code 404, with default header values.

Failed import some users.
*/
type ImportLdapUserNotFound struct {
	Payload []*models.LdapFailedImportUser
}

func (o *ImportLdapUserNotFound) Error() string {
	return fmt.Sprintf("[POST /ldap/users/import][%d] importLdapUserNotFound  %+v", 404, o.Payload)
}
func (o *ImportLdapUserNotFound) GetPayload() []*models.LdapFailedImportUser {
	return o.Payload
}

func (o *ImportLdapUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportLdapUserInternalServerError creates a ImportLdapUserInternalServerError with default headers values
func NewImportLdapUserInternalServerError() *ImportLdapUserInternalServerError {
	return &ImportLdapUserInternalServerError{}
}

/* ImportLdapUserInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImportLdapUserInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ImportLdapUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ldap/users/import][%d] importLdapUserInternalServerError  %+v", 500, o.Payload)
}
func (o *ImportLdapUserInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ImportLdapUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
