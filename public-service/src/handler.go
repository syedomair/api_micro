package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	jwt "github.com/dgrijalva/jwt-go"
	nats "github.com/nats-io/go-nats"
	"github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/public-service/proto"
)

type service struct {
	repo Repository
	nats *nats.Conn
}

func (s *service) Register(ctx context.Context, req *pb.User) (*pb.Response, error) {

	networkId := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"
	userId, err := s.repo.Create(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	responseUserId := map[string]string{"user_id": userId}

	subject := "User.UserCreated"
	err = s.nats.Publish(subject, []byte("Hello NATS"))
	if err != nil {
		log.Printf("Error during publishing: %s", err)
	}
	s.nats.Flush()
	return &pb.Response{Result: common.SUCCESS, Data: responseUserId, Error: nil}, err
}

func (s *service) Authenticate(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {

	networkId := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"

	user, err := s.repo.Authenticate(req, networkId)
	if err != nil {
		return &pb.Response{Result: common.FAILURE, Data: nil, Error: common.DatabaseError()}, nil
	}
	fmt.Println(user)

	type Claims struct {
		CurrentUserId string `json:"current_user_id"`
		NetworkId     string `json:"network_id"`
		IsAdmin       string `json:"is_admin"`
		jwt.StandardClaims
	}

	claims := Claims{
		user.Id,
		user.NetworkId,
		user.IsAdmin,
		jwt.StandardClaims{
			Issuer: "MEEM",
		},
	}

	signingKey := []byte(common.SIGNING_KEY)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, _ := token.SignedString(signingKey)

	tokenStr := map[string]string{"token": signedJwtToken}
	return &pb.Response{Result: common.SUCCESS, Data: tokenStr, Error: nil}, nil
}
