package main

import (
	"FenixTesterGui/grpc_in"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	//	ecpb "github.com/jlambert68/FenixGrpcApi/Client/fenixGuiTestCaseBuilderServerGrpcApi/echo/go_grpc_api"
)

type localreferenceLibraryStruct struct {
	grpc_in *grpc_in.GRPCInStruct
}

type fenixGuiBuilderProxyServerObjectStruct struct {
	logger               *logrus.Logger
	gcpAccessToken       *oauth2.Token
	runAsTrayApplication bool
	localreferencs       localreferenceLibraryStruct
}

// Variable holding everything together
var fenixTesterGuiObject *fenixGuiBuilderProxyServerObjectStruct

//TODO FIXA DENNA PATH, HMMM borde köra i DB framöver
// For now hardcoded MerklePath
//var merkleFilterPath string = //"AccountEnvironment/ClientJuristictionCountryCode/MarketSubType/MarketName/" //SecurityType/"

// Echo gRPC-server
/*
type ecServer struct {
	echo.UnimplementedEchoServer
}


*/

var (
	// Standard gRPC Clientr
	remoteFenixGuiBuilderServerConnection *grpc.ClientConn
	//gRpcClientForFenixGuiBuilderServer fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServicesClient
	FenixGuiBuilderServerAddressToDial string

	fenixGuiBuilderServerGrpcClient fenixGuiTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServicesClient
)

// Bad solution but using temp storage before real variable is initiated
var tempRunAsTrayApplication bool
