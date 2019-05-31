# hc

[![GoDoc Widget]][GoDoc] [![Travis Widget]][Travis]

`hcf` is a lightweight framework to develop HomeKit accessories in Go. This is a fork from github.com/brutella/hc 
It abstracts the **H**omeKit **A**ccessory **P**rotocol (HAP) and makes it easy to work with [services](service/README.md) and [characteristics](characteristic/README.md).

`hcf` handles the underlying communication between HomeKit accessories and clients.
You can focus on implementing the business logic for your accessory, without having to worry about the protocol.

**What is HomeKit?**

[HomeKit][homekit] is a set of protocols and libraries from Apple. It is used by Apple's platforms to communicate with smart home appliances. A non-commercial version of the documentation is now available on the [HomeKit developer website](https://developer.apple.com/homekit/).

HomeKit is fully integrated into iOS since iOS 8. Developers can use [HomeKit.framework](https://developer.apple.com/documentation/homekit) to communicate with accessories using high-level APIs.

<img alt="Home.app" src="_img/home-icon.png?raw=true" width="92" />

[Travis]: https://travis-ci.org/grumpylabs/hcf
[Travis Widget]: https://travis-ci.org/grumpylabs/hcf.svg

## Features

- Full implementation of the HAP in Go
- Supports all HomeKit [services and characteristics](service/README.md)
- Built-in service announcement via DNS-SD using [dnssd](http://github.com/brutella/dnssd)
- Runs on linux and macOS



# License

`hcf` is available under the Apache License 2.0 license. See the LICENSE file for more info.

[homekit]: https://developer.apple.com/homekit/
