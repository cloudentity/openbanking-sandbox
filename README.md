# openbanking-sandbox

![alt Cool Sandbox diagram](https://raw.githubusercontent.com/cloudentity/openbanking-sandbox/master/docs/ob-sandbox-diagram.png)

## Overview

This sandbox demonstrates Cloudentity's ACP capabilities for handling Openbanking scenarios.

The sandbox consists of the following components:

* [ACP - Cloudentity's Authorization Control Plane](https://docs.authorization.cloudentity.com/)
  - single node cockroachdb
  - hazelcast
* tpp - sample third party provider 
* consent-page - sample bank's consent page
* consent-self-service - self service for consent management
* consent-admin - admin portal for tpps consent management
* bank - sample bank APIs

Please follow `Setup` section, TPP app itself provides a description of the demo flow. 

## Setup

We're using docker-compose to spin the sandbox.

### Prerequisites

* docker 19.03.x
* docker-compose 1.25.x
* make
* contact sales@cloudentity.com to get credentials to be able to download ACP docker image

### Commands

To build docker images run:

```sh
make build
```

Than to run the sandbox, execute:

```sh
make run
```

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

If you end up playing with the sandbox you can turn it off:

``` sh
make clean
```
