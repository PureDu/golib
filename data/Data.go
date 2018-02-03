package data

import (
	"unsafe"

	"github.com/PureDu/golib/util"
)

type DATA_TYPE int32

const (
	DT_UNKNOWN = 0
	//bool
	DT_BOOLEAN
	//int
	DT_INT
	//int32
	DT_INT32
	//int64
	DT_INT64
	//float32
	DT_FLOAT32
	//float64
	DT_FLOAT64
	//string
	DT_STRING
	//object(indent + serial)
	DT_OBJECT
	//unsafe.Pointer
	DT_POINTER
	DT_MAX
)

const (
	NULL_BOOLEAN = false
	NULL_INT     = int(0)
	NULL_INT32   = int32(0)
	NULL_INT64   = int64(0)
	NULL_FLOAT32 = float32(0)
	NULL_FLOAT64 = float64(0)
	NULL_STRING  = float64(0)
)

var NULL_GUID GUID = GUID{0, 0}

type Data struct {
	dataType DATA_TYPE
	pointer  unsafe.Pointer
	buffer   [2]uint64
}

func (self *Data) reset(t DATA_TYPE) {
	self.buffer[0] = 0
	self.buffer[1] = 0
	self.pointer = (unsafe.Pointer)(&self.buffer)
	self.dataType = t
}

func (self *Data) GetType() DATA_TYPE {
	return self.dataType
}

func (self *Data) SetDefaultValue(t DATA_TYPE) {

}

func (self *Data) IsNullValue() bool {
	return false
}

func (self *Data) GetBool() bool {
	util.Assert(self.dataType == DT_BOOLEAN, "self.dataType == DT_BOOLEAN")
	return *(*bool)(self.pointer)
}

func (self *Data) GetInt() int {
	util.Assert(self.dataType == DT_INT, "self.dataType == DT_INT")
	return *(*int)(self.pointer)
}

func (self *Data) GetInt32() int32 {
	util.Assert(self.dataType == DT_INT32, "self.dataType == DT_INT32")
	return *(*int32)(self.pointer)
}

func (self *Data) GetInt64() int64 {
	util.Assert(self.dataType == DT_INT64, "self.dataType == DT_INT64")
	return *(*int64)(self.pointer)
}

func (self *Data) GetFloat32() float32 {
	util.Assert(self.dataType == DT_FLOAT32, "self.dataType == DT_FLOAT32")
	return *(*float32)(self.pointer)
}

func (self *Data) GetFloat64() float64 {
	util.Assert(self.dataType == DT_FLOAT64, "self.dataType == DT_FLOAT64")
	return *(*float64)(self.pointer)
}

func (self *Data) GetString() *string {
	util.Assert(self.dataType == DT_STRING, "self.dataType == DT_STRING")
	return (*string)(self.pointer)
}

func (self *Data) GetObject() GUID {
	util.Assert(self.dataType == DT_OBJECT, "self.dataType == DT_OBJECT")
	return GUID{LOW: self.buffer[0], High: self.buffer[1]}
}

func (self *Data) GetPointer() unsafe.Pointer {
	util.Assert(self.dataType == DT_POINTER, "self.dataType == DT_POINTER")
	return self.pointer
}

func (self *Data) SetUnknown() {
	self.reset(DT_UNKNOWN)
}

func (self *Data) SetBool(value bool) {
	self.reset(DT_BOOLEAN)
	*(*bool)(self.pointer) = value
}
