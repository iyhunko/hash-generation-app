export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
proto/hash.go.proto
