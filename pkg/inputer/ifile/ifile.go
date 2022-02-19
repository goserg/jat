package ifile

import (
	"io"
	"os"
)

type IFile struct {
	err       error
	fileNames []string
}

func New(fileNames []string) *IFile {
	return &IFile{
		fileNames: fileNames,
	}
}

func (f *IFile) GetError() error {
	return f.err
}

func (f *IFile) GetInput(data chan<- []byte) {
	go func() {
		defer close(data)

		for _, fileName := range f.fileNames {
			file, err := os.Open(fileName)
			if err != nil {
				f.err = err
				return
			}
			defer file.Close()

			b, err := io.ReadAll(file)
			if err != nil {
				f.err = err
				return
			}
			data <- b
		}
	}()
}
