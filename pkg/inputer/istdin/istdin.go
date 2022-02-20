package istdin

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

type IStdin struct {
	err error
}

func New() *IStdin {
	return &IStdin{}
}

func (s *IStdin) GetInput(data chan<- []byte) {
	go func() {
		defer close(data)
		reader := bufio.NewReader(os.Stdin)
		buf := bytes.Buffer{}
		for {
			input, err := reader.ReadByte()
			if err != nil && err == io.EOF {
				break
			}
			buf.WriteByte(input)
		}
		data <- buf.Bytes()
	}()
}

func (s *IStdin) GetError() error {
	return s.err
}
