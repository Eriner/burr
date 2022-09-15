run-dev:
	go run -mod=vendor -tags=sqlite ./cmd/server --config config.yaml
clean:
	rm -f ./server
test:
	go test -v -mod=vendor -tags=sqlite ./...
build:
	go build -mod=vendor -o burr -tags=sqlite ./cmd/server
generate:
	go generate -mod=mod -tags=sqlite ./...
release:
	go generate -mod=vendor -tags=postgres,redis ./...
	go test -v -mod=vendor -tags=postgres,redis ./...
	git diff-index --quiet --cached HEAD --
	git diff-files --quiet
	git ls-files --others --exclude-standard
	go build -mod=vendor -o burr -tags=postgres,redis ./...
docker:
	docker build -f Dockerfile.test --output type=tar,dest=/dev/null .
	docker build -f Dockerfile -t ghcr.io/eriner/burr:latest .
signoff:
	drone info
	drone lint
	drone sign eriner/burr --save

.PHONY: %
