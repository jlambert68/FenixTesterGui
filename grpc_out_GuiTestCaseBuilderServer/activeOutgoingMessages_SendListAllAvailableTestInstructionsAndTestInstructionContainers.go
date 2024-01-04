package grpc_out_GuiTestCaseBuilderServer

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"context"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// SendListAllAvailableTestInstructionsAndTestInstructionContainers - Get available TestInstructions and TestInstructionContainers
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) SendListAllAvailableTestInstructionsAndTestInstructionContainers(
	gCPAuthenticatedUser string) (
	returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"ID": "57596a36-cdde-45b4-9633-e4ba88904cce",
	}).Debug("Incoming - 'SendListAllAvailableTestInstructionsAndTestInstructionContainers'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"ID": "b2289982-1853-419c-8265-5393f3baa6ad",
	}).Debug("Outgoing - 'SendListAllAvailableTestInstructionsAndTestInstructionContainers'")

	//
	if len(gCPAuthenticatedUser) == 0 {
		// gCPAuthenticatedUser must have a value

		defer sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                   "3d9e4a27-6fa9-4cc3-b147-a49b385258d6",
			"gCPAuthenticatedUser": gCPAuthenticatedUser,
		}).Debug("gCPAuthenticatedUser is missing a value")

		ackNackResponse := &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "gCPAuthenticatedUser is missing a value",
			ErrorCodes: nil,
			ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
				grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
		}

		returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage{
			DomainsThatCanOwnTheTestCase:      nil,
			ImmatureTestInstructions:          nil,
			ImmatureTestInstructionContainers: nil,
			AckNackResponse:                   ackNackResponse,
		}

		return returnMessage
	}

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

		returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage{
			DomainsThatCanOwnTheTestCase:      nil,
			ImmatureTestInstructions:          nil,
			ImmatureTestInstructionContainers: nil,
			AckNackResponse:                   ackNackResponse,
		}

		return returnMessage

	}

	// Create the request message
	userIdentificationMessage := &fenixGuiTestCaseBuilderServerGrpcApi.UserIdentificationMessage{
		UserIdOnComputer:     sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser: sharedCode.CurrentUserAuthenticatedTowardsGCP,
		ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
			grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
	}

	// Do gRPC-call
	//ctx = context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
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

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage{
				DomainsThatCanOwnTheTestCase:      nil,
				ImmatureTestInstructions:          nil,
				ImmatureTestInstructionContainers: nil,
				AckNackResponse:                   ackNackResponse,
			}

			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiTestCaseCaseBuilderServerGrpcClient.ListAllAvailableTestInstructionsAndTestInstructionContainers(ctx, userIdentificationMessage)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "d7235084-33e5-43a2-9fa7-dfb05ec6869e",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendListAllAvailableTestInstructionsAndTestInstructionContainers'")

	} else if returnMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                     "30e6f1ee-202a-47bf-a2c4-5066d0f8cf75",
			"Message from FenixTestGuiBuilderServer": returnMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendListAllAvailableTestInstructionsAndTestInstructionContainers'")
	}

	return returnMessage

}
