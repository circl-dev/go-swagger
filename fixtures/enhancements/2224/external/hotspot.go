// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/protodev-site/validate"
)

// Hotspot hotspot
//
// swagger:model Hotspot
type Hotspot struct {

	// f b ID
	FBID uint64 `json:"FBID,omitempty" gorm:"index"`

	// created at
	// Format: date-time
	CreatedAt *Time `json:"created_at,omitempty"`

	// id
	ID uint64 `json:"id,omitempty" gorm:"primary_key"`

	// updated at
	// Format: date-time
	UpdatedAt *Time `json:"updated_at,omitempty"`

	// version
	Version uint64 `json:"version,omitempty"`

	// access points
	AccessPoints []*AccessPoint `json:"access_points"`

	// type
	// Required: true
	Type *HotspotType `json:"type"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Hotspot) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		FBID uint64 `json:"FBID,omitempty"`

		CreatedAt *Time `json:"created_at,omitempty"`

		ID uint64 `json:"id,omitempty"`

		UpdatedAt *Time `json:"updated_at,omitempty"`

		Version uint64 `json:"version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.FBID = dataAO0.FBID

	m.CreatedAt = dataAO0.CreatedAt

	m.ID = dataAO0.ID

	m.UpdatedAt = dataAO0.UpdatedAt

	m.Version = dataAO0.Version

	// AO1
	var dataAO1 struct {
		AccessPoints []*AccessPoint `json:"access_points"`

		Type *HotspotType `json:"type"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AccessPoints = dataAO1.AccessPoints

	m.Type = dataAO1.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Hotspot) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		FBID uint64 `json:"FBID,omitempty"`

		CreatedAt *Time `json:"created_at,omitempty"`

		ID uint64 `json:"id,omitempty"`

		UpdatedAt *Time `json:"updated_at,omitempty"`

		Version uint64 `json:"version,omitempty"`
	}

	dataAO0.FBID = m.FBID

	dataAO0.CreatedAt = m.CreatedAt

	dataAO0.ID = m.ID

	dataAO0.UpdatedAt = m.UpdatedAt

	dataAO0.Version = m.Version

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	var dataAO1 struct {
		AccessPoints []*AccessPoint `json:"access_points"`

		Type *HotspotType `json:"type"`
	}

	dataAO1.AccessPoints = m.AccessPoints

	dataAO1.Type = m.Type

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hotspot
func (m *Hotspot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hotspot) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if m.CreatedAt != nil {
		if err := m.CreatedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_at")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_at")
			}
			return err
		}
	}

	return nil
}

func (m *Hotspot) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if m.UpdatedAt != nil {
		if err := m.UpdatedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_at")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_at")
			}
			return err
		}
	}

	return nil
}

func (m *Hotspot) validateAccessPoints(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessPoints) { // not required
		return nil
	}

	for i := 0; i < len(m.AccessPoints); i++ {
		if swag.IsZero(m.AccessPoints[i]) { // not required
			continue
		}

		if m.AccessPoints[i] != nil {
			if err := m.AccessPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("access_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Hotspot) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hotspot based on the context it is used
func (m *Hotspot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAccessPoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hotspot) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedAt != nil {
		if err := m.CreatedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_at")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_at")
			}
			return err
		}
	}

	return nil
}

func (m *Hotspot) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdatedAt != nil {
		if err := m.UpdatedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_at")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_at")
			}
			return err
		}
	}

	return nil
}

func (m *Hotspot) contextValidateAccessPoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AccessPoints); i++ {

		if m.AccessPoints[i] != nil {
			if err := m.AccessPoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("access_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Hotspot) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Hotspot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Hotspot) UnmarshalBinary(b []byte) error {
	var res Hotspot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
