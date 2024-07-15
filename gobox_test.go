package main_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BenchmarkMessageBox(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchmarkMessageBox()
	}
}

func benchmarkMessageBox() {
	myBox := app.New()
	myWindow := myBox.NewWindow("Benchmark Test")

	textLabel := canvas.NewText("Benchmarking message box", nil)
	textLabel.TextSize = 16

	pngIcon := canvas.NewImageFromFile("icons/linux/infoLinux.png")
	pngIcon.SetMinSize(fyne.NewSize(50, 45))

	iconAndMessageContainer := container.NewHBox(
		widget.NewLabel(""),
		pngIcon,
		widget.NewLabel(""),
		textLabel,
	)

	closeButton := widget.NewButton("OK", func() { fyne.CurrentApp().Quit() })

	buttonContainer := container.NewHBox(
		layout.NewSpacer(),
		closeButton,
	)

	mainContainer := container.NewVBox(
		iconAndMessageContainer,
		layout.NewSpacer(),
		buttonContainer,
	)

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(350, 110))
	// myWindow.ShowAndRun()
}

//Multiple tests ~ iterations = 400-450 -> 1.3s-1.5s
