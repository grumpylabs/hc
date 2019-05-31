# hcf

[![Codefresh build status]( https://g.codefresh.io/api/badges/pipeline/rbg/grumpylabs%2Fhcf%2Fhcf?branch=develop&key=eyJhbGciOiJIUzI1NiJ9.NTg4OGUxZmZiZWQ4YjYwMTAwOTc2ZTc4.Us8wF_6Xllyq8OWmKUqBGBTxb8i4bLsniV0a2TOLewo&type=cf-1)]( https://g.codefresh.io/pipelines/hcf/builds?repoOwner=grumpylabs&repoName=hcf&serviceName=grumpylabs%2Fhcf&filter=trigger:build~Build;branch:develop;pipeline:5cf1a1d7bbf24f11489a8523~hcf)

[![GoDoc Widget]][GoDoc] 

`hcf` is a lightweight framework to develop HomeKit accessories in Go. This is a fork from github.com/brutella/hc 
It abstracts the **H**omeKit **A**ccessory **P**rotocol (HAP) and makes it easy to work with [services](service/README.md) and [characteristics](characteristic/README.md).

`hcf` handles the underlying communication between HomeKit accessories and clients.
You can focus on implementing the business logic for your accessory, without having to worry about the protocol.


# License

`hcf` is available under the Apache License 2.0 license. See the LICENSE file for more info.

[homekit]: https://developer.apple.com/homekit/
