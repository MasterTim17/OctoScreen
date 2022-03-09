package uiWidgets

import (
	// "fmt"
	"strings"

	"github.com/MasterTim17/OctoScreen/octoprintApis"
	// "github.com/MasterTim17/OctoScreen/octoprintApis/dataModels"
	"github.com/MasterTim17/OctoScreen/utils"
)

type OctoScreenInfoBox struct {
	*SystemInfoBox
}

func CreateOctoScreenInfoBox(
	client				*octoprintApis.Client,
	octoScreenVersion	string,
) *OctoScreenInfoBox {
	logoImage := utils.MustImageFromFile("logos/octoscreen-isometric-90%.png")

	str2 := ""
	str3 := ""
	stringArray := strings.Split(octoScreenVersion, " ")
	if len(stringArray) == 2 {
		str2 = stringArray[0]
		str3 = stringArray[1]
	} else {
		str2 = octoScreenVersion
		str3 = ""
	}

	base := CreateSystemInfoBox(
		client,
		logoImage,
		"OctoScreen",
		str2,
		str3,
	)

	instance := &OctoScreenInfoBox {
		SystemInfoBox:			base,
	}

	return instance
}
