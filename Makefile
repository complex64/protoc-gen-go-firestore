all: lint test
build:
	go build ./...

# Run all tests.
t :test
test:
	go test ./...

# Build and run the aggregated linker.
l: lint
lint: build
	golangci-lint run ./...

# Lint the .proto files with buf.
b: buf
buf:
	cd proto && buf lint
	cd internal/gen/test && buf lint
	cd internal/require && buf lint

# Assumes $GOPATH/bin is in your $PATH!
gen: generate
generate: firestorepb install
	# Files used by tests of the plugin implementation.
	cd internal/gen/test && buf generate

	# Remove code generated from tests.
	find proto -name '*.go' -delete

	# Files used by tests of the internal packages.
	cd internal/require && buf generate

.PHONY: firestorepb
firestorepb:
	# Generate the standalone module and update/lock dependencies.
	cd proto && buf generate
	find proto -type f -name "*.pb.go" -exec mv {} firestorepb \;
	cd firestorepb && go mod tidy

# Install `protoc-gen-go` into $GOPATH/bin.
i: install
install:
	go install

# Remove all generated files.
clean:
	go clean
	find -name '*.pb.go' -delete
	$(MAKE) firestorepb

p: proto
.PHONY: proto
proto:
	buf lint proto
	cd internal/gen/test && buf lint
	cd internal/require && buf lint
