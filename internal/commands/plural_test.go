package commands

import "testing"

func TestPlural(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    int
		want string
	}{
		{name: "zero", n: 0, want: ""},
		{name: "one", n: 1, want: ""},
		{name: "many", n: 2, want: "s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := plural(tt.n); got != tt.want {
				t.Errorf("plural(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}
