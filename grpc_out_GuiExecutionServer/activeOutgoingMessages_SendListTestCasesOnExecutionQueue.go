package grpc_out_GuiExecutionServer

import (
	"FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"fmt"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// ********************************************************************************************************************

// SendListTestCasesOnExecutionQueue - Load all TestCaseExecutions that exists on ExecutionQueue
func (grpcOut *GRPCOutGuiExecutionServerStruct) SendListTestCasesOnExecutionQueue(listTestCasesInExecutionQueueRequest *fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueRequest) (listTestCasesInExecutionQueueResponse *fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueResponse) {

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string
	var err error

	ctx = context.Background()

	// Set up connection to Server
	ctx, err = grpcOut.setConnectionToFenixGuiExecutionMessageServer_new(ctx)
	//grpcOut.setConnectionToFenixGuiTestCaseBuilderServer()
	if err != nil {
		if err != nil {
			// When error
			listTestCasesInExecutionQueueResponse = &fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueResponse{
				TestCasesInExecutionQueue: nil,
				AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
					AckNack:    false,
					Comments:   err.Error(),
					ErrorCodes: nil,
					ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
						GetHighestFenixGuiExecutionServerProtoFileVersion()),
				},
			}
			return listTestCasesInExecutionQueueResponse
		}
	}

	// Do gRPC-call
	//ctx := context.Background()
	/*
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer func() {
			//TODO Fixa så att denna inte görs som allt går bra
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID": "66c2dcf3-0b51-417b-b402-438a60c7eb86",
			}).Error("Running Defer Cancel function")
			cancel()
		}()

	*/

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiExecutionServer == sharedCode.GCP {

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.GcpObject.GenerateGCPAccessToken(ctx, gcp.TargetServerGuiExecutionServer)
		if returnMessageAckNack == false {
			// When error
			listTestCasesInExecutionQueueResponse = &fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueResponse{
				TestCasesInExecutionQueue: nil,
				AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
					AckNack:    false,
					Comments:   returnMessageString,
					ErrorCodes: nil,
					ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
						GetHighestFenixGuiExecutionServerProtoFileVersion()),
				},
			}
			return listTestCasesInExecutionQueueResponse

		}

	}

	// Do the gRPC-call
	listTestCasesInExecutionQueueResponse, err = fenixGuiExecutionServerGrpcClient.ListTestCasesOnExecutionQueue(ctx, listTestCasesInExecutionQueueRequest)

	// Shouldn't happen
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "4cede160-bdca-4357-aab8-ab4f6bedb534",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'ListTestCasesOnExecutionQueue'")

		listTestCasesInExecutionQueueResponse = &fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueResponse{
			TestCasesInExecutionQueue: nil,
			AckNackResponse: &fenixExecutionServerGuiGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   fmt.Sprintf("Problem to do gRPC-call to FenixGuiExecutionServer for 'ListTestCasesOnExecutionQueue'. ErrorMessage: '%s'", err.Error()),
				ErrorCodes: nil,
				ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
					GetHighestFenixGuiExecutionServerProtoFileVersion()),
			},
		}
		return listTestCasesInExecutionQueueResponse

	} else if listTestCasesInExecutionQueueResponse.AckNackResponse.AckNack == false {
		// FenixTestGuiBuilderServer couldn't handle gPRC call
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "8121ea5c-7f61-48c7-8d6e-87c80f6def50",
			"Message from FenixGuiExecutionServer": listTestCasesInExecutionQueueResponse.AckNackResponse.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'ListTestCasesOnExecutionQueue'")
	}

	return listTestCasesInExecutionQueueResponse

}
