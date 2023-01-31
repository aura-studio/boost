package plainest

import (
	"fmt"
	"math/big"

	"github.com/aura-studio/boost/encoding"
)

var (
	json  = encoding.NewJSON()
	proto = encoding.NewProtobuf()
)

// Int is a wrapper for int.
type Int struct {
	_root  *Root
	_key   string
	_value int
}

// Get is a getter for Int
func (f *Int) Get() int {
	return f._value
}

// Set is a setter for Int
func (f *Int) Set(value int) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Int) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Int8 is a wrapper for int8.
type Int8 struct {
	_root  *Root
	_key   string
	_value int8
}

// Get is a getter for Int8
func (f *Int8) Get() int8 {
	return f._value
}

// Set is a setter for Int8
func (f *Int8) Set(value int8) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Int8) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Int8) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Int16 is a wrapper for int16.
type Int16 struct {
	_root  *Root
	_key   string
	_value int16
}

// Get is a getter for Int16
func (f *Int16) Get() int16 {
	return f._value
}

// Set is a setter for Int16
func (f *Int16) Set(value int16) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Int16) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Int32 is a wrapper for int32.
type Int32 struct {
	_root  *Root
	_key   string
	_value int32
}

// Get is a getter for Int32
func (f *Int32) Get() int32 {
	return f._value
}

// Set is a setter for Int32
func (f *Int32) Set(value int32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Int32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Int64 is a wrapper for int64.
type Int64 struct {
	_root  *Root
	_key   string
	_value int64
}

// Get is a getter for Int64
func (f *Int64) Get() int64 {
	return f._value
}

// Set is a setter for Int64
func (f *Int64) Set(value int64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Int64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Uint is a wrapper for uint.
type Uint struct {
	_root  *Root
	_key   string
	_value uint
}

// Get is a getter for Uint
func (f *Uint) Get() uint {
	return f._value
}

// Set is a setter for Uint
func (f *Uint) Set(value uint) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Uint) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Uint) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Uint8 is a wrapper for uint8.
type Uint8 struct {
	_root  *Root
	_key   string
	_value uint8
}

// Get is a getter for Uint8
func (f *Uint8) Get() uint8 {
	return f._value
}

// Set is a getter for Uint8
func (f *Uint8) Set(value uint8) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Uint8) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Uint8) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Uint16 is a wrapper for uint16.
type Uint16 struct {
	_root  *Root
	_key   string
	_value uint16
}

// Get is a getter for Uint16
func (f *Uint16) Get() uint16 {
	return f._value
}

// Set is a setter for Uint16
func (f *Uint16) Set(value uint16) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Uint16) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Uint16) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Uint32 is a wrapper for int8.
type Uint32 struct {
	_root  *Root
	_key   string
	_value uint32
}

// Get is a getter for Uint32
func (f *Uint32) Get() uint32 {
	return f._value
}

// Set is a setter for Uint32
func (f *Uint32) Set(value uint32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Uint32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Uint32) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Uint64 is a wrapper for uint64.
type Uint64 struct {
	_root  *Root
	_key   string
	_value uint64
}

// Get is a getter for Uint64
func (f *Uint64) Get() uint64 {
	return f._value
}

// Set is a setter for Uint64
func (f *Uint64) Set(value uint64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Uint64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Uint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Float32 is a wrapper for float32.
type Float32 struct {
	_root  *Root
	_key   string
	_value float32
}

// Get is a getter for Float32
func (f *Float32) Get() float32 {
	return f._value
}

// Set is a setter for Float32
func (f *Float32) Set(value float32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Float32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Float32) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Float64 is a wrapper for float64.
type Float64 struct {
	_root  *Root
	_key   string
	_value float64
}

// Get is a getter for Float64
func (f *Float64) Get() float64 {
	return f._value
}

// Set is a setter for Float64
func (f *Float64) Set(value float64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Float64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// BigInt is a wrapper for big.Int
type BigInt struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for BigInt
func (f *BigInt) Get() int64 {
	if f._value == "" {
		return 0
	}
	n, ok := new(big.Int).SetString(f._value, 10)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigInt Parse failed, value=%#v",
			"big.Int SetString error", f._value))
	}
	return n.Int64()
}

