package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}
	s := make([]int, size)
	for i := range size {
		s[i] = rand.Int()
	}
	return s
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	return slices.Max(data)
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	var wg sync.WaitGroup
	subSlices := make([][]int, CHUNKS)
	sliceLen := len(data)
	subSliceSize := sliceLen / CHUNKS
	remainder := sliceLen % CHUNKS
	start := 0
	sL := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		end := start + subSliceSize
		if i < remainder {
			end++
		}
		subSlices[i] = data[start:end]
		start = end
		wg.Add(1)

		go func(i int, tS []int) {
			defer wg.Done()
			sL[i] = maximum(tS)
		}(i, subSlices[i])
	}
	wg.Wait()
	return maximum(sL)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	g := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(g)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(g)
	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
