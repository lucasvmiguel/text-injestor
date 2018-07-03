run:
	go run main.go

build:
	go build main.go

test:
	go test -v -cover -race ./...

docker-build:
	docker build -t text-injestor .

docker-run:
	docker run -d -p 8080:8080 text-injestor