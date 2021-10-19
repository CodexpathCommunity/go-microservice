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
