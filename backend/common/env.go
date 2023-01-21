package common

import (
	"database/sql"

	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

var (
	MockModeActivated = false
)

var (
	ApiPort = 8080
	Router  = &gin.Engine{}
)

var (
	Database       = &sql.DB{}
	DatabaseUser   = ""
	DatabasePass   = ""
	DatabaseHost   = ""
	DatabaseName   = ""
	DatabasePort   = ""
	DatabaseSSL    = ""
	DatabaseURL    = ""
	DatabaseDriver = "postgres"
)

var (
	RabbitURL  = ""
	RabbitPort = ""

	RabbitConn    = &amqp.Connection{}
	RabbitChannel = &amqp.Channel{}

	RabbitQueueName = "cbmbr_queue"
	RabbitQueue     = &amqp.Queue{}
)

var (
	GrpcServer           = &grpc.Server{}
	GrpcConn             = &grpc.ClientConn{}
	AccountServiceClient = pb.NewAccountServiceClient(GrpcConn)
)
