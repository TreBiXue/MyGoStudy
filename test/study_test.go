package test

import (
	"fmt"
	"testing"
)

// TestAdd 单元测试
func TestAdd(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		result := add(1, 2)
		t.Log("TEST result:", result)
	})

	t.Run("Test2", func(t *testing.T) {
		result := add(99999, 333)
		fmt.Println("11111111111")
		t.Log("TEST result:", result)
	})
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		add(5, 6)
		// b.Log("TEST result:", result)
		b.StopTimer()
	}

	// 报告内存分配情况
	b.ReportAllocs()
}

func add(x int, y int) int {
	return x + y
}
