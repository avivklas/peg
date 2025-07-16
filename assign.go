package peg

import (
	"reflect"
	"strconv"
	"strings"
)

func assignValue(value reflect.Value, raw string) error {
	switch value.Kind() {
	case reflect.String:
		value.SetString(raw)
	case reflect.Invalid:
	case reflect.Bool:
		bVal, err := strconv.ParseBool(raw)
		if err != nil {
			return err
		}

		value.SetBool(bVal)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return err
		}

		value.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		i, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return err
		}

		value.SetUint(i)
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Complex64:
	case reflect.Complex128:
	case reflect.Array:
	case reflect.Chan:
	case reflect.Func:
	case reflect.Interface:
	case reflect.Map:
	case reflect.Pointer:
	case reflect.Slice:
		elm := value.Type().Elem()
		switch elm.Kind() {
		case reflect.String:
			value.Set(reflect.ValueOf(strings.Split(raw, ",")))
		}
	case reflect.Struct:
	case reflect.UnsafePointer:
	}

	return nil
}
