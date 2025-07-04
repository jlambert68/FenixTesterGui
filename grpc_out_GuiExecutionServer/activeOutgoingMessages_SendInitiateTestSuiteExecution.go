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

// SendInitiateTestSuiteExecution - Initiate a TestSuiteExecution
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendInitiateTestSuiteExecution(
	initiateSingleTestSuiteExecutionRequestMessage *fenixExecutionServerGuiGrpcApi.InitiateSingleTestSuiteExecutionRequestMessage) (
	initiateSingleTestSuiteExecutionResponseMessage *fenixExecutionServerGuiGrpcApi.InitiateSingleTestSuiteExecutionResponseMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "f5641b27-781d-4bdd-b213-28f1e786bcfd",
		"initiateSingleTestSuiteExecutionRequestMessage": initiateSingleTestSuiteExecutionRequestMessage,
	}).Debug("Incoming 'grpcOut - SendInitiateTestCaseExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "245a7e75-0cd2-4394-827e-5b6776751d36",
		"initiateSingleTestSuiteExecutionResponseMessage": &initiateSingleTestSuiteExecutionResponseMessage,
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
		initiateSingleTestSuiteExecutionResponseMessage = &fenixExecutionServerGuiGrpcApi.InitiateSingleTestSuiteExecutionResponseMessage{
			TestCasesInExecutionQueue: nil,
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   err.Error(),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
		}
		return initiateSingleTestSuiteExecutionResponseMessage
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
			initiateSingleTestSuiteExecutionResponseMessage = &fenixExecutionServerGuiGrpcApi.InitiateSingleTestSuiteExecutionResponseMessage{
				TestCasesInExecutionQueue: nil,
				AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
					AckNack:    false,
					Comments:   returnMessageString,
					ErrorCodes: nil,
					ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
						GetHighestFenixGuiExecutionServerProtoFileVersion()),
				},
			}
			return initiateSingleTestSuiteExecutionResponseMessage

		}

	}

	// Do the gRPC-call
	initiateSingleTestSuiteExecutionResponseMessage, err = fenixGuiExecutionServerGrpcClient.InitiateTestSuiteExecution(
		ctx, initiateSingleTestSuiteExecutionRequestMessage)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "136dd79e-bf82-419b-9a7d-56bc46e32984",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendInitiateTestSuiteExecution'")

		initiateSingleTestSuiteExecutionResponseMessage = &fenixExecutionServerGuiGrpcApi.InitiateSingleTestSuiteExecutionResponseMessage{
			TestCasesInExecutionQueue: nil,
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   fmt.Sprintf("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendInitiateTestSuiteExecution'. ErrorMessage: '%s'", err.Error()),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
		}
		return initiateSingleTestSuiteExecutionResponseMessage

	} else if initiateSingleTestSuiteExecutionResponseMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "9346b271-6e13-4013-b248-8c1330be1203",
			"Message from FenixGuiExecutionServer": initiateSingleTestSuiteExecutionResponseMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendInitiateTestSuiteExecution'")
	}

	return initiateSingleTestSuiteExecutionResponseMessage

}
