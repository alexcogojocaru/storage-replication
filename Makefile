all:
	@protoc -I=proto --go_out=. ./proto/request.proto
	@protoc -I=proto --go-grpc_out=. ./proto/request.proto
