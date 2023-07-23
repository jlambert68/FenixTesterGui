package grpc_out_GuiTestCaseBuilderServer

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"context"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// LoadHashesForTestCases - Load all hashes for the TestCases supplied into gRPC-call to TestCaseBuilderServer.
// Hashes are used to check if a TestCase is changed in database compared to TestGui-version of the TestCase.
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) LoadHashesForTestCases(userId string, testCaseUuidList []string) (returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesHashResponse) {

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	ctx = context.Background()

	// Set up connection to Server
	ctx, err = grpcOut.setConnectionToFenixGuiTestCaseBuilderServer_new(ctx)
	//grpcOut.setConnectionToFenixGuiTestCaseBuilderServer()
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

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestCasesHashResponse{
				AckNack:         ackNackResponse,
				TestCasesHashes: nil,
			}
			return returnMessage
		}
	}

	// Create the request message
	var testCasesHashRequest *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesHashRequest
	testCasesHashRequest = &fenixGuiTestCaseBuilderServerGrpcApi.TestCasesHashRequest{
		UserIdentification: &fenixGuiTestCaseBuilderServerGrpcApi.UserIdentificationMessage{
			UserId: userId,
			ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
				grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
		},
		TestCaseUuids: testCaseUuidList,
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

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestCasesHashResponse{
				AckNack:         ackNackResponse,
				TestCasesHashes: nil,
			}
			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiTestCaseCaseBuilderServerGrpcClient.GetTestCasesHashes(ctx, testCasesHashRequest)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "23a33bdf-5f4c-4f57-ba1d-a8bb0bfeb5bb",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'GetTestCasesHashes'")

	} else if returnMessage.AckNack.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                     "93409dc3-d334-420a-85da-68ab1d5100cd",
			"Message from FenixTestGuiBuilderServer": returnMessage.AckNack.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'GetTestCasesHashes'")
	}

	return returnMessage

}
