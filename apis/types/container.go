// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Container Container contains response of Engine API:
// GET "/containers/json""
//
// swagger:model Container

type Container struct {

	// command
	Command string `json:"Command,omitempty"`

	// created
	Created int64 `json:"Created,omitempty"`

	// host config
	HostConfig *HostConfig `json:"HostConfig,omitempty"`

	// ID
	ID string `json:"ID,omitempty"`

	// image
	Image string `json:"Image,omitempty"`

	// image ID
	ImageID string `json:"ImageID,omitempty"`

	// labels
	Labels map[string]string `json:"Labels,omitempty"`

	// names
	Names []string `json:"Names"`

	// pid
	Pid int `json:"Pid,omitempty"`

	// size root fs
	SizeRootFs int64 `json:"SizeRootFs,omitempty"`

	// size rw
	SizeRw int64 `json:"SizeRw,omitempty"`

	// state
	State string `json:"State,omitempty"`

	// status
	Status string `json:"Status,omitempty"`
}

/* polymorph Container Command false */

/* polymorph Container Created false */

/* polymorph Container HostConfig false */

/* polymorph Container ID false */

/* polymorph Container Image false */

/* polymorph Container ImageID false */

/* polymorph Container Labels false */

/* polymorph Container Names false */

/* polymorph Container SizeRootFs false */

/* polymorph Container SizeRw false */

/* polymorph Container State false */

/* polymorph Container Status false */

// Validate validates this container
func (m *Container) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNames(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Container) validateNames(formats strfmt.Registry) error {

	if swag.IsZero(m.Names) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Container) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Container) UnmarshalBinary(b []byte) error {
	var res Container
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
