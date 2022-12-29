package main

import (
	"context"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	sampleV1 "github.com/reacher-jeon/grpc_server_sample/pb/sample/v1"
	typeV1 "github.com/reacher-jeon/grpc_server_sample/pb/type/v1"
	lr "github.com/reacher-jeon/grpc_server_sample/utils/logger"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	HOST = "localhost"
	PORT = 5000
)

type Server struct {
	sampleV1.UnimplementedSampleServiceServer
}

var (
	defaultServer *grpc.Server
	logger        *lr.Logger
)

func main() {
	//서버 부팅 시 로그 출력.
	logoPrint()

	//logger.Printf("Starting server...")
	address := HOST + ":" + strconv.Itoa(PORT)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
	go waitSignal(sigs)

	lis, listenErr := net.Listen("tcp", address)
	if listenErr != nil {
		log.Fatalf("failed to listen: %v", listenErr)
	}

	/*defaultServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpcMiddleware(),
			zeroLog.NewUnaryServerInterceptorWithLogger(logger.GetHandle()),
		)),
	)

	logger.Printf("server listening at %v", lis.Addr())
	if err := defaultServer.Serve(lis); err != nil {
		logger.Fatal().Msgf("failed to serve: %v", err)
	}*/

	grpcServer := grpc.NewServer()

	sampleService := Server{}
	sampleService.RegisterService(grpcServer)

	fmt.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}

	//sample.RegisterSampleServiceServer(defaultServer, &Server{})
	//RegisterService(defaultServer)

}

func (s *Server) RegisterService(mainServer *grpc.Server) error {
	//taskV1.RegisterTaskServiceServer(mainServer, &Server{})
	sampleV1.RegisterSampleServiceServer(mainServer, &Server{})
	fmt.Println("Register Success")

	return nil
}

func (s *Server) GetInfo(ctx context.Context, req *sampleV1.GetInfoInfoRequest) (*sampleV1.GetInfoResponse, error) {
	fmt.Println("===== GetInfo Start =====", ctx)

	requestHeader := req.GetHeader()
	fmt.Println("[Request] GetInfo Header:", requestHeader)

	requestBody := req.Info
	if requestBody != nil {
		fmt.Println("[Request] GetInfo Request Body:", requestBody)
	}

	responseData := &sampleV1.GetInfoResponse{
		Header: &typeV1.Header{
			Version:     "",
			ToIds:       requestHeader.ToIds,
			FromId:      typeV1.RequestNode_REQUEST_NODE_CUSTOM,
			RequesterId: requestHeader.RequesterId,
		},
		Error:   &typeV1.Error{},
		Message: "Hello gRPC",
	}

	return responseData, nil
}

func packageName() string {
	pc, _, _, _ := runtime.Caller(1)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	pkage := ""
	if parts[pl-2][0] == '(' {
		pkage = strings.Join(parts[0:pl-2], ".")
	} else {
		pkage = strings.Join(parts[0:pl-1], ".")
	}
	packagenames := strings.Split(pkage, "/")
	return packagenames[0]
}

func logoPrint() {
	//serviceLogo := `DX-Dev3 ` + packageName()
	serviceLogo := `gRPC Server Example`
	fmt.Println(figure.NewFigure(serviceLogo, "doom", true))
	//start time
	fmt.Println("Started at : ", time.Now().Format(time.RFC3339))
	//add excution evironment
	fmt.Println("Server Address : " + HOST + ":" + fmt.Sprintf("%v", PORT))
}

func grpcMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		debugString := fmt.Sprintf("%v", time.Now().Format(time.RFC3339))
		fmt.Println(debugString)
		return handler(ctx, req)
	}
}

func waitSignal(signals chan os.Signal) {
	s := <-signals
	fmt.Println("Got System signal:", s)
	shutdown()
}

func shutdown() {
	if defaultServer != nil {
		defaultServer.GracefulStop()
		fmt.Printf("Server shutdown at %v\n", time.Now().Format(time.RFC3339))
	}
}
