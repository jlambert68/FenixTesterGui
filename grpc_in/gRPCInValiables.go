package grpc_in

import (
	fenixUserGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixUserGui/fenixUserGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type GRPCInStruct struct {
	logger *logrus.Logger
}

var GrpcIn GRPCInStruct

// gRPC variables
var (
	registerFenixUserGuiServer *grpc.Server
	lis                        net.Listener
)

// gRPC Server
type fenixUserGuiGrpcServicesServer struct {
	fenixUserGuiGrpcApi.UnimplementedFenixUserGuiGrpcServicesServer
}
