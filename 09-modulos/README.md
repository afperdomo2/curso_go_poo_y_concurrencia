# Usando módulos

```sh
# Acceder a la carpeta
cd 09-modulos

# Crear un módulo en este directorio
go mod init github.com/afperdomo2/test

# Instalar un módulo externo de pruebas
go get github.com/donvito/hellomod

# Instalar la versión 2
go get github.com/donvito/hellomod/v2
```

```sh
# Compilar el proyecto
go build main.go

# Ejecutar el proyecto compilado
./main.exe
```

```sh
# Eliminar dependencias que no se están utilizando
go mod tidy
```
