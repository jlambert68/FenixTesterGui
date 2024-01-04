package main

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"FenixTesterGui/grpc_in"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"FenixTesterGui/grpc_out_GuiTestCaseBuilderServer"
	"FenixTesterGui/gui"
	"FenixTesterGui/messageStreamEngine"
	"FenixTesterGui/restAPI"
	"context"
	"errors"
	"fmt"
	uuidGenerator "github.com/google/uuid"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path/filepath"
	"time"
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

	// Create Unique Uuid for run time instance used as identification when communication with GuiExecutionServer
	sharedCode.ApplicationRunTimeUuid = uuidGenerator.New().String()
	fmt.Println("sharedCode.ApplicationRunTimeUuid: " + sharedCode.ApplicationRunTimeUuid)

	// Set up BackendObject
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
	var filePathName string
	filePathName = "" //"logs/mylog008.log"

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

	// Initiate gcp-logger
	gcp.GcpObject.SetLogger(sharedCode.Logger)

	// Set dial address to GuiServer
	fenixTesterGuiObject.subPackageObjects.restAPI.SetDialAddressString(grpc_out_GuiTestCaseBuilderServer.FenixGuiTestCaseBuilderServerAddressToDial)
	fenixTesterGuiObject.subPackageObjects.uiServer.SetDialAddressString(grpc_out_GuiTestCaseBuilderServer.FenixGuiTestCaseBuilderServerAddressToDial)

	// Secure that Access token has been created and 'sharedCode.CurrentUserAuthenticatedTowardsGCP' got a value
	var returnMessageAckNack bool
	var returnMessageString string
	var ctx context.Context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()
	_, returnMessageAckNack, returnMessageString = gcp.GcpObject.GenerateGCPAccessToken(ctx, gcp.TargetServerGuiTestCaseBuilderServer)
	if returnMessageAckNack == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                  "8e3bfe8a-e920-44fc-bdf4-f476bfd048b2",
			"returnMessageString": returnMessageString,
		}).Fatalln("Couldn't generate access token")
	}
	fmt.Println("va fannnnn")

	// Start RestApi-server
	//go fenixTesterGuiObject.subPackageObjects.restAPI.RestAPIServer()

	// Clean up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Start Backend gRPC-server
	go fenixTesterGuiObject.subPackageObjects.grpcIn.InitGrpcServer()

	defer func() {
		// Inform GuiExecutionServer that TesterGui is closing down
		var ackNackResponse *fenixExecutionServerGuiGrpcApi.AckNackResponse
		ackNackResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.SendTesterGuiIsClosingDown()

		if ackNackResponse.AckNack == false {

			errorId := "cb8be454-8b62-4a22-af36-2182d31260fa"
			err := errors.New(fmt.Sprintf("couldn't do 'SendTesterGuiIsClosingDown' to GuiExecutionServer due to error: '%s', {error: %s} [ErrorID: %s]", ackNackResponse.Comments, errorId))

			fmt.Println(err) // TODO Send on Error-channel

			//os.Exit(0)

		}
	}()

	// Inform GuiExecutionServer that TesterGui is starting up
	// Initiate TestCaseExecution
	var ackNackResponse *fenixExecutionServerGuiGrpcApi.AckNackResponse
	ackNackResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.SendTesterGuiIsStartingUp()

	if ackNackResponse.AckNack == false {

		errorId := "fb15f862-7754-4a55-aa16-4c0bda086c4f"
		err := errors.New(fmt.Sprintf("couldn't do 'SendTesterGuiIsStartingUp' to GuiExecutionServer due to error: '%s', {error: %s} [ErrorID: %s]", ackNackResponse.Comments, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		//os.Exit(0)

	} else {

	}

	// Start up MessageStreamEngine (PuBSub) to receive ExecutionStatus from GuiExecutionServer
	messageStreamEngine.InitiateAndStartMessageStreamChannelReader()

	// Start UI Server
	fenixTesterGuiObject.subPackageObjects.uiServer.StartUIServer()

}
