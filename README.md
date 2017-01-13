# Go Dummy VNFM (AMQP)
`go-dummy-vnfm` is a sample Virtual Network Function Manager implementation for [OpenBaton](http://openbaton.github.io), written using Go and [go-openbaton](http://github.com/mcilloni/go-openbaton).

The Go Dummy VNFM behaves similarly to the [Java Dummy AMQP VNFM](https://github.com/openbaton/dummy-vnfm-amqp) and can be considered a drop-in replacement.

The two behave exactly the same. The VNFM will connect through AMQP and will grant every request sent by the NFVO, using the `test` VIM.

## Technical Requirements

You need a fully working NFVO to use the Dummy VNFM, plus either the Java or Go test VimManager running and available to the NFVO; the official Docker container image is recommended.

## How to install `go-dummy-vnfm`

On both *NIX and Windows:
```shell
go get -u github.com/mcilloni/go-dummy-vnfm
```

The `go` tool will automatically fetch and build both the sources and their dependencies, and a `go-dummy-vnfm` binary will be generated in `$GOPATH/bin` (`%GOPATH%\bin` on Windows CMD).

## How to use `go-dummy-vnfm`

 ```bash
 go-dummy-vnfm --cfg /path/to/config.toml
 ```

The VNFM must be configured using a configuration file, specified through the `--cfg` flag (see [the sample configuration for further details](https://raw.githubusercontent.com/mcilloni/go-dummy-vnfm/master/config.toml.sample)).

In case no such flag is specified, the default is to search in the current directory for a file named `config.toml`.

## How to configure `go-dummy-vnfm`

The sample configuration should work straight out of the box with a standard local setup of OpenBaton.

`vnfm.logger.level` allowss to change the default verbosity of the logger, choosing a value between `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL` and `PANIC`.

## How to extend `go-dummy-vnfm`

Feel free to use `go-dummy-vnfm` as a starting point to create new VNFMs, you only need to fill the dummy methods with actual code.

## Issue tracker

Issues and bug reports should be posted to the GitHub Issue Tracker of this project.

## What is Open Baton?

Open Baton is an open source project providing a comprehensive implementation of the ETSI Management and Orchestration (MANO) specification and the TOSCA Standard.

Open Baton provides multiple mechanisms for interoperating with different VNFM vendor solutions. It has a modular architecture which can be easily extended for supporting additional use cases. 

It integrates with OpenStack as its standard de-facto VIM implementation, and provides a driver mechanism for supporting additional VIM types. It supports Network Service management either using the provided Generic VNFM and Juju VNFM, or integrating additional specific VNFMs. It provides several mechanisms (REST or PUB/SUB) for interoperating with external VNFMs. 

It can be combined with additional components (Monitoring, Fault Management, Autoscaling, and Network Slicing Engine) for building a unique MANO comprehensive solution.

## Source Code and documentation

The Source Code of the other Open Baton projects can be found [here][openbaton-github] and the documentation can be found [here][openbaton-doc]

## News and Website

Check the [Open Baton Website][openbaton]

## Licensing and distribution
Apache License, Version 2.0. See LICENSE for further details.
