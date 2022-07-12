package vm

import (
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//Add(ZInt(1), ZInt(1))
		//Add(ZInt(1), ZDouble(1))
		//Add(ZInt(1), ZBool(true))
		//Add(ZDouble(1), ZInt(1))
		Add(nil, ZBool(false), ZBool(false))
		//Add(ZDouble(1), ZBool(true))
		//Add(ZBool(true), ZInt(1))
		//Add(ZBool(true), ZDouble(1))
		//Add(ZBool(true), ZBool(true))
	}
}

func BenchmarkDefaultAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if float64(1)+float64(1) > 0 {

		}
	}
}
