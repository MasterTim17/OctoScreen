package uiWidgets

import (
	// "fmt"

	"github.com/gotk3/gotk3/gtk"

	"github.com/MasterTim17/OctoScreen/logger"
	"github.com/MasterTim17/OctoScreen/octoprintApis"
	"github.com/MasterTim17/OctoScreen/octoprintApis/dataModels"
	"github.com/MasterTim17/OctoScreen/utils"
)


type CommandButton struct {
	*gtk.Button

	client				*octoprintApis.Client
	parentWindow		*gtk.Window
	commandDefinition	*dataModels.CommandDefinition
}

func CreateCommandButton(
	client				*octoprintApis.Client,
	parentWindow		*gtk.Window,
	commandDefinition	*dataModels.CommandDefinition,
	iconName			string,
) *CommandButton {
	style := ""
	if commandRequiresConfirmation(commandDefinition) {
		style = "color-warning-sign-yellow"
	}
	base := utils.MustButtonImageStyle(utils.StrEllipsisLen(commandDefinition.Name, 16), iconName + ".svg", style, nil)

	instance := &CommandButton {
		Button:				base,
		client:				client,
		parentWindow:		parentWindow,
		commandDefinition:	commandDefinition,
	}
	_, err := instance.Button.Connect("clicked", instance.handleClicked)
	if err != nil {
		logger.LogError("PANIC!!! - CreateCommandButton()", "instance.Button.Connect()", err)
		panic(err)
	}

	return instance
}

func commandRequiresConfirmation(commandDefinition *dataModels.CommandDefinition) bool {
	return commandDefinition != nil && len(commandDefinition.Confirm) > 0
}

func (this *CommandButton) handleClicked() {
	if commandRequiresConfirmation(this.commandDefinition) {
		utils.MustConfirmDialogBox(
			this.parentWindow,
			this.commandDefinition.Confirm + "\n\nAre you sure you want to proceed?",
			this.sendCommand,
		)()
	} else {
		this.sendCommand()
	}
}

func (this *CommandButton) sendCommand() {
	logger.Infof("CommandButton.sendCommand(), now sending command %q", this.commandDefinition.Name)

	commandRequest := &octoprintApis.SystemExecuteCommandRequest{
		Source: dataModels.Custom,
		Action: this.commandDefinition.Action,
	}

	err := commandRequest.Do(this.client)
	if err != nil {
		logger.LogError("CommandButton.sendCommand()", "Do(SystemExecuteCommandRequest)", err)
	}
}
