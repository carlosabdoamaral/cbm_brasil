package common

import (
	"database/sql"

	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

// DB_NAME        string = "cbmbr"
var (
	DB_USER          string = "postgres"
	DB_PASS          string = "postgres"
	DB_HOST          string = "localhost"
	DB_NAME          string = "postgres"
	DB_PORT          string = "5432"
	DB_SSL           string = "disable"
	DB_DRIVER        string = "postgres"
	RABBIT_URL       string = "amqp://guest:guest@localhost:5672"
	OCCURRENCE_QUEUE string = "occurrenceQueue"
	MOCK_MODE        string = ""
)

var (
	Database = &sql.DB{}
)

var (
	RabbitConn      = &amqp.Connection{}
	RabbitChannel   = &amqp.Channel{}
	RabbitQueue     = &amqp.Queue{}
	RabbitQueueName = ""
)

var (
	GrpcServer              = &grpc.Server{}
	GrpcConn                = &grpc.ClientConn{}
	AccountServiceClient    = pb.NewAccountServiceClient(GrpcConn)
	TutorialServiceClient   = pb.NewTutorialServiceClient(GrpcConn)
	OccurrenceServiceClient = pb.NewOccurrenceServiceClient(GrpcConn)
)
