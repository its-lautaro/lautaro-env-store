package color

import "testing"

func TestDetect_NoColor(t *testing.T) {
	t.Setenv("NO_COLOR", "")
	if detect() {
		t.Error("detect() = true con NO_COLOR definida (aunque vacía), want false")
	}
}
