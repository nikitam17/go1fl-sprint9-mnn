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
		assert.Len(t, slice, v.want)
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

// test maximum returns the maximum number of elements.
func TestMaxChunks(t *testing.T) {
	requests := []struct {
		want   int // ожидаемый max
		slices []int
	}{
		{0, nil},                                   // len = 0
		{10, []int{10}},                            // len = 1
		{20, []int{5, 3, 7, 20, 9, 4, 6, 2}},       // len = 8
		{20, []int{3, 5, 3, 7, 20, 9, 4, 6, 2}},    // len = 9
		{20, []int{1, 3, 5, 3, 7, 20, 9, 4, 6, 2}}, // len = 10
		{20, []int{1, 1, 3, 3, 4, 5, 3, 2, 2, 7, 20, 9, 4, 6, 2, 4, 5}},    // len = 17
		{20, []int{1, 1, 3, 3, 4, 5, 3, 2, 2, 7, 20, 9, 4, 6, 2, 4, 5, 9}}, // len = 18
	}
	for _, v := range requests {
		assert.Equal(t, v.want, maxChunks(v.slices))
	}
}
