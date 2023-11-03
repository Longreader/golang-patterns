package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Паттер необходим для изолирования бизнес логики. Патер реализует инкапсулирование запроса как самостоятельного обьекта.

	Плюсы:
	- Разделение отправителя и получателя
	- Возможность составления очереди исполнения запросов
	- Возможность отмены запросов
	Минусы:
	- Увеличение сложности кода
*/

type Command interface {
	Execute() string
}

type Send struct {
	receiver Receiver
}

func (s *Send) Execute() string {
	return s.receiver.Sending()
}

type Get struct {
	receiver Receiver
}

func (g *Get) Execute() string {
	return g.receiver.Recieving()
}

type Invoker struct {
	command Command
}

func (i *Invoker) ExecuteCommand() string {
	return i.command.Execute()
}

type Receiver struct{}

func (r *Receiver) Sending() string {
	return "Send sig to service"
}

func (r *Receiver) Recieving() string {
	return "Get sig from service"
}

func Usage4() {
	receiver := &Receiver{}

	concreteCommand1 := &Send{receiver: *receiver}
	invoker1 := &Invoker{command: concreteCommand1}

	concreteCommand2 := &Get{receiver: *receiver}
	invoker2 := &Invoker{command: concreteCommand2}

	result1 := invoker1.ExecuteCommand()
	result2 := invoker2.ExecuteCommand()

	fmt.Println(result1)
	fmt.Println(result2)
}
