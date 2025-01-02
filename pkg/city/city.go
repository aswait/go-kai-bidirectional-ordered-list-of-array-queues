package city

import (
	"fmt"
	"go-kai/pkg/district"
)

type Citier interface {
}

// Объект: Двунаправленный список город
type City struct {
	Name    string    // свойство название
	Head    *CityNode // Ссылка на первый объект списка
	Tail    *CityNode // Ссылка на последний объект списка
	Current *CityNode // Ссылка на текущий объект списка
	Length  int       // свойство длина списка (кол-во узлов)
}

// Ф-ия конструктор объекта City
func NewCity(name string) *City {
	return &City{Name: name}
}

type CityNoder interface {
}

// Объект: Узел двунаправленного списка
// Реализация интерфейса DistrictNoder
type CityNode struct {
	District *district.District // Ссылка на объект района
	Next     *CityNode          // Ссылка на следующий узел
	Prev     *CityNode          // Ссылка на предыдущий узел
}

func (n *CityNode) GetCityDistrcit() *district.District {
	return n.District
}

// Ф-ия конструктор объекта CityNode
func NewDistrictNode(distrcit *district.District) *CityNode {
	return &CityNode{District: distrcit}
}

func (c *City) GetCityName() string {
	return c.Name
}

func (c *City) AddNode(district *district.District) error {
	for node := c.Head; node != nil; node = node.Next {
		if node.District.Name == district.Name {
			return fmt.Errorf("Район с названием '%s' уже существует", district.GetDistrictName())
		}
	}

	// Создаём новый узел
	newNode := &CityNode{District: district}

	if c.Head == nil {
		// Если список пустой
		c.Head = newNode
		c.Tail = newNode
		c.Current = newNode
	} else {
		// Если список не пустой, добавляем узел в конец
		newNode.Prev = c.Tail
		c.Tail.Next = newNode
		c.Tail = newNode
		c.Current = newNode
	}

	// Увеличиваем длину списка
	c.Length++

	return nil
}

func (c *City) FindDistrictByName(name string) *CityNode {
	// Проходим по всему списку
	for node := c.Head; node != nil; node = node.Next {
		if node.District.Name == name {
			return node // Возвращаем найденный узел
		}
	}
	return nil // Если узел не найден, возвращаем nil
}
