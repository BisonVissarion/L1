package main

import (
	"fmt"
)

// Определяем структуру Human
type Human struct {
	Name string
	Age  int
}

// Определяем метод для структуры Human
func (h *Human) SayHello() {
	fmt.Printf("Привет, меня зовут %s и мне %d лет.\n", h.Name, h.Age)
}

// Определяем структуру Action, встраивая в нее структуру Human
type Action struct {
	Human
	ActionDescription string
}

func main() {
	// Создаем объект структуры Action
	action := Action{
		Human: Human{
			Name: "Иван",
			Age:  30,
		},
		ActionDescription: "говорит",
	}

	// Вызываем метод SayHello, который был встроен из структуры Human
	action.SayHello()

	// Добавляем собственное действие
	fmt.Printf("%s %s: Привет миру!\n", action.Name, action.ActionDescription)
}
