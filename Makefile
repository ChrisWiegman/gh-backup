.PHONY: install
install:
	go mod vendor
	go install \
		-ldflags "-s -w -X $(PKG)/internal/cmd.Version=$(VERSION) -X $(PKG)/internal/cmd.Timestamp=$(TIMESTAMP)" \
		./cmd/...
