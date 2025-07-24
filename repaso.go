package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. Declaración de variables
	// Declaración e inicialización de variables
	var x int
	x = 8

	// Declaración e inicialización de otra variable
	y := 7

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// 2. Captura de errores
	newNumber, err := strconv.ParseInt("NaN", 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed number:", newNumber)
	}

	// 3. Uso de Maps
	phoneDirectory := make(map[string]int)
	phoneDirectory["Alice"] = 123456789
	phoneDirectory["Bob"] = 987654321
	fmt.Println("Phone Directory:", phoneDirectory)

	// 4. Uso de Slices (números primos)
	primes := []int{2, 3, 5, 7, 11, 13}

	primes = append(primes, 17, 19) // Agregar más números primos

	for i, prime := range primes {
		fmt.Printf("Prime number at index %d: %d\n", i, prime)
	}
}
