package unpack

import "testing"

func TestStrUnpack(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{name: "Empty string",
			in:      "",
			want:    "",
			wantErr: false},

		{name: "Ordinary string",
			in:      "a4bc2d5e",
			want:    "aaaabccddddde",
			wantErr: false},

		{name: "Non-numeric string",
			in:      "abcd",
			want:    "abcd",
			wantErr: false},

		{name: "Numeric string",
			in:      "45",
			want:    "",
			wantErr: true},

		{name: "Zero string",
			in:      "qwe0",
			want:    "qw",
			wantErr: false},

		{name: "Zero empty string",
			in:      "\\1000",
			want:    "",
			wantErr: false},

		{
			name:    "Non-numeric escape sequence",
			in:      "\\a",
			want:    "a",
			wantErr: false},

		{
			name:    "Leading zeros in number",
			in:      "a001",
			want:    "a",
			wantErr: false},

		{
			name:    "Non-digit string",
			in:      "f33",
			want:    "fffffffffffffffffffffffffffffffff",
			wantErr: false},

		{
			name:    "First symbol digit string error",
			in:      "3f3",
			want:    "",
			wantErr: true},

		{name: "Non-byte rune string",
			in:      "Ц3😇2",
			want:    "ЦЦЦ😇😇",
			wantErr: false},

		{name: "Escape sequences string",
			in:      "qwe\\4\\5",
			want:    "qwe45",
			wantErr: false},

		{name: "Escape sequences string with numbers",
			in:      "qwe\\45",
			want:    "qwe44444",
			wantErr: false},

		{
			name:    "First symbol escape digit string",
			in:      "\\3f3",
			want:    "3fff",
			wantErr: false},

		{
			name:    "Only escape string error",
			in:      "\\",
			want:    "",
			wantErr: true},

		{
			name:    "Last escape string error",
			in:      "f3\\",
			want:    "",
			wantErr: true},

		{
			name:    "Double escape string",
			in:      "f\\\\3",
			want:    "f\\\\\\",
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := StrUnpack(tt.in)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("StrUnpack() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("StrUnpack() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("StrUnpack() = %v, want %v", got, tt.want)
			}
		})
	}
}
