package gobox

import (
	//for os check

	"path/filepath"
	"runtime"

	//all necessary packages from fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	//icons
	Info     *string
	Question *string
	Error    *string
	Warning  *string

	//standard fontsize for every os
	StandardSize float32
	IconHeight   float32
	IconWidth    float32
)

func init() {
	// Check the current OS and set the file paths for the icons
	switch runtime.GOOS {

	//for windows
	case "windows":
		//font size
		StandardSize = 14
		IconHeight = 40
		IconWidth = 35
		i := filepath.Join("icons", "windows", "infoWindows.png")
		q := filepath.Join("icons", "windows", "questionWindows.png")
		e := filepath.Join("icons", "windows", "errorWindows.png")
		w := filepath.Join("icons", "windows", "warningWindows.png")
		Info, Question, Error, Warning = &i, &q, &e, &w

	//for linux
	case "linux":
		//font size
		StandardSize = 16
		IconHeight = 50
		IconWidth = 45
		i := filepath.Join("icons", "linux", "infoLinux.png")
		q := filepath.Join("icons", "linux", "questionLinux.png")
		e := filepath.Join("icons", "linux", "errorLinux.png")
		w := filepath.Join("icons", "linux", "warningLinux.png")
		Info, Question, Error, Warning = &i, &q, &e, &w

	//for mac
	case "darwin":
		//font size
		StandardSize = 16
		IconHeight = 50
		IconWidth = 45
		e := filepath.Join("icons", "macos", "errorMac.png")
		w := filepath.Join("icons", "macos", "warningMac.png")
		Error, Warning = &e, &w
	}
}

// DialogBox displays a modal dialog box with an icon, message, and button.
func DialogBox(title string, icon *string, message string, btntext string, fontsize float32, w float32, h float32) {
	myApp := app.New()
	myWindow := myApp.NewWindow(title)
	myWindow.SetFixedSize(true)
	myWindow.SetMaster()

	// Set window icon for platforms that support it (e.g., Windows)
	if icon != nil {
		// Load the icon image
		iconImage := canvas.NewImageFromFile(*icon)
		myWindow.SetIcon(iconImage.Resource)
	}

	textLabel := widget.NewLabel(message)

	// Display the icon next to the message
	pngIcon := canvas.NewImageFromFile(*icon)

	pngIcon.SetMinSize(fyne.NewSize(IconHeight, IconWidth))

	iconAndMessageContainer := container.NewHBox(
		layout.NewSpacer(),
		pngIcon,
		layout.NewSpacer(),
		textLabel,
		layout.NewSpacer(),
	)

	closeButton := widget.NewButton(btntext, func() { myApp.Quit() })

	buttonsContainer := container.NewHBox(
		layout.NewSpacer(),
		closeButton,
	)

	buttonContainer := container.New(layout.NewGridLayoutWithColumns(2),
		layout.NewSpacer(),
		buttonsContainer,
	)

	// Final container for the dialog box
	mainContainer := container.NewVBox(
		layout.NewSpacer(),
		iconAndMessageContainer,
		layout.NewSpacer(),
		buttonContainer,
	)

	//run
	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(w, h))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
