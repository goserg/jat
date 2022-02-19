package inputer

type Inputer interface {
	GetError() error
	GetInput(data chan<- []byte)
}
