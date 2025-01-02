package ui

import (
	"go-kai/pkg/city"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowMainScreen(window fyne.Window, currentCity *city.City) {
	// Настроенный Label
	cityLabel := widget.NewLabelWithStyle(
		"Город: "+currentCity.GetCityName(),
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true, Italic: true},
	)

	districtsButton := widget.NewButton("Районы города", func() {
		ShowDistrictsScreen(window, currentCity)
	})

	cityInfoButton := widget.NewButton("Актуальная информация", func() {
		// Добавьте свою логику
	})

	changeCityNameButton := widget.NewButton("Изменить название", func() {
		// Добавьте свою логику
	})

	// Обёртка с отступами
	content := container.NewVBox(
		cityLabel,
		districtsButton,
		cityInfoButton,
		changeCityNameButton,
	)

	// Устанавливаем фон и содержимое
	window.SetContent(content)
}

func ShowDistrictsScreen(window fyne.Window, currentCity *city.City) {
	label := widget.NewLabelWithStyle(
		"Районы города "+currentCity.GetCityName(),
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true, Italic: true},
	)

	districts := container.NewVBox()

	for node := currentCity.Head; node != nil; node = node.Next {
		district := node.GetCityDistrcit()
		button := widget.NewButton(district.GetDistrictName(), func() {
			ShowDistrictScreen(window, currentCity, district)
		})
		districts.Add(button)
	}

	backButton := widget.NewButton("Назад", func() {
		ShowMainScreen(window, currentCity)
	})
	backButton.Importance = widget.HighImportance

	addDistrictButton := widget.NewButton("Добавить район", func() {
		AddDistrictScreen(window, currentCity)
	})
	addDistrictButton.Importance = widget.SuccessImportance

	buttons := container.NewVBox(
		addDistrictButton,
		backButton,
	)

	content := container.NewBorder(
		nil,     // Верхний элемент
		buttons, // Нижний элемент
		nil,     // Левый элемент
		nil,     // Правый элемент
		container.NewVBox(
			label,
			districts,
		),
	)

	window.SetContent(content)
}
