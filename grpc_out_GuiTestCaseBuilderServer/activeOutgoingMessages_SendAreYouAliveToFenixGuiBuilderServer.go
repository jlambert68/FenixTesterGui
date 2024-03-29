package grpc_out_GuiTestCaseBuilderServer

import (
	"FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

// ********************************************************************************************************************

// SendAreYouAliveToFenixGuiBuilderServer - Check if FenixGuiBuilderServer is alive
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) SendAreYouAliveToFenixGuiBuilderServer() (returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse) {

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

	// Create the request message
	emptyParameter := &fenixGuiTestCaseBuilderServerGrpcApi.EmptyParameter{

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

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.GcpObject.GenerateGCPAccessToken(ctx, gcp.TargetServerGuiTestCaseBuilderServer)
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

	// Do the gRPC-call
	returnMessage, err = fenixGuiTestCaseCaseBuilderServerGrpcClient.AreYouAlive(ctx, emptyParameter)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "818aaf0b-4112-4be4-97b9-21cc084c7b8b",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendAreYouAliveToFenixGuiBuilderServer'")

	} else if returnMessage.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                     "2ecbc800-2fb6-4e88-858d-a421b61c5529",
			"Message from FenixTestGuiBuilderServer": returnMessage.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendAreYouAliveToFenixGuiBuilderServer'")
	}

	return returnMessage

}
