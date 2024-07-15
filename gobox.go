package gobox

import (
	"embed"
	"path/filepath"
	"runtime"

	// all necessary packages from fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//go:embed icons/windows/*
//go:embed icons/linux/*
//go:embed icons/mac/*
var embeddedFiles embed.FS

var (
	// icons
	Info     string
	Question string
	Error    string
	Warning  string

	// standard fontsize for every OS
	StandardSize float32
	IconHeight   float32
	IconWidth    float32
)

func init() {
	// Check the current OS and set the file paths for the icons
	switch runtime.GOOS {

	// for windows
	case "windows":
		// font size
		StandardSize = 14
		IconHeight = 40
		IconWidth = 35
		Info = "icons/windows/infoWindows.png"
		Question = "icons/windows/questionWindows.png"
		Error = "icons/windows/errorWindows.png"
		Warning = "icons/windows/warningWindows.png"

	// for linux
	case "linux":
		// font size
		StandardSize = 16
		IconHeight = 50
		IconWidth = 45
		Info = "icons/linux/infoLinux.png"
		Question = "icons/linux/questionLinux.png"
		Error = "icons/linux/errorLinux.png"
		Warning = "icons/linux/warningLinux.png"

	// for mac
	case "darwin":
		// font size
		StandardSize = 16
		IconHeight = 50
		IconWidth = 45
		Error = "icons/macos/errorMac.png"
		Warning = "icons/macos/warningMac.png"
	}
}

// DialogBox displays a modal dialog box with an icon, message, and button.
func DialogBox(title string, icon string, message string, btntext string, fontsize float32, w float32, h float32) {
	myApp := app.New()
	myWindow := myApp.NewWindow(title)
	myWindow.SetFixedSize(true)
	myWindow.SetMaster()

	// Set window icon for platforms that support it (e.g., Windows)
	var iconImage *canvas.Image
	if icon != "" {
		// Load the icon image
		iconData, err := embeddedFiles.ReadFile(icon)
		if err == nil {
			iconResource := fyne.NewStaticResource(filepath.Base(icon), iconData)
			iconImage = canvas.NewImageFromResource(iconResource)
			myWindow.SetIcon(iconResource)
		}
	}

	textLabel := widget.NewLabel(message)

	// Display the icon next to the message
	if iconImage != nil {
		iconImage.SetMinSize(fyne.NewSize(IconHeight, IconWidth))
	}

	iconAndMessageContainer := container.NewHBox(
		layout.NewSpacer(),
		iconImage,
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

	// run
	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(w, h))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
