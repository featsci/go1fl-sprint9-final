package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	s := make([]int, size)
	if size == 0 {
		return s
	}
	for i := range size {
		s[i] = rand.Intn(SIZE)
		// s[i] = 0
	}
	return s
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	// Не забудьте учесть все возможные крайние случаи: пустой слайс, слайс с одним элементом и так далее.
	// fmt.Println(len(data))
	if len(data) == 0 || len(data) == 1 {
		return 0
	}
	return slices.Max(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	// создайте переменную типа sync.WaitGroup и используйте её при запуске и ожидании горутин.

	if len(data) == 0 || len(data) == 1 {
		return 0
	}
	// var div [][]int

	for i := 0; i < len(data); i += CHUNKS {
		end := i + CHUNKS

		if end > len(data) {
			end = len(data)
		}

		// div = append(div, data[i:end])
		fmt.Println(slices.Max(data[i:end]))
	}

	// fmt.Printf("%#v\n", div)

	return 0
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	g := generateRandomElements(32)
	fmt.Println(g)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(g)
	elapsed := time.Since(start)

	// fmt.Println(m)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(g)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
