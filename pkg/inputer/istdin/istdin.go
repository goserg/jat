package istdin

import (
	"bufio"
	"bytes"
	"io"
	"log"
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
		info, err := os.Stdin.Stat()
		if err != nil {
			s.err = err
			return
		}

		if info.Mode()&os.ModeCharDevice != 0 {
			log.Fatal("The continious reading from stdin is not supported yet; use with pipe or with filename in args")
		}

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
