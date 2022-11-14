package grpc_out_GuiExecutionServer

import (
	"FenixTesterGui/common_code"
	"crypto/tls"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// ********************************************************************************************************************

// Set upp connection and Dial to FenixGuiExecutionServer
func (grpcOut *GRPCOutGuiExecutionServerStruct) SetConnectionToFenixGuiExecutionServer() (returnMessage *fenixExecutionServerGuiGrpcApi.AckNackResponse) {
	var err error
	var opts []grpc.DialOption

	//When running on GCP then use credential otherwise not
	if sharedCode.ExecutionLocationForFenixGuiExecutionServer == sharedCode.GCP {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})

		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(creds),
		}
	}

	// Set up connection to FenixTestDataSyncServer
	// When run on GCP, use credentials
	if sharedCode.ExecutionLocationForFenixGuiExecutionServer == sharedCode.GCP {
		// Run on GCP
		remoteFenixGuiExecutionServerConnection, err = grpc.Dial(FenixGuiExecutionServerAddressToDial, opts...)
	} else {
		// Run Local
		remoteFenixGuiExecutionServerConnection, err = grpc.Dial(FenixGuiExecutionServerAddressToDial, grpc.WithInsecure())
	}
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "50b59b1b-57ce-4c27-aa84-617f0cde3100",
			"FenixGuiExecutionServerAddressToDial": FenixGuiExecutionServerAddressToDial,
			"error message":                        err,
		}).Error("Did not connect to FenixGuiExecutionServer via gRPC")

		// Create response message for when no success Dial was to be made
		// Set Error codes to return message
		var errorCodes []fenixExecutionServerGuiGrpcApi.ErrorCodesEnum
		var errorCode fenixExecutionServerGuiGrpcApi.ErrorCodesEnum

		errorCode = fenixExecutionServerGuiGrpcApi.ErrorCodesEnum_ERROR_UNSPECIFIED
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixExecutionServerGuiGrpcApi.AckNackResponse{
			AckNack:                      false,
			Comments:                     "Couldn't call FenixGuiExecutionServer",
			ErrorCodes:                   errorCodes,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(GetHighestFenixGuiExecutionServerProtoFileVersion()),
		}

		return returnMessage

	} else {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                                   "0c650bbc-45d0-4029-bd25-4ced9925a059",
			"FenixGuiExecutionServerAddressToDial": FenixGuiExecutionServerAddressToDial,
		}).Info("gRPC connection OK to FenixGuiExecutionServer")

		// Creates a new Clients
		fenixGuiExecutionServerGrpcClient = fenixExecutionServerGuiGrpcApi.NewFenixExecutionServerGuiGrpcServicesClient(remoteFenixGuiExecutionServerConnection)

	}

	return nil
}

// GetHighestFenixGuiServerProtoFileVersion ********************************************************************************************************************
// Get the highest FenixProtoFileVersionEnumeration
func GetHighestFenixGuiExecutionServerProtoFileVersion() int32 {

	// Check if there already is a 'highestFenixGuiExecutionServerProtoFileVersion' saved, if so use that one
	if highestFenixGuiExecutionServerProtoFileVersion != -1 {
		return highestFenixGuiExecutionServerProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestFenixGuiExecutionServerProtoFileVersion = maxValue

	return highestFenixGuiExecutionServerProtoFileVersion
}

// SetLogger
// Set to use the same Logger reference as is used by central part of system
func (grpcOut *GRPCOutGuiExecutionServerStruct) SetLogger(logger *logrus.Logger) {

	//GrpcOutGuiExecutionServerObject = GRPCOutGuiExecutionServerStruct{}

	grpcOut.logger = logger

	return

}

// SetDialAddressString
// Set the Dial Address, which was received from environment variables
func (grpcOut *GRPCOutGuiExecutionServerStruct) SetDialAddressString(dialAddress string) {
	FenixGuiExecutionServerAddressToDial = dialAddress

	return

}
