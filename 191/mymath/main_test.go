package mymath

import (
	"fmt"
	"testing"
)

var xi []int = []int{1, 2, 3, 4, 5, 1000}

func TestCenteredAvg(t *testing.T) {
	n := CenteredAvg(xi)
	wanted := float64(2+3+4+5) / 4
	if n != wanted {
		t.Error("got", n, "want", wanted)
	}
}

func ExampleCenteredAvg() {
	xi := []int{1, 2, 3, 4, 5, 1000}
	fmt.Println(CenteredAvg(xi))
	// Output:
	// 3.5
}

func BenchmarkCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}
