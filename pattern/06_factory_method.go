package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Фабричный метод необходим для создания обьектов класса с помощью внешнего интерфейса.
	Плюсы:
	- Позволяет расширять область создаваемых обьектов добавления новых связей
	- Реализация принципа открытости/закрытости
	Минусы:
	- При расширении приводит к созданию большой иерархии классов, т.к. для каждой сущности необходим класс создателя
*/

type DeliverType int

const (
	Walker DeliverType = iota
	CarOwner
	Other
)

type iDeliver interface {
	getDeliver() string
	setDeliver(string)
	getSalary() float64
	setSalary(float64)
	getType() DeliverType
}

type deliver struct {
	d_type DeliverType
	name   string
	salary float64
	// other fields
}

type CarDeliver struct {
	deliver
}

func (c *CarDeliver) getDeliver() string {
	return c.name
}

func (c *CarDeliver) setDeliver(n string) {
	c.name = n
}

func (c *CarDeliver) getSalary() float64 {
	return c.salary
}

func (c *CarDeliver) setSalary(s float64) {
	c.salary = s
}

func (c *CarDeliver) getType() DeliverType {
	return c.d_type
}

func newCarDeliver() iDeliver {
	return &CarDeliver{
		deliver: deliver{
			d_type: CarOwner,
			name:   "-",
			salary: 15503.21,
		},
	}
}

type WalkDeliver struct {
	deliver
}

func (w *WalkDeliver) getDeliver() string {
	return w.name
}

func (w *WalkDeliver) setDeliver(n string) {
	w.name = n
}

func (w *WalkDeliver) getSalary() float64 {
	return w.salary
}

func (w *WalkDeliver) setSalary(s float64) {
	w.salary = s
}

func (w *WalkDeliver) getType() DeliverType {
	return w.d_type
}

func newWalkDeliver() iDeliver {
	return &CarDeliver{
		deliver: deliver{
			d_type: Walker,
			name:   "-",
			salary: 12645.32,
		},
	}
}

func getDelever(deliverType DeliverType) (iDeliver, error) {
	switch deliverType {
	case Walker:
		return newWalkDeliver(), nil
	case CarOwner:
		return newCarDeliver(), nil
	default:
		fmt.Printf("Error occured, enter correct type of deliver")
		return nil, fmt.Errorf("wrong type of deliver")
	}
}

func showDeliver(d iDeliver) {
	fmt.Println()
	fmt.Println("Type 1 == 'has car' || Type 0 == 'has no car'")
	fmt.Println("Type is", d.getType())
	fmt.Println("Name is", d.getDeliver())
	fmt.Println("Salary is", d.getSalary())
	fmt.Println()
}

func Usage6() {
	deliver1, _ := getDelever(Walker)
	deliver2, _ := getDelever(CarOwner)

	showDeliver(deliver1)
	showDeliver(deliver2)
}
