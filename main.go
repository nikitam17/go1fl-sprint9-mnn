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
		return nil
	}
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Int()
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	l := len(data)
	if l < CHUNKS {
		return maximum(data)
	}
	chunks := make([]int, CHUNKS)
	offSet := l % CHUNKS
	maxCh := l / CHUNKS
	var wg sync.WaitGroup

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		index := i * maxCh
		if offSet != 0 && i == CHUNKS-1 {
			// в последний chunk добавляем оставшиеся элементы слайса
			maxCh += offSet
		}
		chunk := data[index : index+maxCh]
		go func(i int, ch []int) {
			defer wg.Done()
			chunks[i] = maximum(ch)
		}(i, chunk)
	}
	wg.Wait()

	return maximum(chunks)
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
