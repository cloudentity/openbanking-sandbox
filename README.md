# openbanking-sandbox

## Overview
This sandbox demonstrates Cloudentity's ACP capabilities for handling Openbanking scenarios.

The sandbox consists of the following components:

* [ACP - Cloudentity's Authorization Control Plane](https://docs.authorization.cloudentity.com/)
  - single node cockroachdb
  - hazelcast
* tpp - sample third party provider 
* consent-page - sample bank's consent page
* bank - sample bank APIs

## Setup

We're using docker-compose to spin the sandbox.

### Prerequisites

* docker 19.03.x
* docker-compose 1.25.x
* make
* contact sales@cloudentity.com to get credentials to be able to download ACP docker image

### Quickstart

To spin the sandbox run:

```sh
make clean build run
```

Once everything is up, open TPP app:

```
http://localhost:8090
```

Use `user / user` credentials to login in as a sample user.


If you would like to access ACP's admin portal (credentials: `admin / admin`) use:
```
https://localhost:8443/app/default/admin
```

### ACP configuration

By default when you execute `make run`, ACP configuration is automatically imported from `data/seed.yaml` file.
The import file contains tenant, workspaces, services and clients configuration necessary to run the demo flow.

If you would like to configure ACP from scratch please follow the guideliness below:

* create tenant
* create new OpenBanking UK compliant workspace for your customers
* in the new workspace
  * create TPP application with the following settings:
    - grant_types: client_credentials, authorization_code
    - respones_types: token, code
    - request object signing algorithm: RS256
    - Token Endpoint Authentication Method: TLS Client Authentication
    - Certificate metadata: TLS_CLIENT_AUTH_SUBJECT_DN
    - Subject Distinguished Name: << put your TPP cert DN >>, tpp certificate provided in this sandbox: `data/tpp_*.pem` uses `cid1.authorization.cloudentity.com` 
    - JSON WEB KEY SET - here you need to *public* part of your tpp certificate in (jwks format)[https://8gwifi.org/jwkconvertfunctions.jsp]
    - enable Certificate bound access tokens
    - scopes tab:
      * Profile -> openid
      * Openbanking -> accounts
  * create bank resource server application with the following settings:
    - grant_types: client_credentials
    - reponse_type: token
    - request object signing algorithm: none
    - Token Endpoint Authentication Method: TLS Client Authentication
    - Certificate metadata: TLS_CLIENT_AUTH_SUBJECT_DN
    - Subject Distinguished Name: << put your TPP cert DN >>, tpp certificate provided in this sandbox: `data/bank_*.pem` uses `cid2.authorization.cloudentity.com` 
    - enable Certificate bound access tokens
    - scopes tab:
      * Openbanking -> introspect_openbaking_tokens
  * go to settings -> Authorization -> Trusted client certificates and put your tpp client certificate there (in this sandbox `data/ca.pem` used)
* in the `system` workspace
  * create Bank Consent Page application with the following details:
    - grant_types: client_credentials
    - reponse_type: token
    - request object signing algorithm: none
    - Token Endpoint Authentication Method: Client secret post
    - scopes tab:
      * Openbanking -> manage_openbanking_consents
