package main

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"FenixTesterGui/grpc_out_GuiTestCaseBuilderServer"
	"fmt"
	"log"
	"os"
	"strconv"
)

// mustGetEnv is a helper function for getting environment variables.
// Displays a lethal warning if the environment variable is not set.
func mustGetenv(environmentVariable string) string {

	// Verify that Variable exists in Environment Variable Map

	// Create the build variable name
	var buildInjectedVariableNameAsString string
	var buildInjectedVariableValue *string
	buildInjectedVariableNameAsString = "BuildVariable" + environmentVariable

	buildInjectedVariableValue, exist := buildVariablesMap[buildInjectedVariableNameAsString]
	if exist == false {
		// If the 'Build Injected Variable' is missing then end this misery programs life
		log.Fatalln("Environment variable " + buildInjectedVariableNameAsString + " doesn't exist in 'buildVariablesMap'")
	}

	// Extract
	environmentVariableValue := os.Getenv(environmentVariable)
	if environmentVariableValue == "" {
		// No environment variable found so try for build injected variable instead

		// If the 'Build Injected Variable' is empty then end this misery programs life
		if *buildInjectedVariableValue == "" {
			log.Fatalf("Warning: %s environment variable not set by injecting value at build time.\n", environmentVariable)
		} else {
			environmentVariableValue = *buildInjectedVariableValue
		}
	}
	return environmentVariableValue
}

