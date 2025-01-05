package ui

import (
	"fmt"
	"go-kai/pkg/city"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Главный экран
func ShowMainScreen(window fyne.Window, currentCity *city.City) {
	title := Title("Город: " + currentCity.GetCityName())

	districtsButton := widget.NewButton("Районы города", func() {
		ShowDistrictsScreen(window, currentCity)
	})

	cityInfoButton := widget.NewButton("Актуальная информация", func() {
		ShowInfoScreen(window, currentCity)
	})

	changeCityNameButton := widget.NewButton("Изменить", func() {
		ShowChangeCityNameScreen(window, currentCity)
	})
	changeCityNameButton.Importance = widget.WarningImportance

	bottomButtons := container.NewCenter(
		container.NewHBox(changeCityNameButton),
	)

	content := container.NewBorder(
		title,
		bottomButtons,
		nil,
		nil,
		container.NewVBox(
			canvas.NewText("", color.Black),
			districtsButton,
			cityInfoButton,
		),
	)

	// Устанавливаем фон и содержимое
	window.SetContent(content)
}

// Экран с информацией о городе
func ShowInfoScreen(window fyne.Window, currentCity *city.City) {
	structureInfoBox := container.NewVBox()

	for node := currentCity.Head; node != nil; node = node.Next {
		district := node.GetCityDistrcit()
		districtName := canvas.NewText("• "+district.GetDistrictName()+" р-н", color.White)
		districtName.TextSize = 16

		structureInfoBox.Add(districtName)

		for _, street := range district.GetDistrcitStreets() {
			streetName := canvas.NewText("\t• ул. "+street.GetStreetName(), color.White)
			streetName.TextSize = 16

			structureInfoBox.Add(streetName)
		}
	}

	distrcictsCounterText := canvas.NewText(
		fmt.Sprintf("Кол-во райнов: %v", currentCity.GetCityLength()),
		color.White,
	)
	distrcictsCounterText.TextSize = 16

	streetsCounterText := canvas.NewText(
		fmt.Sprintf("Кол-во улиц: %v", currentCity.GetTotalStreets()),
		color.White)
	streetsCounterText.TextSize = 16

	streetsLengthText := canvas.NewText(
		fmt.Sprintf("Общая протяженность улиц: %v км", currentCity.GetTotalLength()),
		color.White)
	streetsLengthText.TextSize = 16

	backButton := widget.NewButton("Назад", func() {
		ShowMainScreen(window, currentCity)
	})
	backButton.Importance = widget.HighImportance

	exportButton := widget.NewButton("Экспорт структуры", func() {
		ExportToJson(window, currentCity)
	})
	exportButton.Importance = widget.LowImportance

	importButton := widget.NewButton("Импорт структуры", func() {
		ImportToJson(window, currentCity)
	})
	importButton.Importance = widget.LowImportance

	bottomButtons := container.NewCenter(
		container.NewHBox(backButton, exportButton, importButton),
	)

	content := container.NewBorder(
		Title("г. "+currentCity.GetCityName()), // Верхний элемент
		bottomButtons, // Нижний элемент
		nil,           // Левый элемент
		nil,           // Правый элемент
		container.NewVBox(
			canvas.NewText("", color.Black),
			distrcictsCounterText,
			streetsCounterText,
			streetsLengthText,
			canvas.NewText("", color.Black),
			structureInfoBox,
		),
	)

	window.SetContent(content)
}

// Экран изменения города
func ShowChangeCityNameScreen(window fyne.Window, currentCity *city.City) {
	backButton := BackButton(func() {
		ShowMainScreen(window, currentCity)
	})

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Новое название города")

	submitButton := widget.NewButton("Готово", func() {
		if entry.Text == "" {
			dialog.ShowError(fmt.Errorf("Название города не может быть пустым"), window)
			return
		}

		err := currentCity.SetCityName(entry.Text)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		dialog.ShowInformation("Успешно", fmt.Sprintf("Название города обновлено"), window)

		ShowMainScreen(window, currentCity)
	})
	submitButton.Importance = widget.SuccessImportance

	buttons := container.NewCenter(
		container.NewHBox(backButton, submitButton),
	)

	content := container.NewBorder(
		nil,     // Верхний элемент
		buttons, // Нижний элемент
		nil,     // Левый элемент
		nil,     // Правый элемент
		container.NewVBox(
			Title("г. "+currentCity.GetCityName()),
			canvas.NewText("", color.Black),
			entry,
		),
	)

	window.SetContent(content)
}
