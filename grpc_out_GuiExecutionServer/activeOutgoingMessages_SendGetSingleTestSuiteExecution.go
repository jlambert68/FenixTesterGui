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

// SendGetSingleTestSuiteExecution - Retrieve a single TestSuiteExecution with all its content
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendGetSingleTestSuiteExecution(
	getSingleTestSuiteExecutionRequest *fenixExecutionServerGuiGrpcApi.GetSingleTestSuiteExecutionRequest) (
	getSingleTestSuiteExecutionResponse *fenixExecutionServerGuiGrpcApi.GetSingleTestSuiteExecutionResponse) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id":                                 "f9cc2562-ab81-4206-ac69-c3694d250d87",
		"getSingleTestSuiteExecutionRequest": getSingleTestSuiteExecutionRequest,
	}).Debug("Incoming 'grpcOut - SendGetSingleTestSuiteExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id":                                  "110cbb59-c789-4a17-b12f-0b9179b2b4c4",
		"getSingleTestSuiteExecutionResponse": &getSingleTestSuiteExecutionResponse,
	}).Debug("Outgoing 'grpcOut - SendGetSingleTestSuiteExecution'")

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	ctx = context.Background()

	// Set up connection to Server
	ctx, err = grpcOut.setConnectionToFenixGuiExecutionMessageServer_new(ctx)
	//grpcOut.setConnectionToFenixGuiTestSuiteBuilderServer()
	if err != nil {
		// When error
		getSingleTestSuiteExecutionResponse = &fenixExecutionServerGuiGrpcApi.GetSingleTestSuiteExecutionResponse{
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   err.Error(),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
		}

		return getSingleTestSuiteExecutionResponse
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
			getSingleTestSuiteExecutionResponse = &fenixExecutionServerGuiGrpcApi.GetSingleTestSuiteExecutionResponse{
				AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
					AckNack:    false,
					Comments:   returnMessageString,
					ErrorCodes: nil,
					ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
						GetHighestFenixGuiExecutionServerProtoFileVersion()),
				},
			}

			return getSingleTestSuiteExecutionResponse

		}

	}

	// Do the gRPC-call
	getSingleTestSuiteExecutionResponse, err = fenixGuiExecutionServerGrpcClient.GetSingleTestSuiteExecution(
		ctx,
		getSingleTestSuiteExecutionRequest)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "7606184b-c67a-4e19-ac19-35f1f3e1fd3c",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendGetSingleTestSuiteExecution'")

		getSingleTestSuiteExecutionResponse = &fenixExecutionServerGuiGrpcApi.GetSingleTestSuiteExecutionResponse{
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   fmt.Sprintf("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendGetSingleTestSuiteExecution'. ErrorMessage: '%s'", err.Error()),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
		}

		return getSingleTestSuiteExecutionResponse

	} else if getSingleTestSuiteExecutionResponse.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "74ba7348-cec6-43b7-b651-2d9e4400304f",
			"Message from FenixGuiExecutionServer": getSingleTestSuiteExecutionResponse.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendGetSingleTestSuiteExecution'")
	}

	return getSingleTestSuiteExecutionResponse

}
