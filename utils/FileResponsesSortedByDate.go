package utils

import (
	// "fmt"
	// "sort"

	// "github.com/MasterTim17/OctoScreen/octoprintApis"
	"github.com/MasterTim17/OctoScreen/octoprintApis/dataModels"
	// "github.com/MasterTim17/OctoScreen/uiWidgets"
)


type FileResponsesSortedByDate []*dataModels.FileResponse

func (this FileResponsesSortedByDate) Len() int {
	 return len(this)
}

func (this FileResponsesSortedByDate) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this FileResponsesSortedByDate) Less(i, j int) bool {
	return this[j].Date.Time.Before(this[i].Date.Time)
}
