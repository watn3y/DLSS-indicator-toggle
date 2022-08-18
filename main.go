package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var indicatorStatusString string
var indicatorStatus int

func main() {

	getKey()
	var indicatorTextLabel *walk.TextLabel
	var windowMain *walk.MainWindow

	MainWindow{
		Title:    "DLSS Indicator Toggle",
		AssignTo: &windowMain,
		Size: Size{
			Width:  300,
			Height: 200,
		},
		Layout: VBox{},
		Children: []Widget{
			TextLabel{
				Text:          "The DLSS Indicator is currently",
				TextAlignment: AlignHCenterVCenter,
			},
			TextLabel{
				Text:          indicatorStatusString,
				AssignTo:      &indicatorTextLabel,
				TextAlignment: AlignHCenterVCenter,
			},

			PushButton{
				Text: "Toggle Indicator",

				OnClicked: func() {
					if indicatorStatus == 1024 {
						setKey(0)
					} else {
						setKey(1)
					}
					getKey()
					indicatorTextLabel.SetText(indicatorStatusString)
				},
			},
		},
	}.Run()
}
