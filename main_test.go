package main

import (
	"testing"

	"github.com/stretchr/testify/require"
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
		require.Equal(t, v.want, len(slice))
	}
}

// test maximum returns the maximum number of elements.
func TestMaximum(t *testing.T) {
	requests := []struct {
		count int // передаваемое значение
		want  int // ожидаемое количество в ответе
	}{
		{0, 0},
		{1, 1},
		{10, 10},
	}
	for _, v := range requests {
		slice := generateRandomElements(v.count)
		max := maximum(slice)
		if v.want == 0 {
			require.Equal(t, v.want, max)
		} else if v.want == 1 {
			require.Equal(t, slice[0], max)
		}
	}
}
