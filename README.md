# Burr

Burr is a **WIP** federated social network server.

## Quick-start

```sh
docker-compose up
```

## Configuration

The `config.yaml` file defines Burr's configuration.

## Frontend

Burr is a backend-only social network server and does not implement a frontend.

The aim of v1.0 is to implement the Pleroma API so that existing UIs "just work".

## Backend

Burr requires on the following external services:

* Vault for secrets management
* Faktory for job queuing
* Postgresql for the database
* S3 (external) for internet-facing file hosting
* S3 (internal) for worker-process file mutations
* Redis for worker cache

### API

To provide compatibility with existing federated UIs, such as Soapbox, Burr will
implement the Pleroma HTTP API.

### Database

ent is used as the database driver and it should be possible to support multiple
database driver implementations, however postgresql is the only supported
production database driver.

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
