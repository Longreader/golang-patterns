package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Шаблон посититель необходим для добавления функционала для существующей иерархии классов без изменения классов
	Плюсы:
	- Позволяет работать с несколькими иерархиями классов
	- Позволяет легко расширять функционал классов
	Минусы:
	- При расширении иерархии класса нарушается структура посетителя
	- Добавление нового метода требует

	Пример: Обработка платежей (картой/наличкой)
*/

type Visitor interface {
	visitCard(c *CardPay)
	visitCash(c *CashPay)
}

type PaymentCase interface {
	Accept(visitor Visitor)
}

type CardPay struct {
	// Данные счета и т.д.
}

func (c *CardPay) Accept(visitor Visitor) {
	visitor.visitCard(c)
}

type CashPay struct {
	// Данные по оплате и т.д.
}

func (c *CashPay) Accept(visitor Visitor) {
	visitor.visitCash(c)
}

type RefundCulc struct {
}

func (r *RefundCulc) visitCard(c *CardPay) {
	fmt.Println("Counting refund for card payments")
}

func (r *RefundCulc) visitCash(c *CashPay) {
	fmt.Println("Counting refund for cash payment")
}

type PayCulc struct {
}

func (p *PayCulc) visitCard(c *CardPay) {
	fmt.Println("Payment calculation for card payments")
}

func (p *PayCulc) visitCash(c *CashPay) {
	fmt.Println("Payment calculation for cash payment")
}

func Usage3() {
	cardPayOne := &CardPay{}
	cashPayOne := &CashPay{}

	payCalculation := &PayCulc{}
	refundCalculation := &RefundCulc{}

	cardPayOne.Accept(payCalculation)
	cardPayOne.Accept(refundCalculation)

	cashPayOne.Accept(payCalculation)
	cashPayOne.Accept(refundCalculation)
}
