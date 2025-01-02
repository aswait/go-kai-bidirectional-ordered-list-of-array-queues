package ui

import (
	"fmt"
	"go-kai/pkg/city"
	"go-kai/pkg/district"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func AddDistrictScreen(window fyne.Window, currentCity *city.City) {
	label := widget.NewLabelWithStyle(
		"Добавление нового района",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true, Italic: true},
	)

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Введите название района")

	submitButton := widget.NewButton("Добавить", func() {
		if entry.Text == "" {
			dialog.ShowError(fmt.Errorf("Название района не может быть пустым"), window)
			return
		}

		district := district.NewDistrict(entry.Text)

		err := currentCity.AddNode(district)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		dialog.ShowInformation("Успешно", fmt.Sprintf("Район '%s' успешно добавлен", district.GetDistrictName()), window)

		ShowDistrictsScreen(window, currentCity)
	})

	backButton := widget.NewButton("Назад", func() {
		ShowDistrictsScreen(window, currentCity)
	})
	backButton.Importance = widget.HighImportance

	buttons := container.NewVBox(
		submitButton,
		backButton,
	)

	topBox := container.NewVBox(
		label,
		entry,
	)

	content := container.NewBorder(
		topBox,  // Верхний элемент
		buttons, // Нижний элемент
		nil,     // Левый элемент
		nil,     // Правый элемент
	)

	window.SetContent(content)
}

func ShowDistrictScreen(window fyne.Window, currentCity *city.City, distrcit *district.District) {
	label := widget.NewLabelWithStyle(
		"Район: "+distrcit.GetDistrictName(),
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true, Italic: true},
	)

	streetsButton := widget.NewButton("Улицы района", func() {
		ShowStreetsScreen(window, currentCity, distrcit)
	})

	backButton := widget.NewButton("Назад", func() {
		ShowDistrictsScreen(window, currentCity)
	})
	backButton.Importance = widget.HighImportance

	buttons := container.NewVBox(
		backButton,
	)

	topBox := container.NewVBox(
		label,
		streetsButton,
	)

	content := container.NewBorder(
		topBox,  // Верхний элемент
		buttons, // Нижний элемент
		nil,     // Левый элемент
		nil,     // Правый элемент
	)

	window.SetContent(content)
}

func ShowStreetsScreen(window fyne.Window, currentCity *city.City, distrcit *district.District) {
	label := widget.NewLabelWithStyle(
		"Улицы района "+distrcit.GetDistrictName(),
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true, Italic: true},
	)

	streets := container.NewVBox()

	for _, street := range distrcit.GetDistrcitStreets() {
		streetName := street.GetStreetName()
		button := widget.NewButton(streetName, func() {
			ShowStreetScreen(window, currentCity, distrcit, &street)
		})
		streets.Add(button)
	}

	addStreetButton := widget.NewButton("Добавить улицу", func() {
		AddStreetScreen(window, currentCity, distrcit)
	})
	addStreetButton.Importance = widget.SuccessImportance

	backButton := widget.NewButton("Назад", func() {
		ShowDistrictScreen(window, currentCity, distrcit)
	})
	backButton.Importance = widget.HighImportance

	buttons := container.NewVBox(
		addStreetButton,
		backButton,
	)

	content := container.NewBorder(
		nil,     // Верхний элемент
		buttons, // Нижний элемент
		nil,     // Левый элемент
		nil,     // Правый элемент
		container.NewVBox(
			label,
			streets,
		),
	)

	window.SetContent(content)
}
