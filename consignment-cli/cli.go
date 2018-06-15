package main

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	pb "shippy/proto/consignment"
)

const (
	ADDRESS           = "localhost:50051"
	DEFAULT_INTO_FILE = "consignment.json"
)

func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)

	if err != nil {
		return nil, err
	}

	return consignment, nil
}

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Connect error: %v", err)
	}

	defer conn.Close()

	client := pb.NewShippingServiceClient(conn)

	consignment, err := parseFile(DEFAULT_INTO_FILE)

	if err != nil {
		log.Fatalf("Parse file error %v", err)
	}

	resp, err := client.CreateConsignment(context.Background(), consignment)

	if err != nil {
		log.Fatalf("Create consignment error %v", err)
	}

	log.Printf("Created: %t", resp.Created)
}
