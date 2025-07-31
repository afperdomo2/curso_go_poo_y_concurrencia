// Este programa implementa un pipeline en Go utilizando canales y goroutines.
// El pipeline consta de tres etapas:
// 1. Generator: genera números del 1 al 10 y los envía por un canal.
// 2. Double: recibe los números, los multiplica por 2 y los envía a otro canal.
// 3. Printer: imprime los resultados finales recibidos del canal.
// Cada etapa se ejecuta en una goroutine diferente y se comunica usando canales.
package main

import "fmt"

// Generator genera números del 1 al 10 y los envía al canal c.
// Al finalizar, cierra el canal para indicar que no habrá más datos.
func Generator(c chan<- int) {
	for i := range 10 {
		c <- i + 1
	}
	close(c) // Cierra el canal después de enviar todos los valores.
}

// Double recibe valores del canal in, los multiplica por 2 y los envía al canal out.
// Al finalizar, cierra el canal out para indicar que no habrá más datos.
func Double(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * 2
	}
	close(out) // Cierra el canal out después de enviar todos los valores.
}

// Printer recibe valores del canal c y los imprime por pantalla.
func Printer(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

// main es la función principal del programa.
// Crea los canales y lanza las goroutines para cada etapa del pipeline.
func main() {
	numbers := make(chan int) // Canal para enviar números.
	doubles := make(chan int) // Canal para enviar números doblados.

	go Generator(numbers)       // Lanza la goroutine para generar números.
	go Double(numbers, doubles) // Lanza la goroutine para doblar los números.
	Printer(doubles)            // Llama a la función Printer para imprimir los números doblados.

	fmt.Println("Todas las tareas han finalizado. 👍")
}
