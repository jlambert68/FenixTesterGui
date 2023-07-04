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
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendListTestCasesWithFinishedExecutions(listTestCasesWithFinishedExecutionsRequest *fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsRequest) (listTestCasesWithFinishedExecutionsResponse *fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsResponse) {

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
		listTestCasesWithFinishedExecutionsResponse = &fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsResponse{
			TestCaseWithFinishedExecution: nil,
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   err.Error(),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
		}
		return listTestCasesWithFinishedExecutionsResponse
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
			listTestCasesWithFinishedExecutionsResponse = &fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsResponse{
				TestCaseWithFinishedExecution: nil,
				AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
					AckNack:    false,
					Comments:   returnMessageString,
					ErrorCodes: nil,
					ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
						GetHighestFenixGuiExecutionServerProtoFileVersion()),
				},
			}
			return listTestCasesWithFinishedExecutionsResponse

		}

	}

	// Do the gRPC-call
	listTestCasesWithFinishedExecutionsResponse, err = fenixGuiExecutionServerGrpcClient.ListTestCasesWithFinishedExecutions(ctx, listTestCasesWithFinishedExecutionsRequest)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "32d6db17-f938-4f97-8289-4bbc347af947",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'ListTestCasesWithFinishedExecutions'")

		listTestCasesWithFinishedExecutionsResponse = &fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsResponse{
			TestCaseWithFinishedExecution: nil,
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   fmt.Sprintf("Problem to do gRPC-call to FenixGuiExecutionServer for 'ListTestCasesWithFinishedExecutions'. ErrorMessage: '%s'", err.Error()),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
		}
		return listTestCasesWithFinishedExecutionsResponse

	} else if listTestCasesWithFinishedExecutionsResponse.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "e52ec64c-ea50-4948-91a9-a972fb4dca9c",
			"Message from FenixGuiExecutionServer": listTestCasesWithFinishedExecutionsResponse.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'ListTestCasesUnderExecution'")
	}

	return listTestCasesWithFinishedExecutionsResponse

}
