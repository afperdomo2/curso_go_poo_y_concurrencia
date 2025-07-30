# Testing

```sh
# Ir a la carpeta
cd 10-testing

# Para poder ejecutar los tests, se necesita inicializar un módulo
go mod init github.com/afperdomo2/testing

# Ejecutar los tests
go test

# Saber qué porcentaje de código está cubierto por tests
go test -cover

# Ejecuta las pruebas y genera un archivo detallado de cobertura de código.
# En windows: --coverprofile
# En Linux: -coverprofile
go test --coverprofile=coverage.out

# Analizar el archivo coverage.out

# Permite identificar rápidamente qué funciones están bien cubiertas por pruebas y cuáles no
# Indica el porcentaje de líneas ejecutadas en cada función y el total del paquete.
go tool cover --func=coverage.out

# Genera un reporte visual en formato HTML mostrando la cobertura de código fuente a partir del archivo coverage.out.
go tool cover --html=coverage.out
```
