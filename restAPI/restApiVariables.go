package restAPI

import (
	"FenixTesterGui/grpc_out_GuiTestCaseBuilderServer"
	"github.com/sirupsen/logrus"
)

type RestApiStruct struct {
	logger                             *logrus.Logger
	GrpcOut                            *grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct
	fenixGuiBuilderServerAddressToDial string
}
