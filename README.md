### reload core package with command

`go clean -modcache && go mod tidy`

cannot? try next

`GOPRIVATE=github.com/weiloon1234/gokit go get -u github.com/weiloon1234/gokit@main`

still cannot?

try `go clean -modcache`, if have error unlinkat directory not empty then 


`go env GOMODCACHE`

This will return something like:

`/Users/your_name/go/pkg/mod`

then change directory to the folder

`cd /Users/your_name/go/pkg/mod`

and run `rm -rf cache`,then retry

## BUILD
`go build -o bin/cli ./cmd/cli`

`go build -o bin/web ./cmd/web`