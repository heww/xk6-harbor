// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// UpdateRobotReader is a Reader for the UpdateRobot structure.
type UpdateRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRobotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRobotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRobotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRobotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateRobotConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRobotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRobotOK creates a UpdateRobotOK with default headers values
func NewUpdateRobotOK() *UpdateRobotOK {
	return &UpdateRobotOK{}
}

/* UpdateRobotOK describes a response with status code 200, with default header values.

Success
*/
type UpdateRobotOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateRobotOK) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotOK ", 200)
}

func (o *UpdateRobotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateRobotBadRequest creates a UpdateRobotBadRequest with default headers values
func NewUpdateRobotBadRequest() *UpdateRobotBadRequest {
	return &UpdateRobotBadRequest{}
}

/* UpdateRobotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateRobotBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRobotBadRequest) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateRobotBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRobotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRobotUnauthorized creates a UpdateRobotUnauthorized with default headers values
func NewUpdateRobotUnauthorized() *UpdateRobotUnauthorized {
	return &UpdateRobotUnauthorized{}
}

/* UpdateRobotUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateRobotUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRobotUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateRobotUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRobotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRobotForbidden creates a UpdateRobotForbidden with default headers values
func NewUpdateRobotForbidden() *UpdateRobotForbidden {
	return &UpdateRobotForbidden{}
}

/* UpdateRobotForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateRobotForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRobotForbidden) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotForbidden  %+v", 403, o.Payload)
}
func (o *UpdateRobotForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRobotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRobotNotFound creates a UpdateRobotNotFound with default headers values
func NewUpdateRobotNotFound() *UpdateRobotNotFound {
	return &UpdateRobotNotFound{}
}

/* UpdateRobotNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateRobotNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRobotNotFound) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotNotFound  %+v", 404, o.Payload)
}
func (o *UpdateRobotNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRobotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRobotConflict creates a UpdateRobotConflict with default headers values
func NewUpdateRobotConflict() *UpdateRobotConflict {
	return &UpdateRobotConflict{}
}

/* UpdateRobotConflict describes a response with status code 409, with default header values.

Conflict
*/
type UpdateRobotConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRobotConflict) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotConflict  %+v", 409, o.Payload)
}
func (o *UpdateRobotConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRobotConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRobotInternalServerError creates a UpdateRobotInternalServerError with default headers values
func NewUpdateRobotInternalServerError() *UpdateRobotInternalServerError {
	return &UpdateRobotInternalServerError{}
}

/* UpdateRobotInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateRobotInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRobotInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateRobotInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRobotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
