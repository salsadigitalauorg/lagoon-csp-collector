# Lagoon CSP collector
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/salsadigitalauorg/lagoon-csp-collector)
[![Go Report Card](https://goreportcard.com/badge/github.com/salsadigitalauorg/lagoon-csp-collector)](https://goreportcard.com/report/github.com/salsadigitalauorg/lagoon-csp-collector)
[![Release](https://img.shields.io/github/v/release/salsadigitalauorg/lagoon-csp-collector)](https://github.com/salsadigitalauorg/lagoon-csp-collector/releases/latest)

## Installation

### Docker

Run directly from a docker image:
```sh
docker run --rm ghcr.io/salsadigitalauorg/lagoon-csp-collector:main lagoon-csp-collector <flags>
```

Or add to your docker image:

```Dockerfile
COPY --from=ghcr.io/salsadigitalauorg/lagoon-csp-collector:main /usr/local/bin/lagoon-csp-collector /usr/local/bin/lagoon-csp-collector
```

## Usage

```
$ lagoon-csp-collector -h
Usage of lagoon-csp-collector:
  -api string
        The endpoint to hydrate the CSP report
  -port string
        Port to run the collector on (default "3000")
  -test-domain string
        A domain to validate in the health check
```

### Flags

#### `api`

An API endpoint to retrieve additional data from for the CSP, this is intended to return a Lagoon project name that matches the domain name for a the CSP report violation.

#### `port`

The port that the service runs on.

#### `test-domain`

A domain that will be sent to `api` during the health check to determine if the service is up and responding correctly.

## Local development

### Build
```sh
git clone git@github.com:salsadigitalauorg/lagoon-csp-collector.git && cd lagoon-csp-collector
go generate ./...
go build -ldflags="-s -w" -o build/lagoon-csp-collector .
go run . -h
```