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

// ModifyAppVersionReader is a Reader for the ModifyAppVersion structure.
type ModifyAppVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyAppVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyAppVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyAppVersionOK creates a ModifyAppVersionOK with default headers values
func NewModifyAppVersionOK() *ModifyAppVersionOK {
	return &ModifyAppVersionOK{}
}

/*ModifyAppVersionOK handles this case with default header values.

A successful response.
*/
type ModifyAppVersionOK struct {
	Payload *models.OpenpitrixModifyAppVersionResponse
}

func (o *ModifyAppVersionOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/app_versions][%d] modifyAppVersionOK  %+v", 200, o.Payload)
}

func (o *ModifyAppVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixModifyAppVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
