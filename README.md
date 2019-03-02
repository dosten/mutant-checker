Mutant Checker
==============

Requisitos
----------

Este proyecto tiene como dependencias:

- Go 1.11
- Redis

Instalacion
-----------

Ejecutar el siguiente comando para compilar el proyecto:

```
$ go build -o build/mutant-checker ./cmd/mutant-checker
```

Una vez compilado el proyecto, y teniendo un server de Redis corriendo, ejecutar:

```
$ REDIS_URL=redis://localhost:6379 ./build/mutant-checker
```

Uso
---

Una vez que la API se encuentre corriendo, se pueden realizar consultas a los dos endpoints disponibles, los cuales poseen las mismas caracteristicas que en el enunciado del challenge.

Tests
-----

Para correr los tests, ejecutar:

```
$ go test ./...
```

Si se quiere observar el coveraga, ejecutar:

```
$ go test -coverprofile cp.out ./...
$ go tool cover -html=cp.out
```

Nota: por cuestion de tiempo decidi solamente enforcarme en tener coverage de la logica de chequeo de mutantes, la cual es del 100%.
