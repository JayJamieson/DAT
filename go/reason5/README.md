# Go tooling

Examples here show case Go tooling

## Formatting Linting S

- `gofmt -w -s -d dat.go` format file
- `gofmt -w -s -d dat.go` format all files recursively
- `go fmt ./...` wrapper command for the two above commands
- `go vet dat.go` - static analysis of file, think of it as psalm

## Adding dependencies

`github.com/alexedwards/argon2id`

- `go get github.com/foo/bar@latest` get last latest version of package
- `go get github.com/foo/bar@afge1c` get package version from commit id

## Cleanup dependencies

- `go mod tidy` makes sure go.mod matches source code, removes unused dependencies
- `go mod verify` make sure dependencies haven't changed since they were downloaded
  - `require github.com/alexedwards/argon2id v0.0.0-20211130144151-3585854a6387 // indirect`
- `go mod why -m github.com/alexedwards/argon2id` explains why this dependency is used

## Documentation

- `go doc strings` simplified documentation of strings package, method signatures and constants
- `go doc -all strings` shows full documentation

## Testing and Test coverage

- `go test .` runs all tests in current directory
- `go test ./...` runs all tests
- `go test -coverprofile=/tmp/profile.out ./...` run tests with test coverage
- `go tool cover -html=/tmp/profile.out` convert test coverage report to html

## References

- <https://www.alexedwards.net/blog/an-overview-of-go-tooling>
