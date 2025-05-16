# API specs

## Schema
The API should follow the OpenAPI 3.0 specification defined [here](./open-api.yaml).
You can see it directly here: https://petstore.swagger.io/?url=https://raw.githubusercontent.com/go-feature-flag/flag-management/refs/heads/main/.specs/open-api.yaml

In this schema you will find all the relevant endpoints and their parameters.

## Configuration
Following the example of GO Feature Flag relay proxy, the configuration must be handled by a file based configuration system.

### Library
We should use `koanf` for configuration management.

### Location
We expect the configuration file to be named, `goff-management.yaml` and to be located in one of those locations:
- `./`
- `/goff/`
- `/etc/opt/goff/`

If the file is not found in any of those locations, the application should exit with an error message.

It should also be possible to provide the configuration file location as a command line argument. The command line argument should take precedence over the default locations.
The command line argument should be `--config`.

## Authentication
The API must be protected by an API key authentication mechanism. The API key should be passed in the request headers for all endpoints that require authentication (everything under the `[Core API]` tag).
The API key is passed in the `Authorization` header as a Bearer token.

We should aim to be compatible with most of the authentication providers such as Okta, Auth0, etc...  

## Database
The API should store the data in a database.
The default database must be Postgres, but it should be possible to use other databases by providing a different connection string in the configuration file.

### Migration
All the database changes should be done using a migration mechanism.
All the migrations should be stored in the [`database_migrations`](../database_migrations) folder.

We expect all the migration to be applied manually by running the `migrate` command, we do not expect the application to apply the migrations automatically.


## Flag history
The API `/v1/flags/{id}/versions` is providing the history of a flag. This endpoint should return the history of the flag, including all the changes made to it. 

The history should be stored in a separate table in the database.
We may use a trigger function to store the history of the flag in the database. The trigger function should be called on every update of the flag.