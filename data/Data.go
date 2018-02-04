package data

import (
	"unsafe"

	"github.com/PureDu/golib/math"
	"github.com/PureDu/golib/util"
)

type Dataer interface {
	CopyFrom(src Dataer)
	Clone() Dataer
	GetType() DATATYPE

	SetDefaultValue(t DATATYPE)

	IsNullValue() bool

	GetBool() bool

	GetInt() int

	GetInt32() int32

	GetInt64() int64

	GetFloat32() float32

	GetFloat64() float64

	GetString() string

	GetObject() GUID

	GetPointer() unsafe.Pointer

	SetUnknown()

	SetBool(value bool)

	SetInt(value int)

	SetInt32(value int32)

	SetInt64(value int64)

	SetFloat32(value float32)

	SetFloat64(value float64)

	SetString(value string)

	SetObject(value GUID)

	SetPointer(value unsafe.Pointer)
}

var _ Dataer = &data{}

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

var dataTytpNameMap = map[DATATYPE]string{
	DTUNKNOWN: "DTUNKNOWN",
	DTBOOLEAN: "DTBOOLEAN",
	DTINT:     "DTINT",
	DTINT32:   "DTINT32",
	DTINT64:   "DTINT64",
	DTFLOAT32: "DTFLOAT32",
	DTFLOAT64: "DTFLOAT64",
	DTSTRING:  "DTSTRING",
	DTOBJECT:  "DTOBJECT",
	DTPOINTER: "DTPOINTER",
}

func (dt DATATYPE) String() string {
	return dataTytpNameMap[dt]
}

const (
	NULLBOOLEAN bool    = false
	NULLINT     int     = 0
	NULLINT32   int32   = 0
	NULLINT64   int64   = 0
	NULLFLOAT32 float32 = 0
	NULLFLOAT64 float64 = 0
	NULLSTRING  string  = ""
)

var NULLGUID GUID = NewGUID(0, 0)

type data struct {
	dataType DATATYPE
	pointer  unsafe.Pointer
	buffer   [2]uint64
}

func NewData() Dataer {
	return newData()
}

func newData() *data {
	d := &data{}
	d.reset(DTUNKNOWN)
	return d
}
func (self *data) reset(t DATATYPE) {
	self.buffer[0] = 0
	self.buffer[1] = 0
	self.pointer = (unsafe.Pointer)(&self.buffer)
	self.dataType = t
}

func (self *data) CopyFrom(src Dataer) {
	srcData := src.(*data)
	self.copyFrom(srcData)
}

func (self *data) copyFrom(src *data) {
	self.dataType = src.dataType
	self.buffer = src.buffer

	if self.dataType == DTSTRING ||
		self.dataType == DTPOINTER {
		self.pointer = src.pointer
	}
}

func (self *data) Clone() Dataer {
	cloneData := newData()
	cloneData.copyFrom(self)
	return cloneData
}
func (self *data) GetType() DATATYPE {
	return self.dataType
}

func (self *data) SetDefaultValue(t DATATYPE) {
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
	case DTPOINTER:
		self.SetPointer(nil)
	default:
		util.Assert(false, "don't have data type")
	}
}

func (self *data) IsNullValue() bool {

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
		return math.IsZeroFloat32(self.GetFloat32())
	case DTFLOAT64:
		return math.IsZeroFloat64(self.GetFloat64())
	case DTSTRING:
		return self.GetString() == NULLSTRING
	case DTOBJECT:
		return self.GetObject() == NULLGUID
	case DTPOINTER:
		return self.GetPointer() == nil
	}
	return false
}

func (self *data) GetBool() bool {
	if !util.Assert(self.dataType == DTBOOLEAN,
		"self.dataType == DT_BOOLEAN") {
		return NULLBOOLEAN
	}
	return *(*bool)(self.pointer)
}

func (self *data) GetInt() int {
	if !util.Assert(self.dataType == DTINT,
		"self.dataType == DT_INT") {
		return NULLINT
	}
	return *(*int)(self.pointer)
}

func (self *data) GetInt32() int32 {
	if !util.Assert(self.dataType == DTINT32,
		"self.dataType == DTINT32") {
		return NULLINT32
	}
	return *(*int32)(self.pointer)
}

func (self *data) GetInt64() int64 {
	if !util.Assert(self.dataType == DTINT64,
		"self.dataType == DTINT64") {
		return NULLINT64
	}
	return *(*int64)(self.pointer)
}

func (self *data) GetFloat32() float32 {
	if !util.Assert(self.dataType == DTFLOAT32,
		"self.dataType == DT_FLOAT32") {
		return NULLFLOAT32
	}
	return *(*float32)(self.pointer)
}

func (self *data) GetFloat64() float64 {
	if !util.Assert(self.dataType == DTFLOAT64,
		"self.dataType == DT_FLOAT64") {
		return NULLFLOAT64
	}
	return *(*float64)(self.pointer)
}

func (self *data) GetString() string {
	if !util.Assert(self.dataType == DTSTRING,
		"self.dataType == DT_STRING") {

		return NULLSTRING
	}
	return *(*string)(self.pointer)
}

func (self *data) GetObject() GUID {
	if !util.Assert(self.dataType == DTOBJECT,
		"self.dataType == DT_OBJECT") {
		return NULLGUID
	}
	return NewGUID(self.buffer[0], self.buffer[1])
}

func (self *data) GetPointer() unsafe.Pointer {
	if !util.Assert(self.dataType == DTPOINTER,
		"self.dataType == DT_POINTER") {
		return nil
	}
	return self.pointer
}

func (self *data) SetUnknown() {
	self.reset(DTUNKNOWN)
}

func (self *data) SetBool(value bool) {
	self.reset(DTBOOLEAN)
	*(*bool)(self.pointer) = value
}

func (self *data) SetInt(value int) {
	self.reset(DTINT)
	*(*int)(self.pointer) = value
}

func (self *data) SetInt32(value int32) {
	self.reset(DTINT32)
	*(*int32)(self.pointer) = value
}

func (self *data) SetInt64(value int64) {
	self.reset(DTINT64)
	*(*int64)(self.pointer) = value
}

func (self *data) SetFloat32(value float32) {
	self.reset(DTFLOAT32)
	*(*float32)(self.pointer) = value
}

func (self *data) SetFloat64(value float64) {
	self.reset(DTFLOAT64)
	*(*float64)(self.pointer) = value
}

func (self *data) SetString(value string) {
	self.reset(DTSTRING)
	self.pointer = unsafe.Pointer(&value)
}

func (self *data) SetObject(value GUID) {
	self.reset(DTOBJECT)
	self.buffer[0] = value.high
	self.buffer[1] = value.low
}

func (self *data) SetPointer(value unsafe.Pointer) {
	self.reset(DTPOINTER)
	self.pointer = value
}
