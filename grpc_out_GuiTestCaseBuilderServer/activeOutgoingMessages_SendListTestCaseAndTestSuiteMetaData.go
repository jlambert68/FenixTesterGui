package grpc_out_GuiTestCaseBuilderServer

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"context"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// SendListTestCaseAndTestSuiteMetaData - Load all TestCaseMetaData that the User can use when creating TestCasesMapPtr
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) SendListTestCaseAndTestSuiteMetaData() (
	returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListTestCaseAndTestSuiteMetaDataResponseMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "c167ad66-144a-4dec-b39f-54ef30a20db4",
	}).Debug("Incoming 'grpcOut - SendListTestCaseAndTestSuiteMetaData'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "9bc0770f-7b48-4159-9cd2-f159965d4ac9",
	}).Debug("Outgoing 'grpcOut - SendListTestCaseAndTestSuiteMetaData'")

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

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListTestCaseAndTestSuiteMetaDataResponseMessage{
				AckNackResponse:                        ackNackResponse,
				TestCaseAndTestSuiteMetaDataForDomains: nil,
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

			returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListTestCaseAndTestSuiteMetaDataResponseMessage{
				AckNackResponse:                        ackNackResponse,
				TestCaseAndTestSuiteMetaDataForDomains: nil,
			}

			return returnMessage
		}

	}

	// Do the gRPC-call
	returnMessage, err = fenixGuiTestCaseCaseBuilderServerGrpcClient.ListTestCaseAndTestSuiteMetaData(ctx, userIdentificationMessage)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "c3df6f17-8e33-410a-b8f2-69185a91b270",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendListTestCaseAndTestSuiteMetaData'")

		// When error
		ackNackResponse := &fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendListTestCaseAndTestSuiteMetaData'",
			ErrorCodes: nil,
			ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
				grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
		}

		returnMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ListTestCaseAndTestSuiteMetaDataResponseMessage{
			AckNackResponse:                        ackNackResponse,
			TestCaseAndTestSuiteMetaDataForDomains: nil,
		}

		return returnMessage

	} else if returnMessage.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                     "75f6764a-c8ea-43d5-ab05-0c824b4905ce",
			"Message from FenixTestGuiBuilderServer": returnMessage.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'SendListTestCaseAndTestSuiteMetaData'")
	}

	return returnMessage

}
