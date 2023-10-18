package main

// Variables that can be populated while building the application
// Used when building a executable to be run by itself

var (
	//GCP
	BuildVariableGCPAuthentication                            string
	BuildVariableUseServiceAccountForGuiExecutionServer       string
	BuildVariableUseServiceAccountForGuiTestCaseBuilderServer string
	BuildVariableAuthClientId                                 string
	BuildVariableAuthClientSecret                             string
	BuildVariableGcpProject                                   string
	BuildVariableTestExecutionStatusPubSubTopicBase           string
	BuildVariableLocalServiceAccountPath                      string

	// GUI-ExecutionServer
	BuildVariableExecutionLocationForFenixGuiExecutionServer string
	BuildVariableFenixGuiExecutionServerAddress              string
	BuildVariableFenixGuiExecutionServerPort                 string

	// GUI-TestCaseBuilderServer
	BuildVariableExecutionLocationForFenixGuiTestCaseBuilderServer string
	BuildVariableFenixGuiTestCaseBuilderServerAddress              string
	BuildVariableFenixGuiTestCaseBuilderServerPort                 string

	// This Application
	BuildVariableExecutionLocationForThisApplication string
	BuildVariableFYNE_SCALE                          string
	BuildVariableRunAsTrayApplication                string
	BuildVariableApplicationGrpcPort                 string
)

var buildVariablesMap = map[string]*string{
	//GCP
	"BuildVariableGCPAuthentication":                            &BuildVariableGCPAuthentication,
	"BuildVariableUseServiceAccountForGuiExecutionServer":       &BuildVariableUseServiceAccountForGuiExecutionServer,
	"BuildVariableUseServiceAccountForGuiTestCaseBuilderServer": &BuildVariableUseServiceAccountForGuiTestCaseBuilderServer,
	"BuildVariableAuthClientId":                                 &BuildVariableAuthClientId,
	"BuildVariableAuthClientSecret":                             &BuildVariableAuthClientSecret,
	"BuildVariableGcpProject":                                   &BuildVariableGcpProject,
	"BuildVariableTestExecutionStatusPubSubTopicBase":           &BuildVariableTestExecutionStatusPubSubTopicBase,
	"BuildVariableLocalServiceAccountPath":                      &BuildVariableLocalServiceAccountPath,

	// GUI-ExecutionServer
	"BuildVariableExecutionLocationForFenixGuiExecutionServer": &BuildVariableExecutionLocationForFenixGuiExecutionServer,
	"BuildVariableFenixGuiExecutionServerAddress":              &BuildVariableFenixGuiExecutionServerAddress,
	"BuildVariableFenixGuiExecutionServerPort":                 &BuildVariableFenixGuiExecutionServerPort,

	// GUI TestCaseBuilderServer
	"BuildVariableExecutionLocationForFenixGuiTestCaseBuilderServer": &BuildVariableExecutionLocationForFenixGuiTestCaseBuilderServer,
	"BuildVariableFenixGuiTestCaseBuilderServerAddress":              &BuildVariableFenixGuiTestCaseBuilderServerAddress,
	"BuildVariableFenixGuiTestCaseBuilderServerPort":                 &BuildVariableFenixGuiTestCaseBuilderServerPort,

	// This Application
	"BuildVariableExecutionLocationForThisApplication": &BuildVariableExecutionLocationForThisApplication,
	"BuildVariableFYNE_SCALE":                          &BuildVariableFYNE_SCALE,
	"BuildVariableRunAsTrayApplication":                &BuildVariableRunAsTrayApplication,
	"BuildVariableApplicationGrpcPort":                 &BuildVariableApplicationGrpcPort,
}
