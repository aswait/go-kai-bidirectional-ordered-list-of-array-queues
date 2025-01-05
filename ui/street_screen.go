package ui

import (
	"fmt"
	"go-kai/pkg/city"
	"go-kai/pkg/district"
	"go-kai/pkg/street"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Экран вывода всех улиц
func ShowStreetsScreen(window fyne.Window, currentCity *city.City, distrcit *district.District) {
	title := Title("Улицы района " + distrcit.GetDistrictName())

	streets := container.NewVBox()

	for _, street := range distrcit.GetDistrcitStreets() {
		streetName := street.GetStreetName()
		button := widget.NewButton(streetName, func() {
			ShowStreetScreen(window, currentCity, distrcit, &street)
		})

		deleteButton := DeleteButton("✖️", func() {
			ShowConfirmDialog(window, "Вы уверены, что хотите удалить этот объект?", func() {
				distrcit.RemoveStreet(streetName)

				dialog.ShowInformation("Удалено", "Объект был успешно удалён", window)

				ShowStreetsScreen(window, currentCity, distrcit)
			})
		})

		entryContainer := container.NewBorder(
			nil,
			nil,
			nil,
			deleteButton,
			button,
		)

		streets.Add(entryContainer)
	}

	addStreetButton := AddButton(func() {
		AddStreetScreen(window, currentCity, distrcit)
	})

	streets.Add(addStreetButton)

	backButton := widget.NewButton("Назад", func() {
		ShowDistrictScreen(window, currentCity, distrcit)
	})
	backButton.Importance = widget.HighImportance

	buttons := container.NewCenter(
		container.NewHBox(backButton),
	)

	content := container.NewBorder(
		nil,     // Верхний элемент
		buttons, // Нижний элемент
		nil,     // Левый элемент
		nil,     // Правый элемент
		container.NewVBox(
			title,
			canvas.NewText("", color.Black),
			streets,
		),
	)

	window.SetContent(content)
}

// Экран добавления новой улицы
func AddStreetScreen(window fyne.Window, currentCity *city.City, distrcit *district.District) {
	title := Title("Добавление новой улицы")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Введите название улицы")

	lengthEntry := widget.NewEntry()
	lengthEntry.SetPlaceHolder("Введите длину улицы (в км)")

	submitButton := widget.NewButton("Готово", func() {
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
	submitButton.Importance = widget.SuccessImportance

	backButton := widget.NewButton("Назад", func() {
		ShowStreetsScreen(window, currentCity, distrcit)
	})
	backButton.Importance = widget.HighImportance

	buttons := container.NewCenter(
		container.NewHBox(
			backButton,
			submitButton,
		),
	)

	topBox := container.NewVBox(
		title,
		canvas.NewText("", color.Black),
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

// Экран отдельной улицы
func ShowStreetScreen(
	window fyne.Window,
	currentCity *city.City,
	distrcit *district.District,
	street *street.Street,
) {
	title := Title("Улица: " + street.GetStreetName())

	changeButton := ChangeButton(func() {
		ShowStreetChangeScreen(window, currentCity, distrcit, street)
	})

	infoButton := widget.NewButton("Актуальная информация", func() {
		ShowStreetInfoScreen(window, currentCity, distrcit, street)
	})

	backButton := widget.NewButton("Назад", func() {
		ShowStreetsScreen(window, currentCity, distrcit)
	})
	backButton.Importance = widget.HighImportance

	deleteButton := DeleteButton("Удалить", func() {
		ShowConfirmDialog(window, "Вы уверены, что хотите удалить этот объект?", func() {
			distrcit.RemoveStreet(street.GetStreetName())

			dialog.ShowInformation("Удалено", "Объект был успешно удалён", window)

			ShowStreetsScreen(window, currentCity, distrcit)
		})
	})
	deleteButton.Importance = widget.DangerImportance

	buttons := container.NewCenter(
		container.NewHBox(backButton, changeButton, deleteButton),
	)

	topBox := container.NewVBox(
		title,
		canvas.NewText("", color.Black),
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

// Экран вывода информации об улице
func ShowStreetInfoScreen(
	window fyne.Window,
	currentCity *city.City,
	distrcit *district.District,
	street *street.Street,
) {
	title := Title("Улица " + street.GetStreetName())

	backButton := widget.NewButton("Назад", func() {
		ShowStreetScreen(window, currentCity, distrcit, street)
	})
	backButton.Importance = widget.HighImportance

	streetNameText := canvas.NewText(
		fmt.Sprintf("Название: %s", street.GetStreetName()),
		color.White)
	streetNameText.TextSize = 16

	streetLengthText := canvas.NewText(
		fmt.Sprintf("Длина: %v", street.GetStreetLength()),
		color.White)
	streetLengthText.TextSize = 16

	buttons := container.NewCenter(
		container.NewHBox(backButton),
	)

	topBox := container.NewVBox(
		title,
	)

	content := container.NewBorder(
		topBox,  // Верхний элемент
		buttons, // Нижний элемент
		nil,     // Левый элемент
		nil,     // Правый элемент
		container.NewVBox(
			canvas.NewText("", color.Black),
			streetNameText,
			streetLengthText,
		),
	)

	window.SetContent(content)
}

// Экран изменения информации об улице
func ShowStreetChangeScreen(
	window fyne.Window,
	currentCity *city.City,
	distrcit *district.District,
	street *street.Street,
) {
	title := Title("Улица " + street.GetStreetName())

	backButton := widget.NewButton("Назад", func() {
		ShowStreetScreen(window, currentCity, distrcit, street)
	})
	backButton.Importance = widget.HighImportance

	nameEntry := EntryBlock("Новое название района")

	lengthEntry := EntryBlock("Новая длина улицы")

	submitButton := widget.NewButton("Готово", func() {
		if nameEntry.Text == "" {
			dialog.ShowError(fmt.Errorf("Название улицы не может быть пустым"), window)
			return
		}

		length, err := strconv.Atoi(lengthEntry.Text)
		if err != nil || length < 0 {
			dialog.ShowError(fmt.Errorf("Некоректные данные для длины улицы"), window)
			return
		}

		err = street.SetName(nameEntry.Text)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		street.SetLength(length)

		dialog.ShowInformation("Успешно", fmt.Sprintf("Улица '%s' успешно добавлен", street.GetStreetName()), window)

		ShowStreetScreen(window, currentCity, distrcit, street)
	})
	submitButton.Importance = widget.SuccessImportance

	buttons := container.NewCenter(
		container.NewHBox(backButton, submitButton),
	)

	topBox := container.NewVBox(
		title,
	)

	content := container.NewBorder(
		topBox,  // Верхний элемент
		buttons, // Нижний элемент
		nil,     // Левый элемент
		nil,     // Правый элемент
		container.NewVBox(
			canvas.NewText("", color.Black),
			nameEntry,
			lengthEntry,
		),
	)

	window.SetContent(content)
}
