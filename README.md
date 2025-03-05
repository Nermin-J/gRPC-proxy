# gRPC-proxy

#### Install the `protoc` plugins for Go:
    > go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    > go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

#### Update your `PATH` so that the `protoc` compiler can find the plugins:
    > export PATH="$PATH:$(go env GOPATH)/bin"

### nginx
    > nginx -t 
    > brew services restart nginx