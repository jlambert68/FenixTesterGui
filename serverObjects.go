package main

import (
	"FenixTesterGui/grpc_in"
	"github.com/sirupsen/logrus"
	//	ecpb "github.com/jlambert68/FenixGrpcApi/Client/fenixGuiTestCaseBuilderServerGrpcApi/echo/go_grpc_api"
)

type localreferenceLibraryStruct struct {
	grpc_in *grpc_in.GRPCInStruct
}

type fenixGuiBuilderProxyServerObjectStruct struct {
	logger               *logrus.Logger
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

// Bad solution but using temp storage before real variable is initiated
var tempRunAsTrayApplication bool
