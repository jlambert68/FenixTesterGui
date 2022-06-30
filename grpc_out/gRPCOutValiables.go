package grpc_out

import (
	"FenixTesterGui/gcp"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPCOutStruct struct {
	logger                             *logrus.Logger
	fenixGuiBuilderServerAddressToDial string
	gcp                                gcp.GcpObjectStruct
}

var grpcOutVaraible GRPCOutStruct

var (
	// Standard gRPC Clientr
	remoteFenixGuiBuilderServerConnection *grpc.ClientConn
	// FenixGuiBuilderServerAddressToDial gRpcClientForFenixGuiBuilderServer fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServicesClient
	FenixGuiBuilderServerAddressToDial string

	fenixGuiBuilderServerGrpcClient fenixGuiTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServicesClient
)

var highestFenixGuiServerProtoFileVersion int32 = -1
