package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Función anónima simple
	x := 5
	y := func() int {
		return x * 2
	}() // Llamada a la función anónima
	fmt.Println("El resultado es:", y)

	// 2. Función anónima con parámetros
	suma := func(a int, b int) int {
		return a + b
	}
	resultadoSuma := suma(3, 4)         // 7
	resultado := suma(resultadoSuma, 6) // 13
	fmt.Println("La suma es:", resultado)

	// 3. Goroutines con funciones anónimas
	c := make(chan string)
	go func() {
		fmt.Println("Iniciando...")
		time.Sleep(2 * time.Second)
		c <- "Trabajo completado"
	}()
	fmt.Println(<-c)
}
