package grpc_out_GuiExecutionServer

import (
	"FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"fmt"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

// ********************************************************************************************************************

// SendTesterGuiIsStartingUp - Inform GuiExecutionServer that this TesterGui is starting up
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendTesterGuiIsStartingUp() (
	ackNackResponse *fenixExecutionServerGuiGrpcApi.AckNackResponse) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "7c17844b-0bda-489c-bf6b-115e72d5a0fb",
	}).Debug("Incoming 'grpcOut - SendTesterGuiIsStartingUp'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "c3fdd3c1-a988-482b-afed-e3baf4253b18",
	}).Debug("Outgoing 'grpcOut - SendTesterGuiIsStartingUp'")

	// Create message to be sent to GuiExecutioNServer
	var userAndApplicationRunTimeIdentificationMessage *fenixExecutionServerGuiGrpcApi.
		UserAndApplicationRunTimeIdentificationMessage
	userAndApplicationRunTimeIdentificationMessage = &fenixExecutionServerGuiGrpcApi.
		UserAndApplicationRunTimeIdentificationMessage{
		ApplicationRunTimeUuid: sharedCode.ApplicationRunTimeUuid,
		UserId:                 sharedCode.CurrentUserId,
		ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
			GetHighestFenixGuiExecutionServerProtoFileVersion()),
	}

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	ctx = context.Background()

	// Set up connection to Server
	ctx, err = grpcOut.setConnectionToFenixGuiExecutionMessageServer_new(ctx)
	//grpcOut.setConnectionToFenixGuiTestCaseBuilderServer()
	if err != nil {
		// When error
		ackNackResponse = &fenixExecutionServerGuiGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   err.Error(),
			ErrorCodes: nil,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				GetHighestFenixGuiExecutionServerProtoFileVersion()),
		}
		return ackNackResponse
	}

	// Set up connection to Server

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiExecutionServer == sharedCode.GCP {

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.GcpObject.GenerateGCPAccessToken(ctx, gcp.TargetServerGuiExecutionServer)
		if returnMessageAckNack == false {
			// When error
			ackNackResponse = &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   returnMessageString,
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			}

			return ackNackResponse
		}
	}

	// Do the gRPC-call
	ackNackResponse, err = fenixGuiExecutionServerGrpcClient.TesterGuiIsStartingUp(ctx, userAndApplicationRunTimeIdentificationMessage)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "ae23568d-1fa2-4ff5-aab8-a74490dd18f5",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendInitiateTestCaseExecution'")

		ackNackResponse = &fenixExecutionServerGuiGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   fmt.Sprintf("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendTesterGuiIsStartingUp'. ErrorMessage: '%s'", err.Error()),
			ErrorCodes: nil,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				GetHighestFenixGuiExecutionServerProtoFileVersion()),
		}

		return ackNackResponse

	} else if ackNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "7827a8d9-959c-43cb-813d-f5ab344aacf4",
			"Message from FenixGuiExecutionServer": ackNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendTesterGuiIsStartingUp'")
	}

	return ackNackResponse

}
