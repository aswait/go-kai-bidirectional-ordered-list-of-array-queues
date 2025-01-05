package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ShowStartScreen отображает начальный экран с вводом названия города
func ShowStartScreen(window fyne.Window, onSubmit func(cityName string)) {
	// Заголовок
	label := canvas.NewText("Введите название города чтобы продолжить", color.White)
	label.TextSize = 20
	label.Alignment = fyne.TextAlignCenter

	// Поле ввода
	entry := EntryBlock("Название города")

	// Кнопка подтверждения
	submitButton := widget.NewButton("Готово", func() {
		onSubmit(entry.Text) // Передаём введённое название города
	})
	submitButton.Importance = widget.SuccessImportance

	// Контейнер нижних кнопок
	bottomButtons := container.NewCenter(
		container.NewHBox(submitButton),
	)

	// Контейнер ввода
	entryContainer := container.NewVBox(
		canvas.NewText("", color.Black),
		entry,
	)

	content := container.NewBorder(
		label,         // Верхний
		bottomButtons, // Нижний
		nil,           // Левый
		nil,           // Правый
		entryContainer,
	)

	// Устанавливаем содержимое окна
	window.SetContent(content)
}
