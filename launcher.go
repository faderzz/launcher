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
	window.SetWindowTitle("Project X Launcher")
	window.SetMinimumSize2(400, 300)

	// Step 2: Create widgets
	label := widgets.NewQLabel2("Welcome to Project X Launcher", nil, 0)
	label.SetAlignment(0x84) // AlignCenter

	// Login Functions
	// Contains Microsoft Auth +
	// Mojang Auth

	// Mojang Auth
	// Email input
	emailInput := widgets.NewQLineEdit(nil)
	emailInput.SetPlaceholderText("Email")
	// Password input
	passwordInput := widgets.NewQLineEdit(nil)
	passwordInput.SetPlaceholderText("Password")
	passwordInput.SetEchoMode(2) // Password
	// Login button
	loginBtn := widgets.NewQPushButton2("Login", nil)
	loginBtn.ConnectClicked(func(bool) {
		// Call the login function here
		fmt.Println("Login button clicked")
	})

	// Microsoft Auth
	// Microsoft Auth button
	microsoftAuthBtn := widgets.NewQPushButton2("Microsoft Auth", nil)
	microsoftAuthBtn.ConnectClicked(func(bool) {
		// Call the Microsoft Auth function here
		fmt.Println("Microsoft Auth button clicked")
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
