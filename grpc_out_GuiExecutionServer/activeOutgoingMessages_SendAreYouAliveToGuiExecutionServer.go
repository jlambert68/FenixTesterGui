package grpc_out_GuiExecutionServer

import (
	"FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// ********************************************************************************************************************

// SendAreYouAliveToGuiExecutionServer - Check if 'GuiExecutionServer' is alive
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendAreYouAliveToGuiExecutionServer() (returnMessage *fenixExecutionServerGuiGrpcApi.AckNackResponse) {

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	ctx = context.Background()

	// Set up connection to Server
	ctx, err = grpcOut.setConnectionToFenixGuiExecutionMessageServer_new(ctx)
	//grpcOut.setConnectionToFenixGuiTestCaseBuilderServer()
	if err != nil {
		if returnMessageAckNack == false {
			// When error
			returnMessage = &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   err.Error(),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			}

			return returnMessage
		}
	}

	// Set up connection to Server
	returnMessage = grpcOut.SetConnectionToFenixGuiExecutionServer()
	// If there was no connection to backend then return that message
	if returnMessage != nil {
		return returnMessage
	}

	// Create the request message
	emptyParameter := &fenixExecutionServerGuiGrpcApi.EmptyParameter{

		ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
			GetHighestFenixGuiExecutionServerProtoFileVersion()),
	}

	// Do gRPC-call
	/*
		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer func() {
			//TODO Fixa så att denna inte görs som allt går bra
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID": "5094f038-ce5b-4374-af59-45a519bffffa",
			}).Error("Running Defer Cancel function")
			cancel()
		}()

	*/

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiTestCaseBuilderServer == sharedCode.GCP {

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.GcpObject.GenerateGCPAccessToken(ctx)
		if returnMessageAckNack == false {
			// When error
			returnMessage = &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   returnMessageString,
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			}

			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiExecutionServerGrpcClient.AreYouAlive(ctx, emptyParameter)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "eacdd7b8-8460-4354-9df3-a5ab8a87d1d0",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendAreYouAliveToGuiExecutionServer'")

	} else if returnMessage.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "70aeb9f6-e612-4e16-8c88-c75e3261f51b",
			"Message from FenixGuiExecutionServer": returnMessage.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendAreYouAliveToGuiExecutionServer'")
	}

	return returnMessage

}
