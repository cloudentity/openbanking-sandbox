# Openbanking Sandbox with Cloudentity ACP

This repository contains an Openbanking Sandbox, showcasing how Cloudentity ACP supports
Openbanking usecases.

There are several applications in this repository.

- Bank

  - Bank APIs
  - Bank Administrative Portal
  - Bank Consumer Portal
  - Bank Openbanking Consent Page

- Fintech apps

  - Financroo - sample fintech aggregator
  - Developer TPP - technical TPP application

- ACP - Cloudentity Authorization Control Plane

Logical system diagram:

![](diagrams/logical.svg)

## Usecases

This sandbox provides solutions for different use-cases for few different
personas.

### Fintech App User

#### Usecase: Fintech aggregator

User has multiple banking accounts, he'd like to see all his finances in one place.

Solution: Thanks to the standards which Openbanking defines, fintech applications
can connect to multiple financial institutions using the same secure APIs.

ACP allows the bank to expose Openbanking compliant APIs securely, make sure
that only trusted, secure clients can access bank's APIs.

Simplified flow

![](diagrams/fintech.svg)

#### Usecase: Consent Self Service

User wants to review what data is he sharing with Fintech Apps (TPPs).

They want to revoke a specific consent they've given to one of the apps

Solution: ACP allows banks to expose consent information in Bank's customer portal

`consent-self-servicce` applications contains a sample bank customer portal.
It lists all TPPs with access to Users data.
It allows to revoke individual consents as well as all consents for a selected TPP.

![](diagrams/consent-self-service.svg)

### Bank Administrator

Usecase: Revoke TPP consents

Admin wants to revoke all consents given to a TPP - due to the client being compromised.

Solution: ACP allows to mange consents given to the TPP.

`consent-admin` application contains a sample bank administration portal.
It lists all TPPs and allows to revoke all consents given to a particular TPP.

![](diagrams/consent-admin.svg)
