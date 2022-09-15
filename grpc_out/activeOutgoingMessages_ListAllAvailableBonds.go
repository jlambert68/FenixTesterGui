package grpc_out

import (
	sharedCode "FenixTesterGui/common_code"
	"context"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// ListAllAvailableBonds - Get all Bonds that can be used within a TestCase
func (grpcOut *GRPCOutStruct) ListAllAvailableBonds(userId string) (returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage) {

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	// Set up connection to Server
	grpcOut.setConnectionToFenixGuiBuilderServer()

	// Create the request message
	userIdentificationMessage := &fenixGuiTestCaseBuilderServerGrpcApi.UserIdentificationMessage{
		UserId: userId,
		ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
			grpcOut.GetHighestFenixGuiServerProtoFileVersion()),
	}

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		//TODO Fixa så att denna inte görs som allt går bra
		grpcOut.logger.WithFields(logrus.Fields{
			"ID": "793227a3-fe75-4e69-9634-e16096038bd1",
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
			ackNackResponse := &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   returnMessageString,
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
					grpcOut.GetHighestFenixGuiServerProtoFileVersion()),
			}

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage{
				ImmatureBonds:   nil,
				AckNackResponse: ackNackResponse,
			}
			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiBuilderServerGrpcClient.ListAllAvailableBonds(ctx, userIdentificationMessage)

	// Shouldn't happen
	if err != nil {
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":    "8d0dc097-420e-447a-8d8f-53ec9f55c53b",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListAllAvailableBonds'")

	} else if returnMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":                                     "83a9cab0-3a1d-41de-a38f-81ddd492092d",
			"Message from FenixTestGuiBuilderServer": returnMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListAllAvailableBonds'")
	}

	return returnMessage

}
