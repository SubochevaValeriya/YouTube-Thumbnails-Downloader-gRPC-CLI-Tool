build:
	docker-compose build YouTube-Thumbnails-Downloader

run:
	docker-compose up YouTube-Thumbnails-Downloader

test:
	go test -v ./...

lint:
	golangci-lint run