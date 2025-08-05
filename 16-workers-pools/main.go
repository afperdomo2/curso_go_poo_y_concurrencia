// Este programa demuestra el patrón Worker Pool en Go utilizando canales y goroutines.
// Se crean múltiples workers (trabajadores) que procesan tareas concurrentemente desde un canal compartido.
// Cada worker calcula la secuencia de Fibonacci del número asignado y envía el resultado a un canal de resultados.
// Este patrón es útil para limitar la concurrencia y distribuir trabajo entre un número fijo de goroutines.
package main

import (
	"fmt"
)

// Worker representa un trabajador que procesa tareas del canal jobs.
// Cada worker tiene un ID único, lee trabajos del canal jobs y envía resultados al canal results.
func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("👷 Worker %d: 🚀 started job %d\n", id, job)
		// time.Sleep(1 * time.Second)
		fib := Fibonacci(job) // Calcula el número de Fibonacci para el trabajo recibido.
		fmt.Printf("👷 Worker %d: ✅ finished job %d. FIB=%d\n", id, job, fib)
		results <- fib // Envía el resultado al canal de resultados.
	}
}

// Fibonacci calcula el n-ésimo número de la secuencia de Fibonacci de forma recursiva.
// Esta función se usa como ejemplo de una tarea computacionalmente intensiva para los workers.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2) // Recursión para calcular Fibonacci.
}

// main es la función principal que implementa el patrón Worker Pool.
// Crea workers, distribuye tareas entre ellos y recolecta los resultados.
func main() {
	tasks := []int{2, 34, 36, 17, 10, 40, 6, 22, 41, 44, 33, 2, 5, 2, 34, 36, 17, 10, 40, 6, 22, 41, 44, 33, 2, 5, 17, 10, 40, 6, 22, 2, 5, 2, 34, 36, 22, 41, 44}
	numWorkers := 6

	jobs := make(chan int, len(tasks))    // Canal para enviar trabajos a los trabajadores.
	results := make(chan int, len(tasks)) // Canal para recibir resultados de los trabajadores.

	// 1. Lanzar trabajadores.
	fmt.Println("1️⃣ Iniciando trabajadores...")
	for i := range numWorkers {
		go Worker(i, jobs, results)
	}

	// 2. Enviar trabajos a los trabajadores.
	fmt.Println("2️⃣ Enviando trabajos a los trabajadores...")
	for _, task := range tasks {
		jobs <- task // Envía cada tarea al canal de trabajos.
	}
	close(jobs) // Cierra el canal de trabajos para indicar que no hay más tareas.

	// 3. Recibir resultados de los trabajadores.
	fmt.Println("3️⃣ Recibiendo resultados de los trabajadores...")
	for range tasks {
		<-results
	}
	close(results) // Cierra el canal de resultados después de recibir todos los resultados.

	fmt.Println("Todas las tareas han finalizado. 👍")
	fmt.Println("Fin del programa.")
}
