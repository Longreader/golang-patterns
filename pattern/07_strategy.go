package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Патерн стратегия позволяет изменить внетреннее поведение алгоритма (в отличае от декоратора, который меняет внешнее окружение) по
	ходу исполнения программы


*/

type Saver interface {
	Save([]int)
}

type Storage struct {
	Saver Saver
}

func (s *Storage) Store(data []int) {
	s.Saver.Save(data)
}

func (s *Storage) setStorage(ss Saver) {
	s.Saver = ss
}

type FileStore struct{}

func (FileStore) Save(data []int) {
	fmt.Println("Saving data to file")
}

type SQLStore struct{}

func (SQLStore) Save(data []int) {
	fmt.Println("Saving data to sql client")
}

func Usage7() {
	storage1 := Storage{FileStore{}}
	storage1.Store(make([]int, 0))

	storage2 := Storage{SQLStore{}}
	storage2.Store(make([]int, 0))
	// Change behavior
	storage2.setStorage(FileStore{})
	storage2.Store(make([]int, 1))
}
