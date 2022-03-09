package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/MasterTim17/OctoScreen/utils"
)

type ToolPrintingButton struct {
	*gtk.Button
}

func CreateToolPrintingButton(
	index			int,
) *ToolPrintingButton {
	imageFileName := ToolImageFileName(index)
	instance := &ToolPrintingButton{
		Button:  utils.MustButtonImage("", imageFileName, nil),
	}

	ctx, _ := instance.GetStyleContext()
	ctx.AddClass("printing-state")

	return instance
}
