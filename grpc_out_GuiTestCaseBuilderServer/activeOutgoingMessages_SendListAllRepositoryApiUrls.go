package grpc_out_GuiTestCaseBuilderServer

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"context"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// SendListAllRepositoryApiUrls - Load all Repository Api Urls for where Templates are stored
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) SendListAllRepositoryApiUrls() (
	returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListAllRepositoryApiUrlsResponseMessage) {

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

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListAllRepositoryApiUrlsResponseMessage{
				AckNackResponse:   ackNackResponse,
				RepositoryApiUrls: nil,
			}

			return returnMessage
		}
	}

	// Create the request message
	var userIdentificationMessage *fenixGuiTestCaseBuilderServerGrpcApi.UserIdentificationMessage
	userIdentificationMessage = &fenixGuiTestCaseBuilderServerGrpcApi.UserIdentificationMessage{
		UserIdOnComputer:     sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser: sharedCode.CurrentUserAuthenticatedTowardsGCP,
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

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListAllRepositoryApiUrlsResponseMessage{
				AckNackResponse:   ackNackResponse,
				RepositoryApiUrls: nil,
			}

			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiTestCaseCaseBuilderServerGrpcClient.ListAllRepositoryApiUrls(ctx, userIdentificationMessage)

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

		returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListAllRepositoryApiUrlsResponseMessage{
			AckNackResponse:   ackNackResponse,
			RepositoryApiUrls: nil,
		}

		return returnMessage

	} else if returnMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                     "c4aaad0d-545e-4fad-b051-1e97746c3a41",
			"Message from FenixTestGuiBuilderServer": returnMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListAllRepositoryApiUrls'")
	}

	return returnMessage

}
