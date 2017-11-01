// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Repo repo
// swagger:model Repo

type Repo struct {

	// created
	// Read Only: true
	Created strfmt.DateTime `json:"created,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	// Max Length: 12
	// Min Length: 12
	// Pattern: repo-[a-zA-Z0-9]{7}
	ID *string `json:"id"`

	// last modified
	// Read Only: true
	LastModified strfmt.DateTime `json:"last_modified,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

/* polymorph Repo created false */

/* polymorph Repo description false */

/* polymorph Repo id false */

/* polymorph Repo last_modified false */

/* polymorph Repo name false */

/* polymorph Repo url false */

// Validate validates this repo
func (m *Repo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Repo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 12); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 12); err != nil {
		return err
	}

	if err := validate.Pattern("id", "body", string(*m.ID), `repo-[a-zA-Z0-9]{7}`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Repo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Repo) UnmarshalBinary(b []byte) error {
	var res Repo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
