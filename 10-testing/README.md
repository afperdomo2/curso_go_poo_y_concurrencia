# Testing

```sh
# Ir a la carpeta
cd 10-testing

# Para poder ejecutar los tests, se necesita inicializar un módulo
go mod init github.com/afperdomo2/testing

# Ejecutar los tests
go test
```

## 📊 Coverage / Cobertura

```sh
# Saber qué porcentaje de código está cubierto por tests
go test -cover

# Ejecuta las pruebas y genera un archivo detallado de cobertura de código.
# En windows: --coverprofile
# En Linux: -coverprofile
go test --coverprofile=coverage.out

# Permite identificar rápidamente qué funciones están bien cubiertas por pruebas y cuáles no
# Indica el porcentaje de líneas ejecutadas en cada función y el total del paquete.
go tool cover --func=coverage.out

# Genera un reporte visual en formato HTML mostrando la cobertura de código fuente a partir del archivo coverage.out.
go tool cover --html=coverage.out
```

## ⚡ CPU

```sh
# Ejecuta los tests y, además de correr las pruebas, genera un archivo llamado cpu.out con un perfil de uso de CPU
go test --cpuprofile=cpu.out

# Abre la herramienta de análisis de perfiles (pprof) de Go usando el archivo cpu.out generado por go test --cpuprofile=cpu.out.
#
# Ejemplo de uso interactivo:
# Dentro de pprof puedes ejecutar comandos como top, list <función>, o web para visualizar los resultados.
go tool pprof cpu.out


# NOTE: Si tienes problemas en windows con graphviz para generar los reportes 'web' o 'pdf', se puede
#
# instalar usando chocolatey. Ejecutar el siguiente comando en PowerShell en modo admin
#
choco install graphviz
```
