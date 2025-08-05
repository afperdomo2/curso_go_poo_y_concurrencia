# 🚀 Curso de Go: Programación Orientada a Objetos y Concurrencia

Este repositorio contiene el material del curso completo de **Go (Golang)** enfocado en **Programación Orientada a Objetos** y **Concurrencia**. El curso está diseñado de manera progresiva, desde conceptos básicos hasta implementaciones avanzadas de patrones de concurrencia.

## 📋 Descripción General

El curso abarca desde los fundamentos de Go hasta conceptos avanzados de concurrencia, incluyendo:

- **Programación Orientada a Objetos** en Go
- **Patrones de Diseño** (Factory, Abstract Factory)
- **Concurrencia Avanzada** (Goroutines, Channels, Worker Pools)
- **Testing y Mocking** con herramientas nativas
- **Gestión de Módulos** y dependencias
- **Proyecto Final** completo con servidor HTTP concurrente

## 🏗️ Estructura del Proyecto

```sh
curso_go_poo_y_concurrencia/
├── 01-repaso/              # Fundamentos y repaso de Go
├── 02-class/               # Estructuras y métodos (Clases en Go)
├── 03-contructor/          # Patrones de construcción
├── 04-herencia/            # Composición (herencia en Go)
├── 05-interfaces/          # Interfaces y polimorfismo
├── 06-abstract-factory/    # Patrón Abstract Factory
├── 07-funciones-anonimas/  # Funciones anónimas y closures
├── 08-functiones-variadicas/ # Funciones con parámetros variables
├── 09-modulos/             # Gestión de módulos y dependencias
├── 10-testing/             # Testing, coverage y profiling
├── 11-mocking/             # Mocking para testing unitario
├── 12-channels/            # Channels y comunicación entre goroutines
├── 13-wait-group/          # Sincronización con WaitGroup
├── 14-semaforo/            # Control de concurrencia con semáforos
├── 15-pipelines/           # Pipelines concurrentes
├── 16-workers-pools/       # Patrón Worker Pool
├── 17-multiplexacion/      # Multiplexación con select
└── proyecto_final/         # Servidor HTTP con Worker Pool
```

## 📚 Contenido Detallado por Módulo

### 🔰 Módulos Básicos (1-5)

#### **01-repaso** - Fundamentos de Go

- Variables y tipos de datos
- Manejo de errores con `error`
- Maps y Slices
- Goroutines básicas y channels
- Punteros y referencias

#### **02-class** - Estructuras y Métodos

- Definición de `struct`
- Métodos con receivers
- Encapsulación de datos
- Getters y Setters

#### **03-contructor** - Patrones de Construcción

- Funciones constructoras
- Inicialización con `new()`
- Diferentes formas de crear instancias
- Mejores prácticas de inicialización

#### **04-herencia** - Composición en Go

- Embedded structs (composición)
- Simulación de herencia
- Acceso a campos embebidos
- Estructuras anidadas

#### **05-interfaces** - Interfaces y Polimorfismo

- Definición e implementación de interfaces
- Polimorfismo dinámico
- Interfaces implícitas
- Duck typing en Go

### 🏭 Módulos de Patrones (6-8)

#### **06-abstract-factory** - Patrón Abstract Factory

- Implementación del patrón Abstract Factory
- Familias de objetos relacionados
- Interfaces para factorías abstractas
- Polimorfismo con factorías

#### **07-funciones-anonimas** - Funciones Anónimas

- Definición y uso de funciones anónimas
- Closures y captura de variables
- Funciones como parámetros
- Goroutines con funciones anónimas

#### **08-functiones-variadicas** - Funciones Variádicas

- Parámetros variables con `...`
- Múltiples valores de retorno
- Named returns
- Desempaquetado de slices

### 📦 Módulos de Gestión (9-11)

#### **09-modulos** - Gestión de Módulos

- Inicialización con `go mod init`
- Importación de módulos externos
- Versionado de dependencias
- Comandos `go get`, `go mod tidy`

#### **10-testing** - Testing Avanzado

- Tests unitarios con `testing`
- Table-driven tests
- Coverage analysis con `go test -cover`
- CPU profiling con `go test -cpuprofile`
- Reportes HTML de coverage

#### **11-mocking** - Mocking para Tests

- Técnicas de mocking sin librerías externas
- Mock de funciones globales
- Testing de integración con mocks
- Aislamiento de dependencias

### ⚡ Módulos de Concurrencia (12-17)

#### **12-channels** - Channels Básicos

- Channels con buffer
- Comunicación síncrona y asíncrona
- Envío y recepción de datos
- Operaciones de channel

#### **13-wait-group** - Sincronización con WaitGroup

- `sync.WaitGroup` para esperar goroutines
- Patrón Add(), Done(), Wait()
- Sincronización de múltiples tareas
- Manejo de concurrencia controlada

#### **14-semaforo** - Semáforos y Control de Concurrencia

- Implementación de semáforos con channels
- Limitación de concurrencia
- Control de recursos compartidos
- Patrón semáforo para rate limiting

#### **15-pipelines** - Pipelines Concurrentes

- Cadenas de procesamiento
- Etapas de pipeline con channels
- Fan-in y fan-out patterns
- Procesamiento de streams de datos

#### **16-workers-pools** - Worker Pool Pattern

- Implementación de pools de workers
- Distribución de tareas
- Balanceador de carga simple
- Cálculo de Fibonacci concurrente

#### **17-multiplexacion** - Multiplexación con Select

- Sentencia `select` para múltiples channels
- Manejo no-bloqueante de comunicación
- Timeouts y casos default
- Coordinación de múltiples goroutines

### 🎯 Proyecto Final

#### **proyecto_final** - Servidor HTTP Concurrente

