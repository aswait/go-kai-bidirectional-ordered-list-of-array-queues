package district

import (
	"fmt"
	"go-kai/pkg/street"
)

// Интерфейс Districter
type Districter interface {
}

// Объект: Район города
// Реализация интерфейса Districter
type District struct {
	Name    string          // свойство Название района
	Streets []street.Street // поле объединяющие информационные объекты "Улица"
}

// Ф-ия конструктор объекта District
func NewDistrict(name string) *District {
	streets := make([]street.Street, 0) // Инициализация списка улиц

	return &District{
		Name:    name,
		Streets: streets,
	}
}

func (d *District) GetDistrictName() string {
	return d.Name
}

func (d *District) GetDistrcitStreets() []street.Street {
	return d.Streets
}

func (d *District) FindStreetByName(name string) *street.Street {
	for i := range d.Streets {
		if d.Streets[i].GetStreetName() == name {
			return &d.Streets[i] // Возвращаем найденную улицу
		}
	}
	return nil
}

func (d *District) AddStreet(street *street.Street) error {
	// Проверяем, существует ли улица с таким названием
	findedStreet := d.FindStreetByName(street.GetStreetName())
	if findedStreet != nil {
		return fmt.Errorf("Район с названием '%s' уже существует", findedStreet.GetStreetName())
	}

	// Если улица не найдена, добавляем её
	d.Streets = append(d.Streets, *street)
	return nil
}
