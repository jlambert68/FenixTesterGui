package grpc_out_GuiTestCaseBuilderServer

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"context"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// SendSaveAllPinnedTestInstructionsAndTestInstructionContainers - Save pinned TestInstructions and TestInstructionContainers
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) SendSaveAllPinnedTestInstructionsAndTestInstructionContainers(pinnedTestInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.SavePinnedTestInstructionsAndPreCreatedTestInstructionContainersMessage) (returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse) {

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
	/*
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer func() {
			//TODO Fixa så att denna inte görs som allt går bra
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID": "2d688330-025f-492a-b318-bb9374bf76ec",
			}).Error("Running Defer Cancel function")
			cancel()
		}()

	*/

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiTestCaseBuilderServer == sharedCode.GCP {

		// Set logger in GCP-package
		gcp.GcpObject.SetLogger(grpcOut.logger)

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.GcpObject.GenerateGCPAccessToken(ctx)
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
	returnMessage, err = fenixGuiTestCaseCaseBuilderServerGrpcClient.SaveAllPinnedTestInstructionsAndTestInstructionContainers(ctx, pinnedTestInstructionsAndTestContainersMessage)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "b0743d37-cdda-425d-b391-74fb0ab0890e",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendSaveAllPinnedTestInstructionsAndTestInstructionContainers'")

	} else if returnMessage.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                     "7aa6164b-9a51-47fb-8279-f4be52ebab3d",
			"Message from FenixTestGuiBuilderServer": returnMessage.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendSaveAllPinnedTestInstructionsAndTestInstructionContainers'")
	}

	return returnMessage

}
