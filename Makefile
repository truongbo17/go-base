.PHONY: build
# build the binary
build:
	go build -v -o build/ggb main.go

.PHONY: format
# Run go fmt against code
format:
	go fmt ./...

.PHONY: fmt
# fmt is an alias for format
fmt: format

.PHONY: cache
# clean cache
clean:
	rm -rf build/ggb

.PHONY: swagger
# swagger docs generation
swagger:
	swag init

.PHONY: tidy
# tidy the go modules
tidy:
	go mod tidy

