package street

// Интерфейс Streeter
type Streeter interface {
}

// Информационный объект: Улица города
// Реализация интерфейса Streeter
type Street struct {
	Name   string // свойство Название улицы
	Length int    // свойство Длина улицы
}

// Ф-ия конструктор объекта Улица
func NewStreet(name string, length int) *Street {
	return &Street{
		Name:   name,
		Length: length,
	}
}

func (s *Street) GetStreetName() string {
	return s.Name
}
