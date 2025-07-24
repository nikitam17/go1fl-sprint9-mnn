package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size < 0 {
		size = 0
	}
	slice := make([]int, size)
	if size <= 0 {
		return slice
	}
	max := size * 10
	for i := range slice {
		slice[i] = rand.Intn(max)
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	max := 0
	for i := range data {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	chunks := make([]int, CHUNKS)
	maxChunks := SIZE / CHUNKS

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		index := i * maxChunks
		chunk := data[index : index+maxChunks]
		go func(ch []int) {
			defer wg.Done()
			m := maximum(ch)
			mu.Lock()
			chunks = append(chunks, m)
			mu.Unlock()
		}(chunk)
	}
	wg.Wait()

	max := maximum(chunks)
	return max
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(slice)
	finish := time.Since(start)
	elapsed := finish.Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(slice)
	finish = time.Since(start)
	elapsed = finish.Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
