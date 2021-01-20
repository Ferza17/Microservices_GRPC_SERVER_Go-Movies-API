package applications

import (
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/controllers/user_controller"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/protos/user_proto"
	"google.golang.org/grpc"
)

func mapRPC(s *grpc.Server) {
	// Register RPC
	user_proto.RegisterUserServiceServer(s, &user_controller.Server{})
}
