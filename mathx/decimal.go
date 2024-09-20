package mathx

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/shopspring/decimal"
)

func toDecimal[T Number](value T) (decimal.Decimal, error) {
	switch v := any(value).(type) {
	case int:
		return decimal.New(int64(v), 0), nil
	case int8:
		return decimal.New(int64(v), 0), nil
	case int16:
		return decimal.New(int64(v), 0), nil
	case int32:
		return decimal.New(int64(v), 0), nil
	case int64:
		return decimal.New(v, 0), nil
	case uint:
		return decimal.New(int64(v), 0), nil
	case uint8:
		return decimal.New(int64(v), 0), nil
	case uint16:
		return decimal.New(int64(v), 0), nil
	case uint32:
		return decimal.New(int64(v), 0), nil
	case uint64:
		return decimal.New(int64(v), 0), nil
	case float32:
		return decimal.NewFromFloat32(v), nil
	case float64:
		return decimal.NewFromFloat(v), nil
	default:
		return decimal.New(0, 0), errors.New("unsupported type")
	}
}

// ---------------------------- Add ----------------------------
func AddE[T1 Number, T2 Number](a T1, b T2) (float64, error) {
	decA, err := toDecimal(a)
	if err != nil {
		return 0, err
	}

	decB, err := toDecimal(b)
	if err != nil {
		return 0, err
	}

	decC, _ := decA.Add(decB).Float64()

	return decC, nil
}

func Add[T1 Number, T2 Number](a T1, b T2) float64 {
	r, err := AddE(a, b)
	if err != nil {
		panic(err)
	}
	return r
}

func SafeAdd[T1 Number, T2 Number](a T1, b T2) float64 {
	r, err := AddE(a, b)
	if err != nil {
		return 0
	}

	return r
}

// ---------------------------- Sub ----------------------------

func SubE[T1 Number, T2 Number](a T1, b T2) (float64, error) {
	decA, err := toDecimal(a)
	if err != nil {
		return 0, err
	}

	decB, err := toDecimal(b)
	if err != nil {
		return 0, err
	}

	decC, _ := decA.Sub(decB).Float64()

	return decC, nil
}

func Sub[T1 Number, T2 Number](a T1, b T2) float64 {
	r, err := SubE(a, b)
	if err != nil {
		panic(err)
	}
	return r
}

func SafeSub[T1 Number, T2 Number](a T1, b T2) float64 {
	r, err := SubE(a, b)
	if err != nil {
		return 0
	}

	return r
}

// ---------------------------- Mul ----------------------------
func MulE[T1 Number, T2 Number](a T1, b T2) (float64, error) {
	decA, err := toDecimal(a)
	if err != nil {
		return 0, err
	}

	decB, err := toDecimal(b)
	if err != nil {
		return 0, err
	}

	decC, _ := decA.Mul(decB).Float64()

	return decC, nil
}

func Mul[T1 Number, T2 Number](a T1, b T2) float64 {
	r, err := MulE(a, b)
	if err != nil {
		panic(err)
	}
	return r
}

func SafeMul[T1 Number, T2 Number](a T1, b T2) float64 {
	r, err := MulE(a, b)
	if err != nil {
		return 0
	}

	return r
}

// ---------------------------- Div ----------------------------
func DivE[T1 Number, T2 Number](a T1, b T2) (float64, error) {
	if b == 0 {
		return 0, errors.New("denominator is 0")
	}

	decA, err := toDecimal(a)
	if err != nil {
		return 0, err
	}

	decB, err := toDecimal(b)
	if err != nil {
		return 0, err
	}

	decC, _ := decA.Div(decB).Float64()

	return decC, nil
}

func Div[T1 Number, T2 Number](a T1, b T2) float64 {
	r, err := DivE(a, b)
	if err != nil {
		panic(err)
	}
	return r
}

func SafeDiv[T1 Number, T2 Number](a T1, b T2) float64 {
	r, err := DivE(a, b)
	if err != nil {
		return 0
	}
	return r
}

// ---------------------------- Others ----------------------------
func Max[T Number](x, y T) T {
	return T(math.Max(float64(x), float64(y)))
}

func Min[T Number](x, y T) T {
	return T(math.Min(float64(x), float64(y)))
}

func Precision[T1 Float, T2 Integer](target T1, prec T2) float64 {
	fmtStr := "%." + strconv.FormatInt(int64(prec), 10) + "f"
	result, err := strconv.ParseFloat(fmt.Sprintf(fmtStr, target), 64)
	if err != nil {
		panic(err)
	}
	return result
}

// FloatEqual guarantees n of effective figure
func FloatEqual(f1, f2 float64, n int) bool {
	min := math.Pow10(-1 * n)
	for math.Abs(f1) > 1 {
		f1 /= 10.0
		f2 /= 10.0
	}
	if f1 > f2 {
		return math.Dim(f1, f2) < min
	} else {
		return math.Dim(f2, f1) < min
	}
}

// FloatEqual2 guarantees diff of two number smaller equal than the defined small number
func FloatEqual2(f1, f2 float64, min float64) bool {
	if f1 > f2 {
		return math.Dim(f1, f2) < min
	} else {
		return math.Dim(f2, f1) < min
	}
}
