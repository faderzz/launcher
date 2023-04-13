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

	// Custom top bar
	topBar := widgets.NewQWidget(nil, 0)
	topBar.SetFixedHeight(30)
	topBar.SetStyleSheet("background-color: #1E1E1E;")

	// Custom top bar buttons
	closeBtn := widgets.NewQPushButton2("X", nil)
	closeBtn.SetStyleSheet("color: #F0F0F0; font: bold 14px Arial;")
	closeBtn.SetFixedSize2(30, 30)
	closeBtn.SetFlat(true)
	closeBtn.ConnectClicked(func(checked bool) {
		window.Close()
	})
	minimizeBtn := widgets.NewQPushButton2("-", nil)
	minimizeBtn.SetStyleSheet("color: #F0F0F0; font: bold 14px Arial;")
	minimizeBtn.SetFixedSize2(30, 30)
	minimizeBtn.SetFlat(true)
	minimizeBtn.ConnectClicked(func(checked bool) {
		window.ShowMinimized()
	})

	// Custom top bar layout
	topBarLayout := widgets.NewQHBoxLayout()
	topBarLayout.AddStretch(0)
	topBarLayout.AddWidget(minimizeBtn, 0, 0)
	topBarLayout.AddWidget(closeBtn, 0, 0)
	topBar.SetLayout(topBarLayout)

	// Custom top bar drag
	topBar.ConnectMousePressEvent(func(event *gui.QMouseEvent) {
		window.MousePressEvent(event)
	})
	topBar.ConnectMouseMoveEvent(func(event *gui.QMouseEvent) {
		window.MouseMoveEvent(event)
	})
	topBar.ConnectMouseReleaseEvent(func(event *gui.QMouseEvent) {
		window.MouseReleaseEvent(event)
	})

	// Hover change color
	closeBtn.ConnectEnterEvent(func(event *core.QEvent) {
		closeBtn.SetStyleSheet("color: #404040; font: bold 14px Arial;")
		// Change mouse cursor to hand
		closeBtn.SetCursor(gui.NewQCursor2(core.Qt__PointingHandCursor))
	})
	minimizeBtn.ConnectEnterEvent(func(event *core.QEvent) {
		minimizeBtn.SetStyleSheet("color: #404040; font: bold 14px Arial;")
		// Change mouse cursor to hand
		minimizeBtn.SetCursor(gui.NewQCursor2(core.Qt__PointingHandCursor))
	})

	// Change back to normal color
	closeBtn.ConnectLeaveEvent(func(event *core.QEvent) {
		closeBtn.SetStyleSheet("color: #F0F0F0; font: bold 14px Arial;")
	})
	minimizeBtn.ConnectLeaveEvent(func(event *core.QEvent) {
		minimizeBtn.SetStyleSheet("color: #F0F0F0; font: bold 14px Arial;")
	})

	// Create widgets
	label := widgets.NewQLabel(nil, 0)
	pixmap := gui.NewQPixmap3("logo.png", "", core.Qt__AutoColor)
	label.SetPixmap(pixmap)
	label.SetAlignment(core.Qt__AlignCenter)
	label.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	label.SetFixedHeight(100)
	// Set label width to fit the pixmap
	label.SetFixedWidth((pixmap.Width() + 200))
	// Set height of label to fit the pixmap
	label.SetFixedHeight((pixmap.Height() + 100))

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
		// Call the Microsoft Auth function in the auth.go file
		go microsoftAuth()
	})

	// Create a vertical layout for window
	vLayout := widgets.NewQVBoxLayout()
	vLayout.AddWidget(topBar, 0, 0)
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
