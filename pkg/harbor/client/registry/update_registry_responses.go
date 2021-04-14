// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// UpdateRegistryReader is a Reader for the UpdateRegistry structure.
type UpdateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateRegistryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRegistryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRegistryOK creates a UpdateRegistryOK with default headers values
func NewUpdateRegistryOK() *UpdateRegistryOK {
	return &UpdateRegistryOK{}
}

/* UpdateRegistryOK describes a response with status code 200, with default header values.

Success
*/
type UpdateRegistryOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateRegistryOK) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] updateRegistryOK ", 200)
}

func (o *UpdateRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateRegistryUnauthorized creates a UpdateRegistryUnauthorized with default headers values
func NewUpdateRegistryUnauthorized() *UpdateRegistryUnauthorized {
	return &UpdateRegistryUnauthorized{}
}

/* UpdateRegistryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateRegistryUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] updateRegistryUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateRegistryUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRegistryForbidden creates a UpdateRegistryForbidden with default headers values
func NewUpdateRegistryForbidden() *UpdateRegistryForbidden {
	return &UpdateRegistryForbidden{}
}

/* UpdateRegistryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateRegistryForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRegistryForbidden) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] updateRegistryForbidden  %+v", 403, o.Payload)
}
func (o *UpdateRegistryForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRegistryNotFound creates a UpdateRegistryNotFound with default headers values
func NewUpdateRegistryNotFound() *UpdateRegistryNotFound {
	return &UpdateRegistryNotFound{}
}

/* UpdateRegistryNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateRegistryNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRegistryNotFound) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] updateRegistryNotFound  %+v", 404, o.Payload)
}
func (o *UpdateRegistryNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRegistryConflict creates a UpdateRegistryConflict with default headers values
func NewUpdateRegistryConflict() *UpdateRegistryConflict {
	return &UpdateRegistryConflict{}
}

/* UpdateRegistryConflict describes a response with status code 409, with default header values.

Conflict
*/
type UpdateRegistryConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRegistryConflict) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] updateRegistryConflict  %+v", 409, o.Payload)
}
func (o *UpdateRegistryConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRegistryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRegistryInternalServerError creates a UpdateRegistryInternalServerError with default headers values
func NewUpdateRegistryInternalServerError() *UpdateRegistryInternalServerError {
	return &UpdateRegistryInternalServerError{}
}

/* UpdateRegistryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateRegistryInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRegistryInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] updateRegistryInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateRegistryInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRegistryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
