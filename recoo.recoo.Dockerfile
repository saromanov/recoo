FROM golang:1.15-alpine
ADD . /app
WORKDIR /app
RUN apk add --upgrade && apk add curl openssl
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GARCH=amd64
RUN go mod download
RUN go build -o /bin/app ./cmd/main.go
CMD [ "/bin/app" ]