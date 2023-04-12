package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Step 1: Create main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Minecraft Launcher")
	window.SetMinimumSize2(400, 300)

	// Step 2: Create widgets
	label := widgets.NewQLabel2("Welcome to Minecraft Launcher", nil, 0)
	label.SetAlignment(0x84) // AlignCenter

	loginBtn := widgets.NewQPushButton2("Login", nil)
	loginBtn.ConnectClicked(func(bool) {
		// Call the login function here
		fmt.Println("Login button clicked")
	})

	playBtn := widgets.NewQPushButton2("Play", nil)
	playBtn.ConnectClicked(func(bool) {
		// Call the play function here
		fmt.Println("Play button clicked")
	})

	// Step 3: Add widgets to window
	window.SetCentralWidget(label)

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(loginBtn, 0, 0)
	layout.AddWidget(playBtn, 0, 0)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	window.SetCentralWidget(widget)

	// Step 4: Show window
	window.Show()

	app.Exec()
}
