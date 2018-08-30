// Code generated by go-swagger; DO NOT EDIT.

package category_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new category manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for category manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateCategory creates category
*/
func (a *Client) CreateCategory(params *CreateCategoryParams) (*CreateCategoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCategoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateCategory",
		Method:             "POST",
		PathPattern:        "/v1/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateCategoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateCategoryOK), nil

}

/*
DeleteCategories deletes categories
*/
func (a *Client) DeleteCategories(params *DeleteCategoriesParams) (*DeleteCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCategoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteCategories",
		Method:             "DELETE",
		PathPattern:        "/v1/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCategoriesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCategoriesOK), nil

}

/*
DescribeCategories describes categories with filter
*/
func (a *Client) DescribeCategories(params *DescribeCategoriesParams) (*DescribeCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeCategoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeCategories",
		Method:             "GET",
		PathPattern:        "/v1/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeCategoriesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeCategoriesOK), nil

}

/*
ModifyCategory modifies category
*/
func (a *Client) ModifyCategory(params *ModifyCategoryParams) (*ModifyCategoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyCategoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyCategory",
		Method:             "PATCH",
		PathPattern:        "/v1/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModifyCategoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyCategoryOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
