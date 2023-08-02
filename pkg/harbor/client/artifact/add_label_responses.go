// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

// AddLabelReader is a Reader for the AddLabel structure.
type AddLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddLabelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddLabelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddLabelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels] addLabel", response, response.Code())
	}
}

// NewAddLabelOK creates a AddLabelOK with default headers values
func NewAddLabelOK() *AddLabelOK {
	return &AddLabelOK{}
}

/*
AddLabelOK describes a response with status code 200, with default header values.

Success
*/
type AddLabelOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this add label o k response has a 2xx status code
func (o *AddLabelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add label o k response has a 3xx status code
func (o *AddLabelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add label o k response has a 4xx status code
func (o *AddLabelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add label o k response has a 5xx status code
func (o *AddLabelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add label o k response a status code equal to that given
func (o *AddLabelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add label o k response
func (o *AddLabelOK) Code() int {
	return 200
}

func (o *AddLabelOK) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelOK ", 200)
}

func (o *AddLabelOK) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelOK ", 200)
}

func (o *AddLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewAddLabelBadRequest creates a AddLabelBadRequest with default headers values
func NewAddLabelBadRequest() *AddLabelBadRequest {
	return &AddLabelBadRequest{}
}

/*
AddLabelBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddLabelBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this add label bad request response has a 2xx status code
func (o *AddLabelBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add label bad request response has a 3xx status code
func (o *AddLabelBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add label bad request response has a 4xx status code
func (o *AddLabelBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add label bad request response has a 5xx status code
func (o *AddLabelBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add label bad request response a status code equal to that given
func (o *AddLabelBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add label bad request response
func (o *AddLabelBadRequest) Code() int {
	return 400
}

func (o *AddLabelBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelBadRequest  %+v", 400, o.Payload)
}

func (o *AddLabelBadRequest) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelBadRequest  %+v", 400, o.Payload)
}

func (o *AddLabelBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddLabelUnauthorized creates a AddLabelUnauthorized with default headers values
func NewAddLabelUnauthorized() *AddLabelUnauthorized {
	return &AddLabelUnauthorized{}
}

/*
AddLabelUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddLabelUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this add label unauthorized response has a 2xx status code
func (o *AddLabelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add label unauthorized response has a 3xx status code
func (o *AddLabelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add label unauthorized response has a 4xx status code
func (o *AddLabelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add label unauthorized response has a 5xx status code
func (o *AddLabelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add label unauthorized response a status code equal to that given
func (o *AddLabelUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add label unauthorized response
func (o *AddLabelUnauthorized) Code() int {
	return 401
}

func (o *AddLabelUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelUnauthorized  %+v", 401, o.Payload)
}

func (o *AddLabelUnauthorized) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelUnauthorized  %+v", 401, o.Payload)
}

func (o *AddLabelUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddLabelForbidden creates a AddLabelForbidden with default headers values
func NewAddLabelForbidden() *AddLabelForbidden {
	return &AddLabelForbidden{}
}

/*
AddLabelForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddLabelForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this add label forbidden response has a 2xx status code
func (o *AddLabelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add label forbidden response has a 3xx status code
func (o *AddLabelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add label forbidden response has a 4xx status code
func (o *AddLabelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add label forbidden response has a 5xx status code
func (o *AddLabelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add label forbidden response a status code equal to that given
func (o *AddLabelForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the add label forbidden response
func (o *AddLabelForbidden) Code() int {
	return 403
}

func (o *AddLabelForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelForbidden  %+v", 403, o.Payload)
}

func (o *AddLabelForbidden) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelForbidden  %+v", 403, o.Payload)
}

func (o *AddLabelForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddLabelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddLabelNotFound creates a AddLabelNotFound with default headers values
func NewAddLabelNotFound() *AddLabelNotFound {
	return &AddLabelNotFound{}
}

/*
AddLabelNotFound describes a response with status code 404, with default header values.

Not found
*/
type AddLabelNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this add label not found response has a 2xx status code
func (o *AddLabelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add label not found response has a 3xx status code
func (o *AddLabelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add label not found response has a 4xx status code
func (o *AddLabelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add label not found response has a 5xx status code
func (o *AddLabelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add label not found response a status code equal to that given
func (o *AddLabelNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add label not found response
func (o *AddLabelNotFound) Code() int {
	return 404
}

func (o *AddLabelNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelNotFound  %+v", 404, o.Payload)
}

func (o *AddLabelNotFound) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelNotFound  %+v", 404, o.Payload)
}

func (o *AddLabelNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddLabelConflict creates a AddLabelConflict with default headers values
func NewAddLabelConflict() *AddLabelConflict {
	return &AddLabelConflict{}
}

/*
AddLabelConflict describes a response with status code 409, with default header values.

Conflict
*/
type AddLabelConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this add label conflict response has a 2xx status code
func (o *AddLabelConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add label conflict response has a 3xx status code
func (o *AddLabelConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add label conflict response has a 4xx status code
func (o *AddLabelConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add label conflict response has a 5xx status code
func (o *AddLabelConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add label conflict response a status code equal to that given
func (o *AddLabelConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add label conflict response
func (o *AddLabelConflict) Code() int {
	return 409
}

func (o *AddLabelConflict) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelConflict  %+v", 409, o.Payload)
}

func (o *AddLabelConflict) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelConflict  %+v", 409, o.Payload)
}

func (o *AddLabelConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddLabelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddLabelInternalServerError creates a AddLabelInternalServerError with default headers values
func NewAddLabelInternalServerError() *AddLabelInternalServerError {
	return &AddLabelInternalServerError{}
}

/*
AddLabelInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type AddLabelInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this add label internal server error response has a 2xx status code
func (o *AddLabelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add label internal server error response has a 3xx status code
func (o *AddLabelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add label internal server error response has a 4xx status code
func (o *AddLabelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add label internal server error response has a 5xx status code
func (o *AddLabelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add label internal server error response a status code equal to that given
func (o *AddLabelInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add label internal server error response
func (o *AddLabelInternalServerError) Code() int {
	return 500
}

func (o *AddLabelInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *AddLabelInternalServerError) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *AddLabelInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddLabelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
