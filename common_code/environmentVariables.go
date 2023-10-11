package sharedCode

// ***********************************************************************************************************
// The following variables receives their values from environment variables

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

	// GCP
	GCPAuthentication                            bool
	UseServiceAccountForGuiExecutionServer       bool
	UseServiceAccountForGuiTestCaseBuilderServer bool
	AuthClientId                                 string
	AuthClientSecret                             string

	// GUI-ExecutionServer
	ExecutionLocationForFenixGuiExecutionServer ExecutionLocationTypeType
	FenixGuiExecutionServerAddress              string
	FenixGuiExecutionServerPort                 int

	// GUI TestCaseBuilderServer
	ExecutionLocationForFenixGuiTestCaseBuilderServer ExecutionLocationTypeType
	FenixGuiTestCaseBuilderServerAddress              string
	FenixGuiTestCaseBuilderServerPort                 int

	// This Application
	ExecutionLocationForThisApplication ExecutionLocationTypeType
	FYNE_SCALE                          string
	//RunAsTrayApplication                bool
	ApplicationGrpcPort int

	// GCP-project
	GcpProject string

	// PubSub-Topic-base for where to receive the 'TestExecutionsStatus'
	TestExecutionStatusPubSubTopicBase string

	// Local path to Service-Account file
	LocalServiceAccountPath string

	//TODO REMOVE FenixGuiTestCaseBuilderProxyServerAddress string
	//TODO REMOVE FenixGuiTestCaseBuilderProxyServerPort    int
)

const numberOfCharactersfromUuid = 8

// ***********************************************************************************************************
