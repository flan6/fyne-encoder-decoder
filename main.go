package main

import (
	"encoding/base64"
	"net/url"

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

	ButtonUrlEncode := widget.NewButton("URL Encode", func() {
		text.SetText(url.QueryEscape(input.Text))
	})

	ButtonUrlDecode := widget.NewButton("URL Decode", func() {
		decoded, err := url.QueryUnescape(input.Text)
		if err != nil {
			dialog.ShowError(err, w)
		}

		text.SetText(decoded)
	})

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			input,
			ButtonB64Encode,
			ButtonB64Decode,
			ButtonUrlEncode,
			ButtonUrlDecode,
			fyne.NewContainerWithLayout(layout.NewVBoxLayout(), text),
		),
	)

	w.ShowAndRun()
}
