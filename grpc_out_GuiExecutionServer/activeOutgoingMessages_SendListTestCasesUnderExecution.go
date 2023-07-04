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

// SendListTestCasesUnderExecution - Load all TestCaseExecutions that is under execution
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendListTestCasesUnderExecution(listTestCasesUnderExecutionRequest *fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionRequest) (listTestCasesUnderExecutionResponse *fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionResponse) {

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
		listTestCasesUnderExecutionResponse = &fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionResponse{
			TestCasesUnderExecution: nil,
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   err.Error(),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
		}
		return listTestCasesUnderExecutionResponse
	}

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
			listTestCasesUnderExecutionResponse = &fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionResponse{
				TestCasesUnderExecution: nil,
				AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
					AckNack:    false,
					Comments:   returnMessageString,
					ErrorCodes: nil,
					ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
						GetHighestFenixGuiExecutionServerProtoFileVersion()),
				},
			}
			return listTestCasesUnderExecutionResponse

		}

	}

	// Do the gRPC-call
	listTestCasesUnderExecutionResponse, err = fenixGuiExecutionServerGrpcClient.ListTestCasesUnderExecution(ctx, listTestCasesUnderExecutionRequest)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "4cede160-bdca-4357-aab8-ab4f6bedb534",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'ListTestCasesOnExecutionQueue'")

		listTestCasesUnderExecutionResponse = &fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionResponse{
			TestCasesUnderExecution: nil,
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   fmt.Sprintf("Problem to do gRPC-call to FenixGuiExecutionServer for 'ListTestCasesUnderExecution'. ErrorMessage: '%s'", err.Error()),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
		}
		return listTestCasesUnderExecutionResponse

	} else if listTestCasesUnderExecutionResponse.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "8121ea5c-7f61-48c7-8d6e-87c80f6def50",
			"Message from FenixGuiExecutionServer": listTestCasesUnderExecutionResponse.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'ListTestCasesUnderExecution'")
	}

	return listTestCasesUnderExecutionResponse

}
