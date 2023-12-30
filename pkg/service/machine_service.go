package service

import (
	"github.com/bsydorenko7/machines-management-validator-backend/pkg/dataaccess/model"
)

type MachineService struct {
}

func NewMachineService() *MachineService {
	return &MachineService{}
}

func (s *MachineService) ValidateMachines(machines []model.Machine) ([]model.Machine, error) {
	return machines, nil
}