func init() {

	//executionLocationForThisApplication := flag.String("startupType", "0", "The application should be started with one of the following: LOCALHOST_NODOCKER, LOCALHOST_DOCKER, GCP")
	//flag.Parse()

	//convertBuildInjectedVariablesToMapStructure()

	var err error

	// *********************** Extract environment variables for 'This Application' ***********************
	// Get Environment variable to tell how 'this program' was started
	var executionLocationForThisApplication = mustGetenv("ExecutionLocationForThisApplication")

	switch executionLocationForThisApplication {
	case "LOCALHOST_NODOCKER":
		sharedCode.ExecutionLocationForThisApplication = sharedCode.LocalhostNoDocker

	case "LOCALHOST_DOCKER":
		sharedCode.ExecutionLocationForThisApplication = sharedCode.LocalhostDocker

	case "GCP":
		sharedCode.ExecutionLocationForThisApplication = sharedCode.GCP

	default:
		fmt.Println("Unknown Execution location for 'This application': " + executionLocationForThisApplication + ". Expected one of the following: LOCALHOST_NODOCKER, LOCALHOST_DOCKER, GCP")
		os.Exit(0)

	}

	// Get Environment variable for how to scale this applications GUI
	fyneScale := mustGetenv("FYNE_SCALE")

	// Secure that it is a number
	_, err = strconv.ParseFloat(fyneScale, 64)
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'FYNE_SCALE' to a numeric, error: ", err)
		os.Exit(0)

	} else {
		sharedCode.FYNE_SCALE = fyneScale
	}

	/*
		// Get Environment variable to tell if the application should run as a Tray Application or not
		var runAsTrayApplication = mustGetenv("RunAsTrayApplication")

		switch runAsTrayApplication {
		case "YES":
			sharedCode.RunAsTrayApplication = true

		case "NO":
			sharedCode.RunAsTrayApplication = false

		default:
			fmt.Println("Unknown value for 'RunAsTrayApplication': " + runAsTrayApplication + ". Expected one of the following: 'YES', 'NO'")
			os.Exit(0)

		}
	*/

	// Get local gRPC port for this application. gRPC is used for checking connection to backend services among other stuff
	sharedCode.ApplicationGrpcPort, err = strconv.Atoi(mustGetenv("ApplicationGrpcPort"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'ApplicationGrpcPort' to an integer, error: ", err)
		os.Exit(0)
	}

	// *********************** Extract environment variables for 'GUI-TestCaseBuilderServer' ***********************
	// Get Environment variable to tell how 'GUI-TestCaseBuilderServer' was started
	var executionLocationForFenixGuiTestCaseBuilderServer = mustGetenv("ExecutionLocationForFenixGuiTestCaseBuilderServer")

	switch executionLocationForFenixGuiTestCaseBuilderServer {
	case "LOCALHOST_NODOCKER":
		sharedCode.ExecutionLocationForFenixGuiTestCaseBuilderServer = sharedCode.LocalhostNoDocker

	case "LOCALHOST_DOCKER":
		sharedCode.ExecutionLocationForFenixGuiTestCaseBuilderServer = sharedCode.LocalhostDocker

	case "GCP":
		sharedCode.ExecutionLocationForFenixGuiTestCaseBuilderServer = sharedCode.GCP

	default:
		fmt.Println("Unknown Execution location for FenixGuiTestCaseBuilderServer: " + executionLocationForFenixGuiTestCaseBuilderServer + ". Expected one of the following: LOCALHOST_NODOCKER, LOCALHOST_DOCKER, GCP")
		os.Exit(0)

	}

	// Get Environment variable for Address to 'GUI-TestCaseBuilderServer'
	sharedCode.FenixGuiTestCaseBuilderServerAddress = mustGetenv("FenixGuiTestCaseBuilderServerAddress")

	// Get Environment variable regarding Port for 'GUI-TestCaseBuilderServer'
	sharedCode.FenixGuiTestCaseBuilderServerPort, err = strconv.Atoi(mustGetenv("FenixGuiTestCaseBuilderServerPort"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'FenixGuiTestCaseBuilderServerPort' to an integer, error: ", err)
		os.Exit(0)

	}

	// Create Dial-address to 'GUI-TestCaseBuilderServer'
	grpc_out_GuiTestCaseBuilderServer.FenixGuiTestCaseBuilderServerAddressToDial = sharedCode.FenixGuiTestCaseBuilderServerAddress + ":" + strconv.Itoa(sharedCode.FenixGuiTestCaseBuilderServerPort)

	// *********************** Extract environment variables for 'GUI-ExecutionServer' ***********************
	// Get Environment variable to tell how 'GUI-ExecutionServer' was started
	var executionLocationForFenixGuiExecutionServer = mustGetenv("ExecutionLocationForFenixGuiExecutionServer")

	switch executionLocationForFenixGuiExecutionServer {
	case "LOCALHOST_NODOCKER":
		sharedCode.ExecutionLocationForFenixGuiExecutionServer = sharedCode.LocalhostNoDocker

	case "LOCALHOST_DOCKER":
		sharedCode.ExecutionLocationForFenixGuiExecutionServer = sharedCode.LocalhostDocker

	case "GCP":
		sharedCode.ExecutionLocationForFenixGuiExecutionServer = sharedCode.GCP

	default:
		fmt.Println("Unknown Execution location for FenixGuiExecutionServer: " + executionLocationForFenixGuiExecutionServer + ". Expected one of the following: LOCALHOST_NODOCKER, LOCALHOST_DOCKER, GCP")
		os.Exit(0)

	}

	// Get Environment variable for Address to 'GUI-ExecutionServer'
	sharedCode.FenixGuiExecutionServerAddress = mustGetenv("FenixGuiExecutionServerAddress")

	// Get Environment variable regarding Port for 'GUI-ExecutionServer'
	sharedCode.FenixGuiExecutionServerPort, err = strconv.Atoi(mustGetenv("FenixGuiExecutionServerPort"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'FenixGuiExecutionServerPort' to an integer, error: ", err)
		os.Exit(0)

	}

	// Create Dial-address to 'GUI-ExecutionServer'
	grpc_out_GuiExecutionServer.FenixGuiExecutionServerAddressToDial = sharedCode.FenixGuiExecutionServerAddress + ":" + strconv.Itoa(sharedCode.FenixGuiExecutionServerPort)

	// Get Environment variable 'GCPAuthentication' to decide if we should use GCP-authenticfication or not
	var tempBoolAsString string
	var tempBool bool
	tempBoolAsString = mustGetenv("GCPAuthentication")
	tempBool, err = strconv.ParseBool(tempBoolAsString)
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'GCPAuthentication' to a boolean, error: ", err)
		os.Exit(0)
	}
	sharedCode.GCPAuthentication = tempBool

	// Get Environment variable 'UseServiceAccountForGuiExecutionServer' to decide if we should use a service account to log into GCP or to use Users login credentials
	tempBoolAsString = mustGetenv("UseServiceAccountForGuiExecutionServer")
	tempBool, err = strconv.ParseBool(tempBoolAsString)
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'UseServiceAccountForGuiExecutionServer' to a boolean, error: ", tempBoolAsString, err)
		os.Exit(0)
	}
	sharedCode.UseServiceAccountForGuiExecutionServer = tempBool

	// Get Environment variable 'UseServiceAccountForGuiTestCaseBuilderServer' to decide if we should use a service account to log into GCP or to use Users login credentials
	tempBoolAsString = mustGetenv("UseServiceAccountForGuiTestCaseBuilderServer")
	tempBool, err = strconv.ParseBool(tempBoolAsString)
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'UseServiceAccountForGuiTestCaseBuilderServer' to a boolean, error: ", tempBoolAsString, err)
		os.Exit(0)
	}
	sharedCode.UseServiceAccountForGuiTestCaseBuilderServer = tempBool

	// Extract OAuth 2.0 Client ID
	sharedCode.AuthClientId = mustGetenv("AuthClientId")

	// Extract OAuth 2.0 Client Secret
	sharedCode.AuthClientSecret = mustGetenv("AuthClientSecret")

	// Extract the GCP-project
	sharedCode.GcpProject = mustGetenv("GcpProject")

	// Extract PubSub-Topic-base for where to send the 'TestExecutionsStatus'
	sharedCode.TestExecutionStatusPubSubTopicBase = mustGetenv("TestExecutionStatusPubSubTopicBase")

	// Extract local path to Service-Account file
	sharedCode.LocalServiceAccountPath = mustGetenv("LocalServiceAccountPath")
	// The only way have an OK space is to replace an existing character
	if sharedCode.LocalServiceAccountPath == "#" {
		sharedCode.LocalServiceAccountPath = ""
	}

	// Extract if Proxy-server should be used for outgoing requests
	tempBoolAsString = mustGetenv("ShouldProxyServerBeUsed")
	tempBool, err = strconv.ParseBool(tempBoolAsString)
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'ShouldProxyServerBeUsed' to a boolean, error: ", tempBoolAsString, err)
		os.Exit(0)
	}
	sharedCode.ShouldProxyServerBeUsed = tempBool

	// Extract URL to Proxy-server for outgoing requests
	sharedCode.ProxyServerURL = mustGetenv("ProxyServerURL")

}
