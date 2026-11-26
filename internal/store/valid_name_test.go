package store

import "testing"

func TestValidName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "upper", input: "API_USER"},
		{name: "lower", input: "api_user"},
		{name: "leading underscore", input: "_X"},
		{name: "single underscore", input: "_"},
		{name: "digits after first", input: "A1_2"},
		{name: "empty", input: "", wantErr: true},
		{name: "leading digit", input: "1A", wantErr: true},
		{name: "dash", input: "A-B", wantErr: true},
		{name: "space", input: "A B", wantErr: true},
		{name: "dot", input: "A.B", wantErr: true},
		{name: "trailing newline", input: "A\n", wantErr: true},
		{name: "non ascii", input: "AÑO", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validName(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