Un servidor HTTP completo que implementa un **Worker Pool** para procesar cálculos de Fibonacci:

**Características principales:**

- 🌐 **Servidor HTTP** en puerto 8081
- 👷 **Worker Pool** con 4 workers concurrentes
- 📦 **Queue System** con capacidad para 20 jobs
- 🔄 **Procesamiento Asíncrono** de trabajos
- ⚡ **API REST** para envío de trabajos
- 🧮 **Cálculo de Fibonacci** como tarea computacional

**Endpoint disponible:**

```sh
POST /fibonacci
Content-Type: application/x-www-form-urlencoded

name=trabajo1&value=35&delay=2s
```

**Arquitectura del sistema:**

- `Job`: Estructura de trabajo con número, nombre y delay
- `Worker`: Procesador concurrente de trabajos
- `Dispatcher`: Coordinador de workers y distribución de tareas
- `RequestHandler`: Manejador HTTP para API REST

## 🚀 Cómo Ejecutar los Ejemplos

### Prerrequisitos

- Go 1.24.3 o superior instalado
- Git para clonar el repositorio

### Ejecución Individual

```bash
# Clonar el repositorio
git clone https://github.com/afperdomo2/curso_go_poo_y_concurrencia.git
cd curso_go_poo_y_concurrencia

# Ejecutar cualquier módulo
cd 01-repaso
go run main.go

# Para módulos con dependencias
cd 09-modulos
go mod download
go run main.go
```

### Testing

```bash
# Ejecutar tests
cd 10-testing
go test

# Tests con coverage
go test -cover

# Generar reporte de coverage
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# CPU profiling
go test -cpuprofile=cpu.out
go tool pprof cpu.out
```

### Proyecto Final

```bash
cd proyecto_final
go run main.go

# En otra terminal, enviar trabajos
curl -X POST http://localhost:8081/fibonacci \
  -d "name=test1&value=35&delay=2s"
```

## 🧠 Conceptos Clave Aprendidos

### Programación Orientada a Objetos en Go

- ✅ **Encapsulación** mediante structs y métodos
- ✅ **Composición** como alternativa a herencia
- ✅ **Polimorfismo** a través de interfaces
- ✅ **Abstracción** con interfaces y factory patterns

### Concurrencia Avanzada

- ✅ **Goroutines** para paralelismo
- ✅ **Channels** para comunicación segura
- ✅ **Select** para multiplexación
- ✅ **Worker Pools** para gestión de recursos
- ✅ **Pipelines** para procesamiento en cadena
- ✅ **Sincronización** con WaitGroups y semáforos

### Testing y Calidad de Código

- ✅ **Unit Testing** con testing package
- ✅ **Table-driven tests** para casos múltiples
- ✅ **Mocking** sin dependencias externas
- ✅ **Coverage analysis** para métricas de calidad
- ✅ **Profiling** para optimización de performance

### Gestión de Proyectos

- ✅ **Go Modules** para dependencias
- ✅ **Versionado** de módulos
- ✅ **Estructura** de proyectos Go
- ✅ **Mejores prácticas** de organización

## 🎓 Progresión del Aprendizaje

El curso está diseñado con una progresión lógica:

1. **Fundamentos** (01-05): Bases de Go y POO
2. **Patrones** (06-08): Patrones de diseño y funciones avanzadas
3. **Gestión** (09-11): Módulos, testing y mocking
4. **Concurrencia** (12-17): De channels básicos a worker pools
5. **Proyecto Final**: Integración de todos los conceptos

## 🔧 Herramientas y Comandos Útiles

### Comandos Go Básicos

```bash
go run main.go           # Ejecutar directamente
go build main.go         # Compilar
go test                  # Ejecutar tests
go mod init <nombre>     # Inicializar módulo
go mod tidy              # Limpiar dependencias
go get <paquete>         # Instalar dependencia
```

### Testing Avanzado

```bash
go test -v               # Tests verbose
go test -cover           # Con coverage
go test -race            # Detectar race conditions
go test -bench=.         # Benchmarks
go test -cpuprofile=cpu.out  # CPU profiling
```

### Herramientas de Coverage

```bash
go tool cover -func=coverage.out     # Coverage por función
go tool cover -html=coverage.out     # Reporte HTML
```

## 📈 Métricas del Proyecto

- **17 módulos** de aprendizaje progresivo
- **1 proyecto final** completo
- **Múltiples patrones** de concurrencia implementados
- **Testing completo** con coverage y profiling
- **Documentación detallada** en cada módulo

## 🎯 Objetivos de Aprendizaje Alcanzados

Al completar este curso, habrás dominado:

- ✅ **Programación Orientada a Objetos** en Go
- ✅ **Concurrencia avanzada** con goroutines y channels
- ✅ **Patrones de diseño** aplicados en Go
- ✅ **Testing profesional** y técnicas de mocking
- ✅ **Gestión de proyectos** con Go modules
- ✅ **Desarrollo de servidores** HTTP concurrentes
- ✅ **Optimización de performance** con profiling
- ✅ **Mejores prácticas** de desarrollo en Go

## 🤝 Contribuciones

Este proyecto es de código abierto y las contribuciones son bienvenidas. Si encuentras errores o quieres agregar ejemplos adicionales, por favor:

1. Fork el repositorio
2. Crea una rama para tu feature
3. Commit tus cambios
4. Push a la rama
5. Abre un Pull Request

## 📞 Contacto

- **Autor**: afperdomo2
- **GitHub**: [https://github.com/afperdomo2](https://github.com/afperdomo2)
- **Repositorio**: [curso_go_poo_y_concurrencia](https://github.com/afperdomo2/curso_go_poo_y_concurrencia)

---

**¡Domina Go desde los fundamentos hasta la concurrencia avanzada!** 🚀
