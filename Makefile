build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/shippy proto/consignment/consignment.proto
