FROM golang:alpine
RUN apk update && apk add git
RUN apk --update add ca-certificates git
RUN apk add build-base

WORKDIR /app
COPY ../go.mod ./
COPY ../go.sum ./
RUN go mod download
COPY . .

RUN  CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o bin/server grpc-app/server.go
EXPOSE 50051

CMD ["/app/bin/server"]