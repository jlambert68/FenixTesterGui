package main

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_in"
	"FenixTesterGui/grpc_out_GuiTestCaseBuilderServer"
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
			"Id": "14dc7af0-5e46-4073-b920-8dffee7ca307",
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

	// Set up BackendObjec 	t
	fenixTesterGuiObject = &fenixGuiBuilderProxyServerObjectStruct{
		subPackageObjects: &referencesStruct{
			grpcIn: &grpc_in.GRPCInStruct{},
			restAPI: &restAPI.RestApiStruct{
				GrpcOut: &grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct{},
			},
			uiServer: &gui.GlobalUIServerStruct{},
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

	// Initiate logging
	fenixTesterGuiObject.InitLogger(filePathName)

	// Store logger reference in shared code 'for all to use'
	sharedCode.Logger = fenixTesterGuiObject.logger

	// Set logger for sub packages
	fenixTesterGuiObject.subPackageObjects.grpcIn.SetLogger(fenixTesterGuiObject.logger)
	fenixTesterGuiObject.subPackageObjects.restAPI.SetLogger(fenixTesterGuiObject.logger)
	fenixTesterGuiObject.subPackageObjects.uiServer.SetLogger(fenixTesterGuiObject.logger)

	// Set dial address to GuiServer
	fenixTesterGuiObject.subPackageObjects.restAPI.SetDialAddressString(grpc_out_GuiTestCaseBuilderServer.FenixGuiTestCaseBuilderServerAddressToDial)
	fenixTesterGuiObject.subPackageObjects.uiServer.SetDialAddressString(grpc_out_GuiTestCaseBuilderServer.FenixGuiTestCaseBuilderServerAddressToDial)

	// Clean up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Start RestApi-server
	//go fenixTesterGuiObject.subPackageObjects.restAPI.RestAPIServer()

	// Start Backend gRPC-server
	go fenixTesterGuiObject.subPackageObjects.grpcIn.InitGrpcServer()

	// Start UI Server
	fenixTesterGuiObject.subPackageObjects.uiServer.StartUIServer()

}
