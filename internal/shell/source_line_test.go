package shell

import "testing"

func TestSourceLine(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		storePath string
		marker    string
		want      string
	}{
		{
			name:      "plain path",
			storePath: "/home/u/.env",
			marker:    "tool",
			want:      `[ -f "/home/u/.env" ] && source "/home/u/.env"  # tool`,
		},
		{
			name:      "path with spaces is quoted",
			storePath: "/home/u/my env",
			marker:    "tool",
			want:      `[ -f "/home/u/my env" ] && source "/home/u/my env"  # tool`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sourceLine(tt.storePath, tt.marker); got != tt.want {
				t.Errorf("sourceLine() = %q, want %q", got, tt.want)
			}
		})
	}
}
