package run

import (
	"fmt"
	"go-kai/pkg/city"
	"go-kai/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Структура города")

	ui.ShowStartScreen(myWindow, func(cityName string) {
		if cityName == "" {
			dialog.ShowError(fmt.Errorf("Название города не может быть пустым"), myWindow)
			return
		}

		// Создаём объект города
		currentCity := city.NewCity(cityName)

		// Переходим на основной экран
		ui.ShowMainScreen(myWindow, currentCity)
	})

	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.ShowAndRun()
}
