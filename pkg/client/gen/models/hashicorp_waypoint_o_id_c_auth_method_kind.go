// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpWaypointOIDCAuthMethodKind hashicorp waypoint o ID c auth method kind
//
// swagger:model hashicorp.waypoint.OIDCAuthMethod.Kind
type HashicorpWaypointOIDCAuthMethodKind string

func NewHashicorpWaypointOIDCAuthMethodKind(value HashicorpWaypointOIDCAuthMethodKind) *HashicorpWaypointOIDCAuthMethodKind {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpWaypointOIDCAuthMethodKind.
func (m HashicorpWaypointOIDCAuthMethodKind) Pointer() *HashicorpWaypointOIDCAuthMethodKind {
	return &m
}

const (

	// HashicorpWaypointOIDCAuthMethodKindUNKNOWN captures enum value "UNKNOWN"
	HashicorpWaypointOIDCAuthMethodKindUNKNOWN HashicorpWaypointOIDCAuthMethodKind = "UNKNOWN"

	// HashicorpWaypointOIDCAuthMethodKindGITHUB captures enum value "GITHUB"
	HashicorpWaypointOIDCAuthMethodKindGITHUB HashicorpWaypointOIDCAuthMethodKind = "GITHUB"

	// HashicorpWaypointOIDCAuthMethodKindGOOGLE captures enum value "GOOGLE"
	HashicorpWaypointOIDCAuthMethodKindGOOGLE HashicorpWaypointOIDCAuthMethodKind = "GOOGLE"
)

// for schema
var hashicorpWaypointOIdCAuthMethodKindEnum []interface{}

func init() {
	var res []HashicorpWaypointOIDCAuthMethodKind
	if err := json.Unmarshal([]byte(`["UNKNOWN","GITHUB","GOOGLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpWaypointOIdCAuthMethodKindEnum = append(hashicorpWaypointOIdCAuthMethodKindEnum, v)
	}
}

func (m HashicorpWaypointOIDCAuthMethodKind) validateHashicorpWaypointOIDCAuthMethodKindEnum(path, location string, value HashicorpWaypointOIDCAuthMethodKind) error {
	if err := validate.EnumCase(path, location, value, hashicorpWaypointOIdCAuthMethodKindEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp waypoint o ID c auth method kind
func (m HashicorpWaypointOIDCAuthMethodKind) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpWaypointOIDCAuthMethodKindEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp waypoint o ID c auth method kind based on context it is used
func (m HashicorpWaypointOIDCAuthMethodKind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}