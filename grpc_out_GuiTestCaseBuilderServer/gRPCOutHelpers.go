package grpc_out_GuiTestCaseBuilderServer

import (
	common_config "FenixTesterGui/common_code"
	"crypto/tls"
	fenixTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// ********************************************************************************************************************

// Set upp connection and Dial to FenixGuiTestCaseBuilderServer
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) setConnectionToFenixGuiTestCaseBuilderServer() (returnMessage *fenixTestCaseBuilderServerGrpcApi.AckNackResponse) {
	var err error
	var opts []grpc.DialOption

	//When running on GCP then use credential otherwise not
	if common_config.ExecutionLocationForFenixGuiTestCaseBuilderServer == common_config.GCP {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})

		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(creds),
		}
	}

	// Set up connection to FenixTestDataSyncServer
	// When run on GCP, use credentials
	if common_config.ExecutionLocationForFenixGuiTestCaseBuilderServer == common_config.GCP {
		// Run on GCP
		remoteFenixGuiTestCaseBuilderServerConnection, err = grpc.Dial(grpcOut.fenixGuiBuilderServerAddressToDial, opts...)
	} else {
		// Run Local
		remoteFenixGuiTestCaseBuilderServerConnection, err = grpc.Dial(grpcOut.fenixGuiBuilderServerAddressToDial, grpc.WithInsecure())
	}
	if err != nil {
		grpcOut.logger.WithFields(logrus.Fields{
			"ID":                                 "50b59b1b-57ce-4c27-aa84-617f0cde3100",
			"fenixGuiBuilderServerAddressToDial": grpcOut.fenixGuiBuilderServerAddressToDial,
			"error message":                      err,
		}).Error("Did not connect to FenixGuiBuilderServer via gRPC")
		//os.Exit(0)

		// Create response message for when no success dail was able to be made

		// Set Error codes to return message
		var errorCodes []fenixTestCaseBuilderServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestCaseBuilderServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestCaseBuilderServerGrpcApi.ErrorCodesEnum_ERROR_UNSPECIFIED
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestCaseBuilderServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Couldn't call GuiServer",
			ErrorCodes: errorCodes,
		}

		return returnMessage

	} else {
		grpcOut.logger.WithFields(logrus.Fields{
			"ID": "0c650bbc-45d0-4029-bd25-4ced9925a059",
			"fenixGuiTestCaseBuilderServer_address_to_dial": grpcOut.fenixGuiBuilderServerAddressToDial,
		}).Info("gRPC connection OK to FenixTestDataSyncServer")

		// Creates a new Clients
		fenixGuiTestCaseCaseBuilderServerGrpcClient = fenixTestCaseBuilderServerGrpcApi.NewFenixTestCaseBuilderServerGrpcServicesClient(remoteFenixGuiTestCaseBuilderServerConnection)

	}

	return nil
}

// GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion ********************************************************************************************************************
// Get the highest FenixProtoFileVersionEnumeration
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion() int32 {

	// Check if there already is a 'highestFenixProtoFileVersion' saved, if so use that one
	if highestFenixGuiTestCaseBuilderServerProtoFileVersion != -1 {
		return highestFenixGuiTestCaseBuilderServerProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestFenixGuiTestCaseBuilderServerProtoFileVersion = maxValue

	return highestFenixGuiTestCaseBuilderServerProtoFileVersion
}

// SetLogger
// Set to use the same Logger reference as is used by central part of system
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) SetLogger(logger *logrus.Logger) {

	//grpcOutGuiTestCaseBuilderServerObject = GRPCOutGuiTestCaseBuilderServerStruct{}

	grpcOut.logger = logger

	return

}

// SetDialAddressString
// Set the Dial Address, which was received from environment variables
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) SetDialAddressString(dialAddress string) {
	grpcOut.fenixGuiBuilderServerAddressToDial = dialAddress

	return

}
