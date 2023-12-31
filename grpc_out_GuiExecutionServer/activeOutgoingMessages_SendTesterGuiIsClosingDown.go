package grpc_out_GuiExecutionServer

import (
	"FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"fmt"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

// ********************************************************************************************************************

// SendTesterGuiIsClosingDown - Inform GuiExecutionServer that this TesterGui is closing down
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendTesterGuiIsClosingDown() (
	ackNackResponse *fenixExecutionServerGuiGrpcApi.AckNackResponse) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "eb17d34c-e24d-4398-a95f-45224116e250",
	}).Debug("Incoming 'grpcOut - SendTesterGuiIsClosingDown'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "38278fe0-e2cd-4b80-ad9b-f97e8c561fe4",
	}).Debug("Outgoing 'grpcOut - SendTesterGuiIsClosingDown'")

	// Create message to be sent to GuiExecutioNServer
	var userAndApplicationRunTimeIdentificationMessage *fenixExecutionServerGuiGrpcApi.
		UserAndApplicationRunTimeIdentificationMessage
	userAndApplicationRunTimeIdentificationMessage = &fenixExecutionServerGuiGrpcApi.
		UserAndApplicationRunTimeIdentificationMessage{
		ApplicationRunTimeUuid: sharedCode.ApplicationRunTimeUuid,
		UserIdOnComputer:       sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser:   sharedCode.CurrentUserAuthenticatedTowardsGCP,
		ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
			GetHighestFenixGuiExecutionServerProtoFileVersion()),
	}

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	ctx = context.Background()

	// Set up connection to Server
	ctx, err = grpcOut.setConnectionToFenixGuiExecutionMessageServer_new(ctx)
	//grpcOut.setConnectionToFenixGuiTestCaseBuilderServer()
	if err != nil {
		// When error
		ackNackResponse = &fenixExecutionServerGuiGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   err.Error(),
			ErrorCodes: nil,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				GetHighestFenixGuiExecutionServerProtoFileVersion()),
		}
		return ackNackResponse
	}

	// Set up connection to Server

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiExecutionServer == sharedCode.GCP {

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.GcpObject.GenerateGCPAccessToken(ctx, gcp.TargetServerGuiExecutionServer)
		if returnMessageAckNack == false {
			// When error
			ackNackResponse = &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   returnMessageString,
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			}

			return ackNackResponse
		}
	}

	// Do the gRPC-call
	ackNackResponse, err = fenixGuiExecutionServerGrpcClient.TesterGuiIsClosingDown(ctx, userAndApplicationRunTimeIdentificationMessage)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "351a26b5-0fc0-467b-bb64-3dd013aa470e",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendTesterGuiIsClosingDown'")

		ackNackResponse = &fenixExecutionServerGuiGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   fmt.Sprintf("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendTesterGuiIsClosingDown'. ErrorMessage: '%s'", err.Error()),
			ErrorCodes: nil,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				GetHighestFenixGuiExecutionServerProtoFileVersion()),
		}

		return ackNackResponse

	} else if ackNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "c05a3be4-ab02-42de-9a51-9253abf9b2c6",
			"Message from FenixGuiExecutionServer": ackNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendTesterGuiIsClosingDown'")
	}

	return ackNackResponse

}
