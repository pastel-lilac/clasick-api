package di

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"github.com/tozastation/clasick-core/application/usecase"
	"github.com/tozastation/clasick-core/domain/service"
	"github.com/tozastation/clasick-core/infrastructure/persistence/repository"
	"github.com/tozastation/clasick-core/interface/handler/gorm"
	"github.com/tozastation/clasick-core/interface/util/guid"
	"github.com/tozastation/clasick-core/interface/util/hash"
	"github.com/tozastation/clasick-core/interface/util/jwt"
	"google.golang.org/grpc"
	"os"
	"time"
)

type diHandler struct {
	gorm.IGormHandler
}

type IDIHandler interface {
	CreateServer() *grpc.Server
	InitializeUser() usecase.IUserUseCase
}

func NewDIHandler(gorm gorm.IGormHandler) IDIHandler {
	return &diHandler{
		gorm,
	}
}

func (di *diHandler) CreateServer() *grpc.Server {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{})

	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	grpc_logrus.ReplaceGrpcLogger(logger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			//grpc_auth.UnaryServerInterceptor(middleware.AuthFunc),
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logger, opts...),
		)),
	)
	fmt.Println("[Done] Initialize gRPC Server")
	return server
}

func (di *diHandler) InitializeUser() usecase.IUserUseCase {
	db, err := di.IGormHandler.CreateConnection()
	if err != nil {
		panic(err)
	}
	jwtUtil := jwt.NewJWTUtil()
	hashUtil := hash.NewHashUtil()
	guidUtil := guid.NewGuidUtil()
	logger := logrus.New()
	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(repo, jwtUtil, hashUtil, guidUtil)
	uc := usecase.NewUserUseCase(srv, logger)
	return uc
}