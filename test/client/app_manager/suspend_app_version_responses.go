// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// SuspendAppVersionReader is a Reader for the SuspendAppVersion structure.
type SuspendAppVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuspendAppVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSuspendAppVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSuspendAppVersionOK creates a SuspendAppVersionOK with default headers values
func NewSuspendAppVersionOK() *SuspendAppVersionOK {
	return &SuspendAppVersionOK{}
}

/*SuspendAppVersionOK handles this case with default header values.

A successful response.
*/
type SuspendAppVersionOK struct {
	Payload *models.OpenpitrixSuspendAppVersionResponse
}

func (o *SuspendAppVersionOK) Error() string {
	return fmt.Sprintf("[POST /v1/app_version/action/suspend][%d] suspendAppVersionOK  %+v", 200, o.Payload)
}

func (o *SuspendAppVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixSuspendAppVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
