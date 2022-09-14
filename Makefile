run:
	go generate -mod=vendor ./...
	go run -mod=vendor ./cmd/server --config config.yaml
clean:
	rm -f ./server
test:
	go generate -mod=vendor ./...
	go test -v -mod=vendor ./...
build:
	go generate -mod=vendor ./...
	go build -mod=vendor -o burr ./cmd/server
docker:
	docker build -f Dockerfile.test --output type=tar,dest=/dev/null .
	docker build -f Dockerfile -t ghcr.io/eriner/burr:latest .
signoff:
	drone info
	drone lint
	drone sign eriner/burr --save

.PHONY: %
