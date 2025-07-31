// Este programa demuestra el uso de WaitGroup para manejar concurrencia en Go.
// Se lanzan 10 goroutines que simulan la ejecución de tareas en paralelo.
// El programa espera a que todas las tareas finalicen antes de terminar.
package main

import (
	"fmt"
	"sync"
	"time"
)

// main es la función principal del programa.
// Crea un WaitGroup y lanza 10 goroutines, cada una ejecutando doSomething.
// Utiliza wg.Wait() para esperar a que todas las goroutines terminen.
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)              // Incrementa el contador del WaitGroup por cada tarea.
		go doSomething(i, &wg) // Lanza una goroutine para ejecutar la tarea.
	}

	wg.Wait() // Espera a que todas las goroutines llamen a Done().

	fmt.Println("Todas las tareas han finalizado. 👍")
}

// doSomething simula una tarea que toma 2 segundos en completarse.
// Al finalizar, llama a wg.Done() para indicar que la goroutine terminó.
func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Marca la tarea como finalizada al salir de la función.
	fmt.Printf("🚀 Doing something: %d...\n", i)
	time.Sleep(2 * time.Second) // Simula trabajo con una pausa de 2 segundos.
	fmt.Printf("✅ Done %d!\n", i)
}
