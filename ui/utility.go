package ui

import (
	"go-kai/pkg/city"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func BackButton(backScreen func()) *widget.Button {
	backButton := widget.NewButton("Назад", func() {
		backScreen()
	})
	backButton.Importance = widget.HighImportance

	return backButton
}

func Title(text string) *canvas.Text {
	label := canvas.NewText(text, color.White)
	label.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	label.TextSize = 24
	label.Alignment = fyne.TextAlignCenter

	return label
}

func AddButton(addFunc func()) *widget.Button {
	addDistrictButton := widget.NewButton("+", func() {
		addFunc()
	})
	addDistrictButton.Importance = widget.SuccessImportance

	return addDistrictButton
}

func EntryBlock(text string) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(text)

	return entry
}

func DeleteButton(text string, delFunc func()) *widget.Button {
	delButton := widget.NewButton(text, func() {
		delFunc()
	})
	delButton.Importance = widget.DangerImportance

	return delButton
}

func ChangeButton(changeFunc func()) *widget.Button {
	changeButton := widget.NewButton("Изменить", func() {
		changeFunc()
	})
	changeButton.Importance = widget.WarningImportance

	return changeButton
}

// Диалог подтверждения удаления
func ShowConfirmDialog(window fyne.Window, message string, onConfirm func()) {
	dialog.ShowConfirm(
		"Подтверждение удаления",
		message,
		func(confirmed bool) {
			if confirmed {
				onConfirm()
			}
		},
		window,
	)
}

func ExportToJson(window fyne.Window, currentCity *city.City) {
	jsonData, err := currentCity.ToJSON()
	if err != nil {
		dialog.ShowError(err, window)
		return
	}

	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, _ error) {
		if writer == nil {
			return
		}
		defer writer.Close()

		_, err := writer.Write([]byte(jsonData))
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		dialog.ShowInformation("Успех", "Файл успешно сохранён!", window)
	}, window)

	saveDialog.SetFilter(
		storage.NewExtensionFileFilter([]string{".json"}),
	)
	saveDialog.Show()
}

func ImportToJson(window fyne.Window, currentCity *city.City) {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil || reader == nil {
			return
		}
		defer reader.Close()

		data, err := os.ReadFile(reader.URI().Path())
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		if err := currentCity.ImportFromJSON(string(data)); err != nil {
			dialog.ShowError(err, window)
		} else {
			dialog.ShowInformation("Успех", "Импорт завершён успешно", window)
		}
	}, window)
}
