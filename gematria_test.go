package gematria

import (
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

	const MaxInt64 = 1<<63 - 1
	const MinInt64 = -1 << 63

	_, ok = add(MaxInt64, 1)
	if ok {
		t.Error(ok)
	}
	_, ok = add(MinInt64, -1)
	if ok {
		t.Error(ok)
	}
}

func TestHebrew(t *testing.T) {
	v, _ := Hebrew(115)
	if v != "קטו" {
		t.Error(v)
	}
	v, _ = Hebrew(744)
	if v != "תשדמ" {
		t.Error(v)
	}

	for i := 1; i < 1000; i++ {
		v, _ := Hebrew(i)
		j, _ := Value(v)
		if i != j {
			t.Error(i)
		}
	}
}

func TestGeresh(t *testing.T) {
	v, _ := Hebrew(1)
	v, _ = AddGeresh(v)
	if v != "א׳" {
		t.Error(v)
	}

	v, _ = Hebrew(115)
	v, _ = AddGeresh(v)
	if v != "קט״ו" {
		t.Error(v)
	}

	v, _ = Hebrew(116)
	v, _ = AddGeresh(v)
	if v != "קט״ז" {
		t.Error(v)
	}

	v, _ = Hebrew(744)
	v, _ = AddGeresh(v)
	if v != "תשד״מ" {
		t.Error(v)
	}

	v, _ = Hebrew(999)
	v, _ = AddGeresh(v)
	if v != "תתקצ״ט" {
		t.Error(v)
	}
}
