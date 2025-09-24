package grpc_out_GuiExecutionServer

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"context"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// SendListTestSuiteExecutionsThatCanBeViewed - List all TestSuiteExecutions that can be views, used for producing a list
// that the user has access to
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendListTestSuiteExecutionsThatCanBeViewed(
	latestUniqueTestSuiteExecutionDatabaseRowId int32,
	onlyRetrieveLimitedSizedBatch bool,
	batchSize int32,
	retrieveAllExecutionsForSpecificTestSuiteUuid bool,
	specificTestSuiteUuid string,
	testSuiteExecutionFromTimeStamp time.Time,
	testSuiteExecutionToTimeStamp time.Time) (
	listTestSuiteExecutionsResponse *fenixExecutionServerGuiGrpcApi.ListTestSuiteExecutionsResponse) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "5a1001e7-a935-4896-bda8-e508a8b4429f",
	}).Debug("Incoming 'grpcOut - SendListTestSuiteExecutionsThatCanBeViewed'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "f9b85668-c92c-46f2-9cdd-c03b951f9757",
	}).Debug("Outgoing 'grpcOut - SendListTestSuiteExecutionsThatCanBeViewed'")

	/*
		ackNackResponse := &fenixExecutionServerGuiGrpcApi.AckNackResponse{
			AckNack:    true,
			Comments:   "",
			ErrorCodes: nil,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				GetHighestFenixGuiExecutionServerProtoFileVersion()),
		}

		listTestSuiteExecutionsResponse = &fenixExecutionServerGuiGrpcApi.ListTestSuiteExecutionsResponse{
			AckNackResponse:                            ackNackResponse,
			TestSuiteExecutionsList:                     nil,
			LatestUniqueTestSuiteExecutionDatabaseRowId: 0,
			MoreRowsExists:                             false,
		}
		return listTestSuiteExecutionsResponse


	*/
	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	ctx = context.Background()

	// Set up connection to Server
	ctx, err = grpcOut.setConnectionToFenixGuiExecutionMessageServer_new(ctx)

	if err != nil {
		if returnMessageAckNack == false {
			// When error
			ackNackResponse := &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   err.Error(),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			}

			listTestSuiteExecutionsResponse = &fenixExecutionServerGuiGrpcApi.ListTestSuiteExecutionsResponse{
				AckNackResponse:                             ackNackResponse,
				TestSuiteExecutionsList:                     nil,
				LatestUniqueTestSuiteExecutionDatabaseRowId: 0,
				MoreRowsExists:                              false,
			}

			return listTestSuiteExecutionsResponse
		}
	}

	// Create the request message
	var listTestSuiteExecutionsRequest *fenixExecutionServerGuiGrpcApi.ListTestSuiteExecutionsRequest
	listTestSuiteExecutionsRequest = &fenixExecutionServerGuiGrpcApi.ListTestSuiteExecutionsRequest{
		UserAndApplicationRunTimeIdentification: &fenixExecutionServerGuiGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
			ApplicationRunTimeUuid: "",
			UserIdOnComputer:       sharedCode.CurrentUserIdLogedInOnComputer,
			GCPAuthenticatedUser:   sharedCode.CurrentUserAuthenticatedTowardsGCP,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				GetHighestFenixGuiExecutionServerProtoFileVersion()),
		},
		LatestUniqueTestSuiteExecutionDatabaseRowId:   latestUniqueTestSuiteExecutionDatabaseRowId,
		OnlyRetrieveLimitedSizedBatch:                 onlyRetrieveLimitedSizedBatch,
		BatchSize:                                     batchSize,
		RetrieveAllExecutionsForSpecificTestSuiteUuid: retrieveAllExecutionsForSpecificTestSuiteUuid,
		SpecificTestSuiteUuid:                         specificTestSuiteUuid,
		TestSuiteExecutionFromTimeStamp:               timestamppb.New(testSuiteExecutionFromTimeStamp),
		TestSuiteExecutionToTimeStamp:                 timestamppb.New(testSuiteExecutionToTimeStamp),
	}

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiExecutionServer == sharedCode.GCP {

		// Set logger in GCP-package
		gcp.GcpObject.SetLogger(grpcOut.logger)

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.GcpObject.GenerateGCPAccessToken(ctx, gcp.TargetServerGuiExecutionServer)
		if returnMessageAckNack == false {
			// When error
			ackNackResponse := &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   returnMessageString,
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			}

			listTestSuiteExecutionsResponse = &fenixExecutionServerGuiGrpcApi.ListTestSuiteExecutionsResponse{
				AckNackResponse:                             ackNackResponse,
				TestSuiteExecutionsList:                     nil,
				LatestUniqueTestSuiteExecutionDatabaseRowId: 0,
				MoreRowsExists:                              false,
			}

			return listTestSuiteExecutionsResponse
		}

	}

	// Do the gRPC-call
	listTestSuiteExecutionsResponse, err = fenixGuiExecutionServerGrpcClient.ListTestSuiteExecutions(
		ctx, listTestSuiteExecutionsRequest)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "102ea536-37e6-430d-8484-874cc1f4be3f",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendListTestSuiteExecutionsThatCanBeViewed'")

		// When error
		ackNackResponse := &fenixExecutionServerGuiGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Problem to do gRPC-call to FenixGuiExecutionServer for 'SendListTestSuiteExecutionsThatCanBeViewed'",
			ErrorCodes: nil,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				GetHighestFenixGuiExecutionServerProtoFileVersion()),
		}

		listTestSuiteExecutionsResponse = &fenixExecutionServerGuiGrpcApi.ListTestSuiteExecutionsResponse{
			AckNackResponse:                             ackNackResponse,
			TestSuiteExecutionsList:                     nil,
			LatestUniqueTestSuiteExecutionDatabaseRowId: 0,
			MoreRowsExists:                              false,
		}

		return listTestSuiteExecutionsResponse

	} else if listTestSuiteExecutionsResponse.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "bd43452d-0aaf-4b82-96f3-6b4361876d36",
			"Message from FenixGuiExecutionServer": listTestSuiteExecutionsResponse.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendListTestSuiteExecutionsThatCanBeViewed'")
	}

	return listTestSuiteExecutionsResponse

}
