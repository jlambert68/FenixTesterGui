package grpc_in

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"fmt"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	fenixUserGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixUserGui/fenixUserGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// GuiExecutionServerAreYouAlive - *********************************************************************
//Anyone can check if 'GuiExecutionServer' is alive with this service
func (s *fenixUserGuiGrpcServicesServer) GuiExecutionServerAreYouAlive(_ context.Context, _ *fenixUserGuiGrpcApi.EmptyParameter) (*fenixUserGuiGrpcApi.AckNackResponse, error) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "a34795a5-0ded-4cdf-99a6-9829f9a887f3",
	}).Debug("Incoming 'gRPC - GuiExecutionServerAreYouAlive'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "694fa3af-8344-4fc6-8f99-5f14fa0c2765",
	}).Debug("Outgoing 'gRPC - GuiExecutionServerAreYouAlive'")

	var responseMessageFromGuiTestCaseBuilderServer *fenixExecutionServerGuiGrpcApi.AckNackResponse

	// Do gRPC-call to 'GuiExecutionServer'
	responseMessageFromGuiTestCaseBuilderServer = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.SendAreYouAliveToGuiExecutionServer()

	// Create response message
	var ackNackResponse *fenixUserGuiGrpcApi.AckNackResponse
	ackNackResponse = &fenixUserGuiGrpcApi.AckNackResponse{
		AckNack:                            responseMessageFromGuiTestCaseBuilderServer.AckNack,
		Comments:                           fmt.Sprintf("Message from GuiExecutionServer: '%s'", responseMessageFromGuiTestCaseBuilderServer.Comments),
		ErrorCodes:                         nil,
		ProtoFileVersionUsedByFenixUserGui: fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum(getHighestFenixUserGuiServerProtoFileVersion()),
	}

	return ackNackResponse, nil
}
