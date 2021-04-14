// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// SetScannerAsDefaultReader is a Reader for the SetScannerAsDefault structure.
type SetScannerAsDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetScannerAsDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetScannerAsDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetScannerAsDefaultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetScannerAsDefaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetScannerAsDefaultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetScannerAsDefaultOK creates a SetScannerAsDefaultOK with default headers values
func NewSetScannerAsDefaultOK() *SetScannerAsDefaultOK {
	return &SetScannerAsDefaultOK{}
}

/* SetScannerAsDefaultOK describes a response with status code 200, with default header values.

Successfully set the specified scanner registration as system default
*/
type SetScannerAsDefaultOK struct {
}

func (o *SetScannerAsDefaultOK) Error() string {
	return fmt.Sprintf("[PATCH /scanners/{registration_id}][%d] setScannerAsDefaultOK ", 200)
}

func (o *SetScannerAsDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetScannerAsDefaultUnauthorized creates a SetScannerAsDefaultUnauthorized with default headers values
func NewSetScannerAsDefaultUnauthorized() *SetScannerAsDefaultUnauthorized {
	return &SetScannerAsDefaultUnauthorized{}
}

/* SetScannerAsDefaultUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SetScannerAsDefaultUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *SetScannerAsDefaultUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /scanners/{registration_id}][%d] setScannerAsDefaultUnauthorized  %+v", 401, o.Payload)
}
func (o *SetScannerAsDefaultUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SetScannerAsDefaultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetScannerAsDefaultForbidden creates a SetScannerAsDefaultForbidden with default headers values
func NewSetScannerAsDefaultForbidden() *SetScannerAsDefaultForbidden {
	return &SetScannerAsDefaultForbidden{}
}

/* SetScannerAsDefaultForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SetScannerAsDefaultForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *SetScannerAsDefaultForbidden) Error() string {
	return fmt.Sprintf("[PATCH /scanners/{registration_id}][%d] setScannerAsDefaultForbidden  %+v", 403, o.Payload)
}
func (o *SetScannerAsDefaultForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SetScannerAsDefaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetScannerAsDefaultInternalServerError creates a SetScannerAsDefaultInternalServerError with default headers values
func NewSetScannerAsDefaultInternalServerError() *SetScannerAsDefaultInternalServerError {
	return &SetScannerAsDefaultInternalServerError{}
}

/* SetScannerAsDefaultInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SetScannerAsDefaultInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *SetScannerAsDefaultInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /scanners/{registration_id}][%d] setScannerAsDefaultInternalServerError  %+v", 500, o.Payload)
}
func (o *SetScannerAsDefaultInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SetScannerAsDefaultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
