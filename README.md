> Consul Watch handler for Slack

[![Build Status](https://img.shields.io/travis/pennsignals/gentry.svg?style=flat-square)](https://travis-ci.org/pennsignals/gentry) [![Coverage Status](https://img.shields.io/coveralls/github/pennsignals/gentry.svg?style=flat-square)](https://coveralls.io/github/pennsignals/gentry)

## Testing

### Unit Testing

    $ go test -v ./...

### System Testing

`gentry` contains a [Docker Compose](https://github.com/pennsignals/gentry/blob/master/docker-compose.yml) file for local testing. The Docker Compose file defines a `consul` service that uses the official [consul](https://hub.docker.com/_/consul/) Docker image. The [Consul Agent](https://www.consul.io/docs/agent/basics.html) configuration registers a `web` service that performs a periodic [health check](https://www.consul.io/docs/agent/checks.html). By default, the health check for the `web` service is designed to fail and the state of the check return `critical`.

To run the Docker Compose file, invoke the following command:

	$ docker-compose up

The Consul Agent configuration enables the web user interface and is available at: [http://localhost:8500/](http://localhost:8500/)

<p align="center">
    <img src="https://user-images.githubusercontent.com/2184329/32913884-b3ffaad4-cae1-11e7-9435-2e409a1031c1.png" alt="Consul Web UI">
</p>
