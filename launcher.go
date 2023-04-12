package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Project X Launcher")
	window.SetMinimumSize2(400, 300)

	// Create widgets
	label := widgets.NewQLabel2("Welcome to Project X Launcher", nil, 0)
	label.SetAlignment(core.Qt__AlignCenter)

	// Microsoft Auth button
	microsoftAuthBtn := widgets.NewQPushButton2("", nil)
	microsoftAuthBtn.SetStyleSheet("background-color: black; color: white; padding: 10px;")
	microsoftAuthBtn.SetSizePolicy2(widgets.QSizePolicy__MinimumExpanding, widgets.QSizePolicy__Fixed)

	icon := gui.NewQIcon5(":/ms.png")
	microsoftAuthBtn.SetIcon(icon)
	microsoftAuthBtn.SetIconSize(core.NewQSize2(25, 25))

	labelText := widgets.NewQLabel2("Microsoft Auth", nil, 0)
	labelText.SetStyleSheet("margin-left: 10px;")

	layout := widgets.NewQHBoxLayout()
	layout.SetContentsMargins(0, 0, 0, 0)
	layout.AddWidget(microsoftAuthBtn, 0, 0)
	layout.AddWidget(labelText, 0, 0)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	microsoftAuthBtn.ConnectClicked(func(bool) {
		// Call the Microsoft Auth function here
		fmt.Println("Microsoft Auth button clicked")
	})

	// Add widgets to window
	window.SetCentralWidget(label)

	window.SetCentralWidget(widget)

	// Show window
	window.Show()

	app.Exec()
}
