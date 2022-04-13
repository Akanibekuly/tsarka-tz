# tsarka-tz

### Env variables:

```
DEBUG=true
LOG_LEVEL=debug
HTTP_LISTEN=:80
HTTP_CORS=false
REDIS_URL=redis:6379
REDIS_PSW=psw
REDIS_DB=0
PG_DSN=postgres://user:pass@host:5432/dbname?sslmode=disable

```

### DB dump:

```
pg_dump --no-owner -Fc -U postgres stg -f ./stg.custom
```

### DB restore:

```
dropdb -U postgres stg
createdb -U postgres stg
pg_restore --no-owner -d stg -U postgres ./stg.custom
```

### Install `migrate` command-tool:

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

### Create new migration:

```
migrate create -ext sql -dir migrations mg_name
```

### Apply migration:

```
migrate -path migrations -database "postgres://localhost:5432/db_name?sslmode=disable" up
```

<br/>

### Install `swagger-cli`:

```
dir=$(mktemp -d) 
git clone https://github.com/go-swagger/go-swagger "$dir" 
cd "$dir"
go install ./cmd/swagger
rm -rf "$dir"
```
