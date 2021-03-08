package word

import (
	"fmt"
	"testing"
)

var s1 string = "A simple string with six words"
var s2 string = "A simple string with two words test and test"

func TestCount(t *testing.T) {
	n := Count(s1)
	if n != 6 {
		t.Error("got", n, "want", 6)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount(s2)

	for k, v := range m {
		switch k {
		case "test":
			if v != 2 {
				t.Error("got", v, "want", 2)
			}
		default:
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		}
	}
}

func ExampleCount() {
	s := "A simple string with six words"
	fmt.Println(Count(s))
	// Output:
	// 6
}

func ExampleUseCount() {
	s := "A simple string with two words test and test"
	fmt.Println(UseCount(s))
	// Output:
	// map[A:1 and:1 simple:1 string:1 test:2 two:1 with:1 words:1]
}

func BenchmarkCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Count(s1)
	}
}

func BenchmarkUseCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UseCount(s2)
	}
}
