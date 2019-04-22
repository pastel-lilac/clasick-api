package usecase

import (
	"context"
	"github.com/sirupsen/logrus"
	isrv "github.com/tozastation/clasick-core/domain/service"
	rpcuser "github.com/tozastation/clasick-core/interface/rpc/user"
)

type IUserUseCase interface {
	SignIn(ctx context.Context, req *rpcuser.RequestSignIn) (*rpcuser.ResponseSignIn, error)
	SignUp(ctx context.Context, req *rpcuser.RequestSignUp) (*rpcuser.ResponseSignUp, error)
	GetSingleUser(ctx context.Context, req *rpcuser.RequestGetSingleUser) (*rpcuser.ResponseGetSingleUser, error)
}

type userUseCase struct {
	isrv.IUserService
	*logrus.Logger
}

func NewUserUseCase(srv isrv.IUserService, logger *logrus.Logger) IUserUseCase {
	return &userUseCase{
		srv,
		logger,
	}
}

func (uc *userUseCase) SignIn(ctx context.Context, req *rpcuser.RequestSignIn) (*rpcuser.ResponseSignIn, error) {
	return uc.IUserService.SignIn(ctx, req)
}

func (uc *userUseCase) SignUp(ctx context.Context, req *rpcuser.RequestSignUp) (*rpcuser.ResponseSignUp, error) {
	return uc.IUserService.SignUp(ctx, req)
}

func (uc *userUseCase) GetSingleUser(ctx context.Context, req *rpcuser.RequestGetSingleUser) (*rpcuser.ResponseGetSingleUser, error) {
	return uc.IUserService.GetSingleUser(ctx, req)
}
