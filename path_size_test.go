package code

import "testing"

func TestGetPathSize_File2b(t *testing.T) {
	path := "testdata/file_2b.txt"
	got, err := GetPathSize(path, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "2B\t" + path
	if got != want {
		t.Fatalf("GetPathSize() want: %q\ngot:  %q", want, got)
	}
}

func TestGetPathSize_DirWithOneFile(t *testing.T) {
	path := "testdata/dir_with_one_file"
	got, err := GetPathSize(path, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "8B\t" + path
	if got != want {
		t.Fatalf("GetPathSize() want: %q\ngot:  %q", want, got)
	}
}

func TestGetPathSize_EmptyDir(t *testing.T) {
	path := "testdata/empty_dir"
	got, err := GetPathSize(path, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "0B\t" + path
	if got != want {
		t.Fatalf("GetPathSize() want %q, got %q", want, got)
	}
}

func TestFormatSize_NoHuman(t *testing.T) {
	tests := []struct {
		name string
		in   int64
		want string
	}{
		{"zero", 0, "0B"},
		{"bytes", 42, "42B"},
		{"1kb_raw", 1024, "1024B"},
		{"24mb_raw", 25165824, "25165824B"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatSize(tt.in, false)
			if got != tt.want {
				t.Fatalf("FormatSize(%d, false): want %q, got %q", tt.in, tt.want, got)
			}
		})
	}
}

func TestFormatSize_Human_AllUnits(t *testing.T) {
	const (
		KB int64 = 1024
		MB       = KB * 1024
		GB       = MB * 1024
		TB       = GB * 1024
		PB       = TB * 1024
		EB       = PB * 1024 // NOTE: 1024^6 = 2^60 = 1_152_921_504_606_846_976
	)

	tests := []struct {
		name string
		in   int64
		want string
	}{
		{"bytes", 42, "42B"},
		{"exact_1kb", 1 * KB, "1.0KB"},
		{"exact_1mb", 1 * MB, "1.0MB"},
		{"exact_1gb", 1 * GB, "1.0GB"},
		{"exact_1tb", 1 * TB, "1.0TB"},
		{"exact_1pb", 1 * PB, "1.0PB"},
		{"exact_1eb", 1 * EB, "1.0EB"},

		{"about_1_2mb", 1234567, "1.2MB"},
		{"exact_24mb", 24 * MB, "24.0MB"},
		{"one_and_half_gb", 1536 * MB, "1.5GB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatSize(tt.in, true)
			if got != tt.want {
				t.Fatalf("FormatSize(%d, true): want %q, got %q", tt.in, tt.want, got)
			}
		})
	}
}
