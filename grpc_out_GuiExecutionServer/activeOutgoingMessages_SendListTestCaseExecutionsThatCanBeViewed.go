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

// SendListTestCaseExecutionsThatCanBeViewed - List all TestCaseExecutions that can be views, used for producing a list
// that the user has access to
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendListTestCaseExecutionsThatCanBeViewed(
	latestUniqueTestCaseExecutionDatabaseRowId int32,
	onlyRetrieveLimitedSizedBatch bool,
	batchSize int32,
	retrieveAllExecutionsForSpecificTestCaseUuid bool,
	specificTestCaseUuid string,
	testCaseExecutionFromTimeStamp time.Time,
	testCaseExecutionToTimeStamp time.Time) (
	listTestCaseExecutionsResponse *fenixExecutionServerGuiGrpcApi.ListTestCaseExecutionsResponse) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "43969a9b-6d51-4979-87d4-d9033a624fda",
	}).Debug("Incoming 'grpcOut - SendListTestCaseExecutionsThatCanBeViewed'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "24dffe32-ffa2-492e-8777-daf15a961ed7",
	}).Debug("Outgoing 'grpcOut - SendListTestCaseExecutionsThatCanBeViewed'")

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

			listTestCaseExecutionsResponse = &fenixExecutionServerGuiGrpcApi.ListTestCaseExecutionsResponse{
				AckNackResponse:                            ackNackResponse,
				TestCaseExecutionsList:                     nil,
				LatestUniqueTestCaseExecutionDatabaseRowId: 0,
				MoreRowsExists:                             false,
			}

			return listTestCaseExecutionsResponse
		}
	}

	// Create the request message
	var listTestCaseExecutionsRequest *fenixExecutionServerGuiGrpcApi.ListTestCaseExecutionsRequest
	listTestCaseExecutionsRequest = &fenixExecutionServerGuiGrpcApi.ListTestCaseExecutionsRequest{
		UserAndApplicationRunTimeIdentification: &fenixExecutionServerGuiGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
			ApplicationRunTimeUuid: "",
			UserIdOnComputer:       sharedCode.CurrentUserIdLogedInOnComputer,
			GCPAuthenticatedUser:   sharedCode.CurrentUserAuthenticatedTowardsGCP,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				GetHighestFenixGuiExecutionServerProtoFileVersion()),
		},
		LatestUniqueTestCaseExecutionDatabaseRowId: latestUniqueTestCaseExecutionDatabaseRowId,
		OnlyRetrieveLimitedSizedBatch:              onlyRetrieveLimitedSizedBatch,
		BatchSize:                                  batchSize,
		TestCaseExecutionFromTimeStamp:             timestamppb.New(testCaseExecutionFromTimeStamp),
		TestCaseExecutionToTimeStamp:               timestamppb.New(testCaseExecutionToTimeStamp),
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

			listTestCaseExecutionsResponse = &fenixExecutionServerGuiGrpcApi.ListTestCaseExecutionsResponse{
				AckNackResponse:                            ackNackResponse,
				TestCaseExecutionsList:                     nil,
				LatestUniqueTestCaseExecutionDatabaseRowId: 0,
				MoreRowsExists:                             false,
			}

			return listTestCaseExecutionsResponse
		}

	}

	// Do the gRPC-call
	listTestCaseExecutionsResponse, err = fenixGuiExecutionServerGrpcClient.ListTestCaseExecutions(
		ctx, listTestCaseExecutionsRequest)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "eac0186c-52e2-4254-bb3a-27edca729258",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendListTestCaseExecutionsThatCanBeViewed'")

		// When error
		ackNackResponse := &fenixExecutionServerGuiGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Problem to do gRPC-call to FenixGuiExecutionServer for 'SendListTestCaseExecutionsThatCanBeViewed'",
			ErrorCodes: nil,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				GetHighestFenixGuiExecutionServerProtoFileVersion()),
		}

		listTestCaseExecutionsResponse = &fenixExecutionServerGuiGrpcApi.ListTestCaseExecutionsResponse{
			AckNackResponse:                            ackNackResponse,
			TestCaseExecutionsList:                     nil,
			LatestUniqueTestCaseExecutionDatabaseRowId: 0,
			MoreRowsExists:                             false,
		}

		return listTestCaseExecutionsResponse

	} else if listTestCaseExecutionsResponse.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "bd94654a-53c4-41b5-bb83-e61aa7fe9fb7",
			"Message from FenixGuiExecutionServer": listTestCaseExecutionsResponse.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendListTestCaseExecutionsThatCanBeViewed'")
	}

	return listTestCaseExecutionsResponse

}
