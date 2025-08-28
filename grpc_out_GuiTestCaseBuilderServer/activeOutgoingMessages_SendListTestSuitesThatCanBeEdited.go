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

// ListTestSuitesThatCanBeEditedResponseMessage - List all TestSuitesMapPtr that can be edited, used for producing a list
// that the used can chose TestCase to edit from
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) ListTestSuitesThatCanBeEditedResponseMessage(
	testCaseUpdatedMinTimeStamp time.Time,
	testCaseExecutionUpdatedMinTimeStamp time.Time) (
	returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListTestSuitesResponseMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "0e4b4049-fec8-46ab-a517-61f5bd669527",
	}).Debug("Incoming 'grpcOut - ListTestSuitesThatCanBeEditedResponseMessage'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "362b08bc-4488-4f01-84bb-0da21f7d1233",
	}).Debug("Outgoing 'grpcOut - ListTestSuitesThatCanBeEditedResponseMessage'")

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

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListTestSuitesResponseMessage{
				AckNackResponse:           ackNackResponse,
				BasicTestSuiteInformation: nil,
			}

			return returnMessage
		}
	}

	// Create the request message
	var listTestSuitesRequestMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListTestSuitesRequestMessage
	listTestSuitesRequestMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListTestSuitesRequestMessage{
		UserIdOnComputer:     sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser: sharedCode.CurrentUserAuthenticatedTowardsGCP,
		ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
			grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
		TestSuiteUpdatedMinTimeStamp:          timestamppb.New(testCaseUpdatedMinTimeStamp),
		TestSuiteExecutionUpdatedMinTimeStamp: timestamppb.New(testCaseExecutionUpdatedMinTimeStamp),
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

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListTestSuitesResponseMessage{
				AckNackResponse:           ackNackResponse,
				BasicTestSuiteInformation: nil,
			}

			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiTestCaseCaseBuilderServerGrpcClient.ListTestSuitesThatCanBeEdited(
		ctx, listTestSuitesRequestMessage)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "fb966a03-5bee-4b6b-8051-c8e0cc36f34e",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListTestSuitesThatCanBeEditedResponseMessage'")

		// When error
		ackNackResponse := &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListTestSuitesThatCanBeEditedResponseMessage'",
			ErrorCodes: nil,
			ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
				grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
		}

		returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListTestSuitesResponseMessage{
			AckNackResponse:           ackNackResponse,
			BasicTestSuiteInformation: nil,
		}

		return returnMessage

	} else if returnMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                     "8c861ebb-0b30-4341-b882-4772e4e22673",
			"Message from FenixTestGuiBuilderServer": returnMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListTestSuitesThatCanBeEditedResponseMessage'")
	}

	return returnMessage

}
