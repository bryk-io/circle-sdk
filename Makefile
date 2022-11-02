.PHONY: *
.DEFAULT_GOAL:=help

# For commands that require a specific package path, default to all local
# subdirectories if no value is provided.
pkg?="..."

help:
	@echo "Commands available"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /' | sort

## deps: Verify dependencies and remove intermediary products
deps:
	go mod tidy
	go clean

## lint: Static analysis
lint:
	# Go code
	golangci-lint run -v ./$(pkg)

## scan-deps: Look for known vulnerabilities in the project dependencies
# https://github.com/sonatype-nexus-community/nancy
scan-deps:
	@go list -mod=readonly -f '{{if not .Indirect}}{{.}}{{end}}' -m all | nancy sleuth --skip-update-check

## test: Run all unitary tests
test:
	# Unit tests
	# -count=1 -p=1 (disable cache and parallel execution)
	go test -race -v -failfast -count=1 -p=1 -coverprofile=coverage.report ./$(pkg)
	go tool cover -html coverage.report -o coverage.html

## updates: List available updates for direct dependencies
# https://github.com/golang/go/wiki/Modules#how-to-upgrade-and-downgrade-dependencies
updates:
	@GOWORK=off go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}} [{{.Version}} -> {{.Update.Version}}]{{end}}' -m all 2> /dev/null
