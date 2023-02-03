#GOOS ?= darwin
#GOARCH ?= aarch64

golangci-lint ?= ~/go/bin/golangci-lint
upx ?= /usr/bin/upx
goyacc ?= ~/go/bin/goyacc

BUILD_DIR ?= bin
TEST_DIR ?= test-results

project_prefix := github.com/kendallgoto/skc
target_prefix := $(project_prefix)/cmd/
targets := $(sort $(notdir $(wildcard cmd/*)))

VERSION ?= $(shell git describe --always)

go_build_flags := -trimpath
go_build_flags += -o "$(BUILD_DIR)/"

go_test_flags :=  -race -v

.PHONY: build
build: $(targets)

.PHONY: test
test:
	go generate ./...
	mkdir -p $(TEST_DIR)
	go test $(go_test_flags) ./... -json | tee $(TEST_DIR)/$@.json | jq '. | select(.Action!="output") | select(.Action!="skip") | select(.Action!="run") | .Action + "\t" + .Package + "\t" + .Test' --raw-output

.PHONY: test-coverage
test-coverage:
	mkdir -p $(TEST_DIR)
	go test $(go_test_flags) -coverprofile=coverage.cov ./...

.PHONY: clean
clean:
	rm -f pkg/lang/parser/skc_*
	rm -f pkg/lang/parser/skcparser_*
	rm -f pkg/lang/parser/SkcParser.interp pkg/lang/parser/SkcParser.tokens
	rm -f pkg/lang/parser/SkcLexer.interp pkg/lang/parser/SkcLexer.tokens
	rm -rf bin/ test-results/

.PHONY: $(targets)
$(targets):
	go generate ./...
	go build $(go_build_flags) $(addprefix $(target_prefix),$@)

.PHONY: lint
lint:
	$(golangci-lint) run
