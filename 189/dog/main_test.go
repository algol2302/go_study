package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := 10
	dn := Years(n)
	if dn != 70 {
		t.Error("got", dn, "want", 70)
	}
}

func TestYearsTwo(t *testing.T) {
	n := 7
	dn := YearsTwo(n)
	if dn != 49 {
		t.Error("got", dn, "want", 49)
	}
}

func ExampleYears() {
	n := 10
	fmt.Println(Years(n))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	n := 7
	fmt.Println(YearsTwo(n))
	// Output:
	// 49
}

var n int = 10

func BenchmarkYears(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Years(n)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		YearsTwo(n)
	}
}
