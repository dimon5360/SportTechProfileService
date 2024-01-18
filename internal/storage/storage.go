package storage

import (
	"app/main/utils"
	"context"
	"database/sql"
	"fmt"
	"log"

	proto "github.com/dimon5360/SportTechProtos/gen/go/proto"
	_ "github.com/lib/pq"
)

type ProfileUsersService struct {
	proto.UnimplementedProfileUsersServiceServer

	db *sql.DB
}

func CreateService() *ProfileUsersService {
	return &ProfileUsersService{}
}

func (s *ProfileUsersService) Init() {

	var connString string

	env := utils.Env()

	connString = fmt.Sprintf("postgresql://%s:%s@%s/%s",
		env.Value("POSTGRES_USER"),
		env.Value("POSTGRES_PASSWORD"),
		env.Value("POSTGRES_HOST"),
		env.Value("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", connString+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	s.db = db
}

func (s *ProfileUsersService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserInfoResponse, error) {
	return &proto.UserInfoResponse{}, nil
}

func (s *ProfileUsersService) AuthUser(ctx context.Context, req *proto.AuthUserRequest) (*proto.UserInfoResponse, error) {
	return &proto.UserInfoResponse{}, nil
}

func (s *ProfileUsersService) CreateUser(ctx context.Context, req *proto.CreateUserRequst) (*proto.UserInfoResponse, error) {
	return &proto.UserInfoResponse{}, nil
}
