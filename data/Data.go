package data

import (
	"unsafe"

	"github.com/PureDu/golib/util"
)

type DATATYPE int32

const (
	DTUNKNOWN DATATYPE = iota
	//bool
	DTBOOLEAN
	//int
	DTINT
	//int32
	DTINT32
	//int64
	DTINT64
	//float32
	DTFLOAT32
	//float64
	DTFLOAT64
	//string
	DTSTRING
	//object(indent + serial)
	DTOBJECT
	//unsafe.Pointer
	DTPOINTER
	DTMAX
)

const (
	NULLBOOLEAN bool    = false
	NULLINT     int     = 0
	NULLINT32   int32   = 0
	NULLINT64   int64   = 0
	NULLFLOAT32 float32 = 0
	NULLFLOAT64 float64 = 0
	NULLSTRING  string  = ""
)

var NULLGUID GUID = GUID{0, 0}

type Data struct {
	dataType DATATYPE
	pointer  unsafe.Pointer
	buffer   [2]uint64
}

func (self *Data) reset(t DATATYPE) {
	self.buffer[0] = 0
	self.buffer[1] = 0
	self.pointer = (unsafe.Pointer)(&self.buffer)
	self.dataType = t
}

func (self *Data) GetType() DATATYPE {
	return self.dataType
}

func (self *Data) SetDefaultValue(t DATATYPE) {
	switch t {
	case DTBOOLEAN:
		self.SetBool(NULLBOOLEAN)
	case DTINT:
		self.SetInt(NULLINT)
	case DTINT32:
		self.SetInt32(NULLINT32)
	case DTINT64:
		self.SetInt64(NULLINT64)
	case DTFLOAT32:
		self.SetFloat32(NULLFLOAT32)
	case DTFLOAT64:
		self.SetFloat64(NULLFLOAT64)
	case DTSTRING:
		self.SetString(NULLSTRING)
	case DTOBJECT:
		self.SetObject(NULLGUID)
	default:
		util.Assert(false, "don't have data type")
	}
}

func (self *Data) IsNullValue() bool {

	switch self.GetType() {
	case DTBOOLEAN:
		return self.GetBool() == NULLBOOLEAN
	case DTINT:
		return self.GetInt() == NULLINT
	case DTINT32:
		   return self.GetInt32() == NULLINT32
	case DTINT64:
		 return self.GetInt64() == NULLINT64
	case DTFLOAT32:
		 return self.GetFloat32() == 
	case DTFLOAT64:
	case DTSTRING:
		self.SetString(NULLSTRING)
	case DTOBJECT:
		self.SetObject(NULLGUID)
	default:
		util.Assert(false, "don't have data type")
	}
}

func (self *Data) GetBool() bool {
	util.Assert(self.dataType == DTBOOLEAN, "self.dataType == DT_BOOLEAN")
	return *(*bool)(self.pointer)
}

func (self *Data) GetInt() int {
	util.Assert(self.dataType == DTINT, "self.dataType == DT_INT")
	return *(*int)(self.pointer)
}

func (self *Data) GetInt32() int32 {
	util.Assert(self.dataType == DTINT32, "self.dataType == DTINT32")
	return *(*int32)(self.pointer)
}

func (self *Data) GetInt64() int64 {
	util.Assert(self.dataType == DTINT64, "self.dataType == DTINT64")
	return *(*int64)(self.pointer)
}

func (self *Data) GetFloat32() float32 {
	util.Assert(self.dataType == DTFLOAT32, "self.dataType == DT_FLOAT32")
	return *(*float32)(self.pointer)
}

func (self *Data) GetFloat64() float64 {
	util.Assert(self.dataType == DTFLOAT64, "self.dataType == DT_FLOAT64")
	return *(*float64)(self.pointer)
}

func (self *Data) GetString() *string {
	util.Assert(self.dataType == DTSTRING, "self.dataType == DT_STRING")
	return (*string)(self.pointer)
}

func (self *Data) GetObject() GUID {
	util.Assert(self.dataType == DTOBJECT, "self.dataType == DT_OBJECT")
	return GUID{LOW: self.buffer[0], High: self.buffer[1]}
}

func (self *Data) GetPointer() unsafe.Pointer {
	util.Assert(self.dataType == DTPOINTER, "self.dataType == DT_POINTER")
	return self.pointer
}

func (self *Data) SetUnknown() {
	self.reset(DTUNKNOWN)
}

func (self *Data) SetBool(value bool) {
	self.reset(DTBOOLEAN)
	*(*bool)(self.pointer) = value
}

func (self *Data) SetInt(value int) {
	self.reset(DTINT)
	*(*int)(self.pointer) = value
}

func (self *Data) SetInt32(value int32) {
	self.reset(DTINT32)
	*(*int32)(self.pointer) = value
}

func (self *Data) SetInt64(value int64) {
	self.reset(DTINT64)
	*(*int64)(self.pointer) = value
}

func (self *Data) SetFloat32(value float32) {
	self.reset(DTFLOAT32)
	*(*float32)(self.pointer) = value
}

func (self *Data) SetFloat64(value float64) {
	self.reset(DTFLOAT64)
	*(*float64)(self.pointer) = value
}

func (self *Data) SetString(value string) {
	self.reset(DTSTRING)
	self.pointer = unsafe.Pointer(&value)
}

func (self *Data) SetObject(value GUID) {
	self.reset(DTOBJECT)
	self.buffer[0] = value.LOW
	self.buffer[1] = value.High
}

func (self *Data) SetPointer(value unsafe.Pointer) {
	self.reset(DTPOINTER)
	self.pointer = value
}
