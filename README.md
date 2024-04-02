# Nombre del Proyecto

Breve descripción del proyecto.

## Requisitos Previos

  ##-Microservicio de Cuentas:

  Permite crear, consultar, actualizar y eliminar cuentas de usuarios.
  Expondrá una API RESTful para la gestión de cuentas.
  Almacenará la información de las cuentas en MongoDB.

  ##- Microservicio de Transacciones:

  Permite realizar transferencias de dinero entre cuentas.
  Se comunicará con el microservicio de cuentas para obtener información de las cuentas.
  Utilizará RabbitMQ para la comunicación entre microservicios.
  Almacenará las transacciones en MongoDB.


## Configuración

Explica cualquier configuración necesaria antes de levantar el proyecto, como variables de entorno o archivos de configuración que deben ser modificados.

## Instalación

### 1. Clonar el Repositorio

```bash
git clone https://github.com/tu_usuario/tu_proyecto.git
```

### 2. Construir las Imágenes Docker

```bash
docker-compose build
```

### Levantar el Proyecto

```bash
docker-compose up
```

## Estructura del Proyecto

Explica la estructura de tu proyecto y qué hace cada parte importante del mismo.

```
/
|-- Microservicio-account
|       |-> controller
|       |       |->accounts.go
|       |       |->message.go
|       |->database
|       |      |->mongo.go
|       |->models
|       |      |->model.go
|       |->rabbit_mq
|       |      |->routes.go
|       |->routes
|       |      |->routes.go
|       |->go.sum
|       |->go.mod
|       |->main.go
|       |->Dockerfile
|
|-- Microservicio-transaction
|       |-> controller
|       |       |->transacction.go
|       |       |->message.go
|       |->database
|       |      |->mongo.go
|       |->models
|       |      |->model.go
|       |->rabbit_mq
|       |      |->routes.go
|       |->routes
|       |      |->routes.go
|       |->go.sum
|       |->go.mod
|       |->main.go
|       |->Dockerfile
|
|->docker-compose.yml
|-- README.md
|-- ...
```
