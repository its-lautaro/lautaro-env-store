package store

import "testing"

func TestFormatLine(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		v     Var
		stamp string
		want  string
	}{
		{
			name:  "plain value",
			v:     Var{Name: "API_USER", Value: "usuario"},
			stamp: "2026-01-02 03:04",
			want:  "export API_USER='usuario'  # saved 2026-01-02 03:04",
		},
		{
			name:  "value with single quote",
			v:     Var{Name: "X", Value: "it's"},
			stamp: "2026-01-02 03:04",
			want:  `export X='it'"'"'s'  # saved 2026-01-02 03:04`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := formatLine(tt.v, tt.stamp); got != tt.want {
				t.Errorf("formatLine() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFormatLine_RoundTripsThroughParseLine(t *testing.T) {
	t.Parallel()
	line := formatLine(Var{Name: "TOKEN", Value: "a b c"}, "2026-11-19 10:00")
	got, ok := parseLine(line)
	if !ok {
		t.Fatalf("parseLine(%q) no reconoció la línea", line)
	}
	want := Entry{Name: "TOKEN", Saved: "2026-11-19 10:00"}
	if got != want {
		t.Errorf("parseLine(formatLine()) = %+v, want %+v", got, want)
	}
}
