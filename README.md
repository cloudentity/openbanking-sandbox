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

Once everything is up, open TPP app:

```
http://localhost:8090
```

Use `user / user` credentials to login in as a sample user.


If you would like to access ACP's admin portal (credentials: `admin / admin`) use:
```
https://localhost:8443/app/default/admin
```

If you end up playing with the sandbox you can turn it off:

``` sh
make clean
```

### ACP configuration

By default when you execute `make run`, ACP configuration is automatically imported from `data/seed.yaml` file.
The import file contains preconfigured tenant, workspaces, services and clients necessary to run the demo flow.

If you would like to configure ACP from scratch please follow the guideliness below:

* create tenant
* create new OpenBanking UK compliant workspace for your customers
* in the new workspace
  * configure custom consent page (this can be done currently via curl by updating authorization server settings)
  * create TPP application with the following settings:
    - grant_types: client_credentials, authorization_code
    - respones_types: token, code
    - request object signing algorithm: RS256
    - Token Endpoint Authentication Method: TLS Client Authentication
    - Certificate metadata: TLS_CLIENT_AUTH_SUBJECT_DN
    - Subject Distinguished Name: << put your TPP cert DN here >>
    - JSON WEB KEY SET - here you need to *public* part of your tpp certificate in jwks format
      * you can use acp cli command to convert pem keys to jwks: 
      ```
      acp jwks convert -k cert-key.pem -c cert.pem -p jwks-public.json -j jwks-private.json
      ```
    - enable Certificate bound access tokens
    - scopes tab:
      * Profile -> openid
      * Openbanking -> accounts
  * once app has been created adjust env variables in docker-compose.yaml file for tpp app
  * create bank resource server application with the following settings:
    - grant_types: client_credentials
    - reponse_type: token
    - request object signing algorithm: none
    - Token Endpoint Authentication Method: TLS Client Authentication
    - Certificate metadata: TLS_CLIENT_AUTH_SUBJECT_DN
    - Subject Distinguished Name: << put your Bank cert DN >>
    - enable Certificate bound access tokens
    - scopes tab:
      * Openbanking -> introspect_openbaking_tokens
  * once app has been created adjust env variables in docker-compose.yaml file for bank app
  * go to settings -> Authorization -> Trusted client certificates and put your tpp client certificate there
* in the `system` workspace
  * create Bank Consent Page application with the following details:
    - grant_types: client_credentials
    - reponse_type: token
    - request object signing algorithm: none
    - Token Endpoint Authentication Method: Client secret post
    - scopes tab:
      * Openbanking -> manage_openbanking_consents
  * once app has been created adjust env variables in docker-compose.yaml file for consent-page app
