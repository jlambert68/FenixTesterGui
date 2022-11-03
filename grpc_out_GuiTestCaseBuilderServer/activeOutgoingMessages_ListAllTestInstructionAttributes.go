package grpc_out_GuiTestCaseBuilderServer

import (
	sharedCode "FenixTesterGui/common_code"
	"context"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// ListAllTestInstructionAttributes - Get all attributes used within TestInstructions that can be used within a TestCase
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) ListAllTestInstructionAttributes(userId string) (returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage) {

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	// Set up connection to Server
	grpcOut.setConnectionToFenixGuiTestCaseBuilderServer()

	// Create the request message
	userIdentificationMessage := &fenixGuiTestCaseBuilderServerGrpcApi.UserIdentificationMessage{
		UserId: userId,
		ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
			grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
	}

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		//TODO Fixa så att denna inte görs som allt går bra
		grpcOut.logger.WithFields(logrus.Fields{
			"ID": "0ae9bacb-de78-48dc-baba-01dbe56349e0",
		}).Error("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiTestCaseBuilderServer == sharedCode.GCP {

		// Set logger in GCP-package
		grpcOut.gcp.SetLogger(grpcOut.logger)

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = grpcOut.gcp.GenerateGCPAccessTokenForAuthorizedUser(ctx)
		if returnMessageAckNack == false {
			// When error
			ackNackResponse := &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   returnMessageString,
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
					grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
			}

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage{
				TestInstructionAttributesList: nil,
				AckNackResponse:               ackNackResponse,
			}
			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiTestCaseCaseBuilderServerGrpcClient.ListAllImmatureTestInstructionAttributes(ctx, userIdentificationMessage)

	// Shouldn't happen
	if err != nil {
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":    "010c58f4-7897-4faa-bb7e-8e529e37dc2a",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListAllImmatureTestInstructionAttributes'")

	} else if returnMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":                                     "a0f27ee6-9137-4a18-83cc-bf6a7bddac8e",
			"Message from FenixTestGuiBuilderServer": returnMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListAllImmatureTestInstructionAttributes'")
	}

	return returnMessage

}
