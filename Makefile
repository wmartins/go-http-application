BINARY_NAME=main
VERSION=2.0.0

all: go-build docker-build

docker-build:
	docker build -t wmartins/go-http-application:$(VERSION) .

go-build:
	CGO_ENABLED=0 GOOS=linux \
	go build -a -tags netgo -ldflags '-w' -o $(BINARY_NAME) .

clean:
	rm -rf main
