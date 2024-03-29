package main

import (
	"app/main/internal/storage"
	"app/main/internal/utils"
	"log"
	"net"

	"github.com/dimon5360/SportTechProtos/gen/go/proto"
	"google.golang.org/grpc"
)

const (
	configPath  = "../SportTechDockerConfig/"
	serviceEnv  = ".env"
	postgresEnv = configPath + "postgres.env"
)

func main() {

	utils.Env().Load(serviceEnv, postgresEnv)

	log.Println("SportTech profile service v." + utils.Env().Value("SERVICE_VERSION"))

	service := storage.CreateService()
	service.Init()

	lis, err := net.Listen("tcp", utils.Env().Value("PROFILE_GRPC_HOST"))
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterProfileUsersServiceServer(grpcServer, service)
	grpcServer.Serve(lis)
}
