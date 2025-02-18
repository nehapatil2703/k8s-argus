// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NOCItemBase n o c item base
// swagger:discriminator NOCItemBase type
type NOCItemBase interface {
	runtime.Validatable

	// type
	// Required: true
	Type() string
	SetType(string)
}

type nOCItemBase struct {
	typeField string
}

// Type gets the type of this polymorphic type
func (m *nOCItemBase) Type() string {
	return "NOCItemBase"
}

// SetType sets the type of this polymorphic type
func (m *nOCItemBase) SetType(val string) {
}

// UnmarshalNOCItemBaseSlice unmarshals polymorphic slices of NOCItemBase
func UnmarshalNOCItemBaseSlice(reader io.Reader, consumer runtime.Consumer) ([]NOCItemBase, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []NOCItemBase
	for _, element := range elements {
		obj, err := unmarshalNOCItemBase(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalNOCItemBase unmarshals polymorphic NOCItemBase
func UnmarshalNOCItemBase(reader io.Reader, consumer runtime.Consumer) (NOCItemBase, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalNOCItemBase(data, consumer)
}

func unmarshalNOCItemBase(data []byte, consumer runtime.Consumer) (NOCItemBase, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "NOCItemBase":
		var result nOCItemBase
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "device":
		var result DeviceNOCItem
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "website":
		var result WebsiteNOCItem
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)
}

// Validate validates this n o c item base
func (m *nOCItemBase) Validate(formats strfmt.Registry) error {
	return nil
}
