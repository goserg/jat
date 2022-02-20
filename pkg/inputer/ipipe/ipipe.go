package ipipe

import (
	"bufio"
	"os"
)

type IPipe struct {
	err error
}

func New() *IPipe {
	return &IPipe{}
}

func (s *IPipe) GetInput(data chan<- []byte) {
	go func() {
		defer close(data)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data <- scanner.Bytes()
		}
		if scanner.Err() != nil {
			s.err = scanner.Err()
			return
		}
	}()
}

func (s *IPipe) GetError() error {
	return s.err
}
