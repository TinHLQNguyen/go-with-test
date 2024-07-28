package poker_test

import (
	"go-with-test/http-server/pkg/poker"
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := poker.CreateTempFile(t, "12345")
	defer clean()

	tape := &poker.Tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContent, _ := io.ReadAll(file)

	got := string(newFileContent)
	want := "abc"

	poker.AssertEqual(t, got, want)
}
