package messageEngine

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"google.golang.org/grpc"
)

type messageEngineServerStruct struct {
}

var messageEngineServerObject messageEngineServerStruct

var (
	// Standard gRPC Clientr
	remoteFenixGuiExecutionServerConnection *grpc.ClientConn
	FenixGuiExecutionServerAddressToDial    string

	fenixGuiExecutionServerGrpcClient fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesClient
)

var highestFenixGuiExecutionServerProtoFileVersion int32 = -1
