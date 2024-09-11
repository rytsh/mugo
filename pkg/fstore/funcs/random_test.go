package funcs

import (
	"testing"

	"github.com/rytsh/mugo/pkg/fstore/funcs/values"
)

func TestRandom(t *testing.T) {
	values.RandomSeed(1)

	r := Random{}

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

	if v := r.Float(1, 10); v != 4.5299456483862235 {
		t.Errorf("Amount() = %v, want %v", v, 4.5299456483862235)
	}
}
