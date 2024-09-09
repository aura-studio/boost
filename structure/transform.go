package structure

import (
	"encoding/json"
	"errors"
)

var (
	ErrSourceTypeError      = errors.New("source type error")
	ErrDestinationTypeError = errors.New("destination type error")
)

type transform struct{}

var Transform = transform{}

func (transform transform) MapToStruct(m any, s any) error {
	if !Assert.IsPointer(s) || !Assert.IsStructStructureValue(s) {
		return ErrDestinationTypeError
	}

	if !Assert.IsMapStructureValue(m) {
		return ErrSourceTypeError
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, s)
	if err != nil {
		return err
	}

	return nil
}

func (transform transform) StructToMap(s any, m any) error {
	if !Assert.IsPointer(s) || !Assert.IsStructStructureValue(s) {
		return ErrSourceTypeError
	}

	if !Assert.IsMapStructureValue(m) {
		return ErrDestinationTypeError
	}

	bytes, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, m)
	if err != nil {
		return err
	}

	return nil
}
