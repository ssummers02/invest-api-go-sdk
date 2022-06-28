
all:
	protoc -I=./contracts --go_out=plugins=grpc:investapi contracts/*.proto
