package grpc_out_GuiExecutionServer

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPCOutGuiExecutionServerStruct struct {
	logger *logrus.Logger
	//gcp    gcp.GcpObjectStruct
}

var GrpcOutGuiExecutionServerObject GRPCOutGuiExecutionServerStruct

var (
	// Standard gRPC Clientr
	remoteFenixGuiExecutionServerConnection *grpc.ClientConn
	FenixGuiExecutionServerAddressToDial    string

	fenixGuiExecutionServerGrpcClient fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClientClient
)

var highestFenixGuiExecutionServerProtoFileVersion int32 = -1
