package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	Паттерн состояние используется для описания обьектов, которые находятся в различном состоянии и для оркестрации данных состояний.
	Ключевая особенность патерна в возможности изменять поведения объекта в зависимости от сотояния

	Плюсы:
	- Объект может изменять свое поведение в ходе программы
	- Изменения состояния одного обьекта не влияет на другой
	Минусы:
	-
*/

type State interface {
	changeState(obj *Object)
	doAction(obj *Object)
}

type Object struct {
	state State
}

func (obj *Object) setState(state State) {
	obj.state = state
}

func (obj *Object) doAction() {
	obj.state.doAction(obj)
}
func (obj *Object) changeState() {
	obj.state.changeState(obj)
}

type SaleOn struct{}

func (on *SaleOn) changeState(obj *Object) {
	obj.setState(&SaleOff{})
}

func (on *SaleOn) doAction(obj *Object) {
	fmt.Println("Calculation during sale")
}

type SaleOff struct{}

func (off *SaleOff) changeState(obj *Object) {
	obj.setState(&SaleOn{})
}

func (off *SaleOff) doAction(obj *Object) {
	fmt.Println("Calculation without sale")
}

func Usage8() {
	shop1 := &Object{
		state: &SaleOff{},
	}

	shop1.doAction()

	shop1.changeState()

	shop1.doAction()
}
