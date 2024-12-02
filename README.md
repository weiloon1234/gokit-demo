### reload core package with command

# TO update GoKit to latest version
`make gokit-update`

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