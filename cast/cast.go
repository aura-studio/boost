package cast

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

type stringer struct{ string }

func (s stringer) String() string {
	return s.string
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a any) any {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirectToStringerOrError returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
// or error,
func indirectToStringerOrError(a any) any {
	if a == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

type decimalParser struct{}

var dec = decimalParser{}

func (p decimalParser) ToInt(s string) (int64, error) {
	switch {
	case strings.ContainsAny(s, "tf"):
		b, err := strconv.ParseBool(s)
		if err != nil {
			return 0, err
		}
		if b {
			return 1, nil
		}
		return 0, nil
	case strings.ContainsAny(s, "(i)"):
		n, err := strconv.ParseComplex(s, 128)
		if err != nil {
			return 0, err
		}
		return int64(real(n)), nil
	case strings.Contains(s, "/"):
		n, ok := big.NewRat(0, 1).SetString(s)
		if !ok {
			return 0, fmt.Errorf("unable to cast %#v of type %T to *big.Rat", s, s)
		}
		f, _ := n.Float64()
		return int64(f), nil
	default:
		s = p.trimPointZeroOfIntString(s)
		n, ok := big.NewInt(0).SetString(s, 0)
		if !ok {
			return 0, fmt.Errorf("unable to cast %#v of type %T to *big.Int", s, s)
		}
		return n.Int64(), nil
	}
}

func (p decimalParser) ToUint(s string) (uint64, error) {
	switch {
	case strings.ContainsAny(s, "tf"):
		b, err := strconv.ParseBool(s)
		if err != nil {
			return 0, err
		}
		if b {
			return 1, nil
		}
		return 0, nil
	case strings.ContainsAny(s, "(i)"):
		n, err := strconv.ParseComplex(s, 128)
		if err != nil {
			return 0, err
		}
		return uint64(real(n)), nil
	case strings.Contains(s, "/"):
		n, ok := big.NewRat(0, 1).SetString(s)
		if !ok {
			return 0, fmt.Errorf("unable to cast %#v of type %T to *big.Rat", s, s)
		}
		f, _ := n.Float64()
		return uint64(f), nil
	default:
		s = p.trimPointZeroOfIntString(s)
		n, ok := big.NewInt(0).SetString(s, 0)
		if !ok {
			return 0, fmt.Errorf("unable to cast %#v of type %T to *big.Int", s, s)
		}
		if n.Sign() < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", s, s)
		}
		return n.Uint64(), nil
	}
}

func (p decimalParser) ToFloat32(s string) (float32, error) {
	switch {
	case strings.ContainsAny(s, "tf"):
		b, err := strconv.ParseBool(s)
		if err != nil {
			return 0, err
		}
		if b {
			return 1, nil
		}
		return 0, nil
	case strings.ContainsAny(s, "(i)"):
		n, err := strconv.ParseComplex(s, 128)
		if err != nil {
			return 0, err
		}
		return float32(real(n)), nil
	case strings.Contains(s, "/"):
		n, ok := big.NewRat(0, 1).SetString(s)
		if !ok {
			return 0, fmt.Errorf("unable to cast %#v of type %T to *big.Rat", s, s)
		}
		f, _ := n.Float32()
		return f, nil
	default:
		f, ok := big.NewFloat(0).SetString(s)
		if !ok {
			return 0, fmt.Errorf("unable to cast %#v of type %T to *big.Float", s, s)
		}
		f32, _ := f.Float32()
		return f32, nil
	}
}

func (p decimalParser) ToFloat64(s string) (float64, error) {
	switch {
	case strings.ContainsAny(s, "tf"):
		b, err := strconv.ParseBool(s)
		if err != nil {
			return 0, err
		}
		if b {
			return 1, nil
		}
		return 0, nil
	case strings.ContainsAny(s, "(i)"):
		n, err := strconv.ParseComplex(s, 128)
		if err != nil {
			return 0, err
		}
		return real(n), nil
	case strings.Contains(s, "/"):
		n, ok := big.NewRat(0, 1).SetString(s)
		if !ok {
			return 0, fmt.Errorf("unable to cast %#v of type %T to *big.Rat", s, s)
		}
		f, _ := n.Float64()
		return f, nil
	default:
		f, ok := big.NewFloat(0).SetString(s)
		if !ok {
			return 0, fmt.Errorf("unable to cast %#v of type %T to *big.Float", s, s)
		}
		f64, _ := f.Float64()
		return f64, nil
	}
}

func (p decimalParser) ToBigInt(s string) (*big.Int, error) {
	switch {
	case strings.ContainsAny(s, "tf"):
		b, err := strconv.ParseBool(s)
		if err != nil {
			return nil, err
		}
		if b {
			return big.NewInt(1), nil
		}
		return big.NewInt(0), nil
	case strings.ContainsAny(s, "(i)"):
		n, err := strconv.ParseComplex(s, 128)
		if err != nil {
			return nil, err
		}
		return big.NewInt(int64(real(n))), nil
	case strings.Contains(s, "/"):
		n, ok := big.NewRat(0, 1).SetString(s)
		if !ok {
			return nil, fmt.Errorf("unable to cast %#v of type %T to *big.Rat", s, s)
		}
		f, _ := n.Float64()
		return big.NewInt(int64(f)), nil
	default:
		s = p.trimPointZeroOfIntString(s)
		n, ok := big.NewInt(0).SetString(s, 0)
		if !ok {
			return nil, fmt.Errorf("unable to cast %#v of type %T to *big.Int", s, s)
		}
		return n, nil
	}
}

func (p decimalParser) ToBigFloat(s string) (*big.Float, error) {
	switch {
	case strings.ContainsAny(s, "tf"):
		b, err := strconv.ParseBool(s)
		if err != nil {
			return nil, err
		}
		if b {
			return big.NewFloat(1), nil
		}
		return big.NewFloat(0), nil
	case strings.ContainsAny(s, "(i)"):
		n, err := strconv.ParseComplex(s, 128)
		if err != nil {
			return nil, err
		}
		return big.NewFloat(real(n)), nil
	case strings.Contains(s, "/"):
		n, ok := big.NewRat(0, 1).SetString(s)
		if !ok {
			return nil, fmt.Errorf("unable to cast %#v of type %T to *big.Rat", s, s)
		}
		f, _ := n.Float64()
		return big.NewFloat(f), nil
	default:
		f, ok := big.NewFloat(0).SetString(s)
		if !ok {
			return nil, fmt.Errorf("unable to cast %#v of type %T to *big.Float", s, s)
		}
		return f, nil
	}
}

func (p decimalParser) ToBigRat(s string) (*big.Rat, error) {
	switch {
	case strings.ContainsAny(s, "tf"):
		b, err := strconv.ParseBool(s)
		if err != nil {
			return nil, err
		}
		if b {
			return big.NewRat(1, 1), nil
		}
		return big.NewRat(0, 1), nil
	case strings.ContainsAny(s, "tf"):
		b, err := strconv.ParseBool(s)
		if err != nil {
			return nil, err
		}
		if b {
			return big.NewRat(1, 1), nil
		}
		return big.NewRat(0, 1), nil
	case strings.ContainsAny(s, "(i)"):
		n, err := strconv.ParseComplex(s, 128)
		if err != nil {
			return nil, err
		}
		return big.NewRat(int64(real(n)), 1), nil
	case strings.Contains(s, "/"):
		n, ok := big.NewRat(0, 1).SetString(s)
		if !ok {
			return nil, fmt.Errorf("unable to cast %#v of type %T to *big.Rat", s, s)
		}
		return n, nil
	default:
		f, ok := big.NewRat(0, 1).SetString(s)
		if !ok {
			return nil, fmt.Errorf("unable to cast %#v of type %T to *big.Rat", s, s)
		}
		return f, nil
	}
}

func (p decimalParser) ToComplex64(s string) (complex64, error) {
	switch {
	case strings.ContainsAny(s, "tf"):
		b, err := strconv.ParseBool(s)
		if err != nil {
			return complex(0, 0), err
		}
		if b {
			return complex(1, 0), nil
		}
		return complex(0, 0), nil
	case strings.ContainsAny(s, "(i)"):
		n, err := strconv.ParseComplex(s, 64)
		if err != nil {
			return 0, err
		}
		return complex64(n), nil
	case strings.Contains(s, "/"):
		n, ok := big.NewRat(0, 1).SetString(s)
		if !ok {
			return 0, fmt.Errorf("unable to cast %#v of type %T to *big.Rat", s, s)
		}
		f, _ := n.Float32()
		return complex(f, 0), nil
	default:
		n, err := strconv.ParseComplex(s, 64)
		if err != nil {
			return 0, err
		}
		return complex64(n), nil
	}
}

func (p decimalParser) ToComplex128(s string) (complex128, error) {
	switch {
	case strings.ContainsAny(s, "tf"):
		b, err := strconv.ParseBool(s)
		if err != nil {
			return complex(0, 0), err
		}
		if b {
			return complex(1, 0), nil
		}
		return complex(0, 0), nil
	case strings.ContainsAny(s, "(i)"):
		n, err := strconv.ParseComplex(s, 128)
		if err != nil {
			return 0, err
		}
		return n, nil
	case strings.Contains(s, "/"):
		n, ok := big.NewRat(0, 1).SetString(s)
		if !ok {
			return 0, fmt.Errorf("unable to cast %#v of type %T to *big.Rat", s, s)
		}
		f, _ := n.Float64()
		return complex(f, 0), nil
	default:
		n, err := strconv.ParseComplex(s, 128)
		if err != nil {
			return 0, err
		}
		return n, nil
	}
}

func (p decimalParser) ToBool(s string) (bool, error) {
	switch {
	case strings.ContainsAny(s, "tf"):
		return strconv.ParseBool(s)
	case strings.ContainsAny(s, "(i)"):
		n, err := strconv.ParseComplex(s, 128)
		if err != nil {
			return false, err
		}
		return real(n) != 0, nil
	case strings.Contains(s, "/"):
		n, ok := big.NewRat(0, 1).SetString(s)
		if !ok {
			return false, fmt.Errorf("unable to cast %#v of type %T to *big.Rat", s, s)
		}
		f, _ := n.Float64()
		return f != 0, nil
	default:
		f, ok := big.NewFloat(0).SetString(s)
		if !ok {
			return false, fmt.Errorf("unable to cast %#v of type %T to *big.Float", s, s)
		}
		f64, _ := f.Float64()
		return f64 != 0, nil
	}
}

func (p decimalParser) trimPointZeroOfIntString(s string) string {
	var foundZero bool
	for i := len(s); i > 0; i-- {
		switch s[i-1] {
		case '.':
			if foundZero {
				s = s[:i-1]
				return s
			}
		case '0':
			foundZero = true
		default:
			return s
		}
	}
	return s
}
