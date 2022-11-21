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

	// Set up connection to Server
	var ackNackresponse *fenixExecutionServerGuiGrpcApi.AckNackResponse
	ackNackresponse = grpcOut.SetConnectionToFenixGuiExecutionServer()
	// If there was no connection to backend then return that message
	if ackNackresponse != nil {
		listTestCasesUnderExecutionResponse = &fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionResponse{
			AckNackResponse:         ackNackresponse,
			TestCasesUnderExecution: nil,
		}

		return listTestCasesUnderExecutionResponse
	}

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		//TODO Fixa så att denna inte görs som allt går bra
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID": "66c2dcf3-0b51-417b-b402-438a60c7eb86",
		}).Error("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiExecutionServer == sharedCode.GCP {

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.GcpObject.GenerateGCPAccessTokenForAuthorizedUser(ctx)
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
