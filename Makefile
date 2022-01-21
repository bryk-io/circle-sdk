.PHONY: *
.DEFAULT_GOAL:=help

help:
	@echo "Commands available"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /' | sort

## deps: Verify dependencies and remove intermediary products
deps:
	@-rm -rf vendor
	go mod tidy
	go mod verify
	go mod download
	go mod vendor

## lint: Static analysis
lint:
	# Go code
	golangci-lint run -v ./$(pkg)

## scan: Look for known vulnerabilities in the project dependencies
# https://github.com/sonatype-nexus-community/nancy
scan:
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
	@go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -mod=mod -m all 2> /dev/null

