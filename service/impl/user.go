package impl

import (
	"context"
	"errors"
	"fmt"
	"mantis-user/models"
	userpb "mantis-user/protos"
)

type UserRpcServer struct {
	userModel *models.MembersModel
}

var (
	ErrNotFound = errors.New("用户不存在")
)

func NewUserRpcServer(userModel *models.MembersModel) *UserRpcServer {
	return &UserRpcServer{
		userModel: userModel,
	}
}

func (s *UserRpcServer) FindByToken(ctx context.Context, req *userpb.FindByTokenRequest, rsp *userpb.UserResponse) error {
	member, err := s.userModel.FindByToken(req.Token)
	if err != nil {
		return ErrNotFound
	}

	rsp.Token = member.Token
	rsp.Id = member.ID
	rsp.Password = member.Password
	return nil
}

func (s *UserRpcServer) FindByID(ctx context.Context, req *userpb.FindByIDRequest, rsp *userpb.UserResponse) error {
	fmt.Println("rilegoul", req.Id)
	member, err := s.userModel.FindByID(req.Id)
	if err != nil {
		return ErrNotFound
	}
	fmt.Println(member)
	rsp.Token = member.Token
	rsp.Id = member.ID
	rsp.Password = member.Password
	return nil
}
