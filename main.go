package main

import (
	"fmt"
	"math/rand"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	s := []int{} // make([]int, size)

	for i := range size {
		_ = i
		s = append(s, rand.Intn(SIZE))
	}

	// fmt.Println(s, rand.Intn(SIZE))
	return s
}

// // maximum returns the maximum number of elements.
// func maximum(data []int) int {
// 	// ваш код здесь
// }

// // maxChunks returns the maximum number of elements in a chunks.
// func maxChunks(data []int) int {
// 	// ваш код здесь
// }

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	generateRandomElements(3)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь

	// fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	// fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// // ваш код здесь

	// fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
