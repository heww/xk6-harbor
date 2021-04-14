// Code generated by go-swagger; DO NOT EDIT.

package systeminfo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// GetSysteminfoReader is a Reader for the GetSysteminfo structure.
type GetSysteminfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSysteminfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSysteminfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetSysteminfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSysteminfoOK creates a GetSysteminfoOK with default headers values
func NewGetSysteminfoOK() *GetSysteminfoOK {
	return &GetSysteminfoOK{}
}

/* GetSysteminfoOK describes a response with status code 200, with default header values.

Get general info successfully.
*/
type GetSysteminfoOK struct {
	Payload *models.GeneralInfo
}

func (o *GetSysteminfoOK) Error() string {
	return fmt.Sprintf("[GET /systeminfo][%d] getSysteminfoOK  %+v", 200, o.Payload)
}
func (o *GetSysteminfoOK) GetPayload() *models.GeneralInfo {
	return o.Payload
}

func (o *GetSysteminfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSysteminfoInternalServerError creates a GetSysteminfoInternalServerError with default headers values
func NewGetSysteminfoInternalServerError() *GetSysteminfoInternalServerError {
	return &GetSysteminfoInternalServerError{}
}

/* GetSysteminfoInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSysteminfoInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetSysteminfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /systeminfo][%d] getSysteminfoInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSysteminfoInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSysteminfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
