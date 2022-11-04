package grpc_in

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiTestCaseBuilderServer"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	fenixUserGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixUserGui/fenixUserGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// GuiTestCaseBuilderServerAreYouAlive - *********************************************************************
//Anyone can check if 'GuiTestCaseBuilderServer' is alive with this service
func (s *fenixUserGuiGrpcServicesServer) GuiTestCaseBuilderServerAreYouAlive(_ context.Context, _ *fenixUserGuiGrpcApi.EmptyParameter) (*fenixUserGuiGrpcApi.AckNackResponse, error) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "927f6d10-3734-4280-903a-881b9aaf9bbc",
	}).Debug("Incoming 'gRPC - GuiTestCaseBuilderServerAreYouAlive'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "24aca96e-e35e-4210-8f03-aa642eaec0df",
	}).Debug("Outgoing 'gRPC - GuiTestCaseBuilderServerAreYouAlive'")

	var responseMessageFromGuiTestCaseBuilderServer *fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse

	// Do gRPC-call to 'GuiTestCaseBuilderServer'
	responseMessageFromGuiTestCaseBuilderServer = grpc_out_GuiTestCaseBuilderServer.GrpcOutGuiTestCaseBuilderServerObject.SendAreYouAliveToFenixGuiBuilderServer()

	// Create response message
	var ackNackResponse *fenixUserGuiGrpcApi.AckNackResponse
	ackNackResponse = &fenixUserGuiGrpcApi.AckNackResponse{
		AckNack:                            responseMessageFromGuiTestCaseBuilderServer.AckNack,
		Comments:                           fmt.Sprintf("Message from GuiTestCaseBuilderServer: '%s'", responseMessageFromGuiTestCaseBuilderServer.Comments),
		ErrorCodes:                         nil,
		ProtoFileVersionUsedByFenixUserGui: fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum(getHighestFenixUserGuiServerProtoFileVersion()),
	}

	return ackNackResponse, nil
}
