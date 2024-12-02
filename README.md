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

# Seeder
```
#create seeder
  1. create file at database/seeds/
  2. register seeder at cmd/cli/main.go
  
#run seeder
go run cmd/cli/main.go seed the_seeder_name
```

# CLI command
```
#create command
  1. create file at commands/
  2. add command at cmd/cli/main.go 

#run command
go run cmd/cli/main.go command_name
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