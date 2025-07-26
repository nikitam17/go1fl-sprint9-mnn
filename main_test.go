package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле

// test generateRandomElements generates random elements.
func TestGenerateRandomElements(t *testing.T) {
	requests := []struct {
		count int // передаваемое значение
		want  int // ожидаемое количество в ответе
	}{
		{-1, 0},
		{0, 0},
		{1, 1},
		{10, 10},
	}
	for _, v := range requests {
		slice := generateRandomElements(v.count)
		assert.Equal(t, v.want, len(slice))
	}
}

// test maximum returns the maximum number of elements.
func TestMaximum(t *testing.T) {
	requests := []struct {
		want   int // ожидаемый max
		slices []int
	}{
		{0, nil},
		{10, []int{10}},
		{20, []int{1, 3, 5, 3, 7, 20, 9, 4, 6, 2}},
	}
	for _, v := range requests {
		assert.Equal(t, v.want, maximum(v.slices))
	}
}
