HOSTNAME=marcus.carey
NAMESPACE=terraform
NAME=postmark
VERSION=1.0

OS_NAME:=$(shell uname -s | tr ‘[:upper:]’ ‘[:lower:]’)
HW_CLASS:=$(shell uname -m)
OS_ARCH=${OS_NAME}_${HW_CLASS}

BINARY=terraform-provider-${NAME}
PLUGIN_DIR=${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
WEBSITE_REPO=github.com/hashicorp/terraform-website
PKG_NAME=github

default: build

tools:
	go install github.com/client9/misspell/cmd/misspell@v0.3.4
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1

build: fmtcheck
	go build ./...

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

lint:
	@echo "==> Checking source code against linters..."
	golangci-lint run ./...

test:
	go test ./...
	# commenting this out for release tooling, please run testacc instead

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-lint:
	@echo "==> Checking website against linters..."
	@misspell -error -source=text website/

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: build test testacc vet fmt fmtcheck lint tools test-compile website website-lint website-test

clean:
	@echo "▶️ Removing the Terraform plugin"
	rm -rf ~/.terraform.d/plugins/${PLUGIN_DIR}
	rm -rf examples/.terraform* || true

install: clean
	@echo "▶️ Building the Terraform binary file"
	go build -gcflags="all=-N -l" -o ${BINARY}
	@echo "▶️ Adding the binary to the plugin directory"
	mkdir -p ~/.terraform.d/plugins/${PLUGIN_DIR}
	mv ${BINARY} ~/.terraform.d/plugins/${PLUGIN_DIR}
	@echo "▶️ Build executed successfully"

#generate:
#	go generate ./...
