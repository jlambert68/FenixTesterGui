package grpc_out_GuiExecutionServer

import (
	"FenixTesterGui/gcp"
	fenixGuiExecutionServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPCOutGuiExecutionServerStruct struct {
	logger                             *logrus.Logger
	fenixGuiBuilderServerAddressToDial string
	gcp                                gcp.GcpObjectStruct
}

var grpcOutGuiExecutionServerObject GRPCOutGuiExecutionServerStruct

var (
	// Standard gRPC Clientr
	remoteFenixGuiExecutionServerConnection *grpc.ClientConn
	// FenixGuiExecutionServerAddressToDial gRpcClientForFenixGuiBuilderServer fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServicesClient
	FenixGuiExecutionServerAddressToDial string

	fenixGuiTestCaseCaseBuilderServerGrpcClient fenixGuiExecutionServerGrpcApi.FenixTestCaseBuilderServerGrpcServicesClient
)

var highestFenixGuiExecutionServerProtoFileVersion int32 = -1
