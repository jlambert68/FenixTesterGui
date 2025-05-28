package grpc_out_GuiTestCaseBuilderServer

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"context"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// ListTestCasesThatCanBeEditedResponseMessage - List all TestCasesMapPtr that can be edited, used for producing a list
// that the used can chose TestCase to edit from
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) ListTestCasesThatCanBeEditedResponseMessage(
	testCaseUpdatedMinTimeStamp time.Time,
	testCaseExecutionUpdatedMinTimeStamp time.Time) (
	returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListTestCasesThatCanBeEditedResponseMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "42fe98c6-2410-4abe-87c4-d79b5b30dfce",
	}).Debug("Incoming 'grpcOut - SendListTestCasesThatCanBeEdited'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "57d234d3-8680-43ab-a02f-dbeda0bc4745",
	}).Debug("Outgoing 'grpcOut - SendListTestCasesThatCanBeEdited'")

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	ctx = context.Background()

	// Set up connection to Server
	ctx, err = grpcOut.setConnectionToFenixGuiTestCaseBuilderServer_new(ctx)

	if err != nil {
		if returnMessageAckNack == false {
			// When error
			ackNackResponse := &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   err.Error(),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
					grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
			}

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListTestCasesThatCanBeEditedResponseMessage{
				AckNackResponse:                ackNackResponse,
				TestCasesThatCanBeEditedByUser: nil,
			}

			return returnMessage
		}
	}

	// Create the request message
	var listTestCasesRequestMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListTestCasesRequestMessage
	listTestCasesRequestMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListTestCasesRequestMessage{
		UserIdOnComputer:     sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser: sharedCode.CurrentUserAuthenticatedTowardsGCP,
		ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
			grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
		TestCaseUpdatedMinTimeStamp:          timestamppb.New(testCaseUpdatedMinTimeStamp),
		TestCaseExecutionUpdatedMinTimeStamp: timestamppb.New(testCaseExecutionUpdatedMinTimeStamp),
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

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListTestCasesThatCanBeEditedResponseMessage{
				AckNackResponse:                ackNackResponse,
				TestCasesThatCanBeEditedByUser: nil,
			}

			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiTestCaseCaseBuilderServerGrpcClient.ListTestCasesThatCanBeEdited(
		ctx, listTestCasesRequestMessage)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "c3df6f17-8e33-410a-b8f2-69185a91b270",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListAllRepositoryApiUrls'")

		// When error
		ackNackResponse := &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListAllRepositoryApiUrls'",
			ErrorCodes: nil,
			ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
				grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
		}

		returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListTestCasesThatCanBeEditedResponseMessage{
			AckNackResponse:                ackNackResponse,
			TestCasesThatCanBeEditedByUser: nil,
		}

		return returnMessage

	} else if returnMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                     "becdfdd1-0e21-4e5b-b18f-b395257b8b85",
			"Message from FenixTestGuiBuilderServer": returnMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListAllRepositoryApiUrls'")
	}

	return returnMessage

}
