package main

import (
	"fmt"
)

func main() {
	q := NewQueue[string]()

	// Добавление элементов
	q.PushFront("1")
	q.PushBack("2")
	q.PushBack("3")

	// Вывод длины
	fmt.Println("Length:", q.Length()) // 3

	// Извлечение элементов
	frontItem, ok := q.PopFront()
	if ok {
		fmt.Println("Popped from front:", frontItem) // 1
	}

	backItem, ok := q.PopBack()
	if ok {
		fmt.Println("Popped from back:", backItem) // 3
	}

	// Вращение элементов
	rotatedItem, ok := q.RotateFrontToBack()
	if ok {
		fmt.Println("Rotated from front to back:", rotatedItem) // 2
	}

	// Повторное извлечение элементов для проверки вращения
	frontItem, ok = q.PopFront()
	if ok {
		fmt.Println("Popped from front after rotation:", frontItem) // 2
	}
}
