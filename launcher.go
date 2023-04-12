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

	// Microsoft Auth
	// Create a container widget for the button with a horizontal layout
	microsoftAuthContainer := widgets.NewQWidget(nil, 0)
	microsoftAuthContainer.SetSizePolicy2(widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Minimum)
	microsoftAuthContainer.SetStyleSheet("background-color: #000000; border-radius: 5px;")

	// Create an image label for the left half of the button
	imageLabel := widgets.NewQLabel(nil, 0)
	imageLabel.SetPixmap(gui.NewQPixmap3(":/ms.png", "", core.Qt__AutoColor))
	imageLabel.SetScaledContents(true)
	imageLabel.SetFixedSize2(25, 25)
	imageLabel.SetAlignment(core.Qt__AlignCenter)

	// Create a label for the right half of the button with the text
	textLabel := widgets.NewQLabel2("Sign in with Microsoft", nil, 0)
	textLabel.SetAlignment(core.Qt__AlignCenter)

	// Add the image label and text label to the container widget
	containerLayout := widgets.NewQHBoxLayout2(microsoftAuthContainer)
	containerLayout.AddWidget(imageLabel, 0, core.Qt__AlignLeft)
	containerLayout.AddWidget(textLabel, 0, core.Qt__AlignRight)
	microsoftAuthContainer.SetLayout(containerLayout)

	// Microsoft Auth button
	microsoftAuthBtn := widgets.NewQPushButton(nil)
	microsoftAuthBtn.SetSizePolicy2(widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Minimum)
	microsoftAuthBtn.SetFlat(true)
	microsoftAuthBtn.SetAutoFillBackground(false)
	microsoftAuthBtn.SetStyleSheet("border: none;")
	microsoftAuthBtn.SetMinimumSize2(200, 50)
	microsoftAuthBtn.SetCursor(gui.NewQCursor(core.Qt__PointingHandCursor))
	microsoftAuthBtn.SetFocusPolicy(core.Qt__NoFocus)

	microsoftAuthBtn.SetIconSize(core.NewQSize2(25, 25))
	microsoftAuthBtn.SetIcon(gui.NewQIcon5(":/ms.png"))

	microsoftAuthBtn.ConnectClicked(func(bool) {
		// Call the Microsoft Auth function here
		fmt.Println("Microsoft Auth button clicked")
	})

	// Add widgets to window
	window.SetCentralWidget(label)

	layout := widgets.NewQVBoxLayout2(nil)
	layout.AddWidget(microsoftAuthBtn, 0, 0)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	window.SetCentralWidget(widget)

	// Show window
	window.Show()

	app.Exec()
}
