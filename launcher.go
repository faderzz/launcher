package main

import (
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
	window.SetMinimumSize2(500, 500)
	window.SetMaximumSize2(1600, 1600)
	// Window background
	window.SetStyleSheet("background-color: #1E1E1E;")

	// Create widgets
	label := widgets.NewQLabel2("Project X Launcher", nil, 0)
	label.SetAlignment(core.Qt__AlignCenter)
	label.SetStyleSheet("background: qlineargradient(x1:0, y1:0, x2:0, y2:1, stop:0 #67CEF6, stop:1 #43EB3D); color: #F0F0F0; font: bold 80px Arial;")
	label.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	label.SetFixedHeight(100)
	label.SetWordWrap(true)

	// Microsoft Auth button
	microsoftAuthBtn := widgets.NewQPushButton2("Sign in with Microsoft", nil)
	microsoftAuthBtn.SetStyleSheet("background-color: #2E2E2E; color: #F0F0F0; border: 1px solid #505050; border-radius: 5px; border-radius: 15px; font: bold 14px Arial;")
	microsoftAuthBtn.SetFixedSize2(200, 30)
	microsoftAuthBtn.SetIcon(gui.NewQIcon5("ms.png"))
	microsoftAuthBtn.SetIconSize(core.NewQSize2(20, 20))
	microsoftAuthBtn.SetFlat(true)

	// Create a vertical layout for window
	vLayout := widgets.NewQVBoxLayout()
	vLayout.AddStretch(1)                                        // Add stretchable space at the top
	vLayout.AddWidget(label, 0, core.Qt__AlignCenter)            // Add label centered at the top
	vLayout.AddStretch(1)                                        // Add stretchable space in the middle
	vLayout.AddWidget(microsoftAuthBtn, 0, core.Qt__AlignCenter) // Add button centered in the middle

	// Create a widget to hold the layout
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(vLayout)

	// Set the widget as the central widget of the window
	window.SetCentralWidget(widget)

	// Show window
	window.Show()

	app.Exec()
}
