package grpc_out_GuiExecutionServer

import (
	"FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"FenixTesterGui/grpcurl"
	"context"
	"crypto/tls"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

// ********************************************************************************************************************

// SetConnectionToFenixExecutionWorkerServer - Set upp connection and Dial to FenixExecutionServer
func (grpcOut *GRPCOutGuiExecutionServerStruct) setConnectionToFenixGuiExecutionMessageServer_new(ctx context.Context) (_ context.Context, err error) {
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
		if sharedCode.ExecutionLocationForFenixGuiExecutionServer == sharedCode.GCP {
			// GuiExecutionServer is running in GCP

			// Should ProxyServer be used for outgoing conenctions
			if sharedCode.ShouldProxyServerBeUsed == true {
				// Use Proxy
				remoteFenixGuiExecutionServerConnection, err = gcp.GRPCDialer("")
				if err != nil {
					sharedCode.Logger.WithFields(logrus.Fields{
						"ID":                 "4a07f2ca-6bcb-4ec3-b2e1-0755e023d6bb",
						"error message":      err,
						"dialAttemptCounter": dialAttemptCounter,
					}).Error("Couldn't generate gRPC-connection to GuiExecutionServer via Proxy Server")
					continue
				}

			} else {
				// Don't use Proxy
				ctx, newGrpcClientConnection = dialFromGrpcurl(ctx, FenixGuiExecutionServerAddressToDial)
				remoteFenixGuiExecutionServerConnection = newGrpcClientConnection
				//remoteFenixExecutionWorkerServerConnection, err = grpc.Dial(common_config.FenixExecutionWorkerAddressToDial, opts...)

			}
		} else {
			// Run Local
			remoteFenixGuiExecutionServerConnection, err = grpc.Dial(FenixGuiExecutionServerAddressToDial, grpc.WithInsecure())
		}
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":                                   "50b59b1b-57ce-4c27-aa84-617f0cde3100",
				"FenixGuiExecutionServerAddressToDial": FenixGuiExecutionServerAddressToDial,
				"error message":                        err,
				"dialAttemptCounter":                   dialAttemptCounter,
			}).Error("Did not connect to FenixExecutionServer via gRPC")

			// Add to counter for how many Dial attempts that have been done
			dialAttemptCounter = dialAttemptCounter + 1

			// Only return the error after last attempt
			if dialAttemptCounter >= numberOfDialAttempts {
				return nil, err
			}

		} else {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":                                   "556f52bf-630c-44e8-b666-cd7c562bc100",
				"FenixGuiExecutionServerAddressToDial": FenixGuiExecutionServerAddressToDial,
			}).Info("gRPC connection OK to FenixGuiExecutionServer")

			// Creates a new Client
			fenixGuiExecutionServerGrpcClient = fenixExecutionServerGuiGrpcApi.NewFenixExecutionServerGuiGrpcServicesForGuiClientClient(remoteFenixGuiExecutionServerConnection)

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
			"ID":                                   "4f90c990-b3c7-4d53-88cf-f354504357a3",
			"FenixGuiExecutionServerAddressToDial": FenixGuiExecutionServerAddressToDial,
		}).Info("gRPC connection OK to FenixGuiExecutionServer")

		// Creates a new Clients
		fenixGuiExecutionServerGrpcClient = fenixExecutionServerGuiGrpcApi.NewFenixExecutionServerGuiGrpcServicesForGuiClientClient(remoteFenixGuiExecutionServerConnection)

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
