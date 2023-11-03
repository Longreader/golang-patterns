package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Паттер позволяет обработать запрос нескольким получателям.

	Плюсы:
	- Отделяет получателя запроса и запрос
	- Сокращает количество ошибок, т.к. позволяет определить лишь единый вызов цепочки
	Минусы:
	- Нет возможность тонкой настройки (обхода получателей по собственному маршруту)
*/

type Service interface {
	Kill(*killSig)
	SetNext(Service)
}

type killSig struct {
	sig string
}

type Service1 struct {
	status bool
	next   Service
}

func (s *Service1) Kill(k *killSig) {
	if s.status {
		fmt.Println("Killing service1 with", k.sig)
		s.status = false
		s.next.Kill(k)
	} else {
		s.next.Kill(k)
	}
}

func (s *Service1) SetNext(serv Service) {
	s.next = serv
}

type Service2 struct {
	status bool
	next   Service
}

func (s *Service2) Kill(k *killSig) {
	if s.status {
		fmt.Println("Killing service 2 with", k.sig)
		s.status = false
		s.next.Kill(k)
	} else {
		s.next.Kill(k)
	}
}

func (s *Service2) SetNext(serv Service) {
	s.next = serv
}

type Service3 struct {
	status bool
	next   Service
}

func (s *Service3) Kill(k *killSig) {
	if s.status {
		fmt.Println("Killing service 3 with", k.sig)
		s.status = false
	}
	fmt.Println("DONE ")
}

func (s *Service3) SetNext(serv Service) {
	s.next = nil
}

func Usage5() {

	killSignal := &killSig{
		sig: "SIGQUIT",
	}

	serv3 := &Service3{status: true}
	serv2 := &Service2{status: true}
	serv2.SetNext(serv3)
	serv1 := &Service1{status: true}
	serv1.SetNext(serv2)
	serv1.Kill(killSignal)
}
