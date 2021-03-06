package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("James")
	expected := "Welcome my dear James"
	if got != expected {
		t.Error("got", got, "expected", expected)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
