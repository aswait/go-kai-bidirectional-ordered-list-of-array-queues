package street

import "fmt"

// Интерфейс Streeter
type Streeter interface {
	// Геттеры объекта (получение значений)
	GetStreetName() string
	GetStreetLength() int

	// Сеттеры объекта (установка значений)
	SetName(name string) error
	SetLength(length int)
}

// Информационный объект: Улица города
// Реализация интерфейса Streeter
type Street struct {
	Name   string `json:"name"`   // свойство Название улицы
	Length int    `json:"length"` // свойство Длина улицы
}

// Ф-ия конструктор объекта Улица
func NewStreet(name string, length int) *Street {
	return &Street{
		Name:   name,
		Length: length,
	}
}

// Метод для получени имени
func (s *Street) GetStreetName() string {
	return s.Name
}

// Метод для получения длины
func (s *Street) GetStreetLength() int {
	return s.Length
}

// Метод для установки имени
func (s *Street) SetName(name string) error {
	if name == s.Name {
		return fmt.Errorf("Новое название улицы должно отличаться от старого")
	}

	s.Name = name
	return nil
}

// Метод для установки длины
func (s *Street) SetLength(length int) {
	s.Length = length
}
