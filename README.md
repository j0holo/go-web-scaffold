# Golang scaffold

A golang scaffold to make logging, connecting to databases and caching servers easier.

The goal is to provide a simple scaffold to start quickly with developing new web services.

## Functional requirements

- Logging module which should be easy and flexible, e.g. multiple log files and easy formatting.
- MySQL module with connection pooling.
- Memcached connection module.
- Config file with paths to common settings (path to SSL certs, database URL, memcached URL, default log format).
- Config parser for parameters of other modules.

## Directory overview

- [server.go](./server.go)
- [README.md](./README.md)
- [controller](./controller)    // Contains http.Handlers and Gorilla mux
  - [controller.go](./controller/controller.go)
- [log](./log)            // Creates new loggers
  - [log.go](./log/log.go)
- [model](./model)              // Creates database connections for MySQL and Memcached
  - [model.go](./model/model.go)
- [parser](./parser)            // Parses `config.json` and contains function to provide settings for other modules
  - [parser.go](./parser/parser.go)
- [view](./view)                // Defines the views of the web framework, imported by the controller
  - [view.go](./view/view.go)
