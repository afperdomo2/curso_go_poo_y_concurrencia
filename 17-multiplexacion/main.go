// Este programa demuestra la multiplexación de canales en Go utilizando la sentencia select.
// Se ejecutan dos tareas concurrentes con diferentes duraciones y se procesan los resultados
// conforme van llegando, sin importar el orden de finalización.
// La multiplexación permite manejar múltiples canales de comunicación simultáneamente.
package main

import (
	"fmt"
	"time"
)

// doSomething simula una tarea que tarda un tiempo específico en completarse.
// Envía el parámetro al canal después de esperar la duración especificada.
func doSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i) // Simula trabajo con una pausa.
	c <- param    // Envía el resultado al canal.
}

// main es la función principal que demuestra la multiplexación con select.
// Lanza dos goroutines con diferentes duraciones y usa select para procesar
// los resultados conforme van llegando.
func main() {
	c1 := make(chan int) // Canal para la primera tarea.
	c2 := make(chan int) // Canal para la segunda tarea.

	d1 := 4 * time.Second // Duración de la primera tarea (4 segundos).
	d2 := 2 * time.Second // Duración de la segunda tarea (2 segundos).

	go doSomething(d1, c1, 1) // Lanza la goroutine para la tarea 1.
	go doSomething(d2, c2, 2) // Lanza la goroutine para la tarea 2.

	// Espera los resultados de ambas tareas usando multiplexación.
	for range 2 {
		select {
		case res := <-c1: // Espera el resultado de la tarea 1.
			fmt.Println("Tarea 1 completada con resultado:", res)
		case res := <-c2: // Espera el resultado de la tarea 2.
			fmt.Println("Tarea 2 completada con resultado:", res)
		}
	}
}
