# ehko

[![CI](https://github.com/mikejoh/ehko/actions/workflows/go.yml/badge.svg)](https://github.com/mikejoh/ehko/actions/workflows/go.yml)
[![Release](https://img.shields.io/github/v/release/mikejoh/ehko)](https://github.com/mikejoh/ehko/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/mikejoh/ehko)](https://goreportcard.com/report/github.com/mikejoh/ehko)

A debugging tool created to help you understand Prometheus and Alertmanager better.

`ehko` is a small HTTP server that echoes back whatever is sent to it. Point Alertmanager's
webhook receiver, or any other HTTP client, at `ehko` and inspect exactly what gets sent -
useful when developing or debugging Alertmanager routes, receivers, and webhook payloads,
or just poking at HTTP clients in general.

## Install

### From source

```
go install github.com/mikejoh/ehko/cmd/ehko@latest
```

### Binary

Download a prebuilt binary from the [releases page](https://github.com/mikejoh/ehko/releases/latest).

### Container image

```
docker pull docker.io/mikejoh/ehko:latest
```

## Usage

Start the server, it listens on `:5001`:

```
ehko
```

### Endpoints

* `GET /` - basic liveness check, responds with `ehko!`.
* `POST /alerts` - decodes the request body as Alertmanager's [webhook payload](https://prometheus.io/docs/alerting/latest/configuration/#webhook_config) and logs it.
* `POST /log` - logs the raw request body as-is, useful for inspecting arbitrary payloads.
* `GET /responder/:code` - responds with the given HTTP status `code`, useful for testing how a client (e.g. Alertmanager) handles specific response codes.
