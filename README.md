<p align="center">
    <a href="https://github.com/pennsignals/gentry"><img src="https://rawgit.com/jasonwalsh/d6ab09f72bafa5774c253c736087c089/raw/f813b33b6ccf7bda410436054b48b0a2756e8238/gentry.svg" width="300"></a>
</p>

> [Consul](https://www.consul.io/) Watch handler for Slack

[![Build Status](https://img.shields.io/travis/pennsignals/gentry.svg?style=flat-square)](https://travis-ci.org/pennsignals/gentry) [![Coverage Status](https://img.shields.io/coveralls/github/pennsignals/gentry.svg?style=flat-square)](https://coveralls.io/github/pennsignals/gentry)

## Usage

> [Watches](https://www.consul.io/docs/agent/watches.html) are a way of specifying a view of data (e.g. list of nodes, KV pairs, health checks) which is monitored for updates.

Consul Watches are executable via the Consul Agent [Configuration](https://www.consul.io/docs/agent/options.html) or using the [`consul watch`](https://www.consul.io/docs/commands/watch.html) command. Below demonstrates both examples in their respective order:

### Agent Configuration

```json
{
    "data_dir": "/var/lib/consul",
    "server": true,
    "watches": [
        {
            "args": [
                "/opt/gentry",
                "-channel",
                "<CHANNEL>",
                "-token",
                "<TOKEN>"
            ],
            "type": "checks"
        }
    ]
}
```

### `consul watch` Command

    $ consul watch -type=checks /opt/gentry -channel=<CHANNEL> -token=<TOKEN>

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
