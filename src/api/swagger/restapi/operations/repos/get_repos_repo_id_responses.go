// Code generated by go-swagger; DO NOT EDIT.

package repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"apphub/src/api/swagger/models"
)

// GetReposRepoIDOKCode is the HTTP code returned for type GetReposRepoIDOK
const GetReposRepoIDOKCode int = 200

/*GetReposRepoIDOK A repo

swagger:response getReposRepoIdOK
*/
type GetReposRepoIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Repo `json:"body,omitempty"`
}

// NewGetReposRepoIDOK creates GetReposRepoIDOK with default headers values
func NewGetReposRepoIDOK() *GetReposRepoIDOK {
	return &GetReposRepoIDOK{}
}

// WithPayload adds the payload to the get repos repo Id o k response
func (o *GetReposRepoIDOK) WithPayload(payload *models.Repo) *GetReposRepoIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos repo Id o k response
func (o *GetReposRepoIDOK) SetPayload(payload *models.Repo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposRepoIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReposRepoIDNotFoundCode is the HTTP code returned for type GetReposRepoIDNotFound
const GetReposRepoIDNotFoundCode int = 404

/*GetReposRepoIDNotFound The repo does not exist.

swagger:response getReposRepoIdNotFound
*/
type GetReposRepoIDNotFound struct {
}

// NewGetReposRepoIDNotFound creates GetReposRepoIDNotFound with default headers values
func NewGetReposRepoIDNotFound() *GetReposRepoIDNotFound {
	return &GetReposRepoIDNotFound{}
}

// WriteResponse to the client
func (o *GetReposRepoIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// GetReposRepoIDInternalServerErrorCode is the HTTP code returned for type GetReposRepoIDInternalServerError
const GetReposRepoIDInternalServerErrorCode int = 500

/*GetReposRepoIDInternalServerError An unexpected error occured.

swagger:response getReposRepoIdInternalServerError
*/
type GetReposRepoIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetReposRepoIDInternalServerError creates GetReposRepoIDInternalServerError with default headers values
func NewGetReposRepoIDInternalServerError() *GetReposRepoIDInternalServerError {
	return &GetReposRepoIDInternalServerError{}
}

// WithPayload adds the payload to the get repos repo Id internal server error response
func (o *GetReposRepoIDInternalServerError) WithPayload(payload *models.Error) *GetReposRepoIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos repo Id internal server error response
func (o *GetReposRepoIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposRepoIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
