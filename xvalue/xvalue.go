package xvalue

import (
	"runtime"
	"unsafe"
	"fmt"
)

type XValue struct {
	noCopy
	kind Kind
	ptr  unsafe.Pointer
}

type ValueError struct {
	Method string
	Kind   Kind
}

func (e *ValueError) Error() string {
	if e.Kind == 0 {
		return "reflect: call of " + e.Method + " on zero Value"
	}
	return "reflect: call of " + e.Method + " on " + e.Kind.String() + " Value"
}

func methodName() string {
	pc, _, _, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	if f == nil {
		return "unknown method"
	}
	return f.Name()
}


func (v *XValue) mustBe(expected Kind) {
	if v.kind != expected {
		panic(&ValueError{methodName(), v.kind})
	}
}

func (v *XValue) mustBeAssignable() {
	if v.kind == Invalid {
		panic(&ValueError{methodName(), Invalid})
	}
}

func NewValue(k Kind) *XValue {
	v := &XValue{kind:k}
	switch k {
	case Bool, Int, Int8, Int16,
	     Int32, Int64, Uint, Uint8,
	     Uint16, Uint32, Uint64, Float32,
	     Float64:
		v.ptr = unsafe.Pointer(new(uint64))
	case String:
		v.ptr = unsafe.Pointer(new(string))
	case Object:
		v.ptr = unsafe.Pointer(new(XGUID))
	case UserData:
		v.ptr = unsafe.Pointer(new(interface{}))
	default:
		panic(fmt.Sprintf("unknown value kind %s", k))
	}
	return v
}

func (v *XValue) IsValid() bool  {
	return v.kind != Invalid
}

func (v *XValue) Kind() Kind  {
	return v.kind
}

func (v *XValue) Bool() bool {
	v.mustBe(Bool)
	return *(*bool)(v.ptr)
}

func (v *XValue) Int() int {
	v.mustBe(Int)
	return *(*int)(v.ptr)
}

func (v *XValue) Int8() int8 {
	v.mustBe(Int8)
	return *(*int8)(v.ptr)
}

func (v *XValue) Int16() int16 {
	v.mustBe(Int16)
	return *(*int16)(v.ptr)
}

func (v *XValue) Int32() int32 {
	v.mustBe(Int32)
	return *(*int32)(v.ptr)
}

func (v *XValue) Int64() int64 {
	v.mustBe(Int64)
	return *(*int64)(v.ptr)
}

func (v *XValue) Uint() uint {
	v.mustBe(Uint)
	return *(*uint)(v.ptr)
}

func (v *XValue) Uint8() uint8 {
	v.mustBe(Uint8)
	return uint8(uintptr(v.ptr))
}

func (v *XValue) Uint16() uint16 {
	v.mustBe(Uint16)
	return *(*uint16)(v.ptr)
}

func (v *XValue) Uint32() uint32 {
	v.mustBe(Uint32)
	return *(*uint32)(v.ptr)
}

func (v *XValue) Uint64() uint64 {
	v.mustBe(Uint64)
	return *(*uint64)(v.ptr)
}

func (v *XValue) Float32() float32 {
	v.mustBe(Float32)
	return *(*float32)(v.ptr)
}

func (v *XValue) Float64() float64 {
	v.mustBe(Float64)
	return *(*float64)(v.ptr)
}

func (v *XValue) String() string {
	v.mustBe(String)
	return *(*string)(v.ptr)
}

func (v *XValue) Object() XGUID  {
	v.mustBe(Object)
	return *(*XGUID)(v.ptr)
}

func (v *XValue) SetBool(x bool) {
	v.mustBe(Bool)
	*(*bool)(v.ptr) = x
}

func (v *XValue) SetInt(x int) {
	v.mustBe(Int)
	*(*int)(v.ptr) = x
}

func (v *XValue) SetInt8(x int8) {
	v.mustBe(Int8)
	*(*int8)(v.ptr) = x
}

func (v *XValue) SetInt16(x int16) {
	v.mustBe(Int16)
	*(*int16)(v.ptr) = x
}

func (v *XValue) SetInt32(x int32) {
	v.mustBe(Int32)
	*(*int32)(v.ptr) = x
}

func (v *XValue) SetInt64(x int64) {
	v.mustBe(Int64)
	*(*int64)(v.ptr) = x
}

func (v *XValue) SetUint(x uint) {
	v.mustBe(Uint64)
	*(*uint)(v.ptr) = x
}

func (v *XValue) SetUint8(x uint8) {
	v.mustBe(Uint8)
	*(*uint8)(v.ptr) = x
}

func (v *XValue) SetUint16(x uint16) {
	v.mustBe(Uint16)
	*(*uint16)(v.ptr) = x
}

func (v *XValue) SetUint32(x uint32) {
	v.mustBe(Uint32)
	*(*uint32)(v.ptr) = x
}

func (v *XValue) SetUint64(x uint64) {
	v.mustBe(Int64)
	*(*uint64)(v.ptr) = x
}

func (v *XValue) SetFloat32(x float32) {
	v.mustBe(Float32)
	*(*float32)(v.ptr) = x
}

func (v *XValue) SetFloat64(x float64) {
	v.mustBe(Float64)
	*(*float64)(v.ptr) = x
}

func (v *XValue) SetString(x string) {
	v.mustBe(String)
	*(*string)(v.ptr) = x
}

func (v *XValue) SetObject(x XGUID)  {
	v.mustBe(Object)
	*(*XGUID)(v.ptr) = x
}

func (v *XValue) SetUserData(x interface{})  {
	v.mustBe(UserData)
	*(*interface{})(v.ptr) = x
}

func (v *XValue) UserDdata() interface{}  {
	v.mustBe(UserData)
	return *(*interface{})(v.ptr)
}

func (v *XValue) ToString() string  {
	switch v.kind {
	case Bool:
		return fmt.Sprintf("%v", v.Bool())
	case Uint, Uint8, Uint16, Uint32, Uint64:
		return fmt.Sprintf("%v", v.Uint64())
	case Int, Int8, Int16, Int32, Int64:
		return fmt.Sprintf("%v", v.Int64())
	case Float32:
		return fmt.Sprintf("%v", v.Float32())
	case Float64:
		return fmt.Sprintf("%v", v.Float64())
	case String:
		return v.String()
	case Object:
		return v.Object().String()
	case UserData:
		fmt.Sprintf("%v", v.UserDdata())
	}
	return "invalid"
}