// Set is a setter for BigInt
func (f *BigInt) Set(value int64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = big.NewInt(value).String()
	f._root._mod[f._key] = f._value
}

// GetBig is a getter for BigInt
func (f *BigInt) GetBig() *big.Int {
	if f._value == "" {
		return big.NewInt(0)
	}
	n, ok := new(big.Int).SetString(f._value, 10)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigInt Parse failed, value=%#v",
			"big.Int SetString error", f._value))
	}
	return n
}

// SetBig is a setter for BigInt
func (f *BigInt) SetBig(n *big.Int) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	if n == nil {
		f._value = ""
	} else {
		f._value = n.String()
	}
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *BigInt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f BigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// BigRat is a wrapper for big.Rat
type BigRat struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for BigFloat
func (f *BigRat) Get() float64 {
	if f._value == "" {
		return 0
	}
	n, ok := new(big.Rat).SetString(f._value)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigRat Parse failed, value=%#v",
			"big.Rat SetString error", f._value))
	}
	v, _ := n.Float64()
	return v
}

// Set is a setter for BigFloat
func (f *BigRat) Set(v float64) {
	rat, _ := big.NewFloat(v).Rat(nil)
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = rat.String()
	f._root._mod[f._key] = f._value
}

// GetBig is a getter for BigRat
func (f *BigRat) GetBig() *big.Rat {
	if f._value == "" {
		return big.NewRat(0, 0)
	}
	n, ok := new(big.Rat).SetString(f._value)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigRat Parse failed, value=%#v",
			"big.Rat SetString error", f._value))
	}
	return n
}

// SetBig is a setter for BigRat
func (f *BigRat) SetBig(n *big.Rat) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	if n == nil {
		f._value = ""
	} else {
		f._value = n.String()
	}
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *BigRat) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f BigRat) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// BigFloat is a wrapper for big.Float
type BigFloat struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for BigFloat
func (f *BigFloat) Get() float64 {
	if f._value == "" {
		return 0
	}
	n, ok := new(big.Float).SetString(f._value)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigFloat Parse failed, value=%#v",
			"big.Float SetString error", f._value))
	}
	f64, _ := n.Float64()
	return f64
}

// Set is a setter for BigFloat
func (f *BigFloat) Set(value float64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = big.NewFloat(value).String()
	f._root._mod[f._key] = f._value
}

// GetBig is a getter for BigFloat
func (f *BigFloat) GetBig() *big.Float {
	if f._value == "" {
		return big.NewFloat(0)
	}
	n, ok := new(big.Float).SetString(f._value)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigFloat Parse failed, value=%#v",
			"big.Float SetString error", f._value))
	}
	return n
}

// SetBig is a setter for BigFloat
func (f *BigFloat) SetBig(n *big.Float) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	if n == nil {
		f._value = ""
	} else {
		f._value = n.String()
	}
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *BigFloat) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f BigFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Bool is a wrapper for bool.
type Bool struct {
	_root  *Root
	_key   string
	_value bool
}

// Get is a getter for Bool
func (f *Bool) Get() bool {
	return f._value
}

// Set is a setter for Bool
func (f *Bool) Set(value bool) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Bool) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// String is a wrapper for bool.
type String struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for String
func (f *String) Get() string {
	return f._value
}

// Set is a setter for String
func (f *String) Set(value string) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *String) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f String) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Time is a wrapper for Unix time
type Time struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for Time
func (f *Time) Get() int64 {
	if f._value == "" {
		return 0
	}
	return timeStringToStamp(f._value)
}

// Set is a setter for Time
func (f *Time) Set(value int64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	if value == 0 {
		f._value = ""
	} else {
		f._value = timeStampToString(value)
	}

	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Time) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// JSON is a wrapper for json
type JSON struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for String
func (f *JSON) Get(n interface{}) {
	if len(f._value) == 0 {
		return
	}
	err := json.Unmarshal([]byte(f._value), n)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree JSON Unmarshal failed, value=%#v",
			err.Error(), f._value))
	}
}

