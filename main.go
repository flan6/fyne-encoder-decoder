package main

import (
	"encoding/base64"
	"encoding/hex"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Encode Decode")

	w.Resize(fyne.NewSize(500, 300))

	input := widget.NewEntry()
	input.PlaceHolder = "Input"

	text := widget.NewMultiLineEntry()

	ButtonB64Encode := widget.NewButton("Base64 Encode", func() {
		text.SetText(base64.StdEncoding.EncodeToString([]byte(input.Text)))
	})

	ButtonB64Decode := widget.NewButton("Base64 Decode", func() {
		decoded, err := base64.StdEncoding.DecodeString(input.Text)
		if err != nil {
			dialog.ShowError(err, w)
		}

		text.SetText(string(decoded))
	})

	ButtonHexEncode := widget.NewButton("Hex Encode", func() {
		text.SetText(hex.EncodeToString([]byte(input.Text)))
	})

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			input,
			ButtonB64Encode,
			ButtonB64Decode,
			ButtonHexEncode,
			fyne.NewContainerWithLayout(layout.NewVBoxLayout(), text),
		),
	)

	w.ShowAndRun()
}
