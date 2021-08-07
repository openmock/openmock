GOLANGCILINT := $(shell command -v golangci-lint 2> /dev/null)

deps:
	@go install github.com/go-swagger/go-swagger/cmd/swagger@v0.27.0
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1

.PHONY: vendor
vendor:
	@go mod tidy
	@go mod vendor

build: build_omctl build_swagger

build_omctl:
	@go build -o $(PWD)/omctl github.com/openmock/openmock/cmd/omctl

build_swagger:
	@echo "Building OpenMock Server to $(PWD)/om-swagger ..."
	@go build -o $(PWD)/om github.com/openmock/openmock/swagger_gen/cmd/open-mock-server

test: lint
	@go test -race -covermode=atomic . ./pkg/admin ./pkg/evaluator

run: build
	OPENMOCK_TEMPLATES_DIR=./demo_templates ./om

lint:
	@golangci-lint run -D errcheck .

#################
# Swagger stuff #
#################

PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)

gen: api_docs swagger

api_docs:
	@echo "Installing swagger-merger" && npm install swagger-merger -g
	@swagger-merger -i $(PWD)/swagger/index.yaml -o $(PWD)/docs/api_docs/bundle.yaml

verify_swagger:
	@echo "Running $@"
	@swagger validate $(PWD)/docs/api_docs/bundle.yaml

# list of files that contain custom edits and shouldn't be overwritten by generation
PROTECTED_FILES := restapi/configure_open_mock.go restapi/server.go cmd/open-mock-server/main.go

swagger: verify_swagger
	@echo "Regenerate swagger files"
	@for file in $(PROTECTED_FILES); do \
		echo $$file ; \
		rm -f /tmp/`basename $$file`; \
		cp $(PWD)/swagger_gen/$$file /tmp/`basename $$file` 2>/dev/null ; \
	done
	@rm -f /tmp/configure_open_mock.go
	@cp $(PWD)/swagger_gen/restapi/configure_open_mock.go /tmp/configure_open_mock.go 2>/dev/null || :
	@rm -rf $(PWD)/swagger_gen
	@mkdir $(PWD)/swagger_gen
	@swagger generate server -t ./swagger_gen -f $(PWD)/docs/api_docs/bundle.yaml
	@for file in $(PROTECTED_FILES); do \
		cp /tmp/`basename $$file` $(PWD)/swagger_gen/$$file 2>/dev/null ; \
	done
