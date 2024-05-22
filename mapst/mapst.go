package mapst

import (
	"errors"
	"reflect"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

var (
	errDstTypeMustBePointerOfStruct             = errors.New("dst type must be pointer of struct")
	errDstTypeMustBePointerOfMapStringInterface = errors.New("dst type must be pointer of map[string]interface{}")
	errTypeConvertionNotSupported               = errors.New("type conversion not supported")
)

func Convert(src any, dst any) {
	if err := ConvertE(src, dst); err != nil {
		panic(err)
	}
}

func ConvertE(src any, dst any) error {
	if isMapStringInterface(src) || isPointerOfMapStringInterface(src) {
		if isPointerOfStruct(dst) {
			return convertMapToStruct(src, dst)
		} else {
			return errDstTypeMustBePointerOfStruct
		}
	}

	if isStruct(src) || isPointerOfStruct(src) {
		switch {
		case isMapStringInterface(dst):
			return convertStructToMap(src, getPointerOfMapStringInterface(dst))
		case isPointerOfMapStringInterface(dst):
			return convertStructToMap(src, dst)
		default:
			return errDstTypeMustBePointerOfMapStringInterface
		}
	}

	return errTypeConvertionNotSupported
}

func convertMapToStruct(m any, s any) error {
	return mapstructure.Decode(m, s)
}

func convertStructToMap(s any, m any) error {
	return mapstructure.Decode(structs.Map(s), m)
}

func isMapStringInterface(v any) bool {
	typeTo := reflect.TypeOf(v)
	if typeTo.Kind() == reflect.Map {
		if typeTo.Key().Kind() == reflect.String && typeTo.Elem().Kind() == reflect.Interface {
			return true
		}
	}
	return false
}

func isPointerOfMapStringInterface(v any) bool {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Map &&
		t.Elem().Key().Kind() == reflect.String && t.Elem().Elem().Kind() == reflect.Interface {
		return true
	}
	return false
}

func isStruct(v any) bool {
	t := reflect.TypeOf(v)
	return t.Kind() == reflect.Struct
}

func isPointerOfStruct(v any) bool {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct {
		return true
	}
	return false
}

func getPointerOfMapStringInterface(v any) any {
	typeTo := reflect.TypeOf(v)
	return reflect.New(typeTo.Elem()).Interface()
}
