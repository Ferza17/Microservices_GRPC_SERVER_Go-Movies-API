package applications

import (
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/logger_utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartApplication() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	mapRPC(server)
	reflection.Register(server)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logger_utils.Error("Error while connect to tcp", err)
	}

	if err := server.Serve(lis); err != nil {
		logger_utils.Error("Unable to serve ", err)
	}


}
