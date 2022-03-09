package ui

import (
	// "github.com/MasterTim17/OctoScreen/interfaces"
	"github.com/MasterTim17/OctoScreen/logger"
	"github.com/MasterTim17/OctoScreen/octoprintApis"
	// "github.com/MasterTim17/OctoScreen/octoprintApis/dataModels"
	"github.com/MasterTim17/OctoScreen/uiWidgets"
	// "github.com/MasterTim17/OctoScreen/utils"
)


var temperaturePresetsPanelInstance *temperaturePresetsPanel

type temperaturePresetsPanel struct {
	CommonPanel

	selectHotendStepButton	*uiWidgets.SelectToolStepButton

}

func TemperaturePresetsPanel(
	ui						*UI,
	selectHotendStepButton	*uiWidgets.SelectToolStepButton,
) *temperaturePresetsPanel {
	if temperaturePresetsPanelInstance == nil {
		instance := &temperaturePresetsPanel {
			CommonPanel:			NewCommonPanel("temperaturePresetsPanel", ui),
			selectHotendStepButton:	selectHotendStepButton,
		}
		instance.initialize()
		temperaturePresetsPanelInstance = instance
	}

	return temperaturePresetsPanelInstance
}

func (this *temperaturePresetsPanel) initialize() {
	defer this.Initialize()
	this.createAllOffButton()
	this.createTemperaturePresetButtons()
}

func (this *temperaturePresetsPanel) createAllOffButton() {
	allOffButton := uiWidgets.CreateCoolDownButton(this.UI.Client, this.UI.GoToPreviousPanel)
	this.AddButton(allOffButton)
}

func (this *temperaturePresetsPanel) createTemperaturePresetButtons() {
	settings, err := (&octoprintApis.SettingsRequest{}).Do(this.UI.Client)
	if err != nil {
		logger.LogError("TemperaturePresetsPanel.getTemperaturePresets()", "Do(SettingsRequest)", err)
		return
	}

	// 12 (max) - Back button - All Off button = 10 available slots to display.
	const maxSlots = 10

	count := 0
	for _, temperaturePreset := range settings.Temperature.TemperaturePresets {
		if count < maxSlots {
			temperaturePresetButton := uiWidgets.CreateTemperaturePresetButton(
				this.UI.Client,
				this.selectHotendStepButton,
				"heat-up.svg",
				temperaturePreset,
				this.UI.GoToPreviousPanel,
			)
			this.AddButton(temperaturePresetButton)
			count++
		}
	}
}
