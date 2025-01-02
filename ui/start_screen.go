package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ShowStartScreen отображает начальный экран с вводом названия города
func ShowStartScreen(window fyne.Window, onSubmit func(cityName string)) {
	// Поле ввода
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Введите название города")

	// Кнопка подтверждения
	submitButton := widget.NewButton("Подтвердить", func() {
		onSubmit(entry.Text) // Передаём введённое название города
	})

	// Размещение элементов на экране
	content := container.NewVBox(
		widget.NewLabel("Введите название города:"),
		entry,
		submitButton,
	)

	// Устанавливаем содержимое окна
	window.SetContent(content)
}
