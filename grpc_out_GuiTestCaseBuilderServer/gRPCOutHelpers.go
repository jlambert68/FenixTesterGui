package grpc_out_GuiTestCaseBuilderServer

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpcurl"
	"context"
	"crypto/tls"
	fenixTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

// ********************************************************************************************************************

// SetConnectionToFenixExecutionWorkerServer - Set upp connection and Dial to FenixExecutionServer
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) setConnectionToFenixGuiTestCaseBuilderServer_new(ctx context.Context) (_ context.Context, err error) {
	/* THis is done in the new
	var opts []grpc.DialOption

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		panic(fmt.Sprintf("cannot load root CA certs, err: %s", err))
	}

	//When running on GCP then use credential otherwise not

	if common_config.ExecutionLocationForFenixExecutionWorkerServer == common_config.GCP {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            systemRoots,
		})

		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(creds),
		}
	}
	*/

	// slice with sleep time, in milliseconds, between each attempt to Dial to Worker
	var sleepTimeBetweenDialAttempts []int
	sleepTimeBetweenDialAttempts = []int{100, 100, 200, 200, 300, 300, 500, 500, 600, 1000} // Total: 3.6 seconds

	// Do multiple attempts to do connection to Execution Worker
	var numberOfDialAttempts int
	var dialAttemptCounter int
	numberOfDialAttempts = len(sleepTimeBetweenDialAttempts)
	dialAttemptCounter = 0

	for {

		// Set up connection to Fenix Execution Worker
		// When run on GCP, use credentials
		var newGrpcClientConnection *grpc.ClientConn
		if sharedCode.ExecutionLocationForFenixGuiTestCaseBuilderServer == sharedCode.GCP {
			// Run on GCP
			ctx, newGrpcClientConnection = dialFromGrpcurl(ctx, FenixGuiTestCaseBuilderServerAddressToDial)
			remoteFenixGuiTestCaseBuilderServerConnection = newGrpcClientConnection
			//remoteFenixExecutionWorkerServerConnection, err = grpc.Dial(common_config.FenixExecutionWorkerAddressToDial, opts...)
		} else {
			// Run Local
			remoteFenixGuiTestCaseBuilderServerConnection, err = grpc.Dial(FenixGuiTestCaseBuilderServerAddressToDial, grpc.WithInsecure())
		}
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID": "edd4d8fd-c15d-48ec-abdd-ddeb302c3674",
				"FenixGuiTestCaseBuilderServerAddressToDial": FenixGuiTestCaseBuilderServerAddressToDial,
				"error message":      err,
				"dialAttemptCounter": dialAttemptCounter,
			}).Error("Did not connect to FenixGuiTestCaseBuilderServer via gRPC")

			// Add to counter for how many Dial attempts that have been done
			dialAttemptCounter = dialAttemptCounter + 1

			// Only return the error after last attempt
			if dialAttemptCounter >= numberOfDialAttempts {
				return nil, err
			}

		} else {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID": "d5cccdd1-a72a-4697-924a-d517816e3083",
				"FenixGuiTestCaseBuilderServerAddressToDial": FenixGuiTestCaseBuilderServerAddressToDial,
			}).Info("gRPC connection OK to FenixGuiExecutionServer")

			// Creates a new Client
			fenixGuiTestCaseCaseBuilderServerGrpcClient = fenixTestCaseBuilderServerGrpcApi.NewFenixTestCaseBuilderServerGrpcServicesClient(remoteFenixGuiTestCaseBuilderServerConnection)

			break
		}

		// Sleep for some time before retrying to connect
		time.Sleep(time.Millisecond * time.Duration(sleepTimeBetweenDialAttempts[dialAttemptCounter-1]))

	}

	return ctx, err
}

var (
	isUnixSocket func() bool
)

func dialFromGrpcurl(ctx context.Context, target string) (context.Context, *grpc.ClientConn) {

	//target := grpc_out_GuiExecutionServer.FenixGuiExecutionServerAddressToDial

	dialTime := 10 * time.Second

	ctx, cancel := context.WithTimeout(ctx, dialTime)
	defer cancel()
	var opts []grpc.DialOption

	var creds credentials.TransportCredentials

	var tlsConf *tls.Config

	creds = credentials.NewTLS(tlsConf)

	grpcurlUA := "FenixCAConnector"
	//if grpcurl.version == grpcurl.no_version {
	//	grpcurlUA = "grpcurl/dev-build (no version set)"
	//}

	opts = append(opts, grpc.WithUserAgent(grpcurlUA))
	//opts = append(opts, grpc.WithNoProxy())

	network := "tcp"
	if isUnixSocket != nil && isUnixSocket() {
		network = "unix"
	}

	cc, err := grpcurl.BlockingDial(ctx, network, target, creds, opts...)
	if err != nil {
		log.Panicln("Failed to Dial, ", target, err.Error())
	}
	return ctx, cc

}

// Set upp connection and Dial to FenixGuiTestCaseBuilderServer
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) setConnectionToFenixGuiTestCaseBuilderServer() (returnMessage *fenixTestCaseBuilderServerGrpcApi.AckNackResponse) {
	var err error
	var opts []grpc.DialOption

	//When running on GCP then use credential otherwise not
	if sharedCode.ExecutionLocationForFenixGuiTestCaseBuilderServer == sharedCode.GCP {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})

		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(creds),
		}
	}

	// Set up connection to FenixTestDataSyncServer
	// When run on GCP, use credentials
	if sharedCode.ExecutionLocationForFenixGuiTestCaseBuilderServer == sharedCode.GCP {
		// Run on GCP
		remoteFenixGuiTestCaseBuilderServerConnection, err = grpc.Dial(FenixGuiTestCaseBuilderServerAddressToDial, opts...)
	} else {
		// Run Local
		remoteFenixGuiTestCaseBuilderServerConnection, err = grpc.Dial(FenixGuiTestCaseBuilderServerAddressToDial, grpc.WithInsecure())
	}
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID": "50b59b1b-57ce-4c27-aa84-617f0cde3100",
			"FenixGuiTestCaseBuilderServerAddressToDial": FenixGuiTestCaseBuilderServerAddressToDial,
			"error message": err,
		}).Error("Did not connect to FenixGuiTestCaseBuilderServer via gRPC")
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
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID": "526e3dad-3534-49ab-b107-9c26d1d45d0e",
			"FenixGuiTestCaseBuilderServerAddressToDial": FenixGuiTestCaseBuilderServerAddressToDial,
		}).Info("gRPC connection OK to FenixGuiTestCaseBuilderServer")

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

	//GrpcOutGuiTestCaseBuilderServerObject = GRPCOutGuiTestCaseBuilderServerStruct{}

	grpcOut.logger = logger

	return

}

// SetDialAddressString
// Set the Dial Address, which was received from environment variables
func (grpcOut *GRPCOutGuiTestCaseBuilderServerStruct) SetDialAddressString(dialAddress string) {
	FenixGuiTestCaseBuilderServerAddressToDial = dialAddress

	return

}
