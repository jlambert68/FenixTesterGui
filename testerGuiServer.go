package main

import (
	"FenixTesterGui/grpc_in"
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
		fenixTesterGuiObject.localreferencs.grpc_in.StopGrpcServer()

		//log.Println("Close DB_session: %v", DB_session)
		//DB_session.Close()
	}
}

func fenixGuiTestCaseBuilderServerMain() {

	// Connect to CloudDB
	//fenixSyncShared.ConnectToDB()

	// Set up BackendObject
	fenixTesterGuiObject = &fenixGuiBuilderProxyServerObjectStruct{
		localreferencs: localreferenceLibraryStruct{
			grpc_in: &grpc_in.GrpcIn,
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

	// Clean up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Start RestApi-server
	go grpc_in.fenixGuiBuilderProxyServerObject.restAPIServer()

	// Start Backend gRPC-server
	grpc_in.fenixGuiBuilderProxyServerObject.InitGrpcServer()

}
