# Akash Provider Daemon

[![tests](https://github.com/akash-network/provider/actions/workflows/tests.yaml/badge.svg)](https://github.com/akash-network/provider/actions/workflows/tests.yaml)

This folder contains the Akash Provider Daemon. This piece of software listens to events emitted from the Akash blockchain (code in `../app/app.go`) and takes actions on a connected Kubernetes cluster to provision compute capacity based on the bids that the configured provider key wins. The following are the pieces of the daemon:

## Development environment

[This doc](https://github.com/akash-network/node/blob/master/_docs/development-environment.md) guides through setting up local development environment

## Structure

### [`bidengine`](./bidengine)

The bid engine queries for any existing orders on chain, and based on the on-chain provider configuration, places bids on behalf of the configured provider based on configured selling prices for resources. The daemon listens for changes in the configuration so users can use automation tooling to dynamically change the prices they are charging w/o restarting the daemon. You can see the key management code for `provider` tx signing in `cmd/run.go`.

### [`cluster`](./cluster)

The cluster package contains the necessary code for interacting with clusters of compute that a `provider` is offering on the open marketplace to deploy orders on behalf of users creating `deployments` based on `manifest`s. Right now only `kubernetes` is supported as a backend, but `providers` could easily implement other cluster management solutions such as OpenStack, VMWare, OpenShift, etc...

### [`cmd`](./cmd)

The `cobra` command line utility that wraps the rest of the code here and is buildable.

### [`event`](./event)

Declares the pubsub events that the `provider` needs to take action on won leases and received manifests.

### [`gateway`](./gateway)

Contains hanlder code for the rest server exposed by the `provider`

### [`manifest`](./manifest)

# Archived Repository

This repository was initially used for educational and testing purposes during our early exploration of compute infrastructure, including a period of collaboration and experimentation with the Akash Network under the Apache 2.0 license.

We have now migrated to a new, independently developed production codebase under a new GitHub organization:
👉 [https://github.com/spheron-core/](https://github.com/spheron-core/)

This migration aligns with our roadmap toward TGE, Foundation-based governance, and long-term code maintainability.

Note:

- This repository is now deprecated and archived.
- It is no longer maintained and is not used in any part of the production infrastructure.
- All original attributions have been preserved in compliance with the Apache 2.0 license.
- No active development will occur on this repository moving forward.

For the latest updates and active development, please refer to the new organization above.
