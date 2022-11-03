package grpc_out_GuiTestCaseBuilderServer

import (
	"FenixTesterGui/gcp"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPCOutGuiTestCaseBuilderServerStruct struct {
	logger                             *logrus.Logger
	fenixGuiBuilderServerAddressToDial string
	gcp                                gcp.GcpObjectStruct
}

var grpcOutGuiTestCaseBuilderServerObject GRPCOutGuiTestCaseBuilderServerStruct

var (
	// Standard gRPC Clientr
	remoteFenixGuiTestCaseBuilderServerConnection *grpc.ClientConn
	// FenixGuiTestCaseBuilderServerAddressToDial gRpcClientForFenixGuiBuilderServer fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServicesClient
	FenixGuiTestCaseBuilderServerAddressToDial string

	fenixGuiTestCaseCaseBuilderServerGrpcClient fenixGuiTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServicesClient
)

var highestFenixGuiTestCaseBuilderServerProtoFileVersion int32 = -1
