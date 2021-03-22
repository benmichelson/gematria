package gematria

import (
	"math"
	"testing"
)

func TestValue(t *testing.T) {
	v, _ := Value("אמת")
	if v != 441 {
		t.Error(v)
	}
	v, _ = Value("non hebrew text")
	if v != 0 {
		t.Error(v)
	}

	v, _ = Value("אמת ויציב")
	if v != 559 {
		t.Error(v)
	}

	v, _ = Value("מים")
	if v != 90 {
		t.Error(v)
	}

	v, _ = Value("some hebrew text: שמחה")
	if v != 353 {
		t.Error(v)
	}

}
func TestAdd(t *testing.T) {
	v, ok := add(100, 100)
	if v != 200 || !ok {
		t.Error(v, ok)
	}
	v, ok = add(-100, -100)
	if v != -200 || !ok {
		t.Error(v, ok)
	}
	v, ok = add(-100, 100)
	if v != 0 || !ok {
		t.Error(v, ok)
	}
	v, ok = add(math.MaxInt64, 1)
	if ok {
		t.Error(ok)
	}
	v, ok = add(math.MinInt64, -1)
	if ok {
		t.Error(ok)
	}
}
