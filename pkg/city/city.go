package city

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-kai/pkg/district"
	"go-kai/pkg/street"
)

type Citier interface {
	// Геттеры
	GetCityName() string
	GetCityLength() int
	GetTotalStreets() int
	GetTotalLength() int

	// Сеттеры
	SetCityName(name string) error

	// Другие необходимые методы
	AddNode(district *district.District) error
	FindDistrictByName(name string) *CityNode
	RemoveNode(name string) bool
}

// Объект: Двунаправленный список город
type City struct {
	Name    string    `json:"name"`              // свойство название
	Head    *CityNode `json:"head,omitempty"`    // Ссылка на первый объект списка
	Tail    *CityNode `json:"tail,omitempty"`    // Ссылка на последний объект списка
	Current *CityNode `json:"current,omitempty"` // Ссылка на текущий объект списка
	Length  int       `json:"length"`            // свойство длина списка (кол-во узлов)
}

// Ф-ия конструктор объекта City
func NewCity(name string) *City {
	return &City{Name: name}
}

type CityNoder interface {
	GetCityDistrcit() *district.District
}

// Объект: Узел двунаправленного списка
// Реализация интерфейса DistrictNoder
type CityNode struct {
	District *district.District `json:"district"`       // Ссылка на объект района
	Next     *CityNode          `json:"next,omitempty"` // Ссылка на следующий узел
	Prev     *CityNode          `json:"prev,omitempty"` // Ссылка на предыдущий узел
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

	newNode := &CityNode{District: district}

	if c.Head == nil {
		c.Head = newNode
		c.Tail = newNode
		c.Current = newNode
	} else {
		newNode.Prev = c.Tail
		c.Tail.Next = newNode
		c.Tail = newNode
		c.Current = newNode
	}

	c.Length++

	return nil
}

func (c *City) FindDistrictByName(name string) *CityNode {
	for node := c.Head; node != nil; node = node.Next {
		if node.District.Name == name {
			return node
		}
	}
	return nil
}

func (c *City) SetCityName(name string) error {
	if name == c.GetCityName() {
		return fmt.Errorf("Новое название города должно отличаться от старого")
	}

	c.Name = name
	return nil
}

func (c *City) RemoveNode(name string) bool {
	node := c.FindDistrictByName(name)
	if node == nil {
		return false
	}

	if node == c.Head {
		c.Head = node.Next
		if c.Head != nil {
			c.Head.Prev = nil
		}
	}

	if node == c.Tail {
		c.Tail = node.Prev
		if c.Tail != nil {
			c.Tail.Next = nil
		}
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if c.Current == node {
		c.Current = c.Head
	}

	c.Length--
	return true
}

func (c *City) GetCityLength() int {
	return c.Length
}

func (c *City) GetTotalStreets() int {
	var streetsCounter int

	for node := c.Head; node != nil; node = node.Next {
		streetsCounter += node.District.GetLength()
	}

	return streetsCounter
}

func (c *City) GetTotalLength() int {
	var totalLength int

	for node := c.Head; node != nil; node = node.Next {
		totalLength += node.District.GetTotalStreetsLength()
	}

	return totalLength
}

type SerializableCity struct {
	Name      string               `json:"name"`
	Districts []*district.District `json:"districts"`
}

func (c *City) ToSerializable() SerializableCity {
	var districts []*district.District
	current := c.Head

	for current != nil {
		districts = append(districts, current.District)
		current = current.Next
	}

	return SerializableCity{
		Name:      c.Name,
		Districts: districts,
	}
}

func (c *City) ToJSON() (string, error) {
	serializableCity := c.ToSerializable()
	jsonData, err := json.MarshalIndent(serializableCity, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (c *City) ImportFromJSON(jsonData string) error {
	var importCity SerializableCity
	err := json.Unmarshal([]byte(jsonData), &importCity)
	if err != nil {
		return err
	}

	c.Name = importCity.Name

	for _, impDistrict := range importCity.Districts {
		if c.FindDistrictByName(impDistrict.Name) != nil {
			return errors.New("район с именем " + impDistrict.Name + " уже существует")
		}

		newDistrict := district.NewDistrict(impDistrict.Name)

		for _, streetName := range impDistrict.Streets {
			err := newDistrict.AddStreet(street.NewStreet(streetName.Name, streetName.Length))
			if err != nil {
				return err
			}
		}

		c.AddNode(newDistrict)
	}

	return nil
}
