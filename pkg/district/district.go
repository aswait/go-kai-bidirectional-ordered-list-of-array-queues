package district

import (
	"fmt"
	"go-kai/pkg/street"
)

// Интерфейс Districter
type Districter interface {
	// Геттеры
	GetDistrictName() string
	GetLength() int
	GetDistrcitStreets() []street.Street
	GetTotalStreetsLength() int

	// Сеттеры
	SetDistrictName(name string) error

	// Другие необходимые методы
	FindStreetByName(name string) *street.Street
	AddStreet(street *street.Street) error
	RemoveStreet(name string) bool
}

// Объект: Район города
// Реализация интерфейса Districter
type District struct {
	Name    string          `json:"name"`    // свойство Название района
	Streets []street.Street `json:"streets"` // поле объединяющие информационные объекты "Улица"
}

// Ф-ия конструктор объекта District
func NewDistrict(name string) *District {
	streets := make([]street.Street, 0) // Инициализация списка улиц

	return &District{
		Name:    name,
		Streets: streets,
	}
}

// Метод для получения имени
func (d *District) GetDistrictName() string {
	return d.Name
}

// Метод для получения улиц
func (d *District) GetDistrcitStreets() []street.Street {
	return d.Streets
}

// Метод для получения кол-ва улиц
func (d *District) GetLength() int {
	return len(d.Streets)
}

// Метод для получения общей длины всех улиц района
func (d *District) GetTotalStreetsLength() int {
	var streetsLength int

	for _, street := range d.Streets {
		streetsLength += street.GetStreetLength()
	}

	return streetsLength
}

// Метод для установки имени
func (d *District) SetDistrictName(name string) error {
	if name == d.Name {
		return fmt.Errorf("Новое название района должно отличаться от старого")
	}

	d.Name = name
	return nil
}

// Метод для поиска улицы по имени
func (d *District) FindStreetByName(name string) *street.Street {
	for i := range d.Streets {
		if d.Streets[i].GetStreetName() == name {
			return &d.Streets[i] // Возвращаем найденную улицу
		}
	}
	return nil
}

// Метод для добавления улицы
func (d *District) AddStreet(street *street.Street) error {
	findedStreet := d.FindStreetByName(street.GetStreetName())
	if findedStreet != nil {
		return fmt.Errorf("Район с названием '%s' уже существует", findedStreet.GetStreetName())
	}

	d.Streets = append(d.Streets, *street)
	return nil
}

// Метод для удаления улицы
func (d *District) RemoveStreet(name string) bool {
	for i, street := range d.Streets {
		if street.Name == name {
			d.Streets = append(d.Streets[:i], d.Streets[i+1:]...)
			return true
		}
	}
	return false
}
