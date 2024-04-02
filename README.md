# PRUEBA TECNICA TACH

Se me encomendo crear dos proyectos independientes con GO utilizando algun framework, uno para cada microservicio.
Definir las APIs RESTful para cada microservicio.
Implementar la lógica de negocio para cada microservicio.
Utilizar Docker para contenerizar cada microservicio.
Configurar MongoDB para el almacenamiento de datos.
Conectar los microservicios a MongoDB.
Implementar la mensajería con RabbitMQ para el microservicio de transacciones.
Probar los microservicios de forma individual y conjunta.

## Requisitos Previos

  ## Skills
  - Go
  - Fiber framework go
  - Mongodb
  - Docker
  - Conocimientos en arquitectura de microservicios orientada a eventos
  - Experiencia con RESTful APIs
  - Conocimientos en mensajería con RabbitMQ

  ##-Microservicio de Cuentas:

  Permite crear, consultar, actualizar y eliminar cuentas de usuarios.
  Expondrá una API RESTful para la gestión de cuentas.
  Almacenará la información de las cuentas en MongoDB.

  ##- Microservicio de Transacciones:

  Permite realizar transferencias de dinero entre cuentas.
  Se comunicará con el microservicio de cuentas para obtener información de las cuentas.
  Utilizará RabbitMQ para la comunicación entre microservicios.
  Almacenará las transacciones en MongoDB.

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
|->Desafio-pdf
|->docker-compose.yml
|-- README.md
|-- ...
```
