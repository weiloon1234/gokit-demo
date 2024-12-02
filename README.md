### reload core package with command

# TO update GoKit to latest version
`make gokit-update`

# Migration
```
#create new migration
go run cmd/cli/main.go migrate create the_migration_name

#migrate
go run cmd/cli/main.go migrate up

#rollback
go run cmd/cli/main.go migrate down

#migrate (force)
go run cmd/cli/main.go migrate force
```

# BUILD -> CLI
`go build -o bin/cli ./cmd/cli` or `make build-cli`

# BUILD -> WEB
`go build -o bin/web ./cmd/web` or `make build-web`

# BUILD ALL
`make build-all`

# RUN SERVER (WITHOUT BUILD)
`make run-server`

# #BUILD AND RUN SERVER
`make run-server-build`