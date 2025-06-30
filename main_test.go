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

func TestMaximum_Success(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Пустой слайс",
			args: args{
				data: []int{},
			},
			want: 0,
		},
		{
			name: "Положительное число в середине",
			args: args{
				data: []int{0, 1060, 45},
			},
			want: 1060,
		},
		{
			name: "Число 23 в конце",
			args: args{
				data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
			},
			want: 23,
		},
		{
			name: "Число 100 вначале",
			args: args{
				data: []int{100, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.args.data); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxChunks_Success(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Пустой слайс",
			args: args{
				data: []int{},
			},
			want: 0,
		},
		{
			name: "Положительное число в середине",
			args: args{
				data: []int{0, 1060, 45},
			},
			want: 1060,
		},
		{
			name: "Число 23 в конце",
			args: args{
				data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
			},
			want: 23,
		},
		{
			name: "Число 100 вначале",
			args: args{
				data: []int{100, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.args.data); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
