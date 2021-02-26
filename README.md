# openbanking-sandbox

![alt Cool Sandbox diagram](https://raw.githubusercontent.com/cloudentity/openbanking-sandbox/master/docs/ob-sandbox-diagram.png)

## Overview

This sandbox demonstrates Cloudentity's ACP capabilities for handling Openbanking scenarios.

The sandbox consists of the following components:

* [ACP - Cloudentity's Authorization Control Plane](https://docs.authorization.cloudentity.com/)
  - single node cockroachdb
  - hazelcast
* tpp - technical third party provider
* financroo - advanced third party provider which allows to connect multiple banks
* consent-page - sample bank's consent page
* consent-self-service - self service for consent management
* consent-admin - admin portal for tpps consent management
* bank - sample bank APIs

## Setup

We're using docker-compose to spin the sandbox.

### Prerequisites

* docker 19.03.x
* docker-compose 1.25.x
* make
* contact sales@cloudentity.com to get credentials to be able to download ACP docker image

For Windows you need to use WSL.

### Quickstart

To run the sandbox use the following command:

```sh
make run
```

It will download prebuild images from the docker hub.

If you finish playing with the sandbox you can turn it off:

``` sh
make clean
```

### Credentials

Once everything is up, open TPP technical app:

```
https://localhost:8090
```

Use one of the following users to log in:
- `user / p@ssw0rd! `
- `user2 / p@ssw0rd! `
- `user3 / p@ssw0rd! `

You can also use a Financroo app (use `test / p@ssw0rd!` to log in):

```
https://localhost:8091
```

Consent self service (use the same credentials as for TPP):

```
https://localhost:8085
```

Consent admin (credentials: `admin / p@ssw0rd!`):

```
https://localhost:8086
```

If you would like to access ACP's admin portal (credentials: `admin / p@ssw0rd!`) use:

```
https://localhost:8443/app/default/admin
```

### Developer commands

To build docker images locally run:

```sh
make build
```

To run the sandbox with local images, use:

```sh
make run-dev
```

### Run e2e tests locally

``` sh
cd tests
npm i
yarn run cypress open
```
