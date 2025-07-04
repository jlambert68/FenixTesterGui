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

// SendInitiateTestCaseExecution - Initiate a TestCaseExecution
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendInitiateTestCaseExecution(
	initiateSingleTestCaseExecutionRequestMessage *fenixExecutionServerGuiGrpcApi.InitiateSingleTestCaseExecutionRequestMessage) (
	initiateSingleTestCaseExecutionResponseMessage *fenixExecutionServerGuiGrpcApi.InitiateSingleTestCaseExecutionResponseMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "f5672b96-c58b-45eb-a2bb-bf399a6de9e6",
		"initiateSingleTestCaseExecutionRequestMessage": initiateSingleTestCaseExecutionRequestMessage,
	}).Debug("Incoming 'grpcOut - SendInitiateTestCaseExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "694fa3af-8344-4fc6-8f99-5f14fa0c2765",
		"initiateSingleTestCaseExecutionResponseMessage": &initiateSingleTestCaseExecutionResponseMessage,
	}).Debug("Outgoing 'grpcOut - SendInitiateTestCaseExecution'")

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
		initiateSingleTestCaseExecutionResponseMessage = &fenixExecutionServerGuiGrpcApi.InitiateSingleTestCaseExecutionResponseMessage{
			TestCasesInExecutionQueue: nil,
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   err.Error(),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
		}
		return initiateSingleTestCaseExecutionResponseMessage
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
			initiateSingleTestCaseExecutionResponseMessage = &fenixExecutionServerGuiGrpcApi.InitiateSingleTestCaseExecutionResponseMessage{
				TestCasesInExecutionQueue: nil,
				AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
					AckNack:    false,
					Comments:   returnMessageString,
					ErrorCodes: nil,
					ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
						GetHighestFenixGuiExecutionServerProtoFileVersion()),
				},
			}
			return initiateSingleTestCaseExecutionResponseMessage

		}

	}

	// Do the gRPC-call
	initiateSingleTestCaseExecutionResponseMessage, err = fenixGuiExecutionServerGrpcClient.InitiateTestCaseExecution(ctx, initiateSingleTestCaseExecutionRequestMessage)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "258310aa-1162-48ff-9a8c-3efc3d7d5b7c",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendInitiateTestCaseExecution'")

		initiateSingleTestCaseExecutionResponseMessage = &fenixExecutionServerGuiGrpcApi.InitiateSingleTestCaseExecutionResponseMessage{
			TestCasesInExecutionQueue: nil,
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   fmt.Sprintf("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendInitiateTestCaseExecution'. ErrorMessage: '%s'", err.Error()),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
		}
		return initiateSingleTestCaseExecutionResponseMessage

	} else if initiateSingleTestCaseExecutionResponseMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "36e0cf63-6c2f-4cc8-8a1d-8cb25a8c3236",
			"Message from FenixGuiExecutionServer": initiateSingleTestCaseExecutionResponseMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendInitiateTestCaseExecution'")
	}

	return initiateSingleTestCaseExecutionResponseMessage

}
