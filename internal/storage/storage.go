package storage

import (
	"app/main/internal/utils"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/dimon5360/SportTechProtos/gen/go/proto"
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func invalidProfileResponse(err error) (*proto.UserProfileResponse, error) {
	return &proto.UserProfileResponse{}, err
}

func (s *ProfileUsersService) GetProfile(ctx context.Context, req *proto.GetProfileRequest) (*proto.UserProfileResponse, error) {

	profile, err := s.GetProfileByIdFromDatabase(req.UserId)
	if err != nil {
		return invalidProfileResponse(err)
	}

	return &proto.UserProfileResponse{
		Id:        profile.Id,
		Username:  profile.Username,
		Firstname: profile.Firstname,
		Lastname:  profile.Lastname,
		CreatedAt: timestamppb.New(profile.Created_at),
		UpdatedAt: timestamppb.New(profile.Updated_at),
	}, nil
}

func (s *ProfileUsersService) CreateProfile(ctx context.Context, req *proto.CreateProfileRequst) (*proto.UserProfileResponse, error) {

	profile, err := s.AddProfileToDatabase(req.Username, req.Firstname, req.Lastname, req.UserId)
	if err != nil {
		return invalidProfileResponse(fmt.Errorf("%s: %v", "Failed handling profile data", err))
	}

	return &proto.UserProfileResponse{
		Id:        profile.Id,
		Username:  profile.Username,
		Firstname: profile.Firstname,
		Lastname:  profile.Lastname,
	}, err
}
