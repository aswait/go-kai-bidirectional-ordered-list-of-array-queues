package ui

import (
	"fmt"
	"go-kai/pkg/city"
	"go-kai/pkg/district"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Экран вывода всех районов города
func ShowDistrictsScreen(window fyne.Window, currentCity *city.City) {
	districts := container.NewVBox()

	for node := currentCity.Head; node != nil; node = node.Next {
		district := node.GetCityDistrcit()
		button := widget.NewButton(district.GetDistrictName(), func() {
			ShowDistrictScreen(window, currentCity, district)
		})

		deleteButton := DeleteButton("✖️", func() {
			ShowConfirmDialog(window, "Вы уверены, что хотите удалить этот объект?", func() {
				currentCity.RemoveNode(district.GetDistrictName())

				dialog.ShowInformation("Удалено", "Объект был успешно удалён", window)

				ShowDistrictsScreen(window, currentCity)
			})
		})

		entryContainer := container.NewBorder(
			nil,
			nil,
			nil,
			deleteButton,
			button,
		)

		districts.Add(entryContainer)
	}

	addDistrictButton := AddButton(func() {
		AddDistrictScreen(window, currentCity)
	})
	districts.Add(addDistrictButton)

	backButton := widget.NewButton("Назад", func() {
		ShowMainScreen(window, currentCity)
	})
	backButton.Importance = widget.HighImportance

	bottomButtons := container.NewCenter(
		container.NewHBox(backButton),
	)

	content := container.NewBorder(
		nil,           // Верхний элемент
		bottomButtons, // Нижний элемент
		nil,           // Левый элемент
		nil,           // Правый элемент
		container.NewVBox(
			Title("Районы города "+currentCity.GetCityName()),
			canvas.NewText("", color.Black),
			districts,
		),
	)

	window.SetContent(content)
}

// Экран добавления нового района
func AddDistrictScreen(window fyne.Window, currentCity *city.City) {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Введите название района")

	submitButton := widget.NewButton("Готово", func() {
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
	submitButton.Importance = widget.SuccessImportance

	backButton := widget.NewButton("Назад", func() {
		ShowDistrictsScreen(window, currentCity)
	})
	backButton.Importance = widget.HighImportance

	buttons := container.NewCenter(
		container.NewHBox(backButton, submitButton),
	)

	topBox := container.NewVBox(
		Title("Добавление нового района"),
		canvas.NewText("", color.Black),
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

// Экран отдельного района
func ShowDistrictScreen(window fyne.Window, currentCity *city.City, distrcit *district.District) {
	title := Title("Район: " + distrcit.GetDistrictName())

	backButton := BackButton(func() {
		ShowDistrictsScreen(window, currentCity)
	})

	deleteButton := DeleteButton("Удалить", func() {
		ShowConfirmDialog(window, "Вы уверены, что хотите удалить этот объект?", func() {
			currentCity.RemoveNode(distrcit.GetDistrictName())

			dialog.ShowInformation("Удалено", "Объект был успешно удалён", window)

			ShowDistrictsScreen(window, currentCity)
		})
	})

	changeButton := ChangeButton(func() {
		ShowChangeDistrictScreen(window, currentCity, distrcit)
	})

	buttons := container.NewCenter(
		container.NewHBox(backButton, changeButton, deleteButton),
	)

	streetsButton := widget.NewButton("Улицы района", func() {
		ShowStreetsScreen(window, currentCity, distrcit)
	})

	infoButton := widget.NewButton("Актуальная информация", func() {
		ShowDistrictInfoScreen(window, currentCity, distrcit)
	})

	topBox := container.NewVBox(
		title,
		canvas.NewText("", color.Black),
		streetsButton,
		infoButton,
	)

	content := container.NewBorder(
		topBox,  // Верхний элемент
		buttons, // Нижний элемент
		nil,     // Левый элемент
		nil,     // Правый элемент
	)

	window.SetContent(content)
}

// Экран вывода информации о районе
func ShowDistrictInfoScreen(window fyne.Window, currentCity *city.City, distrcit *district.District) {
	title := Title("Район: " + distrcit.GetDistrictName())

	backButton := BackButton(func() {
		ShowDistrictScreen(window, currentCity, distrcit)
	})

	streetsCounterText := canvas.NewText(
		fmt.Sprintf("Кол-во улиц: %v", distrcit.GetLength()),
		color.White)
	streetsCounterText.TextSize = 16

	streetsLengthText := canvas.NewText(
		fmt.Sprintf("Общая протяженность улиц: %v км", distrcit.GetTotalStreetsLength()),
		color.White)
	streetsLengthText.TextSize = 16

	structureInfoBox := container.NewVBox()

	for _, street := range distrcit.GetDistrcitStreets() {
		streetName := canvas.NewText("• ул. "+street.GetStreetName(), color.White)
		streetName.TextSize = 16

		structureInfoBox.Add(streetName)
	}

	topBox := container.NewVBox(
		title,
		canvas.NewText("", color.Black),
	)

	buttons := container.NewCenter(
		container.NewHBox(backButton),
	)

	infoContainer := container.NewVBox(
		streetsCounterText,
		streetsLengthText,
		canvas.NewText("", color.Black),
		structureInfoBox,
	)

	content := container.NewBorder(
		topBox,  // Верхний элемент
		buttons, // Нижний элемент
		nil,     // Левый элемент
		nil,     // Правый элемент
		infoContainer,
	)

	window.SetContent(content)
}

// Экран изменения информации о городе
func ShowChangeDistrictScreen(window fyne.Window, currentCity *city.City, distrcit *district.District) {
	title := Title("Район: " + distrcit.GetDistrictName())

	entry := EntryBlock("Новое название района")

	submitButton := widget.NewButton("Готово", func() {
		if entry.Text == "" {
			dialog.ShowError(fmt.Errorf("Название района не может быть пустым"), window)
			return
		}

		err := distrcit.SetDistrictName(entry.Text)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		dialog.ShowInformation("Успешно", fmt.Sprintf("Название района обновлено"), window)

		ShowDistrictScreen(window, currentCity, distrcit)
	})
	submitButton.Importance = widget.SuccessImportance

	backButton := BackButton(func() {
		ShowDistrictScreen(window, currentCity, distrcit)
	})

	buttons := container.NewCenter(
		container.NewHBox(backButton, submitButton),
	)

	topBox := container.NewVBox(
		title,
		canvas.NewText("", color.Black),
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
