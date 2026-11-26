package store

import (
	"slices"
	"testing"
)

func TestStore_readLines(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		content []string // nil = el archivo no existe
		want    []string
	}{
		{name: "missing file", content: nil, want: nil},
		{name: "empty file", content: []string{""}, want: nil},
		{name: "only newlines", content: []string{"\n\n"}, want: nil},
		{name: "trailing newlines trimmed", content: []string{"a\nb\n\n"}, want: []string{"a", "b"}},
		{name: "no trailing newline", content: []string{"a\nb"}, want: []string{"a", "b"}},
		{name: "inner blank lines kept", content: []string{"a\n\nb\n"}, want: []string{"a", "", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := newTestStore(t, tt.content...)
			got, err := s.readLines()
			if err != nil {
				t.Fatalf("readLines() error = %v", err)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("readLines() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStore_readLines_Unreadable(t *testing.T) {
	t.Parallel()
	s := Store{Path: t.TempDir()} // un directorio no se puede leer como archivo
	if _, err := s.readLines(); err == nil {
		t.Error("readLines() sobre un directorio: error = nil, want error")
	}
}
