run:
	go run main.go

build:
	go build main.go

test-unit:
	GOCACHE=off go test -v -cover -race ./...

test-e2e:
	timeout 20s make run &
	cd _tests_/e2e && GOCACHE=off go test -v -cover -race

docker-build:
	docker build -t text-injestor .

docker-run:
	docker run -d -p 8080:8080 text-injestor