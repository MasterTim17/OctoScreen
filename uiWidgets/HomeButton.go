package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/MasterTim17/OctoScreen/logger"
	"github.com/MasterTim17/OctoScreen/octoprintApis"
	"github.com/MasterTim17/OctoScreen/octoprintApis/dataModels"
	"github.com/MasterTim17/OctoScreen/utils"
)

type HomeButton struct {
	*gtk.Button

	client				*octoprintApis.Client
	axes				[]dataModels.Axis
}

func CreateHomeButton(
	client				*octoprintApis.Client,
	buttonLabel			string,
	imageFileName		string,
	axes				...dataModels.Axis,
) *HomeButton {
	base := utils.MustButtonImageStyle(buttonLabel, imageFileName, "", nil)

	instance := &HomeButton {
		Button:				base,
		client:				client,
		axes:				axes,
	}
	_, err := instance.Button.Connect("clicked", instance.handleClicked)
	if err != nil {
		logger.LogError("PANIC!!! - CreateHomeButton()", "instance.Button.Connect()", err)
		panic(err)
	}

	return instance
}

func (this *HomeButton) handleClicked() {
	cmd := &octoprintApis.PrintHeadHomeRequest{Axes: this.axes}
	logger.Infof("Homing the print head in %s axes", this.axes)
	err := cmd.Do(this.client);
	if err != nil {
		logger.LogError("HomeButton.handleClicked()", "Do(PrintHeadHomeRequest)", err)
	}
}
