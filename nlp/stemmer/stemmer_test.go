package stemmer

import "testing"

func TestStem(t *testing.T) {
	stem := Stem("going")
	println(stem)
}
