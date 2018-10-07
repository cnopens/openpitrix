// Code generated by go-swagger; DO NOT EDIT.

package account_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DeleteUsersReader is a Reader for the DeleteUsers structure.
type DeleteUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUsersOK creates a DeleteUsersOK with default headers values
func NewDeleteUsersOK() *DeleteUsersOK {
	return &DeleteUsersOK{}
}

/*DeleteUsersOK handles this case with default header values.

A successful response.
*/
type DeleteUsersOK struct {
	Payload *models.OpenpitrixDeleteUsersResponse
}

func (o *DeleteUsersOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/users][%d] deleteUsersOK  %+v", 200, o.Payload)
}

func (o *DeleteUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDeleteUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
