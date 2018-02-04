package data

import (
	"testing"
	"unsafe"

	"github.com/PureDu/golib/math"

	"github.com/stretchr/testify/assert"
)

func TestData(t *testing.T) {
	data := NewData()
	assert := assert.New(t)
	assert.True(data.GetType() == DTUNKNOWN, "data.GetType() == DTUNKNOWN")

	for t := DTBOOLEAN; t < DTMAX; t++ {
		data.SetDefaultValue(t)
		assert.True(data.GetType() == t, t)
		assert.True(data.IsNullValue(), t.String())
	}

	data.SetBool(true)
	assert.True(data.GetBool() == true)
	data.SetBool(false)
	assert.True(data.GetBool() == false)

	data.SetInt(100)
	assert.True(data.GetInt() == 100)

	data.SetInt32(132)
	assert.True(data.GetInt32() == 132)
	data.SetInt64(164)
	assert.True(data.GetInt64() == 164)

	data.SetFloat32(32.0)
	assert.True(math.IsFloat32Equal(data.GetFloat32(), 32.0))

	data.SetFloat64(64.0)
	assert.True(math.IsFloat64Equal(data.GetFloat64(), 64.0))

	data.SetString("Hello")
	assert.True(data.GetString() == "Hello")

	object := NewGUID(100, 100)
	data.SetObject(object)
	assert.True(data.GetObject() == object)

	data.SetInt(100)
	cloneData := data.Clone()
	assert.True(data.GetType() == cloneData.GetType())
	assert.True(data.GetInt() == cloneData.GetInt())
	data.SetInt(1)
	assert.True(data.GetType() == cloneData.GetType())
	assert.True(data.GetInt() != cloneData.GetInt())
	data.SetBool(true)
	assert.True(data.GetBool() == true)
	assert.True(data.GetType() != cloneData.GetType())

	copyData := NewData()
	data.SetString("Hello world")
	copyData.CopyFrom(data)
	assert.True(data.GetType() == copyData.GetType())
	assert.True(data.GetString() == copyData.GetString())
	copyData.SetString("Hello")
	assert.True(data.GetType() == copyData.GetType())
	assert.True(data.GetString() != copyData.GetString())

	pointer := unsafe.Pointer(&object)
	i := 0
	pointer2 := unsafe.Pointer(&i)
	data.SetPointer(pointer)
	assert.True(data.GetPointer() == pointer)
	copyData.CopyFrom(data)
	assert.True(data.GetType() == copyData.GetType())
	assert.True(data.GetPointer() == copyData.GetPointer())

	copyData.SetPointer(pointer2)
	assert.True(data.GetType() == copyData.GetType())
	assert.True(data.GetPointer() != copyData.GetPointer())

	data = NewData()

	assert.True(data.GetBool() == false)
	assert.True(data.GetInt() == 0)
	assert.True(data.GetInt32() == 0)
	assert.True(data.GetInt64() == 0)
	assert.True(math.IsZeroFloat32(data.GetFloat32()))
	assert.True(math.IsZeroFloat64(data.GetFloat64()))
	assert.True(data.GetObject() == NewGUID(0, 0))
	assert.True(data.GetString() == "")
	assert.True(data.GetPointer() == nil)
}

func BenchmarkSetBool(b *testing.B) {
	data := NewData()
	boolean := b.N%2 == 0
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data.SetBool(boolean)
	}
}

func BenchmarkGetBool(b *testing.B) {
	data := NewData()
	boolean := b.N%2 == 0
	data.SetBool(boolean)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data.GetBool()
	}
}

func BenchmarkSetInt64(b *testing.B) {
	data := NewData()
	i64 := int64(b.N)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data.SetInt64(i64)
	}
}

func BenchmarkGetInt64(b *testing.B) {
	data := NewData()
	i64 := int64(b.N)
	data.SetInt64(i64)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data.GetInt64()
	}
}

func BenchmarkSetString(b *testing.B) {
	data := NewData()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data.SetString("Hello Wrold")
	}
}

func BenchmarkGetString(b *testing.B) {
	data := NewData()
	data.SetString("Hello")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data.GetString()
	}
}
