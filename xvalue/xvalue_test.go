package xvalue

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
)

func TestValue(t *testing.T)  {
	as := assert.New(t)
	v := &XValue{}
	as.True(v.Kind() == Invalid, "kind is invalid")

	v = NewValue(Bool)
	as.True(v.Bool() == false)
	v.SetBool(true)
	as.True(v.Bool() == true)
	v.SetBool(false)
	as.True(v.Bool() == false)

	v = NewValue(Int64)
	v.SetInt64(math.MaxInt64)
	as.True(v.Int64() == math.MaxInt64)

	v = NewValue(UserData)
	v.SetUserData(100)
	as.True(v.UserDdata().(int) == 100)
}

func BenchmarkValue_Bool(b *testing.B) {
	v := NewValue(Bool)
	boolean := b.N % 2 == 0
	v.SetBool(boolean)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v.Bool()
	}
}