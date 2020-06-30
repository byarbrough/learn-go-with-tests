package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expect := 4

	if sum != expect {
		t.Errorf("got %d expect %d", sum, expect)
	}
}

func ExampleAdd() {
	sum := Add(10, 6)
	fmt.Println(sum)
	// Output: 16
}
