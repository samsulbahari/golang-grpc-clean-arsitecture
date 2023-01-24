package main

import (
	"fmt"
	"grpc/config/database"
	"grpc/customer/delivery"
	customerRepo "grpc/customer/repository"
	customerService "grpc/customer/service"
	"grpc/middleware"
	"net"

	"google.golang.org/grpc"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

const (
	port = ":50051"
)

func main() {

	// //testproduct := &pb.Products{}
	// for _, products := range products.Data {
	// 	fmt.Println(products.GetCategory().GetId())
	// }

	db := database.Connectdb()
	cr := customerRepo.NewCustomerRepository(db)
	customerService := customerService.NewCustomerService(cr)
	netlister, _ := net.Listen("tcp", port)
	gcpserver := grpc.NewServer(grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(middleware.ExampleAuthFunc)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(middleware.ExampleAuthFunc)))

	delivery.NewCustomerHandler(customerService, gcpserver)
	err := gcpserver.Serve(netlister)
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}

}
