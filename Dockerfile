FROM golang:1.18-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o youtube-thumbnails-downloader ./server/server.go

CMD ["./youtube-thumbnails-downloader"]