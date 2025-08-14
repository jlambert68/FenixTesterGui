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

// SendInitiateTestSuiteExecutionWithAllTestDataSets - Initiate a TestSuiteExecution with all its TestDataSets
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendInitiateTestSuiteExecutionWithAllTestDataSets(
	initiateTestSuiteExecutionWithAllTestDataSetsRequestMessage *fenixExecutionServerGuiGrpcApi.InitiateTestSuiteExecutionWithAllTestDataSetsRequestMessage) (
	initiateSingleTestSuiteExecutionResponseMessage *fenixExecutionServerGuiGrpcApi.InitiateSingleTestSuiteExecutionResponseMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "ec2eeb8f-02fb-4b45-89e6-51d46911aad9",
		"initiateTestSuiteExecutionWithAllTestDataSetsRequestMessage": initiateTestSuiteExecutionWithAllTestDataSetsRequestMessage,
	}).Debug("Incoming 'grpcOut - SendInitiateTestSuiteExecutionWithAllTestDataSets'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "0c782d61-02b6-4b6d-8732-7d65bf9494a1",
		"initiateSingleTestSuiteExecutionResponseMessage": &initiateSingleTestSuiteExecutionResponseMessage,
	}).Debug("Outgoing 'grpcOut - SendInitiateTestSuiteExecutionWithAllTestDataSets'")

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
	initiateSingleTestSuiteExecutionResponseMessage, err = fenixGuiExecutionServerGrpcClient.InitiateTestSuiteExecutionWithAllTestDataSets(
		ctx, initiateTestSuiteExecutionWithAllTestDataSetsRequestMessage)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "f1a4bbf0-40c1-4cc6-ab0e-a212bbb4d007",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendInitiateTestSuiteExecutionWithAllTestDataSets'")

		initiateSingleTestSuiteExecutionResponseMessage = &fenixExecutionServerGuiGrpcApi.InitiateSingleTestSuiteExecutionResponseMessage{
			TestCasesInExecutionQueue: nil,
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   fmt.Sprintf("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendInitiateTestSuiteExecutionWithAllTestDataSets'. ErrorMessage: '%s'", err.Error()),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
		}
		return initiateSingleTestSuiteExecutionResponseMessage

	} else if initiateSingleTestSuiteExecutionResponseMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "c4c5fa21-76ef-4327-8777-290e881b3491",
			"Message from FenixGuiExecutionServer": initiateSingleTestSuiteExecutionResponseMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendInitiateTestSuiteExecutionWithAllTestDataSets'")
	}

	return initiateSingleTestSuiteExecutionResponseMessage

}
