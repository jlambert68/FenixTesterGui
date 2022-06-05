package restAPI

import (
	"FenixTesterGui/grpc_out"
	"github.com/sirupsen/logrus"
)

type RestApiStruct struct {
	logger                             *logrus.Logger
	GrpcOut                            *grpc_out.GRPCOutStruct
	fenixGuiBuilderServerAddressToDial string
}
