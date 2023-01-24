package gorgeouslasagna

import "testing"

func TestRemainingOvenTime(t *testing.T) {
	got := RemainingOvenTime(30)
	expected := 10

	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}

}

func TestPreparationTime(t *testing.T) {
	got := PreparationTime(2)
	expected := 4

	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}

}

func TestElapsedTime(t *testing.T) {
	got := ElapsedTime(3, 20)
	expected := 26

	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
}
