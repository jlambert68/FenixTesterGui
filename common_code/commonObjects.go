package sharedCode

// ***********************************************************************************************************
// The following variables receives their values from environment variables

// Where is the client running
var ExecutionLocation ExecutionLocationTypeType

// ExecutionLocationForFenixGuiTestCaseBuilderServer - Where is the FenixGuiTestCaseBuilderServer running
var ExecutionLocationForFenixGuiTestCaseBuilderServer ExecutionLocationTypeType

// ExecutionLocationTypeType - Definition type for where the different programs are running
type ExecutionLocationTypeType int

// Constants used for where stuff is running
const (
	LocalhostNoDocker ExecutionLocationTypeType = iota
	LocalhostDocker
	GCP
)

// enixGuiBuilderProxyServer
var LocationForFenixGuiBuilderProxyServerTypeMapping = map[ExecutionLocationTypeType]string{
	LocalhostNoDocker: "LOCALHOST_NODOCKER",
	LocalhostDocker:   "LOCALHOST_DOCKER",
	GCP:               "GCP",
}

// Address and port to 'FenixGuiTestCaseBuilderServer' and 'FenixGuiExecutionServer', will get their values from Environment variables at startup
var (
	FenixGuiTestCaseBuilderServerAddress string //
	FenixGuiTestCaseBuilderServerPort    int    // TODO remove,

	FenixGuiBuilderProxyServerAddress string
	FenixGuiBuilderProxyServerPort    int
)

const numberOfCharactersfromUuid = 8

// ***********************************************************************************************************
