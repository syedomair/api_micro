package main

import (
	"context"
	"testing"
	"time"

	common "github.com/syedomair/api_micro/common"
	pb "github.com/syedomair/api_micro/public-service/proto"
	testdata "github.com/syedomair/api_micro/testdata"
	"google.golang.org/grpc/metadata"
)

func TestCreate(t *testing.T) {
	env := Env{repo: &mockDB{}, nats: &mockNATS{}, logger: common.GetLogger()}
	md := metadata.New(map[string]string{"authorization": testdata.TestValidPublicToken})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	user := &pb.User{
		FirstName: testdata.ValidFirstName,
		LastName:  testdata.ValidLastName,
		Email:     testdata.ValidEmail,
		Password:  testdata.ValidPassword}
	response, _ := env.Register(ctx, user)

	//TEST 1 correct authorization
	expected := "success"
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}

	md = metadata.New(map[string]string{"authorization": testdata.TestInValidPublicToken})
	ctx = metadata.NewIncomingContext(context.Background(), md)

	response, _ = env.Register(ctx, user)

	//TEST 2 incorrect authorization
	expected = "failure"
	if expected != response.Result {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, response.Result)
	}
}

type mockNATS struct {
}

func (mnats *mockNATS) PublishRegisterEvent(userId string, networkId string) error {
	return nil
}
func (mnats *mockNATS) PublishAuthEvent(userId string, token string) error {
	return nil
}

type mockDB struct {
}

func (mdb *mockDB) Create(user *pb.User, networkId string) (string, error) {
	return testdata.UserId, nil
}
func (mdb *mockDB) Authenticate(user *pb.LoginRequest, networkId string) (*pb.User, error) {
	return &pb.User{
		Id:        testdata.UserId,
		NetworkId: testdata.NetworkId,
		FirstName: testdata.ValidFirstName,
		LastName:  testdata.ValidLastName,
		Email:     testdata.ValidEmail,
		Password:  testdata.ValidPassword,
		IsAdmin:   testdata.IsAdmin,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339)}, nil
}
func (mdb *mockDB) GetNetworkFromApiKey(apiKey string) (*pb.Network, error) {
	return &pb.Network{
		Id:        testdata.NetworkId,
		Name:      testdata.NetworkName,
		ApiKey:    testdata.ApiKey,
		Secret:    testdata.ApiSecret,
		Status:    testdata.NetworkStatus,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339)}, nil

}

/*
func (mdb *mockDB) initMockDb() ([]*pb.User, error) {
	users := make([]*pb.User, 0)
	users = append(users, &pb.User{Id: "04b58e6e-f910-4ff0-83f1-27fbfa85dc2f", NetworkId: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", FirstName: "First Name 1", LastName: "Last Name 1", Email: "email1@gmail.com", Password: "123", IsAdmin: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})
	users = append(users, &pb.User{Id: "04b58e6e-f910-4ff0-83f1-27fbfa85dc2f", NetworkId: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", FirstName: "First Name 2", LastName: "Last Name 2", Email: "email2@gmail.com", Password: "123", IsAdmin: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})
	users = append(users, &pb.User{Id: "04b58e6e-f910-4ff0-83f1-27fbfa85dc2f", NetworkId: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", FirstName: "First Name 3", LastName: "Last Name 3", Email: "email3@gmail.com", Password: "123", IsAdmin: "1", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})
	return users, nil
}
*/
