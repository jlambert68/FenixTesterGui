package grpc_out

import (
	sharedCode "FenixTesterGui/common_code"
	"context"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// SendSaveFullTestCase - Save full TestCase to database
func (grpcOut *GRPCOutStruct) SendSaveFullTestCase(gRPCFullTestCaseMessageToSend *fenixGuiTestCaseBuilderServerGrpcApi.FullTestCaseMessage) (returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse) {

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	// Set up connection to Server
	grpcOut.setConnectionToFenixGuiBuilderServer()

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		//TODO Fixa så att denna inte görs som allt går bra
		grpcOut.logger.WithFields(logrus.Fields{
			"ID": "8f935725-745f-4aa9-b647-335cba045b08",
		}).Error("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiServer == sharedCode.GCP {

		// Set logger in GCP-package
		grpcOut.gcp.SetLogger(grpcOut.logger)

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = grpcOut.gcp.GenerateGCPAccessTokenForAuthorizedUser(ctx)
		if returnMessageAckNack == false {
			// When error
			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   returnMessageString,
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
					grpcOut.GetHighestFenixGuiServerProtoFileVersion()),
			}

			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiBuilderServerGrpcClient.SaveFullTestCase(ctx, gRPCFullTestCaseMessageToSend)

	// Shouldn't happen
	if err != nil {
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":    "dc065973-cbab-4cf9-8091-a0f687cd2d36",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SaveFullTestCase'")

	} else if returnMessage.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":                                     "146af5d6-336c-4653-b8a1-b59f809a7b8b",
			"Message from FenixTestGuiBuilderServer": returnMessage.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SaveFullTestCase'")
	}

	return returnMessage

}
