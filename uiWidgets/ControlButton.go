package uiWidgets

import (
	// "fmt"

	"github.com/gotk3/gotk3/gtk"

	"github.com/MasterTim17/OctoScreen/logger"
	"github.com/MasterTim17/OctoScreen/octoprintApis"
	"github.com/MasterTim17/OctoScreen/octoprintApis/dataModels"
	"github.com/MasterTim17/OctoScreen/utils"
)


type ControlButton struct {
	*gtk.Button

	client				*octoprintApis.Client
	parentWindow		*gtk.Window
	controlDefinition	*dataModels.ControlDefinition
}

func CreateControlButton(
	client				*octoprintApis.Client,
	parentWindow		*gtk.Window,
	controlDefinition	*dataModels.ControlDefinition,
	iconName			string,
) *ControlButton {
	style := ""
	if controlRequiresConfirmation(controlDefinition) {
		style = "color-warning-sign-yellow"
	}
	base := utils.MustButtonImageStyle(utils.StrEllipsisLen(controlDefinition.Name, 16), iconName + ".svg", style, nil)

	instance := &ControlButton {
		Button:				base,
		client:				client,
		parentWindow:		parentWindow,
		controlDefinition:	controlDefinition,
	}
	_, err := instance.Button.Connect("clicked", instance.handleClicked)
	if err != nil {
		logger.LogError("PANIC!!! - CreateControlButton()", "instance.Button.Connect()", err)
		panic(err)
	}

	return instance
}

func controlRequiresConfirmation(controlDefinition *dataModels.ControlDefinition) bool {
	return controlDefinition != nil && len(controlDefinition.Confirm) > 0
}

func (this *ControlButton) handleClicked() {
	if controlRequiresConfirmation(this.controlDefinition) {
		utils.MustConfirmDialogBox(
			this.parentWindow,
			this.controlDefinition.Confirm + "\n\nAre you sure you want to proceed?",
			this.sendCommand,
		)()
	} else {
		this.sendCommand()
	}
}

func (this *ControlButton) sendCommand() {
	logger.Infof("ControlButton.sendCommand(), now sending command %q", this.controlDefinition.Name)

	commandRequest := &octoprintApis.CommandRequest{
		Commands: this.controlDefinition.Commands,
	}

	if len(this.controlDefinition.Command) != 0 {
		commandRequest.Commands = []string{this.controlDefinition.Command}
	}

	err := commandRequest.Do(this.client)
	if err != nil {
		logger.LogError("ControlButton.sendCommand()", "Do(CommandRequest)", err)
	}
}
