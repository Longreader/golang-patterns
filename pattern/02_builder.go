package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Шаблон строитель позволяет определять разные конфигурации одного обьекта с условием остутсвия незавершенной постройки.
	Это означает что обьект не может быть "недостроеным"
	Плюсы:
	- Продоставляет осуществлять управление процессом постройки
	- Позволяет определять вариации одного обьекта
	Минусы:
	- Доступность классов
	- Констуктор каждого продукта
	- Затрудненное внедрение зависимостей

	Пример: Запуск сервиса
*/

type Setup interface {
	setDB()
	setLogs()
	setBroker()
	setServer()
	getController() controller
}

func getSetup(setupType string) Setup {
	if setupType == "soft" {
		return getSoftStart()
	}
	if setupType == "prod" {
		return getProdStart()
	}
	return nil
}

type controller struct {
	db     string
	log    string
	broker string
	server string
}

type softStart struct {
	dbSetup     string
	logSetup    string
	brokerSetup string
	serverSetup string
}

func getSoftStart() *softStart {
	return &softStart{}
}

func (ss *softStart) setDB() {
	ss.dbSetup = "SQLite"
}

func (ss *softStart) setLogs() {
	ss.logSetup = "filename"
}

func (ss *softStart) setBroker() {
	ss.brokerSetup = "gRPC"
}

func (ss *softStart) setServer() {
	ss.serverSetup = "localhost"
}

func (ss *softStart) getController() controller {
	return controller{
		db:     ss.dbSetup,
		log:    ss.logSetup,
		broker: ss.brokerSetup,
		server: ss.serverSetup,
	}
}

type prodStart struct {
	dbSetup     string
	logSetup    string
	brokerSetup string
	serverSetup string
}

func getProdStart() *prodStart {
	return &prodStart{}
}

func (ps *prodStart) setDB() {
	ps.dbSetup = "PostgreSQL"
}

func (ps *prodStart) setLogs() {
	ps.logSetup = "Log to server"
}

func (ps *prodStart) setBroker() {
	ps.brokerSetup = "NATS"
}

func (ps *prodStart) setServer() {
	ps.serverSetup = "host"
}

func (ps *prodStart) getController() controller {
	return controller{
		db:     ps.dbSetup,
		log:    ps.logSetup,
		broker: ps.brokerSetup,
		server: ps.serverSetup,
	}
}

type appHandler struct {
	setup Setup
}

func NewAppHandler(s Setup) *appHandler {
	return &appHandler{
		setup: s,
	}
}

func (ah *appHandler) ConfigApp() controller {
	ah.setup.setDB()
	ah.setup.setLogs()
	ah.setup.setBroker()
	ah.setup.setServer()
	return ah.setup.getController()
}

func Usage2() {
	ss := getSetup("prod")

	handler := NewAppHandler(ss)
	app := handler.ConfigApp()

	fmt.Printf("Database configured: %s\n", app.db)
	fmt.Printf("Server host configured: %s\n", app.server)
	fmt.Printf("Inner brocker : %s\n", app.broker)

}
