  <img src="https://raw.githubusercontent.com/openbaton/openbaton.github.io/master/images/openBaton.png" width="250"/>

  Copyright © 2015-2016 [Open Baton](http://openbaton.org).
  Licensed under [Apache v2 License](http://www.apache.org/licenses/LICENSE-2.0).

# Go Dummy VNFM (AMQP)
`go-dummy-vnfm` is a sample Virtual Network Function Manager implementation for [Open Baton][openbaton], written using Go and [go-openbaton].

The Go Dummy VNFM behaves similarly to the [Java Dummy AMQP VNFM][dummy-vnfm-amqp] and can be considered a drop-in replacement.

The two behave exactly the same. The VNFM will connect through AMQP and will grant every request sent by the NFVO, using the `test` VIM driver.

## Technical Requirements

You need a fully working NFVO to use the Dummy VNFM, plus either the Java or Go test VimManager running and available to the NFVO; the official Docker container image is recommended.

## How to install go-dummy-vnfm

On both *NIX and Windows:
```shell
go get -u github.com/openbaton/go-dummy-vnfm
```

The `go` tool will automatically fetch and build both the sources and their dependencies, and a `go-dummy-vnfm` binary will be generated in `$GOPATH/bin` (`%GOPATH%\bin` on Windows CMD).

## How to use go-dummy-vnfm

 ```bash
 go-dummy-vnfm --cfg /path/to/config.toml
 ```

The VNFM must be configured using a configuration file, specified through the `--cfg` flag (see [the sample configuration for further details]()).

In case no such flag is specified, the default behaviour is to search in the current directory for a file named `config.toml`.

Follow the [tutorial] to learn how to use the dummy VNFM.

## How to configure go-dummy-vnfm

The sample configuration should work straight out of the box with a standard local setup of OpenBaton.

`vnfm.logger.level` can be used to change the default verbosity of the logger, choosing a value between `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL` and `PANIC`.

## How to extend go-dummy-vnfm

Feel free to use `go-dummy-vnfm` as a starting point to create new VNFMs, you only need to fill the dummy methods with actual code.

## Issue tracker

Issues and bug reports should be posted to the GitHub Issue Tracker of this project.
# What is Open Baton?

Open Baton is an open source project providing a comprehensive implementation of the ETSI Management and Orchestration (MANO) specification and the TOSCA Standard.

Open Baton provides multiple mechanisms for interoperating with different VNFM vendor solutions. It has a modular architecture which can be easily extended for supporting additional use cases. 

It integrates with OpenStack as standard de-facto VIM implementation, and provides a driver mechanism for supporting additional VIM types. It supports Network Service management either using the provided Generic VNFM and Juju VNFM, or integrating additional specific VNFMs. It provides several mechanisms (REST or PUB/SUB) for interoperating with external VNFMs. 

It can be combined with additional components (Monitoring, Fault Management, Autoscaling, and Network Slicing Engine) for building a unique MANO comprehensive solution.

## Source Code and documentation

The Source Code of the other Open Baton projects can be found [here][openbaton-github] and the documentation can be found [here][openbaton-doc] .

## News and Website

Check the [Open Baton Website][openbaton]
Follow us on Twitter @[openbaton][openbaton-twitter].

## Licensing and distribution
Copyright © [2015-2017] Open Baton project

Licensed under the Apache License, Version 2.0 (the "License");

you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

## Support
The Open Baton project provides community support through the Open Baton Public Mailing List and through StackOverflow using the tags openbaton.

## Supported by
  <img src="https://raw.githubusercontent.com/openbaton/openbaton.github.io/master/images/fokus.png" width="250"/><img src="https://raw.githubusercontent.com/openbaton/openbaton.github.io/master/images/tu.png" width="150"/>

[dummy-vnfm-amqp]: https://github.com/openbaton/dummy-vnfm-amqp
[fokus-logo]: https://raw.githubusercontent.com/openbaton/openbaton.github.io/master/images/fokus.png
[go-openbaton]: http://github.com/openbaton/go-openbaton
[openbaton]: http://openbaton.org
[openbaton-doc]: http://openbaton.org/documentation
[openbaton-github]: http://github.org/openbaton
[openbaton-logo]: https://raw.githubusercontent.com/openbaton/openbaton.github.io/master/images/openBaton.png
[openbaton-mail]: mailto:users@openbaton.org
[openbaton-twitter]: https://twitter.com/openbaton
[sample-conf]: https://raw.githubusercontent.com/openbaton/go-dummy-vnfm/master/config.toml.sample
[tub-logo]: https://raw.githubusercontent.com/openbaton/openbaton.github.io/master/images/tu.png
[tutorial]: https://openbaton.github.io/documentation/dummy-NSR