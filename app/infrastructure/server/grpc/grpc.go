package grpc

import (
	"fmt"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/infrastructure/proto/pb"
	"github.com/evenyosua18/oauth/app/infrastructure/server/grpc/middleware"
	accessTokenSvc "github.com/evenyosua18/oauth/app/infrastructure/server/grpc/service/accessToken"
	authenticationSvc "github.com/evenyosua18/oauth/app/infrastructure/server/grpc/service/authentication"
	endpointSvc "github.com/evenyosua18/oauth/app/infrastructure/server/grpc/service/endpoint"
	"github.com/evenyosua18/oauth/app/infrastructure/server/grpc/service/registration"
	"github.com/evenyosua18/oauth/app/usecase"
	accessTokenUC "github.com/evenyosua18/oauth/app/usecase/accessToken"
	authenticationUC "github.com/evenyosua18/oauth/app/usecase/authentication"
	endpointUC "github.com/evenyosua18/oauth/app/usecase/endpoint"
	registrationUC "github.com/evenyosua18/oauth/app/usecase/registration"
	"github.com/evenyosua18/oauth/config"
	"github.com/evenyosua18/oauth/util/logger"
	"github.com/evenyosua18/oauth/util/tracer"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() {
	/*start service*/

	//get config
	cfg := config.GetConfig().Server

	//setup jaeger
	jaegerModel := tracer.SetupJaeger{
		ServiceName: cfg.ServiceName,
		Endpoints:   cfg.Tracer.Endpoint,
	}

	jaegerModel.SetAttribute("environment", cfg.Tracer.Env)

	tp, err := tracer.New(jaegerModel)
	if err != nil {
		panic(err)
	}
	otel.SetTracerProvider(tp)

	//config grpc server
	var options []grpc.ServerOption

	maxIdle, err := time.ParseDuration(cfg.Grpc.MaxIdle + "h")

	if err != nil {
		panic(err)
	}

	maxAge, err := time.ParseDuration(cfg.Grpc.MaxAge + "m")

	if err != nil {
		panic(err)
	}

	maxAgeGrace, err := time.ParseDuration(cfg.Grpc.MaxAge + "s")

	if err != nil {
		panic(err)
	}

	options = append(options, grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     maxIdle,
		MaxConnectionAge:      maxAge,
		MaxConnectionAgeGrace: maxAgeGrace,
	}))

	options = append(options, grpc.UnaryInterceptor(middleware.ChainUnaryServer(middleware.PanicRecovery())))

	//create grpc server
	grpcServer := grpc.NewServer(options...)

	//setup dependency injection for use case (interaction)
	ctn := usecase.NewContainer()

	//register grpc server
	Apply(grpcServer, ctn)
	reflection.Register(grpcServer)

	//run grpc server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(`%s:%d`, cfg.Grpc.Host, cfg.Grpc.Port))

		if err != nil {
			logger.Log(logger.Fatal, "", err)
		}

		if err = grpcServer.Serve(lis); err != nil {
			logger.Log(logger.Fatal, "failed to start grpc server", err)
		}
	}()

	log.Println(fmt.Sprintf("grpc server is running at %s:%d", cfg.Grpc.Host, cfg.Grpc.Port))

	//get signal when server interrupted
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	log.Fatalf("process killed with signal: %s", sig.String())
}

func Apply(server *grpc.Server, ctn *usecase.Container) {
	pb.RegisterEndpointServiceServer(server, endpointSvc.NewServiceEndpoint(ctn.Resolve(string(constant.EndpointCTN)).(*endpointUC.InteractionEndpoint)))
	pb.RegisterAuthenticationServer(server, authenticationSvc.NewServiceAuthentication(ctn.Resolve(string(constant.AuthenticationCTN)).(*authenticationUC.InteractionAuthentication)))
	pb.RegisterAccessTokenServer(server, accessTokenSvc.NewServiceAccessToken(ctn.Resolve(string(constant.AccessTokenCTN)).(*accessTokenUC.InteractionAccessToken)))
	pb.RegisterRegistrationUserServiceServer(server, registration.NewServiceRegistration(ctn.Resolve(string(constant.RegistrationCTN)).(*registrationUC.InteractionRegistration)))
}
