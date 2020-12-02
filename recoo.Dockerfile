FROM golang:1.15-alpine
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o /bin/app ./cmd/main.go
ENTRYPOINT [ /bin/app ]