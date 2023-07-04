package grpc_out_GuiTestCaseBuilderServer

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
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

	ctx = context.Background()

	// Set up connection to Server
	ctx, err = grpcOut.setConnectionToFenixGuiTestCaseBuilderServer_new(ctx)
	//grpcOut.setConnectionToFenixGuiTestCaseBuilderServer()
	if err != nil {
		// When error
		ackNackResponse := &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   err.Error(),
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
		cancel()
	}()

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiTestCaseBuilderServer == sharedCode.GCP {

		// Set logger in GCP-package
		gcp.GcpObject.SetLogger(grpcOut.logger)

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.GcpObject.GenerateGCPAccessToken(ctx, gcp.TargetServerGuiTestCaseBuilderServer)
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
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "010c58f4-7897-4faa-bb7e-8e529e37dc2a",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListAllImmatureTestInstructionAttributes'")

	} else if returnMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                     "a0f27ee6-9137-4a18-83cc-bf6a7bddac8e",
			"Message from FenixTestGuiBuilderServer": returnMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListAllImmatureTestInstructionAttributes'")
	}

	return returnMessage

}
