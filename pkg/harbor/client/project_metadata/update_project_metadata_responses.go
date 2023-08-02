// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

// UpdateProjectMetadataReader is a Reader for the UpdateProjectMetadata structure.
type UpdateProjectMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateProjectMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateProjectMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateProjectMetadataConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateProjectMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}] updateProjectMetadata", response, response.Code())
	}
}

// NewUpdateProjectMetadataOK creates a UpdateProjectMetadataOK with default headers values
func NewUpdateProjectMetadataOK() *UpdateProjectMetadataOK {
	return &UpdateProjectMetadataOK{}
}

/*
UpdateProjectMetadataOK describes a response with status code 200, with default header values.

Success
*/
type UpdateProjectMetadataOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this update project metadata o k response has a 2xx status code
func (o *UpdateProjectMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update project metadata o k response has a 3xx status code
func (o *UpdateProjectMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project metadata o k response has a 4xx status code
func (o *UpdateProjectMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project metadata o k response has a 5xx status code
func (o *UpdateProjectMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update project metadata o k response a status code equal to that given
func (o *UpdateProjectMetadataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update project metadata o k response
func (o *UpdateProjectMetadataOK) Code() int {
	return 200
}

func (o *UpdateProjectMetadataOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataOK ", 200)
}

func (o *UpdateProjectMetadataOK) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataOK ", 200)
}

func (o *UpdateProjectMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateProjectMetadataBadRequest creates a UpdateProjectMetadataBadRequest with default headers values
func NewUpdateProjectMetadataBadRequest() *UpdateProjectMetadataBadRequest {
	return &UpdateProjectMetadataBadRequest{}
}

/*
UpdateProjectMetadataBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateProjectMetadataBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update project metadata bad request response has a 2xx status code
func (o *UpdateProjectMetadataBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project metadata bad request response has a 3xx status code
func (o *UpdateProjectMetadataBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project metadata bad request response has a 4xx status code
func (o *UpdateProjectMetadataBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project metadata bad request response has a 5xx status code
func (o *UpdateProjectMetadataBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update project metadata bad request response a status code equal to that given
func (o *UpdateProjectMetadataBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update project metadata bad request response
func (o *UpdateProjectMetadataBadRequest) Code() int {
	return 400
}

func (o *UpdateProjectMetadataBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProjectMetadataBadRequest) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProjectMetadataBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateProjectMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateProjectMetadataUnauthorized creates a UpdateProjectMetadataUnauthorized with default headers values
func NewUpdateProjectMetadataUnauthorized() *UpdateProjectMetadataUnauthorized {
	return &UpdateProjectMetadataUnauthorized{}
}

/*
UpdateProjectMetadataUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateProjectMetadataUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update project metadata unauthorized response has a 2xx status code
func (o *UpdateProjectMetadataUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project metadata unauthorized response has a 3xx status code
func (o *UpdateProjectMetadataUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project metadata unauthorized response has a 4xx status code
func (o *UpdateProjectMetadataUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project metadata unauthorized response has a 5xx status code
func (o *UpdateProjectMetadataUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update project metadata unauthorized response a status code equal to that given
func (o *UpdateProjectMetadataUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update project metadata unauthorized response
func (o *UpdateProjectMetadataUnauthorized) Code() int {
	return 401
}

func (o *UpdateProjectMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateProjectMetadataUnauthorized) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateProjectMetadataUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateProjectMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateProjectMetadataForbidden creates a UpdateProjectMetadataForbidden with default headers values
func NewUpdateProjectMetadataForbidden() *UpdateProjectMetadataForbidden {
	return &UpdateProjectMetadataForbidden{}
}

/*
UpdateProjectMetadataForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateProjectMetadataForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update project metadata forbidden response has a 2xx status code
func (o *UpdateProjectMetadataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project metadata forbidden response has a 3xx status code
func (o *UpdateProjectMetadataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project metadata forbidden response has a 4xx status code
func (o *UpdateProjectMetadataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project metadata forbidden response has a 5xx status code
func (o *UpdateProjectMetadataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update project metadata forbidden response a status code equal to that given
func (o *UpdateProjectMetadataForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update project metadata forbidden response
func (o *UpdateProjectMetadataForbidden) Code() int {
	return 403
}

func (o *UpdateProjectMetadataForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectMetadataForbidden) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectMetadataForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateProjectMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateProjectMetadataNotFound creates a UpdateProjectMetadataNotFound with default headers values
func NewUpdateProjectMetadataNotFound() *UpdateProjectMetadataNotFound {
	return &UpdateProjectMetadataNotFound{}
}

/*
UpdateProjectMetadataNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateProjectMetadataNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update project metadata not found response has a 2xx status code
func (o *UpdateProjectMetadataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project metadata not found response has a 3xx status code
func (o *UpdateProjectMetadataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project metadata not found response has a 4xx status code
func (o *UpdateProjectMetadataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project metadata not found response has a 5xx status code
func (o *UpdateProjectMetadataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update project metadata not found response a status code equal to that given
func (o *UpdateProjectMetadataNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update project metadata not found response
func (o *UpdateProjectMetadataNotFound) Code() int {
	return 404
}

func (o *UpdateProjectMetadataNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectMetadataNotFound) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectMetadataNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateProjectMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateProjectMetadataConflict creates a UpdateProjectMetadataConflict with default headers values
func NewUpdateProjectMetadataConflict() *UpdateProjectMetadataConflict {
	return &UpdateProjectMetadataConflict{}
}

/*
UpdateProjectMetadataConflict describes a response with status code 409, with default header values.

Conflict
*/
type UpdateProjectMetadataConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update project metadata conflict response has a 2xx status code
func (o *UpdateProjectMetadataConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project metadata conflict response has a 3xx status code
func (o *UpdateProjectMetadataConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project metadata conflict response has a 4xx status code
func (o *UpdateProjectMetadataConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project metadata conflict response has a 5xx status code
func (o *UpdateProjectMetadataConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update project metadata conflict response a status code equal to that given
func (o *UpdateProjectMetadataConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update project metadata conflict response
func (o *UpdateProjectMetadataConflict) Code() int {
	return 409
}

func (o *UpdateProjectMetadataConflict) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataConflict  %+v", 409, o.Payload)
}

func (o *UpdateProjectMetadataConflict) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataConflict  %+v", 409, o.Payload)
}

func (o *UpdateProjectMetadataConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateProjectMetadataConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateProjectMetadataInternalServerError creates a UpdateProjectMetadataInternalServerError with default headers values
func NewUpdateProjectMetadataInternalServerError() *UpdateProjectMetadataInternalServerError {
	return &UpdateProjectMetadataInternalServerError{}
}

/*
UpdateProjectMetadataInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateProjectMetadataInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update project metadata internal server error response has a 2xx status code
func (o *UpdateProjectMetadataInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project metadata internal server error response has a 3xx status code
func (o *UpdateProjectMetadataInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project metadata internal server error response has a 4xx status code
func (o *UpdateProjectMetadataInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project metadata internal server error response has a 5xx status code
func (o *UpdateProjectMetadataInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update project metadata internal server error response a status code equal to that given
func (o *UpdateProjectMetadataInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update project metadata internal server error response
func (o *UpdateProjectMetadataInternalServerError) Code() int {
	return 500
}

func (o *UpdateProjectMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateProjectMetadataInternalServerError) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateProjectMetadataInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateProjectMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
