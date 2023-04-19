# Go: Rest & Websockets

REST API with Go: User authentication, post publishing and websockets to read posts in real time.

## Prerequisites

- [Go dev](https://go.dev/)
- [Docker](https://www.docker.com/)

## Tech Stack

- Go
- Docker
- Postgresql
- Websocket

## Deployment

To deploy this project run

### Git clone

```bash
  git clone https://github.com/eafit-201620030010/go-rest-websockets.git
```

### Docker build database

```bash
cd database/

docker build . -t jjchavarrg-ws-rest-db

docker run -p 54321:5432 jjchavarrg-ws-rest-db
```

### Run server

```bash
go run main.go
```

### Docker build app

```bash
docker build . -t jjchavarrg-rest-ws-app

docker run -p 5050:5050 jjchavarrg-rest-ws-app
```

## API Reference

#### Get welcome-api

```http
  GET /
```

#### AUTH

```http
  POST /api/v1/signup
  POST /api/v1/login
  GET  /api/v1/me
```

#### POST

```http
  POST /api/v1/posts
  GET  /api/v1/posts?page=${number}
  PUT  /api/v1/posts/${id}
  GET  /api/v1/posts/${id}
  DEL  /api/v1/posts/${id}
```

#### WebSocket: ws-posts

```http
  ws localhost:5050/ws
```

## Postman colletion

[Colletion Json](./postman)

## Documentation

[Documentation](https://platzi.com/cursos/go-rest-websockets/)

## License

[MIT](https://choosealicense.com/licenses/mit/)
