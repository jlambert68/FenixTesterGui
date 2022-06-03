package grpc_out

import (
	common_config "FenixTesterGui/common_code"
	"crypto/tls"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// ********************************************************************************************************************

// SetConnectionToFenixTestDataSyncServer - Set upp connection and Dial to FenixTestDataSyncServer
func (GrpcOut *GRPCOutStruct) setConnectionToFenixGuiBuilderServer() {

	var err error
	var opts []grpc.DialOption

	//When running on GCP then use credential otherwise not
	if common_config.ExecutionLocationForFenixGuiServer == common_config.GCP {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})

		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(creds),
		}
	}

	// Set up connection to FenixTestDataSyncServer
	// When run on GCP, use credentials
	if common_config.ExecutionLocationForFenixGuiServer == common_config.GCP {
		// Run on GCP
		remoteFenixGuiBuilderServerConnection, err = grpc.Dial(FenixGuiBuilderServerAddressToDial, opts...)
	} else {
		// Run Local
		remoteFenixGuiBuilderServerConnection, err = grpc.Dial(FenixGuiBuilderServerAddressToDial) //, grpc.WithInsecure())
	}
	if err != nil {
		GrpcOut.logger.WithFields(logrus.Fields{
			"ID":                                 "50b59b1b-57ce-4c27-aa84-617f0cde3100",
			"fenixGuiBuilderServerAddressToDial": FenixGuiBuilderServerAddressToDial,
			"error message":                      err,
		}).Error("Did not connect to FenixGuiBuilderServer via gRPC")
		//os.Exit(0)
	} else {
		GrpcOut.logger.WithFields(logrus.Fields{
			"ID": "0c650bbc-45d0-4029-bd25-4ced9925a059",
			"fenixGuiTestCaseBuilderServer_address_to_dial": FenixGuiBuilderServerAddressToDial,
		}).Info("gRPC connection OK to FenixTestDataSyncServer")

		// Creates a new Clients
		fenixGuiBuilderServerGrpcClient = fenixGuiTestCaseBuilderServerGrpcApi.NewFenixTestCaseBuilderServerGrpcServicesClient(remoteFenixGuiBuilderServerConnection)

	}
}

// GetHighestFenixGuiServerProtoFileVersion ********************************************************************************************************************
// Get the highest FenixProtoFileVersionEnumeration
func (GrpcOut *GRPCOutStruct) GetHighestFenixGuiServerProtoFileVersion() int32 {

	// Check if there already is a 'highestFenixProtoFileVersion' saved, if so use that one
	if highestFenixGuiServerProtoFileVersion != -1 {
		return highestFenixGuiServerProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestFenixGuiServerProtoFileVersion = maxValue

	return highestFenixGuiServerProtoFileVersion
}
