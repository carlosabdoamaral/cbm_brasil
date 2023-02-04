package grpc

import (
	"fmt"
	"net"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
	"google.golang.org/grpc"
)

type AccountServer struct {
	pb.UnimplementedAccountServiceServer
}

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

type OccurrenceServer struct {
	pb.UnimplementedOccurreceServiceServer
}

func ServeGrpcServer() {
	fmt.Println("[*] Starting GRPC")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("[!] %s", err.Error())
		return
	}

	common.GrpcServer = grpc.NewServer()
	pb.RegisterAccountServiceServer(common.GrpcServer, &AccountServer{})
	pb.RegisterAuthServiceServer(common.GrpcServer, &AuthServer{})
	pb.RegisterOccurreceServiceServer(common.GrpcServer, &OccurrenceServer{})

	fmt.Printf("[*] Server listening on %v", lis.Addr())
	err = common.GrpcServer.Serve(lis)
	if err != nil {
		fmt.Printf("[!] Failed to serve, error: %s", err.Error())
		return
	}
}
