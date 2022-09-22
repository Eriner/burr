web:
	go run -mod=vendor ./cmd/server --config config.yaml
worker:
	go run -mod=vendor ./cmd/server worker --config config.yaml
clean:
	rm -f ./server
test:
	go test -v -mod=vendor ./...
build:
	go build -mod=vendor -o burr ./cmd/server
generate:
	go generate -mod=mod ./...
licenses:
	go generate -tags=licenses ./licenses.go
release:
	go generate -mod=vendor  ./...
	go test -v -mod=vendor  ./...
	# go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...
	git diff-index --quiet --cached HEAD --
	git diff-files --quiet
	git ls-files --others --exclude-standard
	go build -mod=vendor -o burr  ./cmd/server
docker:
	docker build -f Dockerfile.test --output type=tar,dest=/dev/null .
	docker build -f Dockerfile -t ghcr.io/eriner/burr:latest .
signoff:
	drone info
	drone lint
	drone sign eriner/burr --save

.PHONY: % *
