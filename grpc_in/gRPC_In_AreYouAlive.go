package grpc_in

import (
	sharedCode "FenixTesterGui/common_code"
	fenixUserGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixUserGui/fenixUserGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// AreYouAlive - *********************************************************************
//Anyone can check if Fenix TestCase Builder server is alive with this service
func (s *fenixUserGuiGrpcServicesServer) AreYouAlive(_ context.Context, _ *fenixUserGuiGrpcApi.EmptyParameter) (*fenixUserGuiGrpcApi.AckNackResponse, error) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "689ff51c-0ef7-406c-b4d6-7b5946a39f9a",
	}).Debug("Incoming 'gRPC - AreYouAlive'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "c0f7b368-674a-4cfa-a3fa-aa71294e7b8d",
	}).Debug("Outgoing 'gRPC - AreYouAlive'")

	// Create response message
	var ackNackResponse *fenixUserGuiGrpcApi.AckNackResponse
	ackNackResponse = &fenixUserGuiGrpcApi.AckNackResponse{
		AckNack:                            true,
		Comments:                           "I'am alive, from 'Fenix Inception - GUI'",
		ErrorCodes:                         nil,
		ProtoFileVersionUsedByFenixUserGui: fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum(getHighestFenixUserGuiServerProtoFileVersion()),
	}

	return ackNackResponse, nil
}
