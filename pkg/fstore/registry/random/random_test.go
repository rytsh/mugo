package random

import (
	"math"
	"math/rand"
	"testing"
)

func TestRandom(t *testing.T) {
	r := Random{
		Random: rand.New(rand.NewSource(1)),
	}

	if v := r.Intn(1, 10); v != 6 {
		t.Errorf("Intn() = %v, want %v", v, 6)
	}

	if v := r.Alpha(10); v != "MVjwGebzRv" {
		t.Errorf("Alpha() = %v, want %v", v, "MVjwGebzRv")
	}

	if v := r.AlphaNum(10); v != "s8csTJLqDD" {
		t.Errorf("AlphaNum() = %v, want %v", v, "s8csTJLqDD")
	}

	if v := r.Ascii(10); v != "E#-8M(]kbO" {
		t.Errorf("Ascii() = %v, want %v", v, "E#-8M(]kbO")
	}

	if v := r.Numeric(10); v != "5834232601" {
		t.Errorf("Numeric() = %v, want %v", v, "5834232601")
	}

	if v := r.Float(1, 10); math.Floor(v) != float64(4) {
		t.Errorf("Amount() = %v, want %v", math.Floor(v), 4)
	}
}
