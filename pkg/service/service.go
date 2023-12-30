package service

import "github.com/bsydorenko7/machines-management-validator-backend/pkg/dataaccess/model"

type Machine interface {
	ValidateMachines([]model.Machine) ([]model.Machine, error)
}

type Service struct {
	Machine
}

func NewService() *Service {
	return &Service{
		Machine: NewMachineService(),
	}
}
