package grpc_out_GuiTestCaseBuilderServer

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"context"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// SendListAllAvailablePinnedTestInstructionsAndTestInstructionContainers - Get pinned TestInstructions and TestInstructionContainers
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) SendListAllAvailablePinnedTestInstructionsAndTestInstructionContainers(userId string) (returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

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

		returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage{
			AvailablePinnedTestInstructions:                    nil,
			AvailablePinnedPreCreatedTestInstructionContainers: nil,
			AckNackResponse: ackNackResponse,
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

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage{
				AvailablePinnedTestInstructions:                    nil,
				AvailablePinnedPreCreatedTestInstructionContainers: nil,
				AckNackResponse: ackNackResponse,
			}

			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiTestCaseCaseBuilderServerGrpcClient.ListAllAvailablePinnedTestInstructionsAndTestInstructionContainers(ctx, userIdentificationMessage)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "7bff3257-a193-4d07-83aa-f106f6f734a0",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendListAllAvailablePinnedTestInstructionsAndTestInstructionContainers'")

	} else if returnMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                     "72f79764-2549-4ce7-867e-16cd0f414dff",
			"Message from FenixTestGuiBuilderServer": returnMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendListAllAvailablePinnedTestInstructionsAndTestInstructionContainers'")
	}

	return returnMessage

}
