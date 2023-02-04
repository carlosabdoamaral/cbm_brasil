## What should I do when creating a new proto?

<br/><hr/><br/>

#### **Files**
Nos arquivos abaixo, você deve adicionar algo similar a:
`common.AccountServiceClient = pb.NewAccountServiceClient(conn)`

1. backend/internal/grpc/connect.go
1. backend/internal/grpc/serve.go