// Set is a setter for String
func (f *JSON) Set(value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree JSON Marshal failed, value=%#v",
			err.Error(), value))
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = string(b)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *JSON) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f JSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Proto is a wrapper for json
type Proto struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for String
func (f *Proto) Get(n interface{}) {
	err := proto.Unmarshal([]byte(f._value), n)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree Proto Unmarshal failed", err.Error()))
	}
}

// Set is a setter for String
func (f *Proto) Set(value interface{}) {
	b, err := proto.Marshal(value)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree Proto Marshal failed, value=%#v",
			err.Error(), value))
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = string(b)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Proto) UnmarshalJSON(data []byte) error {
	return nil
}

// MarshalJSON implements json.Marshal
func (f Proto) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d bytes binary", len(f._value))), nil
}

// SliceInt is a wrapper for []int.
type SliceInt struct {
	_root  *Root
	_key   string
	_value []int
}

// Get is a getter for SliceInt
func (f *SliceInt) Get() []int {
	value := make([]int, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceInt
func (f *SliceInt) Set(value []int) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]int, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceInt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceInt8 is a wrapper for []int8.
type SliceInt8 struct {
	_root  *Root
	_key   string
	_value []int8
}

// Get is a getter for SliceInt8
func (f *SliceInt8) Get() []int8 {
	value := make([]int8, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceInt8
func (f *SliceInt8) Set(value []int8) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]int8, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceInt8) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceInt8) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceInt16 is a wrapper for []int16.
type SliceInt16 struct {
	_root  *Root
	_key   string
	_value []int16
}

// Get is a getter for SliceInt16
func (f *SliceInt16) Get() []int16 {
	value := make([]int16, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceInt16
func (f *SliceInt16) Set(value []int16) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]int16, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceInt16) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceInt16) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceInt32 is a wrapper for []int32.
type SliceInt32 struct {
	_root  *Root
	_key   string
	_value []int32
}

