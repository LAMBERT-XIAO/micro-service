build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shippy proto/consignment/consignment.proto
