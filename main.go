package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

const (
	SIZE   = 1000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	result := make([]int, size)
	if size == 0 {
		return result
	}
	workers := runtime.NumCPU() // Количество доступных CPU
	if (size < 100) || (workers == 1) {
		for i := 0; i < size; i++ {
			result[i] = rand.Int()
		}
	} else {
		// генерим значения в несколько потоков
		var wg sync.WaitGroup
		wg.Add(workers)
		for w := 0; w < workers; w++ {
			go func(worker int) {
				defer wg.Done()
				// Каждый worker обрабатывает свою часть слайса
				start := size / workers * worker
				end := size / workers * (worker + 1)
				if worker == workers-1 { // последний
					end = size
				}
				fmt.Printf("Генерим с %d по %d элементы\n", start, end)
				for i := start; i < end; i++ {
					result[i] = rand.Int()
				}
			}(w)
		}
		wg.Wait()
	}
	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	return len(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	return len(data)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	elements := generateRandomElements(SIZE)
	fmt.Println(elements)

	fmt.Println("Ищем максимальное значение в один поток")
	fmt.Println(maximum(elements))
	fmt.Println(maxChunks(elements))
	// ваш код здесь

	//fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	//fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь

	//fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
