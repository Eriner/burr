run:
	go run -mod=vendor ./cmd/server --config config.yaml
clean:
	rm -f ./server
test:
	go test -v -mod=vendor ./...
build:
	go build -mod=vendor -o burr ./cmd/server
generate:
	go generate -mod=mod ./...
docker:
	docker build -f Dockerfile.test --output type=tar,dest=/dev/null .
	docker build -f Dockerfile -t ghcr.io/eriner/burr:latest .
signoff:
	drone info
	drone lint
	drone sign eriner/burr --save

.PHONY: %
