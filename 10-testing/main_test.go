package main

import "testing"

func TestSum(t *testing.T) {
	// 1. Test básico
	// Verifica que la función Sum suma correctamente dos números.
	total := Sum(10, 5)
	if total != 15 {
		t.Errorf("Sum(10, 5) = %d; want 15", total)
	}

	// 2. Test con casos de prueba
	// Utiliza una tabla de casos de prueba para verificar diferentes entradas y salidas.
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 2, want: 3},
		{a: 2, b: 3, want: 5},
		{a: 3, b: 4, want: 7},
		{a: -1, b: 1, want: 0},
		{a: -5, b: -5, want: -10},
	}
	for _, tc := range testCases {
		t.Run("Sum", func(t *testing.T) {
			if got := Sum(tc.a, tc.b); got != tc.want {
				t.Errorf("Sum(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestGetMax(t *testing.T) {
	// 1. Test con casos de prueba
	// Utiliza una tabla de casos de prueba para verificar diferentes entradas y salidas.
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 2, want: 2},
		// {a: 3, b: 2, want: 3},
		// {a: -1, b: -2, want: -1},
		// {a: -5, b: -5, want: -5},
	}
	for _, tc := range testCases {
		t.Run("GetMax", func(t *testing.T) {
			if got := GetMax(tc.a, tc.b); got != tc.want {
				t.Errorf("GetMax(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestGetMin(t *testing.T) {
	// 1. Test con casos de prueba
	// Utiliza una tabla de casos de prueba para verificar diferentes entradas y salidas.
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 2, want: 1},
		{a: 3, b: 2, want: 2},
		{a: -1, b: -2, want: -2},
		{a: -5, b: -5, want: -5},
	}
	for _, tc := range testCases {
		t.Run("GetMin", func(t *testing.T) {
			if got := GetMin(tc.a, tc.b); got != tc.want {
				t.Errorf("GetMin(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
