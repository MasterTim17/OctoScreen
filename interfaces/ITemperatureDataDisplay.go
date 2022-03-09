package interfaces

import (
	// "github.com/MasterTim17/OctoScreen/octoprintApis"
	"github.com/MasterTim17/OctoScreen/octoprintApis/dataModels"
)

type ITemperatureDataDisplay interface {
	UpdateTemperatureData(temperatureData map[string]dataModels.TemperatureData)
}
