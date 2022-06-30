package grpc_out

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
func (grpcOut *GRPCOutStruct) SendAreYouAliveToFenixGuiBuilderServer() (returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse) {

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	// Set up connection to Server
	returnMessage = grpcOut.setConnectionToFenixGuiBuilderServer()
	// If there was no connection to backend then return that message
	if returnMessage != nil {
		return returnMessage
	}

	// Create the request message
	emptyParameter := &fenixGuiTestCaseBuilderServerGrpcApi.EmptyParameter{

		ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
			grpcOut.GetHighestFenixGuiServerProtoFileVersion()),
	}

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		//TODO Fixa så att denna inte görs som allt går bra
		grpcOut.logger.WithFields(logrus.Fields{
			"ID": "88bbfecb-b9a4-4e2f-92e0-cabbbdf75dc8",
		}).Error("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if common_config.ExecutionLocationForFenixGuiServer == common_config.GCP {

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.Gcp.GenerateGCPAccessTokenForAuthorizedUser(ctx)
		if returnMessageAckNack == false {
			// When error
			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   returnMessageString,
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
					grpcOut.GetHighestFenixGuiServerProtoFileVersion()),
			}

			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiBuilderServerGrpcClient.AreYouAlive(ctx, emptyParameter)

	// Shouldn't happen
	if err != nil {
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":    "818aaf0b-4112-4be4-97b9-21cc084c7b8b",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendAreYouAliveToFenixGuiBuilderServer'")

	} else if returnMessage.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":                                     "2ecbc800-2fb6-4e88-858d-a421b61c5529",
			"Message from FenixTestGuiBuilderServer": returnMessage.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendAreYouAliveToFenixGuiBuilderServer'")
	}

	return returnMessage

}

// ********************************************************************************************************************

// SendListAllAvailableTestInstructionsAndTestInstructionContainers - Get available TestInstructions and TestInstructionContainers
func (grpcOut *GRPCOutStruct) SendListAllAvailableTestInstructionsAndTestInstructionContainers(userId string) (returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

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
			"ID": "797c00e1-510d-4cbe-a48b-dc63828ecd7e",
		}).Error("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if common_config.ExecutionLocationForFenixGuiServer == common_config.GCP {

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

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage{
				ImmatureTestInstructions:          nil,
				ImmatureTestInstructionContainers: nil,
				AckNackResponse:                   ackNackResponse,
			}

			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiBuilderServerGrpcClient.ListAllAvailableTestInstructionsAndTestInstructionContainers(ctx, userIdentificationMessage)

	// Shouldn't happen
	if err != nil {
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":    "d7235084-33e5-43a2-9fa7-dfb05ec6869e",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendListAllAvailableTestInstructionsAndTestInstructionContainers'")

	} else if returnMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":                                     "30e6f1ee-202a-47bf-a2c4-5066d0f8cf75",
			"Message from FenixTestGuiBuilderServer": returnMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendListAllAvailableTestInstructionsAndTestInstructionContainers'")
	}

	return returnMessage

}

// SendListAllAvailablePinnedTestInstructionsAndTestInstructionContainers - Get pinned TestInstructions and TestInstructionContainers
func (grpcOut *GRPCOutStruct) SendListAllAvailablePinnedTestInstructionsAndTestInstructionContainers(userId string) (returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

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
			"ID": "c5ba19bd-75ff-4366-818d-745d4d7f1a52",
		}).Error("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if common_config.ExecutionLocationForFenixGuiServer == common_config.GCP {

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.Gcp.GenerateGCPAccessTokenForAuthorizedUser(ctx)
		if returnMessageAckNack == false {
			// When error
			ackNackResponse := &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   returnMessageString,
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
					grpcOut.GetHighestFenixGuiServerProtoFileVersion()),
			}

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage{
				AvailablePinnedTestInstructions:                    nil,
				AvailablePinnedPreCreatedTestInstructionContainers: nil,
				AckNackResponse: ackNackResponse,
			}

			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiBuilderServerGrpcClient.ListAllAvailablePinnedTestInstructionsAndTestInstructionContainers(ctx, userIdentificationMessage)

	// Shouldn't happen
	if err != nil {
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":    "7bff3257-a193-4d07-83aa-f106f6f734a0",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendListAllAvailablePinnedTestInstructionsAndTestInstructionContainers'")

	} else if returnMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":                                     "72f79764-2549-4ce7-867e-16cd0f414dff",
			"Message from FenixTestGuiBuilderServer": returnMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendListAllAvailablePinnedTestInstructionsAndTestInstructionContainers'")
	}

	return returnMessage

}

// SendSaveAllPinnedTestInstructionsAndTestInstructionContainers - Save pinned TestInstructions and TestInstructionContainers
func (grpcOut *GRPCOutStruct) SendSaveAllPinnedTestInstructionsAndTestInstructionContainers(pinnedTestInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.SavePinnedTestInstructionsAndPreCreatedTestInstructionContainersMessage) (returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse) {

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	// Set up connection to Server
	grpcOut.setConnectionToFenixGuiBuilderServer()

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		//TODO Fixa så att denna inte görs som allt går bra
		grpcOut.logger.WithFields(logrus.Fields{
			"ID": "2d688330-025f-492a-b318-bb9374bf76ec",
		}).Error("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if common_config.ExecutionLocationForFenixGuiServer == common_config.GCP {

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.Gcp.GenerateGCPAccessTokenForAuthorizedUser(ctx)
		if returnMessageAckNack == false {
			// When error
			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   returnMessageString,
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
					grpcOut.GetHighestFenixGuiServerProtoFileVersion()),
			}

			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiBuilderServerGrpcClient.SaveAllPinnedTestInstructionsAndTestInstructionContainers(ctx, pinnedTestInstructionsAndTestContainersMessage)

	// Shouldn't happen
	if err != nil {
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":    "b0743d37-cdda-425d-b391-74fb0ab0890e",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendSaveAllPinnedTestInstructionsAndTestInstructionContainers'")

	} else if returnMessage.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":                                     "7aa6164b-9a51-47fb-8279-f4be52ebab3d",
			"Message from FenixTestGuiBuilderServer": returnMessage.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendSaveAllPinnedTestInstructionsAndTestInstructionContainers'")
	}

	return returnMessage

}
