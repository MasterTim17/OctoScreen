package ui

import (
	// "github.com/MasterTim17/OctoScreen/interfaces"
	// "github.com/MasterTim17/OctoScreen/octoprintApis"
	"github.com/MasterTim17/OctoScreen/octoprintApis/dataModels"
	// "github.com/MasterTim17/OctoScreen/uiWidgets"
)


type customItemsPanel struct {
	CommonPanel
	items			[]dataModels.MenuItem
}

func CustomItemsPanel(
	ui				*UI,
	items			[]dataModels.MenuItem,
) *customItemsPanel {
	instance := &customItemsPanel {
		CommonPanel: NewCommonPanel("CustomItemsPanel", ui),
		items:       items,
	}
	instance.initialize()

	return instance
}

func (this *customItemsPanel) initialize() {
	defer this.Initialize()
	this.arrangeMenuItems(this.grid, this.items, 4)
}
