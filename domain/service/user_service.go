package service

import (
	"context"
	irepo "github.com/tozastation/clasick-core/domain/repository"
	"github.com/tozastation/clasick-core/infrastructure/persistence/model/db"
	rpcuser "github.com/tozastation/clasick-core/interface/rpc/user"
	"github.com/tozastation/clasick-core/interface/util/guid"
	"github.com/tozastation/clasick-core/interface/util/hash"
	"github.com/tozastation/clasick-core/interface/util/jwt"
	"os"
)

type IUserService interface {
	SignIn(ctx context.Context, req *rpcuser.RequestSignIn) (*rpcuser.ResponseSignIn, error)
	SignUp(ctx context.Context, req *rpcuser.RequestSignUp) (*rpcuser.ResponseSignUp, error)
	GetSingleUser(ctx context.Context, req *rpcuser.RequestGetSingleUser) (*rpcuser.ResponseGetSingleUser, error)
}

type userService struct {
	irepo.IUserRepository
	jwt.IJWTUtil
	hash.IHashUtil
	guid.IGuidUtil
}

func NewUserService(repo irepo.IUserRepository, jwt jwt.IJWTUtil, hash hash.IHashUtil, guid guid.IGuidUtil) IUserService {
	return &userService{
		repo,
		jwt,
		hash,
		guid,
	}
}

func (srv *userService) SignIn(ctx context.Context, req *rpcuser.RequestSignIn) (*rpcuser.ResponseSignIn, error) {
	resultSelectUser, err := srv.IUserRepository.SelectExistUser(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	err = srv.IHashUtil.CheckHash(resultSelectUser.Password, req.Password)
	if err != nil {
		return nil, err
	}
	result := &rpcuser.ResponseSignIn{
		AccessToken: resultSelectUser.AccessToken,
	}
	return result, nil
}

func (srv *userService) SignUp(ctx context.Context, req *rpcuser.RequestSignUp) (*rpcuser.ResponseSignUp, error) {
	pwHashed, err := srv.IHashUtil.CreateHash(req.Password)
	if err != nil {
		return nil, err
	}
	claim := jwt.Claim{
		Issuer: os.Getenv("ISSUER"),
		Subject: os.Getenv("SUBJECT"),
		Audience: os.Getenv("AUDIENCE"),
		ExpirationTime: 3600,//uint32(os.Getenv("EXPIRATION_TIME")),
		NotBefore: 3600,
		IssuedAt: 3600,
		JWTID: srv.IGuidUtil.CreateGuid(),
		Type: os.Getenv("TYPE"),
	}
	accessToken, err := srv.IJWTUtil.CreateAccessToken(claim)
	if err != nil {
		return nil, err
	}
	dbUser := &db.User{
		Name: req.Name,
		Password: pwHashed,
		Email: "",
		PhoneNum: "",
		AccessToken: accessToken,
		IconPath: "",
		Introduction: "",
	}
	err = srv.IUserRepository.InsertUser(ctx, dbUser)
	if err != nil {
		return nil, err
	}
	result := &rpcuser.ResponseSignUp{
		AccessToken: accessToken,
	}
	return result, nil
}

func (srv *userService) GetSingleUser(ctx context.Context, req *rpcuser.RequestGetSingleUser) (*rpcuser.ResponseGetSingleUser, error) {
	resultSelectUser, err := srv.IUserRepository.SelectUserFromID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	result := &rpcuser.ResponseGetSingleUser{
		User: &rpcuser.ResponseUser{
			UserId: int32(resultSelectUser.ID),
			UserName: resultSelectUser.Name,
			IconPath: resultSelectUser.IconPath,
			Introduction: resultSelectUser.Introduction,
		},
	}
	return result, nil
}