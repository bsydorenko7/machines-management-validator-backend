package service

type Machine interface {
}

type Service struct {
	Machine
}

func NewService() *Service {
	return &Service{}
}
