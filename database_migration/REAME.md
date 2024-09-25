# Database Migration process
This document describes the process of migrating the database from one version to another.
We are using the [`go-migrate` cli](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) to manage the database migrations.


```shell
migrate -source "file:///Users/thomas.poignant/dev/thomaspoignant/app-api/database_migration" \
    -database "postgres://goff-user:my-secret-pw@localhost:5432/gofeatureflag?sslmode=disable" down
```
