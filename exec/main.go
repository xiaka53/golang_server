package main

import (
	"github.com/xiaka53/golang_server/pkg/pkg"
	"github.com/xiaka53/golang_server/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	accountServer := server.NewAccountServer()
	pkg.RegisterAccountServer(srv, accountServer)

	moneyServer := server.NewMoneyServer()
	pkg.RegisterMoneyServer(srv, moneyServer)

	recommenderServer := server.NewRecommenderServer()
	pkg.RegisterRecommenderServer(srv, recommenderServer)

	recordServer := server.NewRecordServer()
	pkg.RegisterRecordServer(srv, recordServer)

	vipServer := server.NewVipServer()
	pkg.RegisterVipServer(srv, vipServer)

	log.Println("gRPC server running on port 50051")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
