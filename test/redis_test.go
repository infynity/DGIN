package test

import (
	"dcgin/model"

	"testing"
)
// 660           4210038 ns/op            9393 B/op         27 allocs/op
func BenchmarkGetTest1(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		model.GetTest1()
	}
}

// 1684           1652576 ns/op              71 B/op          4 allocs/op
func BenchmarkGetTest2(b *testing.B) {
	model.Init()
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		model.GetTest2()
	}
}
