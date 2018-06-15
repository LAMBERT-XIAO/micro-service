package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "shippy/proto/consignment"
)

const (
	PORT = ":50051"
)

// 定义一个仓库接口
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
}

// 实现上面的接口
type Repository struct {
	consignment []*pb.Consignment
}

// 实现Create方法
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignment = append(repo.consignment, consignment)
	return consignment, nil
}

func (repo *Repository) getAll() []*pb.Consignment {
	return repo.consignment
}

// 定义一个service操作repository，service中的方法跟proto文件中service是对应的
type service struct {
	repo Repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	log.Printf("Create a consignment")
	// 接收承运的货物
	consignment, err := s.repo.Create(req)

	if err != nil {
		return nil, err
	}

	resp := &pb.Response{Created: true, Consignment: consignment}
	return resp, nil
}

func main() {
	// 创建一个tcp的链接
	listenter, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("Failed to sever: %v", err)
	}

	log.Printf("Listen on: %s", PORT)

	// 创建一个grpc的server
	server := grpc.NewServer()
	repo := Repository{}

	// 将service对象代理到server上
	pb.RegisterShippingServiceServer(server, &service{repo})

	if err := server.Serve(listenter); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
