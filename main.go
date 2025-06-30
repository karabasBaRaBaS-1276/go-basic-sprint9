package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	SIZE              = 100_000_000
	CHUNKS            = 8
	MIN_PARALLEL_SIZE = 1000 // Минимальный размер для параллельной обработки
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	result := make([]int, size)
	if size == 0 {
		return result
	}
	workers := runtime.NumCPU() // Количество доступных CPU
	//workers := CHUNKS
	if (size < MIN_PARALLEL_SIZE) || (workers == 1) {
		log.Printf("Генерим элементы в один поток\n")
		for i := 0; i < size; i++ {
			result[i] = rand.Int()
		}
	} else {
		// генерим значения в несколько потоков
		log.Printf("Кол-во потоков для генерации: %d\n", workers)
		var wg sync.WaitGroup
		wg.Add(workers)
		for w := 0; w < workers; w++ {
			go func(worker int, data []int) {
				defer wg.Done()
				// Каждый worker обрабатывает свою часть слайса
				start := size / workers * worker
				end := size / workers * (worker + 1)
				if worker == workers-1 { // последний
					end = size
				}
				log.Printf("Генерим с %d по %d элементы\n", start, end)
				for i := start; i < end; i++ {
					data[i] = rand.Int()
				}
			}(w, result)
		}
		wg.Wait()
	}
	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	result := data[0]
	for i := 1; i < len(data); i++ {
		if result < data[i] {
			result = data[i]
		}
	}

	return result
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	workers := CHUNKS
	size := len(data)
	chunkMaxValues := make([]int, CHUNKS)

	// --- запускаем горутину, которая будет из канала вычитывать
	//     значение и складывать в слайс для окончательной обработки
	ch := make(chan int, 1) // буфер = 1

	var wg0 sync.WaitGroup
	wg0.Add(1)
	go func(chunkMaxValues []int, ch <-chan int) {
		defer wg0.Done()
		i := 0
		for maxValue := range ch {
			chunkMaxValues[i] = maxValue
			i++
		}
	}(chunkMaxValues, ch)

	// --- Запускаем горутины для поиска максимума в составных частях
	var wg1 sync.WaitGroup
	wg1.Add(workers)
	for w := 0; w < workers; w++ {
		start := size / workers * w
		end := size / workers * (w + 1)
		if w == workers-1 { // последний
			end = size
		}

		go func(dataChunck []int, ch chan<- int) {
			defer wg1.Done()
			ch <- maximum(dataChunck)
		}(data[start:end], ch)
	}
	wg1.Wait() // Ждем окончания работы составных частей
	close(ch)  // закрываем канал
	wg0.Wait() // Ждем слайс для финальной обработки

	return maximum(chunkMaxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	elements := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	timeStart := time.Now()
	max := maximum(elements)
	elapsed := time.Since(timeStart).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	timeStart = time.Now()
	max = maxChunks(elements)
	elapsed = time.Since(timeStart).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
