package grpc_in

import (
	"FenixTesterGui/common_code"
	fenixUserGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixUserGui/fenixUserGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

// InitGrpcServer - Set up and start Backend gRPC-server
func (grpcIn *GRPCInStruct) InitGrpcServer() {

	var err error

	// Find first non allocated port from defined start port
	grpcIn.logger.WithFields(logrus.Fields{
		"Id": "054bc0ef-93bb-4b75-8630-74e3823f71da",
	}).Info("Backend Server tries to start")

	grpcIn.logger.WithFields(logrus.Fields{
		"Id":                                  "ca3593b1-466b-4536-be91-5e038de178f4",
		"common_config.ApplicationGrpcPort: ": sharedCode.ApplicationGrpcPort,
	}).Info("Start listening on:")
	lis, err = net.Listen("tcp", ":"+strconv.Itoa(sharedCode.ApplicationGrpcPort))

	if err != nil {
		grpcIn.logger.WithFields(logrus.Fields{
			"Id":    "ad7815b3-63e8-4ab1-9d4a-987d9bd94c76",
			"err: ": err,
		}).Error("failed to listen:")
	} else {
		grpcIn.logger.WithFields(logrus.Fields{
			"Id":                                  "ba070b9b-5d57-4c0a-ab4c-a76247a50fd3",
			"common_config.ApplicationGrpcPort: ": sharedCode.ApplicationGrpcPort,
		}).Info("Success in listening on port:")

	}

	// Creates a new RegisterWorkerServer gRPC server
	//go func() {
	grpcIn.logger.WithFields(logrus.Fields{
		"Id": "b0ccffb5-4367-464c-a3bc-460cafed16cb",
	}).Info("Starting Backend gRPC Server")

	registerFenixUserGuiServer = grpc.NewServer()
	fenixUserGuiGrpcApi.RegisterFenixUserGuiGrpcServicesServer(registerFenixUserGuiServer, &fenixUserGuiGrpcServicesServer{})

	// Register Reflection for this gRPC-server.
	reflection.Register(registerFenixUserGuiServer)

	grpcIn.logger.WithFields(logrus.Fields{
		"Id":                                  "e843ece9-b707-4c60-b1d8-14464305e68f",
		"common_config.ApplicationGrpcPort: ": sharedCode.ApplicationGrpcPort,
	}).Info("'RegisterFenixTestCaseBuilderServerGrpcServicesServer' for FenixGuiBuilderProxyServer")
	err = registerFenixUserGuiServer.Serve(lis)
	if err != nil {
		grpcIn.logger.WithFields(logrus.Fields{
			"Id":  "d0f4de61-32fd-4ca9-868b-bdc4765e2d44",
			"err": err,
		}).Error("Couldn't register gRPC-server: 'registerFenixUserGuiServer'")
	}

	//}()

}

// StopGrpcServer - Stop Backend gRPC-server
func (grpcIn *GRPCInStruct) StopGrpcServer() {

	grpcIn.logger.WithFields(logrus.Fields{}).Info("Gracefully stop for: 'registerFenixUserGuiServer'")
	registerFenixUserGuiServer.GracefulStop()

	grpcIn.logger.WithFields(logrus.Fields{
		"common_config.ApplicationGrpcPort: ": sharedCode.ApplicationGrpcPort,
	}).Info("Close net.Listing")
	_ = lis.Close()

}
