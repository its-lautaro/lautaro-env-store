package store

import (
	"slices"
	"testing"
)

func TestStore_Delete(t *testing.T) {
	t.Parallel()
	const initial = "# header\nexport A='1'\nexport B='2'\nexport C='3'\n"
	tests := []struct {
		name        string
		names       []string
		wantRemoved []string
		wantMissing []string
		wantFile    string
	}{
		{
			name:        "removes existing",
			names:       []string{"A", "C"},
			wantRemoved: []string{"A", "C"},
			wantFile:    "# header\nexport B='2'\n",
		},
		{
			name:        "reports missing",
			names:       []string{"B", "Z"},
			wantRemoved: []string{"B"},
			wantMissing: []string{"Z"},
			wantFile:    "# header\nexport A='1'\nexport C='3'\n",
		},
		{
			name:        "all missing leaves file intact",
			names:       []string{"Y", "Z"},
			wantMissing: []string{"Y", "Z"},
			wantFile:    initial,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := newTestStore(t, initial)
			removed, missing, err := s.Delete(tt.names)
			if err != nil {
				t.Fatalf("Delete() error = %v", err)
			}
			if !slices.Equal(removed, tt.wantRemoved) {
				t.Errorf("removed = %q, want %q", removed, tt.wantRemoved)
			}
			if !slices.Equal(missing, tt.wantMissing) {
				t.Errorf("missing = %q, want %q", missing, tt.wantMissing)
			}
			if got := readFile(t, s.Path); got != tt.wantFile {
				t.Errorf("contenido = %q, want %q", got, tt.wantFile)
			}
		})
	}
}

func TestStore_Delete_Unreadable(t *testing.T) {
	t.Parallel()
	s := Store{Path: t.TempDir()}
	if _, _, err := s.Delete([]string{"A"}); err == nil {
		t.Error("Delete() sobre un directorio: error = nil, want error")
	}
}
