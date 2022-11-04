package main

import (
	"FenixTesterGui/grpc_in"
	"FenixTesterGui/gui"
	"FenixTesterGui/restAPI"
	"github.com/sirupsen/logrus"
)

type referencesStruct struct {
	grpcIn   *grpc_in.GRPCInStruct
	restAPI  *restAPI.RestApiStruct
	uiServer *gui.GlobalUIServerStruct
}

/*
type FenixGuiBuilderProxyServerObjectStruct struct {
	LocalReferences referencesStruct
}

*/

type fenixGuiBuilderProxyServerObjectStruct struct {
	logger               *logrus.Logger
	runAsTrayApplication bool
	subPackageObjects    *referencesStruct
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
//var tempRunAsTrayApplication bool

// Create address for FenixGuiServer to call
var fenixGuiBuilderServerAddressToDial string
