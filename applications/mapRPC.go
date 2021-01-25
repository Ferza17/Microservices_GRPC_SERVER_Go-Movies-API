package applications

import (
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/controllers/user"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/protos/server/user_proto"
	"google.golang.org/grpc"
)

func mapRPC(s *grpc.Server) {
	// Register RPC
	user_proto.RegisterUserServiceServer(s, &user.Server{})
}