// Get is a getter for SliceInt32
func (f *SliceInt32) Get() []int32 {
	value := make([]int32, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceInt32
func (f *SliceInt32) Set(value []int32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]int32, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceInt32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceInt64 is a wrapper for []int64.
type SliceInt64 struct {
	_root  *Root
	_key   string
	_value []int64
}

// Get is a getter for SliceInt64
func (f *SliceInt64) Get() []int64 {
	value := make([]int64, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceInt64
func (f *SliceInt64) Set(value []int64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]int64, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceInt64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceUint is a wrapper for []uint.
type SliceUint struct {
	_root  *Root
	_key   string
	_value []uint
}

// Get is a getter for SliceUint
func (f *SliceUint) Get() []uint {
	value := make([]uint, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceUint
func (f *SliceUint) Set(value []uint) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]uint, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceUint) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceUint) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceUint is a wrapper for []uint8.
type SliceUint8 struct {
	_root  *Root
	_key   string
	_value []uint8
}

// Get is a getter for SliceUint8
func (f *SliceUint8) Get() []uint8 {
	value := make([]uint8, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceUint8
func (f *SliceUint8) Set(value []uint8) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]uint8, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceUint8) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceUint8) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceUint16 is a wrapper for []uint16.
type SliceUint16 struct {
	_root  *Root
	_key   string
	_value []uint16
}

// Get is a getter for SliceUint16
func (f *SliceUint16) Get() []uint16 {
	value := make([]uint16, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceUint16
func (f *SliceUint16) Set(value []uint16) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]uint16, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceUint16) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceUint16) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceUint32 is a wrapper for []uint32.
type SliceUint32 struct {
	_root  *Root
	_key   string
	_value []uint32
}

// Get is a getter for SliceUint32
func (f *SliceUint32) Get() []uint32 {
	value := make([]uint32, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceUint32
func (f *SliceUint32) Set(value []uint32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]uint32, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceUint32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceUint32) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceUint64 is a wrapper for []int64.
type SliceUint64 struct {
	_root  *Root
	_key   string
	_value []uint64
}

// Get is a getter for SliceUint64
func (f *SliceUint64) Get() []uint64 {
	value := make([]uint64, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceUint64
func (f *SliceUint64) Set(value []uint64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]uint64, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceUint64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceUint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceFloat32 is a wrapper for []float32.
type SliceFloat32 struct {
	_root  *Root
	_key   string
	_value []float32
}

// Get is a getter for SliceFloat32
func (f *SliceFloat32) Get() []float32 {
	value := make([]float32, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceFloat32
func (f *SliceFloat32) Set(value []float32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]float32, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceFloat32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceFloat32) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceFloat64 is a wrapper for []float64.
type SliceFloat64 struct {
	_root  *Root
	_key   string
	_value []float64
}

// Get is a getter for SliceFloat64
func (f *SliceFloat64) Get() []float64 {
	value := make([]float64, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceFloat64
func (f *SliceFloat64) Set(value []float64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]float64, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceFloat64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceBigInt is a wrapper for []big.Int.
type SliceBigInt struct {
	_root  *Root
	_key   string
	_value []*big.Int
}

// Get is a getter for SliceBigInt
func (f *SliceBigInt) Get() []*big.Int {
	value := make([]*big.Int, len(f._value))
	for i, v := range f._value {
		n := new(big.Int)
		value[i] = n.Add(v, n)
	}
	return value
}

// Set is a setter for SliceBigInt
func (f *SliceBigInt) Set(value []*big.Int) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]*big.Int, len(value))
	for i, v := range value {
		n := new(big.Int)
		f._value[i] = n.Add(v, n)
	}
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceBigInt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceBigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceBigRat is a wrapper for []big.Rat.
type SliceBigRat struct {
	_root  *Root
	_key   string
	_value []*big.Rat
}

// Get is a getter for SliceBigRat
func (f *SliceBigRat) Get() []*big.Rat {
	value := make([]*big.Rat, len(f._value))
	for i, v := range f._value {
		n := new(big.Rat)
		value[i] = n.Add(v, n)
	}
	return value
}

// Set is a setter for SliceBigRat
func (f *SliceBigRat) Set(value []*big.Rat) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]*big.Rat, len(value))
	for i, v := range value {
		n := new(big.Rat)
		f._value[i] = n.Add(v, n)
	}
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceBigRat) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceBigRat) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceBigFloat is a wrapper for []big.Int.
type SliceBigFloat struct {
	_root  *Root
	_key   string
	_value []*big.Float
}

// Get is a getter for SliceBigFloat
func (f *SliceBigFloat) Get() []*big.Float {
	value := make([]*big.Float, len(f._value))
	for i, v := range f._value {
		n := new(big.Float)
		value[i] = n.Add(v, n)
	}
	return value
}

// Set is a setter for SliceBigFloat
func (f *SliceBigFloat) Set(value []*big.Float) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]*big.Float, len(value))
	for i, v := range value {
		n := new(big.Float)
		f._value[i] = n.Add(v, n)
	}
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceBigFloat) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceBigFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceTime is a wrapper for []Unix time
type SliceTime struct {
	_root  *Root
	_key   string
	_value []string
}

// Get is a getter for SliceTime
func (f *SliceTime) Get() []int64 {
	var ns []int64
	for _, s := range f._value {
		if s == "" {
			ns = append(ns, 0)
		} else {
			t := timeStringToStamp(s)
			ns = append(ns, t)
		}
	}
	return ns
}

// Set is a setter for SliceTime
func (f *SliceTime) Set(ns []int64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	var ss []string
	for _, n := range ns {
		if n == 0 {
			ss = append(ss, "")
		} else {
			ss = append(ss, timeStampToString(n))
		}
	}
	f._value = ss

	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceTime) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceBool is a wrapper for []bool
type SliceBool struct {
	_root  *Root
	_key   string
	_value []bool
}

// Get is a getter for SliceBool
func (f *SliceBool) Get() []bool {
	value := make([]bool, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceBool
func (f *SliceBool) Set(value []bool) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]bool, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceBool) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceString is a wrapper for []string
type SliceString struct {
	_root  *Root
	_key   string
	_value []string
}

// Get is a getter for SliceString
func (f *SliceString) Get() []string {
	value := make([]string, len(f._value))
	copy(value, f._value)
	return value
}

// Set is a setter for SliceString
func (f *SliceString) Set(value []string) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]string, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceString) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f SliceString) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}
