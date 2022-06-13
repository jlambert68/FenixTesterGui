package main

import (
	"FenixTesterGui/grpc_in"
	"FenixTesterGui/grpc_out"
	"FenixTesterGui/gui"
	"FenixTesterGui/restAPI"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path/filepath"
)

// Used for only process cleanup once
var cleanupProcessed = false

func cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		fenixTesterGuiObject.logger.WithFields(logrus.Fields{
			"Id": "9b1527db-b215-45cd-ac1e-dd37b276edb5",
		}).Info("Clean up and shut down servers")

		// Stop Backend gRPC Server
		fenixTesterGuiObject.subPackageObjects.grpcIn.StopGrpcServer()

		//log.Println("Close DB_session: %v", DB_session)
		//DB_session.Close()
	}
}

func fenixGuiBuilderServerMain() {

	// Connect to CloudDB
	//fenixSyncShared.ConnectToDB()

	// Set up BackendObject
	fenixTesterGuiObject = &fenixGuiBuilderProxyServerObjectStruct{
		subPackageObjects: &referencesStruct{
			grpcIn: &grpc_in.GRPCInStruct{},
			restAPI: &restAPI.RestApiStruct{
				GrpcOut: &grpc_out.GRPCOutStruct{},
			},
			uiServer: &gui.UIServerStruct{},
		},
	}

	// Init logger
	// When application is run as tray application then use text file as log
	var filePathName = ""
	var err error

	if fenixTesterGuiObject.runAsTrayApplication == true {
		// Get path for this application

		logfilename := "mylog.log"
		filePathName, err = filepath.Abs(logfilename)
		if err != nil {
			log.Println("Couldn't generate filePathName for log: ", err)
			os.Exit(0)
		}
	}

	fenixTesterGuiObject.InitLogger(filePathName)

	// Set logger for sub packages
	fenixTesterGuiObject.subPackageObjects.grpcIn.SetLogger(fenixTesterGuiObject.logger)
	fenixTesterGuiObject.subPackageObjects.restAPI.SetLogger(fenixTesterGuiObject.logger)
	fenixTesterGuiObject.subPackageObjects.uiServer.SetLogger(fenixTesterGuiObject.logger)

	// Set dial address to GuiServer
	fenixTesterGuiObject.subPackageObjects.restAPI.SetDialAddressString(fenixGuiBuilderServerAddressToDial)
	fenixTesterGuiObject.subPackageObjects.uiServer.SetDialAddressString(fenixGuiBuilderServerAddressToDial)

	// Clean up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Start RestApi-server
	go fenixTesterGuiObject.subPackageObjects.restAPI.RestAPIServer()

	// Start Backend gRPC-server
	go fenixTesterGuiObject.subPackageObjects.grpcIn.InitGrpcServer()

	// Start UI Server
	//fenixTesterGuiObject.subPackageObjects.uiServer. StartUIServer()
	fenixTesterGuiObject.subPackageObjects.uiServer.StartUIServer()

}
