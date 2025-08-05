// Package main implementa un servidor HTTP que procesa trabajos de cálculo de Fibonacci
// utilizando un patrón Worker Pool para manejar la concurrencia de manera eficiente.
//
// El servidor expone un endpoint HTTP que acepta trabajos para calcular números de Fibonacci
// y los procesa de forma asíncrona usando múltiples workers.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Job representa un trabajo que debe ser procesado por un worker.
// Contiene el número para calcular su Fibonacci, un nombre identificativo
// y un delay para simular procesamiento.
type Job struct {
	Number int           // Número para el cual calcular el Fibonacci
	Name   string        // Nombre identificativo del trabajo
	Delay  time.Duration // Tiempo de espera para simular procesamiento
}

// Worker representa un trabajador que procesa jobs de forma concurrente.
// Cada worker tiene su propio canal de trabajos y se comunica con el dispatcher
// a través del WorkerPool para recibir trabajos y reportar su disponibilidad.
type Worker struct {
	Id         int           // Identificador único del worker
	JobQueue   chan Job      // Canal para recibir trabajos específicos de este worker
	WorkerPool chan chan Job // Canal compartido para reportar disponibilidad al pool
	QuitChan   chan bool     // Canal para recibir señales de parada
}

// NewWorker crea una nueva instancia de Worker con el ID especificado.
// El worker se registra automáticamente en el pool de workers proporcionado.
//
// Parámetros:
//   - id: Identificador único para el worker
//   - workerPool: Canal compartido donde el worker reportará su disponibilidad
//
// Retorna:
//   - *Worker: Nueva instancia de worker configurada
func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		WorkerPool: workerPool,
		JobQueue:   make(chan Job),
		QuitChan:   make(chan bool),
	}
}

// Start inicia el worker en una goroutine separada.
// El worker entra en un bucle donde se registra en el pool de workers
// y espera a recibir trabajos o señales de parada.
//
// El método no bloquea y el worker continuará ejecutándose hasta
// que reciba una señal de parada a través de QuitChan.
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue // Registra el canal de trabajo del trabajador en el pool.
			select {
			case job := <-w.JobQueue: // Espera a recibir un trabajo del canal de trabajo.
				fmt.Printf("👷 Worker %d received job: %s with number: %d\n", w.Id, job.Name, job.Number)
				fib := Fibonacci(job.Number) // Calcula el número de Fibonacci para el trabajo recibido.
				time.Sleep(job.Delay)        // Simula el procesamiento del trabajo con un retraso.
				fmt.Printf("✅ Worker %d processed job: %s with number: %d → Fibonacci: %d\n", w.Id, job.Name, job.Number, fib)
			case <-w.QuitChan: // Escucha si se recibe una señal para detener el trabajador.
				fmt.Printf("🛑 Worker %d is stopping.\n", w.Id)
				return // Sale de la goroutine y detiene el trabajador.
			}
		}
	}()
}

// Stop envía una señal de parada al worker.
// El método no bloquea y el worker se detendrá de forma asíncrona.
// También cierra el canal de trabajos del worker para liberar recursos.
func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true // Envía una señal para detener el trabajador.
		close(w.JobQueue)  // Cierra el canal de trabajo del trabajador.
		fmt.Printf("🛑 Worker %d has stopped.\n", w.Id)
	}()
}

// Dispatcher gestiona un pool de workers y distribuye trabajos entre ellos.
// Actúa como coordinador central que recibe trabajos del JobQueue
// y los asigna a workers disponibles a través del WorkerPool.
type Dispatcher struct {
	MaxWorkers int           // Número máximo de workers en el pool
	WorkerPool chan chan Job // Canal para comunicación con workers disponibles
	JobQueue   chan Job      // Canal para recibir trabajos a procesar
}

// NewDispatcher crea una nueva instancia de Dispatcher.
//
// Parámetros:
//   - jobQueue: Canal donde se recibirán los trabajos a procesar
//   - maxWorkers: Número máximo de workers que manejará el dispatcher
//
// Retorna:
//   - *Dispatcher: Nueva instancia de dispatcher configurada
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: make(chan chan Job, maxWorkers),
	}
}

// Dispatch procesa trabajos del JobQueue y los distribuye a workers disponibles.
// Este método bloquea y debe ejecutarse en una goroutine separada.
// Continúa procesando trabajos hasta que se cierre el JobQueue.
func (d *Dispatcher) Dispatch() {
	for job := range d.JobQueue { // Itera sobre los trabajos recibidos en el canal de trabajos.
		go func(job Job) {
			workerJobQueue := <-d.WorkerPool // Obtiene un canal de trabajo de un trabajador disponible.
			workerJobQueue <- job            // Envía el trabajo al canal del trabajador.
		}(job)
	}
}

// Run inicializa y pone en funcionamiento el dispatcher.
// Crea el número especificado de workers, los inicia y comienza
// a despachar trabajos. Este método no bloquea.
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool) // Crea un nuevo trabajador.
		worker.Start()                       // Inicia el trabajador.
	}
	go d.Dispatch() // Comienza a despachar trabajos a los trabajadores.
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// RequestHandler maneja las solicitudes HTTP para crear trabajos de Fibonacci.
// Acepta solicitudes POST con parámetros de formulario y crea trabajos
// que se envían al canal de trabajos para ser procesados por los workers.
//
// Parámetros esperados en la solicitud:
//   - delay: Duración del delay de procesamiento (ej: "2s", "500ms")
//   - value: Número entero para calcular su Fibonacci
//   - name: Nombre identificativo del trabajo
func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid delay parameter", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid value parameter", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	job := Job{
		Number: value,
		Name:   name,
		Delay:  delay,
	}
	jobQueue <- job

	w.WriteHeader(http.StatusCreated)
}

// El servidor:
//   - Crea un pool de 4 workers
//   - Configura una cola de trabajos con capacidad para 20 trabajos
//   - Inicia un servidor HTTP en el puerto 8081
//   - Expone el endpoint POST /fibonacci para recibir trabajos
func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)

	jobQueue := make(chan Job, maxQueueSize) // Canal para recibir trabajos.

	dispatcher := NewDispatcher(jobQueue, maxWorkers) // Crea un despachador con el canal de trabajos y el número máximo de trabajadores.
	dispatcher.Run()                                  // Inicia el despachador.

	fmt.Println("🚀 Starting server on port", port)
	http.HandleFunc("/fibonacci", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue) // Maneja las solicitudes HTTP para crear trabajos.
	})

	// Inicia el servidor HTTP y registra cualquier error fatal
	log.Fatal(http.ListenAndServe(port, nil))
}
