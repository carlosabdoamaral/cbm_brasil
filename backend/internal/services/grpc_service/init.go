package grpcservice

import (
	"log"
	"net"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
	"google.golang.org/grpc"
)

type AccountServer struct {
	pb.UnimplementedAccountServiceServer
}

type OccurrenceServer struct {
	pb.UnimplementedOccurrenceServiceServer
}

type TutorialServer struct {
	pb.UnimplementedTutorialServiceServer
}

func Init() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		common.LogError(err.Error())
		return
	}

	common.GrpcServer = grpc.NewServer()
	pb.RegisterAccountServiceServer(common.GrpcServer, &AccountServer{})
	pb.RegisterOccurrenceServiceServer(common.GrpcServer, &OccurrenceServer{})
	pb.RegisterTutorialServiceServer(common.GrpcServer, &TutorialServer{})

	log.Printf("Server listening on %v", lis.Addr())
	if err := common.GrpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
