package main

import "fmt"

func main() {
	// Channel con un buffer de tamaño 3
	// Permite enviar hasta 3 valores sin bloquear
	// Luego, se pueden recibir esos valores en orden
	c := make(chan int, 3)

	c <- 25
	c <- 30
	c <- 35

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
