package nlp

import (
	"reflect"
	"testing"
)

// The Go Tool will remove the "*_test.go" from the executable in the end

//It is a test because:
// - It begins with "Test*" prefix
// - It has a param of a type [*testing.T]
func TestTokenize(t *testing.T) {
	text := "a b c going"
	expected := []string{"a", "b", "c", "go"}
	actual := Tokenize(text)

	//there is no comparator for slices in Go

	if !reflect.DeepEqual(expected, actual) {
		//t.Fail()
		t.Fatal("asd")
	}

}

//Test that uses a lot of data
func TestTokenize2(t *testing.T) {
	testCases := []struct {
		text     string
		expected []string
	}{
		{"", nil},
		{"hi", []string{"hi"}},
		{"HI", []string{"hi"}},
		{"a b c", []string{"a", "b", "c"}},
	}

	for _, tc := range testCases {
		// you can use [t.Run] in order to break these into separate tests
		actual := Tokenize(tc.text)
		expected := tc.expected
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("Expected: %v, Actual: %v", expected, actual)
		}
	}
}
