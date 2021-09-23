run:
	protoc --proto_path=./pb --proto_path=. --go_out=plugins=grpc:. ./pb/user.proto
	protoc --proto_path=./pb --proto_path=. --grpc-gateway_out=logtostderr=true:. ./pb/user.proto