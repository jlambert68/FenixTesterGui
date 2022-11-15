package main

// Variables that can be populated while building the application
// Used when building a Tray Application

var (
	//GCP
	BuildVariableGCPAuthentication                      string
	BuildVariableUseServiceAccountForGuiExecutionServer string

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

var buildVariablesMap = map[string]string{
	//GCP
	"BuildVariableGCPAuthentication":                      "string",
	"BuildVariableUseServiceAccountForGuiExecutionServer": "string",

	// GUI-ExecutionServer
	"BuildVariableExecutionLocationForFenixGuiExecutionServer": "string",
	"BuildVariableFenixGuiExecutionServerAddress":              "string",
	"BuildVariableFenixGuiExecutionServerPort":                 "string",

	// GUI TestCaseBuilderServer
	"BuildVariableExecutionLocationForFenixGuiTestCaseBuilderServer": "string",
	"BuildVariableFenixGuiTestCaseBuilderServerAddress":              "string",
	"BuildVariableFenixGuiTestCaseBuilderServerPort":                 "string",

	// This Application
	"BuildVariableExecutionLocationForThisApplication": "string",
	"BuildVariableFYNE_SCALE":                          "string",
	"BuildVariableRunAsTrayApplication":                "string",
	"BuildVariableApplicationGrpcPort":                 "string",
}
