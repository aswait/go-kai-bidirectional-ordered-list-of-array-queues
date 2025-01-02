package ui

import (
	"fmt"
	"go-kai/pkg/city"
	"go-kai/pkg/district"
	"go-kai/pkg/street"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func AddStreetScreen(window fyne.Window, currentCity *city.City, distrcit *district.District) {
	label := widget.NewLabelWithStyle(
		"Добавление новой улицы",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true, Italic: true},
	)

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Введите название улицы")

	lengthEntry := widget.NewEntry()
	lengthEntry.SetPlaceHolder("Введите длину улицы (в км)")

	submitButton := widget.NewButton("Добавить", func() {
		if nameEntry.Text == "" {
			dialog.ShowError(fmt.Errorf("Название улицы не может быть пустым"), window)
			return
		}

		length, err := strconv.Atoi(lengthEntry.Text)
		if err != nil || length < 0 {
			dialog.ShowError(fmt.Errorf("Некоректные данные для длины улицы"), window)
			return
		}

		street := street.NewStreet(nameEntry.Text, length)
		err = distrcit.AddStreet(street)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		dialog.ShowInformation("Успешно", fmt.Sprintf("Улица '%s' успешно добавлен", street.GetStreetName()), window)

		ShowStreetsScreen(window, currentCity, distrcit)
	})

	backButton := widget.NewButton("Назад", func() {
		ShowStreetsScreen(window, currentCity, distrcit)
	})
	backButton.Importance = widget.HighImportance

	buttons := container.NewVBox(
		submitButton,
		backButton,
	)

	topBox := container.NewVBox(
		label,
		nameEntry,
		lengthEntry,
	)

	content := container.NewBorder(
		topBox,  // Верхний элемент
		buttons, // Нижний элемент
		nil,     // Левый элемент
		nil,     // Правый элемент
	)

	window.SetContent(content)
}

func ShowStreetScreen(
	window fyne.Window,
	currentCity *city.City,
	distrcit *district.District,
	street *street.Street,
) {
	label := widget.NewLabelWithStyle(
		"Улица: "+street.GetStreetName(),
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true, Italic: true},
	)

	backButton := widget.NewButton("Назад", func() {
		ShowStreetsScreen(window, currentCity, distrcit)
	})
	backButton.Importance = widget.HighImportance

	buttons := container.NewVBox(
		backButton,
	)

	topBox := container.NewVBox(
		label,
	)

	content := container.NewBorder(
		topBox,  // Верхний элемент
		buttons, // Нижний элемент
		nil,     // Левый элемент
		nil,     // Правый элемент
	)

	window.SetContent(content)
}
