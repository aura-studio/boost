package structure

import (
	"encoding/json"
	"reflect"
)

type assert struct{}

var Assert = assert{}

func (assert) IsPointer(a any) bool {
	return reflect.TypeOf(a).Kind() == reflect.Ptr
}

func (assert) IsJSON(a string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(a), &js) == nil
}

func (assert assert) IsMapStructureValue(a any) bool {
	val := reflect.ValueOf(a)
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool, reflect.String:
		return true
	case reflect.Map:
		if val.Type().Key().Kind() != reflect.String {
			return false
		}
		for _, key := range val.MapKeys() {
			if !assert.IsMapStructureValue(val.MapIndex(key).Interface()) {
				return false
			}
		}
		return true
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			if !assert.IsMapStructureValue(val.Index(i).Interface()) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func (assert assert) IsStructStructureValue(a any) bool {
	val := reflect.ValueOf(a)
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool, reflect.String:
		return true
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if !assert.IsStructStructureValue(field.Interface()) {
				return false
			}
		}
		return true
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			if !assert.IsStructStructureValue(val.Index(i).Interface()) {
				return false
			}
		}
		return true
	default:
		return false
	}
}
