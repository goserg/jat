package istdin

import (
	"log"
)

type IStdin struct {
	err error
}

func New() *IStdin {
	return &IStdin{}
}

func (s *IStdin) GetInput(data chan<- []byte) {
	log.Fatal("stdin input is not implemented yet: specify the file names to open") // TODO
}

func (s *IStdin) GetError() error {
	return s.err
}
