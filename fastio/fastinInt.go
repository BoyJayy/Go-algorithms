package main

import (
	"bufio"
	"os"
)

type FS struct {
	r *bufio.Reader
}

func NewFS() *FS {
	return &FS{r: bufio.NewReaderSize(os.Stdin, 1<<20)}
}

func (fs *FS) NextInt64() int64 {
	var c byte
	for {
		b, err := fs.r.ReadByte()
		if err != nil {
			return 0
		}
		c = b
		if c > ' ' {
			break
		}
	}
	sgn := int64(1)
	if c == '-' {
		sgn = -1
		c, _ = fs.r.ReadByte()
	}
	var x int64
	for c > ' ' {
		x = x*10 + int64(c-'0')
		b, err := fs.r.ReadByte()
		if err != nil {
			break
		}
		c = b
	}
	return x * sgn
}
func main() {

}
