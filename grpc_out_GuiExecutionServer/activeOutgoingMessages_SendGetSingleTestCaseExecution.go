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

// SendGetSingleTestCaseExecution - Retrieve a single TestCaseExecution with all its content
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendGetSingleTestCaseExecution(
	getSingleTestCaseExecutionRequest *fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionRequest) (
	getSingleTestCaseExecutionResponse *fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionResponse) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id":                                "f7138122-9982-4db8-8288-07e2b472bd38",
		"getSingleTestCaseExecutionRequest": getSingleTestCaseExecutionRequest,
	}).Debug("Incoming 'grpcOut - SendGetSingleTestCaseExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id":                                 "269ef318-5995-447d-aee1-4bce50b5365c",
		"getSingleTestCaseExecutionResponse": &getSingleTestCaseExecutionResponse,
	}).Debug("Outgoing 'grpcOut - SendGetSingleTestCaseExecution'")

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
		getSingleTestCaseExecutionResponse = &fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionResponse{
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   err.Error(),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
			TestCaseExecutionResponse: nil,
		}

		return getSingleTestCaseExecutionResponse
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
			getSingleTestCaseExecutionResponse = &fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionResponse{
				AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
					AckNack:    false,
					Comments:   returnMessageString,
					ErrorCodes: nil,
					ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
						GetHighestFenixGuiExecutionServerProtoFileVersion()),
				},
				TestCaseExecutionResponse: nil,
			}

			return getSingleTestCaseExecutionResponse

		}

	}

	// Do the gRPC-call
	getSingleTestCaseExecutionResponse, err = fenixGuiExecutionServerGrpcClient.GetSingleTestCaseExecution(
		ctx,
		getSingleTestCaseExecutionRequest)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "cdaa21c7-cc81-4676-951b-6c0dd43079ac",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendGetSingleTestCaseExecution'")

		getSingleTestCaseExecutionResponse = &fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionResponse{
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   fmt.Sprintf("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendGetSingleTestCaseExecution'. ErrorMessage: '%s'", err.Error()),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
			TestCaseExecutionResponse: nil,
		}

		return getSingleTestCaseExecutionResponse

	} else if getSingleTestCaseExecutionResponse.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "74ba7348-cec6-43b7-b651-2d9e4400304f",
			"Message from FenixGuiExecutionServer": getSingleTestCaseExecutionResponse.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendGetSingleTestCaseExecution'")
	}

	return getSingleTestCaseExecutionResponse

}
