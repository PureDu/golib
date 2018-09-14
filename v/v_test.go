package v

import "testing"
import "github.com/PureDu/golib/xvalue"
import (
	"github.com/PureDu/golib/data"
)

func BenchmarkValue_Bool(b *testing.B) {
	v := xvalue.NewValue(xvalue.Bool)
	boolean := b.N % 2 == 0
	v.SetBool(boolean)
	b.ResetTimer()
	b.ReportAllocs()
	a := true
	for i := 0; i < b.N; i++ {
		a = v.Bool()
	}
	_ = a
}

func BenchmarkGetBool(b *testing.B) {
	data := data.NewData()
	boolean := b.N%2 == 0
	data.SetBool(boolean)
	b.ResetTimer()
	b.ReportAllocs()
	a := true
	for i := 0; i < b.N; i++ {
		a = data.GetBool()
	}
	_ = a
}

func BenchmarkBool(b *testing.B) {

	tb := b.N%2 == 0
	b.ResetTimer()
	b.ReportAllocs()
	a := true
	for i := 0; i < b.N; i++ {
		a = tb
	}
	_ = a
}