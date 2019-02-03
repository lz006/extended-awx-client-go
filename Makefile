all: verify check build examples
.PHONY: all

verify:
	hack/verify.sh
.PHONY: verify

check: verify
	go test ./eawx
.PHONY: check

build:
	go build ./eawx
.PHONY: build

examples: build
	hack/build-examples.sh
.PHONY: examples
