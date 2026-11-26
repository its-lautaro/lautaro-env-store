package store

import (
	"slices"
	"testing"
)

func TestStore_Entries(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		content []string
		want    []Entry
	}{
		{name: "missing file", content: nil, want: nil},
		{
			name:    "lists entries in file order",
			content: []string{"# header\nexport B='2'  # saved s2\nexport A='1'  # saved s1\n"},
			want:    []Entry{{Name: "B", Saved: "s2"}, {Name: "A", Saved: "s1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := newTestStore(t, tt.content...).Entries()
			if err != nil {
				t.Fatalf("Entries() error = %v", err)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("Entries() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestStore_Entries_Unreadable(t *testing.T) {
	t.Parallel()
	if _, err := (Store{Path: t.TempDir()}).Entries(); err == nil {
		t.Error("Entries() sobre un directorio: error = nil, want error")
	}
}
