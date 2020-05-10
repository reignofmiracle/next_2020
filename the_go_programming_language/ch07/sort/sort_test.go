package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestStringSlice(t *testing.T) {
	names := []string{"z", "a", "b", "c"}
	sort.Sort(StringSlice(names))
	fmt.Println(names)

	expected := []string{"a", "b", "c", "z"}
	if names[0] != expected[0] {
		t.Errorf("%v = %v", names, expected)
	}
}
