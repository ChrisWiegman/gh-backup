PKG          := github.com/ChrisWiegman/gh-backup
VERSION      := $(shell git describe --tags || echo "0.0.1")
TIMESTAMP    := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

.PHONY: install
install:
	go mod vendor
	go install \
		-ldflags "-s -w -X $(PKG)/internal/backup.Version=$(VERSION) -X $(PKG)/internal/backup.Timestamp=$(TIMESTAMP)" \
		./cmd/...
