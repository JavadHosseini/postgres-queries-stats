package server

type IServer interface {
	Start(port string) error
}
