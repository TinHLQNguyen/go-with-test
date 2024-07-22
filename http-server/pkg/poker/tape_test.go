package poker

import (
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContent, _ := io.ReadAll(file)

	got := string(newFileContent)
	want := "abc"

	AssertEqual(t, got, want)
}
