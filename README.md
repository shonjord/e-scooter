Scootin' Aboot
===============

<img width="200" align="center" alt="Flash" src="https://www.pngkey.com/png/full/90-906654_kick-scooter-png-transparent-image-scooter-transparent-background.png" />

scootin' aboot is a micro-service written in [go](https://golang.org/) responsible to manage scooter's trips.

## Table of Contents

* [Maintainers](#maintainers)
* [Getting started](#getting-started)
    * [Prerequisites](#prerequisites)
    * [Development](#development)
    * [Testing](#testing)
* [Simulation](#simulation)
* [API docs](#api-docs)

## Maintainers

* [Albert Agelviz] aagelviz@gmail.com

[[table of contents]](#table-of-contents)

## Getting started

### Prerequisites

**Required GO Path**
- [GoPath](https://github.com/golang/go/wiki/SettingGOPATH)
    * This PATH is mandatory to be able to work with this service

**Required Tools**
- [Docker](https://docs.docker.com/docker-for-mac/install/)
    * Also available via HomeBrew, `brew install docker docker-compose && brew cask install docker`

[[table of contents]](#table-of-contents)

### Development

Clone this repository:
```bash
git clone git@github.com:shonjord/scooter.git
```

install a daemon in case this package is missing.
````bash
make daemon
````

Pull and start containers by executing:
```bash
make container-up
```

Run migrations if it's the first time running this application.
```bash
make migration
```

While developing is very common and useful to tail the logs, for this you can execute the following command on your CLI:
```bash
docker-compose logs -f {container}
```

[[table of contents]](#table-of-contents)

### Testing

#### Unit Tests

To run unit tests, execute the following make target:
```bash
make test-unit
```

## Simulation
if you want to make a simulation (mobile client connecting to scooter, notifying the current location, and then disconnecting from the trip), you can run:
```bash
make simulation
```

## API docs
the API docs are located under web/api.html open it with your favorite browser.
