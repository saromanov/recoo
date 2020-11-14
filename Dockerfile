FROM golang:1.15-buster
RUN ls -la
ADD . /app
WORKDIR /app
RUN ls -la
RUN go mod download
RUN go build -o /bin/app ./cmd/main.go
ENTRYPOINT [ /bin/app ]