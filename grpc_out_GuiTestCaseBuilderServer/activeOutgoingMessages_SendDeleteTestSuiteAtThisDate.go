package grpc_out_GuiTestCaseBuilderServer

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"context"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// SendDeleteTestSuiteAtThisDate - Marks a TestSuite as deleted in the database
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) SendDeleteTestSuiteAtThisDate(
	gRPCDeleteTestSuiteAtThisDateRequest *fenixGuiTestCaseBuilderServerGrpcApi.DeleteTestSuiteAtThisDateRequest) (
	returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse) {

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
		returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   err.Error(),
			ErrorCodes: nil,
			ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
				grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
		}

		return returnMessage

	}

	// Do gRPC-call
	//ctx := context.Background()
	timeOutDuration := time.Now().Add(30 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), timeOutDuration)
	defer func() {
		cancel()
	}()

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiTestCaseBuilderServer == sharedCode.GCP {

		// Set logger in GCP-package
		gcp.GcpObject.SetLogger(grpcOut.logger)

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.GcpObject.GenerateGCPAccessToken(
			ctx, gcp.TargetServerGuiTestCaseBuilderServer)
		if returnMessageAckNack == false {
			// When error
			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   returnMessageString,
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
					grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
			}

			return returnMessage
		}

	}

	// Add user info
	var userIdentification *fenixGuiTestCaseBuilderServerGrpcApi.UserIdentificationMessage
	userIdentification = &fenixGuiTestCaseBuilderServerGrpcApi.UserIdentificationMessage{
		UserIdOnComputer:     sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser: sharedCode.CurrentUserAuthenticatedTowardsGCP,
		ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.
			CurrentFenixTestCaseBuilderProtoFileVersionEnum(grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
	}
	gRPCDeleteTestSuiteAtThisDateRequest.UserIdentification = userIdentification

	// Do the gRPC-call
	returnMessage, err = fenixGuiTestCaseCaseBuilderServerGrpcClient.DeleteTestSuiteAtThisDate(
		ctx, gRPCDeleteTestSuiteAtThisDateRequest)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "71895152-f5e0-493f-8de4-ca568f74fa43",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendDeleteTestSuiteAtThisDate'")

	} else if returnMessage.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                     "b7e4c6af-57c5-4252-91c1-5d5a200fb008",
			"Message from FenixTestGuiBuilderServer": returnMessage.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendDeleteTestSuiteAtThisDate'")
	}

	return returnMessage

}
