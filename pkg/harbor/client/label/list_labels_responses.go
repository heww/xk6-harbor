// Code generated by go-swagger; DO NOT EDIT.

package label

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// ListLabelsReader is a Reader for the ListLabels structure.
type ListLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListLabelsOK creates a ListLabelsOK with default headers values
func NewListLabelsOK() *ListLabelsOK {
	return &ListLabelsOK{}
}

/* ListLabelsOK describes a response with status code 200, with default header values.

Get successfully.
*/
type ListLabelsOK struct {

	/* Link to previous page and next page
	 */
	Link string

	/* The total count of available items
	 */
	XTotalCount int64

	Payload []*models.Label
}

func (o *ListLabelsOK) Error() string {
	return fmt.Sprintf("[GET /labels][%d] listLabelsOK  %+v", 200, o.Payload)
}
func (o *ListLabelsOK) GetPayload() []*models.Label {
	return o.Payload
}

func (o *ListLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLabelsBadRequest creates a ListLabelsBadRequest with default headers values
func NewListLabelsBadRequest() *ListLabelsBadRequest {
	return &ListLabelsBadRequest{}
}

/* ListLabelsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListLabelsBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListLabelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /labels][%d] listLabelsBadRequest  %+v", 400, o.Payload)
}
func (o *ListLabelsBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListLabelsUnauthorized creates a ListLabelsUnauthorized with default headers values
func NewListLabelsUnauthorized() *ListLabelsUnauthorized {
	return &ListLabelsUnauthorized{}
}

/* ListLabelsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListLabelsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /labels][%d] listLabelsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListLabelsUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListLabelsInternalServerError creates a ListLabelsInternalServerError with default headers values
func NewListLabelsInternalServerError() *ListLabelsInternalServerError {
	return &ListLabelsInternalServerError{}
}

/* ListLabelsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListLabelsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /labels][%d] listLabelsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListLabelsInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
