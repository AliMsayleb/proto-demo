set -xe
go generate
rm -rf ./rpc/*.pb.go
protoc --go_out=./ --go-grpc_out=./ ./proto/*.proto
go build -o ./main main.go
