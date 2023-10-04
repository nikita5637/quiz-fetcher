LOCAL_BIN:=$(CURDIR)/bin
GOIMPORTS_BIN:=$(LOCAL_BIN)/goimports
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
XO_BIN:=$(LOCAL_BIN)/xo

.PHONY: .install-bin-deps
.install-bin-deps:
ifeq ($(wildcard $(GOIMPORTS_BIN)),)
	$(info Installing binary dependency goimports)
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports 
endif

.PHONY: .install-lint
.install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint)
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint
	GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

.PHONY: .install-xo
.install-xo:
ifeq ($(wildcard $(XO_BIN)),)
	$(info Downloading xo)
	GOBIN=$(LOCAL_BIN) go install github.com/xo/xo
	XO_BIN:=$(LOCAL_BIN)/xo
endif

.PHONY: build
build:
	go build -o fetcher ./cmd/fetcher

.PHONY: clean
clean:
	rm fetcher

.PHONY: gen-db
gen-db: .install-xo
	$(XO_BIN) schema \
	--schema="" \
	--go-import="github.com/nikita5637/quiz-fetcher/internal/pkg/logger" \
	--go-import="github.com/nikita5637/quiz-fetcher/internal/pkg/tx" \
	--out internal/pkg/storage/mysql \
	--src templates \
	--template go \
	mysql://$(USERNAME):$(DATABASE_PASSWORD)@$(DBADDR)/$(DBNAME)
	rm -rf ./internal/pkg/storage/mysql/goosedbversion.xo.go ./internal/pkg/storage/mysql/db.xo.go

.PHONY: generate
generate: .install-bin-deps
	$(GOIMPORTS_BIN) -w ./

.PHONY: go-generate
go-generate:
	go generate ./...

.PHONY: lint
lint: .install-lint
	$(info Running lint...)
	$(GOLANGCI_BIN) run --config=.golangci.pipeline.yaml ./...

.PHONY: migrations
migrations:
	goose -dir migrations/ mysql "$(USERNAME):$(DATABASE_PASSWORD)@tcp($(DBADDR))/$(DBNAME)" up

.PHONY: run
run:
	go run ./cmd/fetcher --config ./config.toml

.PHONY: test
test:
	go test -v ./...
