package poker

import (
	"io"
	"os"
)

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0) // clear file
	t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}
