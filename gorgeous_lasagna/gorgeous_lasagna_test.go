package gorgeouslasagna

import "testing"

func TestRemainingOvenTime(t *testing.T) {
	got := RemainingOvenTime(30)
	expected := 11

	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}

}
