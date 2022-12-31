package main

import (
	"flag"
	"log"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/database"
	grpcservice "github.com/carlosabdoamaral/cbm_brasil/backend/internal/services/grpc_service"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/utils"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	utils.ReadEnvFile()
	database.Connect()
	ConnectToGrpcServer()
	grpcservice.Init()
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
