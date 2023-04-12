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
	window.SetMinimumSize2(1250, 800)
	window.SetMaximumSize2(1250, 800)
	// Window background
	window.SetStyleSheet("background-color: #1E1E1E;")

	// Custom window buttons
	window.SetWindowFlags(core.Qt__FramelessWindowHint)
	// Exit button

	// Create widgets
	label := widgets.NewQLabel(nil, 0)
	pixmap := gui.NewQPixmap3("logo.png", "", core.Qt__AutoColor)
	label.SetPixmap(pixmap)
	label.SetAlignment(core.Qt__AlignCenter)
	label.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	label.SetFixedHeight(100)
	// Set label width to fit the pixmap
	label.SetFixedWidth(pixmap.Width())
	// Set height of label to fit the pixmap
	label.SetFixedHeight(pixmap.Height())

	// Microsoft Auth button
	microsoftAuthBtn := widgets.NewQPushButton2("Sign in with Microsoft", nil)
	microsoftAuthBtn.SetStyleSheet("background-color: #2E2E2E; color: #F0F0F0; border: 1px solid #505050; border-radius: 5px; border-radius: 5px; font: bold 14px Arial; margin-bottom: 10px;")
	microsoftAuthBtn.SetFixedSize2(400, 60)
	microsoftAuthBtn.SetIcon(gui.NewQIcon5("ms.png"))
	microsoftAuthBtn.SetIconSize(core.NewQSize2(40, 40))
	microsoftAuthBtn.SetFlat(true)

	// Hover effect
	microsoftAuthBtn.ConnectEnterEvent(func(event *core.QEvent) {
		// Change cursor to hand
		microsoftAuthBtn.SetCursor(gui.NewQCursor2(core.Qt__PointingHandCursor))
		microsoftAuthBtn.SetStyleSheet("background-color: #2E2E2E; color: #F0F0F0; border: 1px solid #505050; border-radius: 5px; border-radius: 5px; font: bold 14px Arial; margin-bottom: 10px; background-color: rgba(255, 255, 255, 0.1);")
	})

	// Hover leave effect
	microsoftAuthBtn.ConnectLeaveEvent(func(event *core.QEvent) {
		// Change color
		microsoftAuthBtn.SetStyleSheet("background-color: #2E2E2E; color: #F0F0F0; border: 1px solid #505050; border-radius: 5px; border-radius: 5px; font: bold 14px Arial; margin-bottom: 10px;")
	})

	// Click function
	microsoftAuthBtn.ConnectClicked(func(checked bool) {
		// TODO: Implement Microsoft Auth
		println("Microsoft Auth")
	})

	// Create a vertical layout for window
	vLayout := widgets.NewQVBoxLayout()
	vLayout.AddStretch(1)
	vLayout.AddWidget(label, 0, core.Qt__AlignCenter)
	vLayout.AddStretch(1)
	vLayout.AddWidget(microsoftAuthBtn, 0, core.Qt__AlignCenter)

	// Create a widget to hold the layout
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(vLayout)

	// Set the widget as the central widget of the window
	window.SetCentralWidget(widget)

	// Show window
	window.Show()

	app.Exec()
}
