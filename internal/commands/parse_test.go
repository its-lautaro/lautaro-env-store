package commands

import (
	"slices"
	"strings"
	"testing"

	"github.com/its-lautaro/lautaro-env-store/internal/store"
)

func TestParse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		args        []string
		want        []store.Var
		wantErrPart string
	}{
		{name: "single", args: []string{"A=1"}, want: []store.Var{{Name: "A", Value: "1"}}},
		{name: "empty value", args: []string{"A="}, want: []store.Var{{Name: "A", Value: ""}}},
		{name: "value with equals", args: []string{"A=b=c"}, want: []store.Var{{Name: "A", Value: "b=c"}}},
		{
			name: "several keep order",
			args: []string{"B=2", "A=1"},
			want: []store.Var{{Name: "B", Value: "2"}, {Name: "A", Value: "1"}},
		},
		{name: "missing equals", args: []string{"A=1", "nope", "otro"}, wantErrPart: "no entiendo: nope otro"},
		{name: "duplicate name", args: []string{"A=1", "A=2"}, wantErrPart: "'A' está repetida"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := parse(tt.args)
			if tt.wantErrPart != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErrPart) {
					t.Fatalf("parse(%q) error = %v, want que contenga %q", tt.args, err, tt.wantErrPart)
				}
				return
			}
			if err != nil {
				t.Fatalf("parse(%q) error = %v", tt.args, err)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("parse(%q) = %+v, want %+v", tt.args, got, tt.want)
			}
		})
	}
}
