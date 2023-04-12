package main

import (
	"fmt"
	"image"
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
	microsoftAuthBtn := widgets.NewQPushButton(nil)
	microsoftAuthBtn.SetSizePolicy2(widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Minimum)
	microsoftAuthBtn.SetFlat(true)
	microsoftAuthBtn.SetAutoFillBackground(true)
	microsoftAuthBtn.SetStyleSheet("background-color: black; color: white; border: none; padding: 0;")
	microsoftAuthBtn.SetMinimumSize2(200, 50)

	// Load the image from file
	imageFile, err := os.Open("ms.png")
	if err != nil {
		fmt.Println("Failed to open image:", err)
		return
	}
	defer imageFile.Close()

	image, _, err := image.Decode(imageFile)
	if err != nil {
		fmt.Println("Failed to decode image:", err)
		return
	}

	// Convert the image to a QPixmap
	pixmap := gui.NewQPixmap3(image)
	if pixmap.IsNull() {
		fmt.Println("Failed to create QPixmap from image")
		return
	}

	// Create a QLabel for the image on the left half
	imageLabel := widgets.NewQLabel(nil, 0)
	imageLabel.SetPixmap(pixmap)
	imageLabel.SetScaledContents(true)

	// Create a QLabel for the text on the right half
	textLabel := widgets.NewQLabel2("Microsoft Auth", nil, 0)
	textLabel.SetAlignment(core.Qt__AlignCenter)

	// Create a QHBoxLayout for the button's layout
	layout := widgets.NewQHBoxLayout()
	layout.SetContentsMargins(0, 0, 0, 0)
	layout.AddWidget(imageLabel, 0, 0)
	layout.AddWidget(textLabel, 0, 0)

	// Set the layout for the button
	buttonWidget := widgets.NewQWidget(nil, 0)
	buttonWidget.SetLayout(layout)
	microsoftAuthBtn.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	microsoftAuthBtn.SetLayout(layout)

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
