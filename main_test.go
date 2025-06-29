package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements_EmptySlice(t *testing.T) {
	// Arange
	// Acc
	result := generateRandomElements(0)
	// Assert
	assert.Equal(t, 0, len(result), "Ожидаем пустой слайс при нулевом размере")
}

func TestGenerateRandomElements_SmallSize(t *testing.T) {
	// Arange
	size := MIN_PARALLEL_SIZE - 1
	// Act
	result := generateRandomElements(size)
	// Assert
	assert.Equal(t, size, len(result), "Ожидаем слайс длиной %d", size)
	for _, value := range result {
		assert.NotEmpty(t, value, "Ожидаем непустое значение")
	}
}

func TestGenerateRandomElements_BigSize(t *testing.T) {
	// Arange
	size := MIN_PARALLEL_SIZE + 1
	// Act
	result := generateRandomElements(size)
	// Assert
	assert.Equal(t, size, len(result), "Ожидаем слайс длиной %d", size)
	for _, value := range result {
		assert.NotEmpty(t, value, "Ожидаем непустое значение")
	}
}
