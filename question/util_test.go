package question

import "testing"

func BenchmarkPrint2lSlice(b *testing.B) {
	arrList := [][]int{
		//{1, 2},
		//{3, 6},
		{8, 10},
	}
	for i := 0; i < b.N; i++ {
		Print2lSlice(arrList)
	}
}

func BenchmarkInitNodeList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InitNodeList(10)
	}
}
