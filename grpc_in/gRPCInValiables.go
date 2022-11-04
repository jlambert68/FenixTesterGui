package grpc_in

import (
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
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
	registerFenixGuiBuilderServer *grpc.Server
	lis                           net.Listener
)

// gRPC Server used for register clients Name, Ip and Port and Clients Test Environments and Clients Test Commands
type fenixGuiTestCaseBuilderGrpcServicesServer struct {
	fenixGuiTestCaseBuilderServerGrpcApi.UnimplementedFenixTestCaseBuilderServerGrpcServicesServer
}
