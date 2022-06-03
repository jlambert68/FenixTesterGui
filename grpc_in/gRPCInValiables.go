package grpc_in

import (
	"FenixTesterGui/fenixTestGuiObject"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type GRPCInStruct struct {
	Logger                        *logrus.Logger
	FenixTesterGuiObjectReference *fenixTestGuiObject.FenixGuiBuilderProxyServerObjectStruct
}

var GrpcIn GRPCInStruct

// gRPC variables
var (
	registerfenixGuiBuilderProxyServerServer *grpc.Server
	lis                                      net.Listener
)

// gRPC Server used for register clients Name, Ip and Port and Clients Test Environments and Clients Test Commands
type fenixGuiTestCaseBuilderGrpcServicesServer struct {
	fenixGuiTestCaseBuilderServerGrpcApi.UnimplementedFenixTestCaseBuilderServerGrpcServicesServer
}
