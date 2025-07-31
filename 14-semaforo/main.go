// Este programa demuestra cómo limitar la concurrencia de goroutines en Go utilizando un semáforo
// implementado con un canal. Se lanzan 10 tareas concurrentes, pero solo 3 pueden ejecutarse al mismo
// tiempo gracias al semáforo. El programa espera a que todas las tareas finalicen antes de terminar.
package main

import (
	"fmt"
	"sync"
	"time"
)

// main es la función principal del programa.
// Crea un semáforo (canal con capacidad 3) para limitar la concurrencia y un WaitGroup para esperar a que
// todas las tareas terminen.
func main() {
	semaphore := make(chan int, 3) // Canal con capacidad 3 que actúa como semáforo para limitar la concurrencia.
	var wg sync.WaitGroup          // WaitGroup para esperar a que todas las goroutines terminen.

	for i := range 10 {
		semaphore <- i                    // Ocupa un lugar en el semáforo antes de lanzar la goroutine.
		wg.Add(1)                         // Incrementa el contador del WaitGroup por cada tarea.
		go doSomething(i, &wg, semaphore) // Lanza una goroutine para ejecutar la tarea.
	}

	wg.Wait() // Espera a que todas las goroutines llamen a Done().
	fmt.Println("Todas las tareas han finalizado. 👍")
}

// doSomething simula una tarea que toma 2 segundos en completarse.
// Al finalizar, libera un lugar en el semáforo y marca la tarea como finalizada en el WaitGroup.
func doSomething(i int, wg *sync.WaitGroup, semaforo chan int) {
	defer wg.Done() // Marca la tarea como finalizada al salir de la función.

	fmt.Println("🚀 Doing something:", i)
	time.Sleep(2 * time.Second) // Simula trabajo con una pausa de 2 segundos.
	fmt.Println("✅ Done", i)

	<-semaforo // Libera un lugar en el semáforo para que otra goroutine pueda ejecutarse.
}
