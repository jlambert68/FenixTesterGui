package main

import (
	"FenixTesterGui/fenixTestGuiObject"
	"github.com/sirupsen/logrus"
)

type fenixGuiBuilderProxyServerObjectStruct struct {
	logger               *logrus.Logger
	runAsTrayApplication bool
	subPackageObjects    *fenixTestGuiObject.FenixGuiBuilderProxyServerObjectStruct
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
