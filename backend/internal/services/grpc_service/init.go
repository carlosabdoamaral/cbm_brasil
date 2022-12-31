package grpcservice

import (
	"flag"
	"log"
	"net"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func ConnectToGrpcServer() *grpc.ClientConn {
	flag.Parse()

	addr := flag.String("addr", "localhost:50051", "the address to connect to")
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to %v: %v", addr, err)
	}

	common.GrpcConn = conn
	common.AccountServiceClient = pb.NewAccountServiceClient(conn)
	common.OccurrenceServiceClient = pb.NewOccurrenceServiceClient(conn)
	common.TutorialServiceClient = pb.NewTutorialServiceClient(conn)

	return conn
}
