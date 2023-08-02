// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

// StartReplicationReader is a Reader for the StartReplication structure.
type StartReplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartReplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewStartReplicationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStartReplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStartReplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStartReplicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartReplicationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /replication/executions] startReplication", response, response.Code())
	}
}

// NewStartReplicationCreated creates a StartReplicationCreated with default headers values
func NewStartReplicationCreated() *StartReplicationCreated {
	return &StartReplicationCreated{}
}

/*
StartReplicationCreated describes a response with status code 201, with default header values.

Created
*/
type StartReplicationCreated struct {

	/* The location of the resource
	 */
	Location string

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this start replication created response has a 2xx status code
func (o *StartReplicationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start replication created response has a 3xx status code
func (o *StartReplicationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start replication created response has a 4xx status code
func (o *StartReplicationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this start replication created response has a 5xx status code
func (o *StartReplicationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this start replication created response a status code equal to that given
func (o *StartReplicationCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the start replication created response
func (o *StartReplicationCreated) Code() int {
	return 201
}

func (o *StartReplicationCreated) Error() string {
	return fmt.Sprintf("[POST /replication/executions][%d] startReplicationCreated ", 201)
}

func (o *StartReplicationCreated) String() string {
	return fmt.Sprintf("[POST /replication/executions][%d] startReplicationCreated ", 201)
}

func (o *StartReplicationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewStartReplicationBadRequest creates a StartReplicationBadRequest with default headers values
func NewStartReplicationBadRequest() *StartReplicationBadRequest {
	return &StartReplicationBadRequest{}
}

/*
StartReplicationBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type StartReplicationBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this start replication bad request response has a 2xx status code
func (o *StartReplicationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start replication bad request response has a 3xx status code
func (o *StartReplicationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start replication bad request response has a 4xx status code
func (o *StartReplicationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this start replication bad request response has a 5xx status code
func (o *StartReplicationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this start replication bad request response a status code equal to that given
func (o *StartReplicationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the start replication bad request response
func (o *StartReplicationBadRequest) Code() int {
	return 400
}

func (o *StartReplicationBadRequest) Error() string {
	return fmt.Sprintf("[POST /replication/executions][%d] startReplicationBadRequest  %+v", 400, o.Payload)
}

func (o *StartReplicationBadRequest) String() string {
	return fmt.Sprintf("[POST /replication/executions][%d] startReplicationBadRequest  %+v", 400, o.Payload)
}

func (o *StartReplicationBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StartReplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartReplicationUnauthorized creates a StartReplicationUnauthorized with default headers values
func NewStartReplicationUnauthorized() *StartReplicationUnauthorized {
	return &StartReplicationUnauthorized{}
}

/*
StartReplicationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StartReplicationUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this start replication unauthorized response has a 2xx status code
func (o *StartReplicationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start replication unauthorized response has a 3xx status code
func (o *StartReplicationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start replication unauthorized response has a 4xx status code
func (o *StartReplicationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this start replication unauthorized response has a 5xx status code
func (o *StartReplicationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this start replication unauthorized response a status code equal to that given
func (o *StartReplicationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the start replication unauthorized response
func (o *StartReplicationUnauthorized) Code() int {
	return 401
}

func (o *StartReplicationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /replication/executions][%d] startReplicationUnauthorized  %+v", 401, o.Payload)
}

func (o *StartReplicationUnauthorized) String() string {
	return fmt.Sprintf("[POST /replication/executions][%d] startReplicationUnauthorized  %+v", 401, o.Payload)
}

func (o *StartReplicationUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StartReplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartReplicationForbidden creates a StartReplicationForbidden with default headers values
func NewStartReplicationForbidden() *StartReplicationForbidden {
	return &StartReplicationForbidden{}
}

/*
StartReplicationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StartReplicationForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this start replication forbidden response has a 2xx status code
func (o *StartReplicationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start replication forbidden response has a 3xx status code
func (o *StartReplicationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start replication forbidden response has a 4xx status code
func (o *StartReplicationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this start replication forbidden response has a 5xx status code
func (o *StartReplicationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this start replication forbidden response a status code equal to that given
func (o *StartReplicationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the start replication forbidden response
func (o *StartReplicationForbidden) Code() int {
	return 403
}

func (o *StartReplicationForbidden) Error() string {
	return fmt.Sprintf("[POST /replication/executions][%d] startReplicationForbidden  %+v", 403, o.Payload)
}

func (o *StartReplicationForbidden) String() string {
	return fmt.Sprintf("[POST /replication/executions][%d] startReplicationForbidden  %+v", 403, o.Payload)
}

func (o *StartReplicationForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StartReplicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartReplicationInternalServerError creates a StartReplicationInternalServerError with default headers values
func NewStartReplicationInternalServerError() *StartReplicationInternalServerError {
	return &StartReplicationInternalServerError{}
}

/*
StartReplicationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type StartReplicationInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this start replication internal server error response has a 2xx status code
func (o *StartReplicationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start replication internal server error response has a 3xx status code
func (o *StartReplicationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start replication internal server error response has a 4xx status code
func (o *StartReplicationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this start replication internal server error response has a 5xx status code
func (o *StartReplicationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this start replication internal server error response a status code equal to that given
func (o *StartReplicationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the start replication internal server error response
func (o *StartReplicationInternalServerError) Code() int {
	return 500
}

func (o *StartReplicationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /replication/executions][%d] startReplicationInternalServerError  %+v", 500, o.Payload)
}

func (o *StartReplicationInternalServerError) String() string {
	return fmt.Sprintf("[POST /replication/executions][%d] startReplicationInternalServerError  %+v", 500, o.Payload)
}

func (o *StartReplicationInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StartReplicationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
