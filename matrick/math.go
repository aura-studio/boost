package matrick

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func Max[T Number](x, y T) T {
	return T(math.Max(float64(x), float64(y)))
}

func Min[T Number](x, y T) T {
	return T(math.Min(float64(x), float64(y)))
}

func Div[T1 Number, T2 Number](a T1, b T2) T1 {
	if b == 0 {
		return 0
	}
	return T1(float64(a) / float64(b))
}

func DivE[T1 Number, T2 Number](a T1, b T2) (T1, error) {
	if b == 0 {
		return 0, errors.New("denominator is 0")
	}
	return T1(float64(a) / float64(b)), nil
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
