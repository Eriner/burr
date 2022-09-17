# Burr

Burr is an opinionated **WIP** federated social networking server.

## Quick-start

```sh
docker-compose up
```

## Configuration

The `config.yaml` file defines Burr's configuration.

## Frontend

Burr is a backend-only social network server and does not implement a frontend.

## Backend

In production, Burr requires on the following external services:

* Vault for secrets management
* Faktory for job queuing
* Postgresql for the database
* S3 (external) for internet-facing file hosting
* S3 (internal) for worker-process file mutations
* Redis for worker cache

For development and very small instances, in-memory and on-disk implementations
are included within the app so that it may be run with zero dependencies. However,
only the production "stack" of dependencies is supported. Performance issues,
instability, race conditions, and other issues are present within the builtin
implementation that do not exist within the full production stack.

### API

To provide compatibility with existing federated UIs, such as [Soapbox](https://soapbox.pub/),
Burr will implement the Mastodon and Pleroma HTTP APIs.

The aim of v1.0 is to implement the [Mastodon API](https://docs.joinmastodon.org/api/).

v1.1 will implement the [Pleroma API](https://docs-develop.pleroma.social/backend/development/API/pleroma_api/).

v1.2 will implement a GraphQL API so that *new* frontends can be created more easily.

Due to deficiencies and the [W3C's abandonment of the spec](https://activitypub.rocks),
this project does not seek to implement the C2S ActivityPub spec at this time.

### Database

[ent](https://entgo.io/) is used as the database driver. Multiple database drivers
exist, however postgresql is the only supported production database driver.

I've done my best to implement opportunistic locking for sqlite where possible,
but for both performance and reliability reasons only postgresql is supported.

## Development

**Test:**

```sh
make test
```

**Run:**
```sh
make run
```

**Build `./burr`:**
```sh
make build
```

**Build (and test) `ghcr.io/eriner/burr:latest`:**
```sh
make docker
```

## History

* Sep 13, 2022: [The start](https://noagendasocial.com/@eriner/108993718610367766)
