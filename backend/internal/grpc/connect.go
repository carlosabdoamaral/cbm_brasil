package grpc

import (
	"flag"
	"log"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectToGRPCServer() *grpc.ClientConn {
	var Addr = flag.String("addr", "localhost:50051", "the address to connect to")

	flag.Parse()

	conn, err := grpc.Dial(*Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to %v: %v", Addr, err)
	}

	common.GrpcConn = conn
	common.AccountServiceClient = pb.NewAccountServiceClient(conn)

	return conn
}
