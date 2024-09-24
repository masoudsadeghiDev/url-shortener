build-react:
	cd ui && npm install && npm run build

build-go:
	go build -o url-shortener cmd/app/main.go

build: build-react build-go

run:
	go run cmd/app/main.go

docker:
	docker build -t url-shortener .

clean:
	go clean
	rm -f url-shortener

.PHONY: build-react build-go build run docker clean
