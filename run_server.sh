server=$1
echo $server

go run grpc-app/server_$server/server_$server.go