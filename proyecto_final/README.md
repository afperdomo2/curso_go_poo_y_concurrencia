# 🧮 Servidor de Cálculo de Fibonacci con Worker Pool

Este proyecto implementa un servidor HTTP concurrente en Go que procesa trabajos de cálculo de números de Fibonacci utilizando el patrón **Worker Pool**. Es el proyecto final del curso de Go POO y Concurrencia.

## 📋 Descripción

El servidor expone un endpoint HTTP que acepta solicitudes para calcular números de Fibonacci. Los trabajos se procesan de forma asíncrona utilizando múltiples workers que trabajan en paralelo, demostrando conceptos avanzados de concurrencia en Go como:

- 🔄 **Goroutines y Channels**
- 👷 **Worker Pool Pattern**
- 🚀 **Procesamiento Asíncrono**
- 🌐 **Servidor HTTP Concurrente**
- ⚡ **Gestión de Recursos y Load Balancing**

## 🏗️ Arquitectura del Sistema

### Componentes Principales

#### 1. **Job** 📦

Estructura que representa un trabajo a procesar:

```go
type Job struct {
    Number int           // Número para calcular Fibonacci
    Name   string        // Identificador del trabajo
    Delay  time.Duration // Tiempo de procesamiento simulado
}
```

#### 2. **Worker** 👷

Trabajador que procesa jobs de forma concurrente:

- Cada worker tiene su propio canal de trabajos
- Se registra en el pool cuando está disponible
- Procesa trabajos de forma independiente
- Puede ser detenido de forma controlada

#### 3. **Dispatcher** 🎯

Coordinador central que:

- Gestiona el pool de workers
- Distribuye trabajos entre workers disponibles
- Mantiene la comunicación entre componentes

#### 4. **RequestHandler** 🌐

Manejador HTTP que:

- Valida parámetros de entrada
- Crea trabajos a partir de solicitudes HTTP
- Envía trabajos al canal de procesamiento

### Flujo de Trabajo

```text
[HTTP Request] → [RequestHandler] → [JobQueue] → [Dispatcher] → [Available Worker] → [Fibonacci Calculation] → [Console Output]
```

1. **Recepción**: El servidor recibe una solicitud HTTP POST
2. **Validación**: Se validan los parámetros (delay, value, name)
3. **Creación**: Se crea un Job con los parámetros
4. **Encolado**: El Job se envía al JobQueue
5. **Distribución**: El Dispatcher asigna el trabajo a un Worker disponible
6. **Procesamiento**: El Worker calcula el Fibonacci y simula procesamiento
7. **Resultado**: Se muestra el resultado en la consola del servidor

## 🚀 Uso del Sistema

### Iniciar el Servidor

```bash
go run main.go
```

El servidor se iniciará en el puerto `8081` con:

- ✅ 4 workers activos
- ✅ Cola de trabajos con capacidad para 20 jobs
- ✅ Endpoint `/fibonacci` disponible

### Enviar Trabajos

**Endpoint:** `POST http://localhost:8081/fibonacci`

**Parámetros (form-data):**

- `name`: Nombre identificativo del trabajo (requerido)
- `value`: Número para calcular Fibonacci (requerido, entero)
- `delay`: Tiempo de procesamiento simulado (requerido, formato: "2s", "500ms", etc.)

### Ejemplos de Uso

#### Con curl

```bash
# Ejemplo básico
curl -X POST http://localhost:8081/fibonacci \
  -d "name=test1" \
  -d "value=10" \
  -d "delay=2s"

# Trabajo más complejo
curl -X POST http://localhost:8081/fibonacci \
  -d "name=fibonacci_35" \
  -d "value=35" \
  -d "delay=1s"

# Múltiples trabajos rápidos
curl -X POST http://localhost:8081/fibonacci -d "name=job1&value=20&delay=500ms"
curl -X POST http://localhost:8081/fibonacci -d "name=job2&value=25&delay=1s"
curl -X POST http://localhost:8081/fibonacci -d "name=job3&value=30&delay=800ms"
```

#### Con Postman

1. Método: `POST`
2. URL: `http://localhost:8081/fibonacci`
3. Body: `form-data`
   - `name`: `mi_trabajo`
   - `value`: `20`
   - `delay`: `1s`

## 📊 Ejemplo de Salida

```text
🚀 Starting server on port :8081
👷 Worker 0 received job: test1 with number: 10
👷 Worker 1 received job: fibonacci_35 with number: 35
✅ Worker 0 processed job: test1 with number: 10 → Fibonacci: 55
👷 Worker 2 received job: job1 with number: 20
✅ Worker 1 processed job: fibonacci_35 with number: 35 → Fibonacci: 9227465
✅ Worker 2 processed job: job1 with number: 20 → Fibonacci: 6765
```

## ⚙️ Configuración

### Parámetros Configurables (en `main.go`)

```go
const (
    maxWorkers   = 4     // Número de workers concurrentes
    maxQueueSize = 20    // Capacidad máxima de la cola de trabajos
    port         = ":8081" // Puerto del servidor
)
```

### Personalización

- **Más Workers**: Aumenta `maxWorkers` para mayor paralelismo
- **Cola Mayor**: Incrementa `maxQueueSize` para manejar más trabajos simultáneos
- **Puerto Diferente**: Cambia `port` según necesidades

## 🧠 Conceptos Demostrados

### 1. **Concurrencia con Goroutines**

- Cada worker ejecuta en su propia goroutine
- Procesamiento simultáneo de múltiples trabajos
- Comunicación segura mediante channels

### 2. **Worker Pool Pattern**

- Pool fijo de workers reutilizables
- Distribución eficiente de carga de trabajo
- Prevención de creación excesiva de goroutines

### 3. **Channel Communication**

- Comunicación tipo-segura entre goroutines
- Sincronización sin locks explícitos
- Patrón productor-consumidor

### 4. **Gestión de Recursos**

- Control de concurrencia limitando workers
- Cola con buffer para evitar bloqueos
- Parada controlada de workers

### 5. **HTTP Server Concurrente**

- Manejo simultáneo de múltiples requests
- Procesamiento no-bloqueante de trabajos
- Integración web con backend concurrente

## 🔧 Dependencias

Este proyecto utiliza únicamente la biblioteca estándar de Go:

- `fmt` - Formateo y salida
- `log` - Logging de errores
- `net/http` - Servidor HTTP
- `strconv` - Conversión de strings
- `time` - Manejo de tiempo y duraciones

**No requiere dependencias externas** 🎉

## 📝 Notas Técnicas

### Implementación de Fibonacci

- Utiliza recursión simple (intencionalmente ineficiente)
- Demuestra procesamiento computacionalmente intensivo
- Para valores grandes (>40) puede ser muy lento

### Consideraciones de Performance

- Worker pool previene sobrecarga del sistema
- Canal con buffer evita bloqueos en alta carga
- Cada worker procesa trabajos secuencialmente

### Casos de Error Manejados

- ❌ Método HTTP incorrecto (solo POST)
- ❌ Parámetros faltantes o inválidos
- ❌ Formato de duración incorrecto
- ❌ Valores no numéricos

## 🎯 Objetivos de Aprendizaje Alcanzados

- ✅ Implementación práctica de concurrencia en Go
- ✅ Uso avanzado de channels y goroutines
- ✅ Patrón Worker Pool para gestión de recursos
- ✅ Integración de concurrencia con servicios web
- ✅ Manejo de errores y validación de entrada
- ✅ Arquitectura escalable y mantenible

---

**Proyecto desarrollado como parte del Curso de Go POO y Concurrencia** 🚀
