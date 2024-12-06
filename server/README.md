# GO Feature Flag API - API to configure your feature flag
![WIP](https://img.shields.io/badge/status-%E2%9A%A0%EF%B8%8FWIP-red)
[![Build](https://github.com/go-feature-flag/flag-management/actions/workflows/ci.yaml/badge.svg)](https://github.com/go-feature-flag/app-api/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/github/go-feature-flag/app-api/graph/badge.svg?token=oqi5Ncgefx)](https://codecov.io/github/go-feature-flag/flag-management/)

This repository is a work in progress initiative to create an API to manage your feature flags.

## Goals
- [x] Create an API to manage your feature flags
- [ ] API should allow to add, modify and delete feature flags.
- [x] Use a database to store the feature flags.
- [ ] This API is created to integrate a front end application to manage the feature flags.
- [ ] We should manage authentication and authorization to access the API.
  - [ ] Authentication should be generic enough to be integrated with any authentication provider.
- [ ] We should be able to provide history of a flag to see when it was created, modified and deleted.

## Tech stack
- GO API using echo
- Postgres database using `sqlx` and `pq` as driver.


## Contributing
⚠️ Since this it is a work in progress initiative please come to the [Slack channel](https://gofeatureflag.org/slack) first before contributing. 

### How to start the project.
After cloning the project you can start the database _(using docker)_:
```shell
make setup-env
```
It will start an instance of postgres with the following credentials:
- user: `goff-user`
- password: `my-secret-pw`

And it will apply the database migrations to your environment.

To start the API:
```shell
make build
./out/bin/goff-api
```

When started you can access the swagger UI at [http://localhost:3001/swagger/](http://localhost:3001/swagger/).
