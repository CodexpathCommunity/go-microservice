# Go Microservice

A simple microservice written in Go.

## Structure

```txt
├── cmd
│   └── postservice
├── internal
│   └── apps
│       ├── authservice
│       │   ├── api
│       │   ├── clients
│       │   ├── config
│       │   ├── db
│       │   │   ├── migrations
│       │   │   └── seeds
│       │   ├── rpc
│       │   └── service
│       ├── postservice
│       │   ├── api
│       │   ├── clients
│       │   ├── config
│       │   ├── db
│       │   │   ├── migrations
│       │   │   └── seeds
│       │   ├── rpc
│       │   └── service
│       │       └── repository
│       └── userservice
│           ├── api
│           ├── clients
│           ├── config
│           ├── db
│           │   ├── migrations
│           │   └── seeds
│           ├── rpc
│           └── service
├── pkg
│   ├── database
│   └── kv
├── scripts
└── web
```

## To Do

- [ ] Microservices (Currently in progress)
  - [ ] Authentication Service
  - [ ] Post Service
  - [ ] User Service
- [ ] API Gateway (Kong / Tyk)
- [ ] Authentication (JWT / OAuth)
- [ ] Message Broker (RabbitMQ / Kafka)
- [ ] OpenTracing (Jaeger)
- [ ] Centralized Configuration (Consul)
