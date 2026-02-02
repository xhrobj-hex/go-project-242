package code

import "testing"

func TestGetPathSize_File2b(t *testing.T) {
	path := "testdata/file_2b.txt"
	got, err := GetPathSize(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "2\t" + path
	if got != want {
		t.Fatalf("want: %q\ngot:  %q", want, got)
	}
}

func TestGetPathSize_DirWithOneFile(t *testing.T) {
	path := "testdata/dir_with_one_file"
	got, err := GetPathSize(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "8\t" + path
	if got != want {
		t.Fatalf("want: %q\ngot:  %q", want, got)
	}
}

func TestGetPathSize_EmptyDir(t *testing.T) {
	path := "testdata/empty_dir"
	got, err := GetPathSize(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "0\t" + path
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}
