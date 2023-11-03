package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Шаблон фасад позволяет сокрыть внутренний функционал и сложные связи от клиента.
	Плюсы:
	- Упрощение работы клиентов
	- Сокрытие сложных зависимостей
	- Уменьшение риска ошибок
	Минусы:
	- Ограничение функционала функционалом фасада
	- Нет тонкой настройки системы (black box)

	Пример: Работа с заказом
*/

// Фасадный тип, обеспечивающий взаимодействие
// с подсистемой
type Order struct {
	payment  *Payment
	delivery *Delivery
	store    *StoreActions
}

func NewOrder() *Order {
	return &Order{
		payment:  NewPayment(),
		delivery: NewDelivery(),
		store:    NewStoreAction(),
	}
}

func (o *Order) HandleOrder() {
	o.payment.pay()
	o.store.removeFromStore()
	o.delivery.deliver()
}

// Класс оплаты по заказу
type Payment struct {
}

func NewPayment() *Payment {
	return &Payment{}
}

func (p *Payment) pay() {
	fmt.Println("::withdraw money")
	fmt.Println("::send check")
}

// Класс доставки заказа
type Delivery struct {
}

func NewDelivery() *Delivery {
	return &Delivery{}
}

func (d *Delivery) deliver() {
	fmt.Println("::send order to post station")
	fmt.Println("::change status to deliver")
}

// Класс взаимодействия склад - заказ
type StoreActions struct {
}

func NewStoreAction() *StoreActions {
	return &StoreActions{}
}

func (sa *StoreActions) removeFromStore() {
	fmt.Println("::remove item from store")
}

func Usage1() {
	ord1 := NewOrder()
	ord1.HandleOrder()
}
