package service

import (
	"github.com/bsydorenko7/machines-management-validator-backend/pkg/dataaccess/model"
	"github.com/bsydorenko7/machines-management-validator-backend/utils"
	"github.com/oskanberg/eif-go"
)

// the values closer to 1 are considered Anomalous
// and the values that are <0.5 are considered to be "normal"
const NORMAL_SCORE float64 = 0.5

type MachineService struct {
}

func NewMachineService() *MachineService {
	return &MachineService{}
}

// validate machines using Extended Isolation Forest algorithm
func (s *MachineService) ValidateMachines(machines []model.Machine) ([]model.Machine, error) {
	var inputData [][]float64
	var err error
	var ageInNanoseconds float64

	for _, machine := range machines {
		ageInNanoseconds, err = utils.GetAgeInNanoseconds(machine.Age)
		if err != nil {
			return nil, err
		}
		inputData = append(inputData, []float64{ageInNanoseconds})
	}
	f := eif.NewForest(inputData)

	var outlierMachines []model.Machine
	for _, machine := range machines {
		ageInNanoseconds, err = utils.GetAgeInNanoseconds(machine.Age)
		if err != nil {
			return nil, err
		}

		if f.Score([]float64{ageInNanoseconds}) >= NORMAL_SCORE {
			outlierMachines = append(outlierMachines, machine)
		}
	}

	return outlierMachines, err
}
