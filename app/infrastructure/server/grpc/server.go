package grpc

import (
	"fmt"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/infrastructure/server/grpc/proto/pb"
	endpointSvc "github.com/evenyosua18/oauth/app/infrastructure/server/grpc/service/endpoint"
	oauthSvc "github.com/evenyosua18/oauth/app/infrastructure/server/grpc/service/oauth"
	"github.com/evenyosua18/oauth/app/usecase"
	accessTokenUC "github.com/evenyosua18/oauth/app/usecase/accessToken"
	endpointUC "github.com/evenyosua18/oauth/app/usecase/endpoint"
	"github.com/evenyosua18/oauth/config"
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

	//setup jaeger
	jaegerModel := tracer.SetupJaeger{
		ServiceName: "oauth-service",
		Endpoints:   "http://localhost:14268/api/traces",
	}

	jaegerModel.SetAttribute("environment", "development")

	tp, err := tracer.New(jaegerModel)
	if err != nil {
		panic(err)
	}
	otel.SetTracerProvider(tp)

	//config grpc server
	var options []grpc.ServerOption
	options = append(options, grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     1 * time.Hour,
		MaxConnectionAge:      5 * time.Minute,
		MaxConnectionAgeGrace: 5 * time.Second,
	}))

	//create grpc server
	grpcServer := grpc.NewServer(options...)

	//setup dependency injection for use case (interaction)
	ctn := usecase.NewContainer()

	//register grpc server
	Apply(grpcServer, ctn)
	reflection.Register(grpcServer)

	//run grpc server
	grpcConfig := config.GetConfig().Server.Grpc
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(`%s:%d`, grpcConfig.Host, grpcConfig.Port))

		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()

	log.Println(fmt.Sprintf("grpc server is running at %s:%d", grpcConfig.Host, grpcConfig.Port))

	//get signal when server interrupted
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	log.Fatalf("process killed with signal: %s", sig.String())
}

func Apply(server *grpc.Server, ctn *usecase.Container) {
	pb.RegisterEndpointServiceServer(server, endpointSvc.NewServiceEndpoint(ctn.Resolve(string(constant.EndpointCTN)).(*endpointUC.InteractionEndpoint)))
	pb.RegisterAuthenticationServer(server, oauthSvc.NewServiceAuthentication())
	pb.RegisterAccessTokenServer(server, oauthSvc.NewServiceAccessToken(ctn.Resolve(string(constant.AccessTokenCTN)).(*accessTokenUC.InteractionAccessToken)))
}
