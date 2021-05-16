// Code generated by go-swagger; DO NOT EDIT.

package label

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// DeleteLabelReader is a Reader for the DeleteLabel structure.
type DeleteLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLabelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLabelOK creates a DeleteLabelOK with default headers values
func NewDeleteLabelOK() *DeleteLabelOK {
	return &DeleteLabelOK{}
}

/* DeleteLabelOK describes a response with status code 200, with default header values.

Success
*/
type DeleteLabelOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteLabelOK) Error() string {
	return fmt.Sprintf("[DELETE /labels/{label_id}][%d] deleteLabelOK ", 200)
}

func (o *DeleteLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteLabelBadRequest creates a DeleteLabelBadRequest with default headers values
func NewDeleteLabelBadRequest() *DeleteLabelBadRequest {
	return &DeleteLabelBadRequest{}
}

/* DeleteLabelBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteLabelBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteLabelBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /labels/{label_id}][%d] deleteLabelBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteLabelBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteLabelUnauthorized creates a DeleteLabelUnauthorized with default headers values
func NewDeleteLabelUnauthorized() *DeleteLabelUnauthorized {
	return &DeleteLabelUnauthorized{}
}

/* DeleteLabelUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteLabelUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteLabelUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /labels/{label_id}][%d] deleteLabelUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteLabelUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteLabelNotFound creates a DeleteLabelNotFound with default headers values
func NewDeleteLabelNotFound() *DeleteLabelNotFound {
	return &DeleteLabelNotFound{}
}

/* DeleteLabelNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteLabelNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteLabelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /labels/{label_id}][%d] deleteLabelNotFound  %+v", 404, o.Payload)
}
func (o *DeleteLabelNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteLabelInternalServerError creates a DeleteLabelInternalServerError with default headers values
func NewDeleteLabelInternalServerError() *DeleteLabelInternalServerError {
	return &DeleteLabelInternalServerError{}
}

/* DeleteLabelInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteLabelInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteLabelInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /labels/{label_id}][%d] deleteLabelInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteLabelInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteLabelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
